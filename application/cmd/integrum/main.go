//  Copyright 2020 Insolar Network Ltd.
//  All rights reserved.
//  This material is licensed under the Insolar License version 1.0,
//  available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg/v9"
	"net/url"
	"os"
	"time"

	"github.com/insolar/insolar/api/requester"
	"github.com/insolar/insolar/configuration"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/secrets"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/log"
	"github.com/insolar/insolar/log/logadapter"
	crypto "github.com/insolar/x-crypto"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const ApplicationShortDescription string = "integrum is tool for creating additional deposit in the members wallet at Insolar Platform"

var (
	waitTime         int64
	apiURL           string
	dbURL            string
	publicURL        string
	memberKeysPath   string
	verbose          bool
	memberPrivateKey crypto.PrivateKey
)

func parseInputParams(cmd *cobra.Command) {
	flags := cmd.Flags()
	flags.StringVarP(&memberKeysPath, "memberkeys", "k", "", "Path to a key pair")
	flags.StringVarP(&dbURL, "dbUrl", "u", "", "URL with credential to the DB")
	flags.StringVarP(&publicURL, "publicUrl", "p", "", "URL to the public Insolar API")
	flags.Int64VarP(&waitTime, "waitTime", "w", 500, "Milliseconds to wait after each transfer")
	flags.BoolVarP(&verbose, "verbose", "v", false, "Print request information")
	flags.BoolP("help", "h", false, "Help for integrum")
}

type DepositMember struct {
	DepositReference []byte `sql:"deposit_ref"`
	MemberReference  []byte `sql:"member_ref"`
}

func main() {
	err := DepositCreatorCommand().Execute()
	if err != nil {
		log.Fatal("integrum execution failed:", err)
	}
}

func getMemberReference(ctx context.Context, userConfig *requester.UserConfigJSON) (string, error) {
	getMemberRequest := &requester.ContractRequest{
		Request: requester.Request{
			Version: "2.0",
			ID:      1,
			Method:  "contract.call",
		},
		Params: requester.Params{
			CallSite:  "member.get",
			PublicKey: userConfig.PublicKey,
		},
	}
	getmemberResponse, err := requester.Send(ctx, publicURL, userConfig, &getMemberRequest.Params)
	if err != nil {
		return "", err
	}
	resp := requester.ContractResponse{}
	err = json.Unmarshal(getmemberResponse, &resp)
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", resp.Error
	}

	if resp.Result == nil {
		return "", errors.New("Error and result are nil")
	}
	return resp.Result.CallResult.(map[string]interface{})["reference"].(string), nil
}

func verifyParams() error {
	// verify that the member keys paramsFile is exist
	if !isFileExists(memberKeysPath) {
		return errors.New("Member keys does not exists")
	}

	// try to read keys
	keys, err := secrets.ReadXCryptoKeysFile(memberKeysPath, false)
	if err != nil {
		return errors.Wrap(err, "Cannot parse member keys. ")
	}
	memberPrivateKey = keys.Private
	return nil
}

func isUrl(str string) (bool, error) {
	parsedUrl, err := url.Parse(str)
	return err == nil && parsedUrl.Scheme != "" && parsedUrl.Host != "", err
}

func isFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getContextWithLogger() context.Context {
	cfg := configuration.NewLog()
	ctx := context.Background()
	cfg.Formatter = "text"
	if verbose {
		cfg.Level = insolar.DebugLevel.String()
	} else {
		cfg.Level = insolar.WarnLevel.String()
	}

	defaultCfg := logadapter.DefaultLoggerSettings()
	defaultCfg.Instruments.CallerMode = insolar.NoCallerField
	defaultCfg.Instruments.MetricsMode = insolar.NoLogMetrics
	logger, _ := log.NewLogExt(cfg, defaultCfg, 0)
	ctx = inslogger.SetLogger(ctx, logger)
	log.SetGlobalLogger(logger)

	return ctx
}

func createUserConfig(privateKey crypto.PrivateKey) (*requester.UserConfigJSON, error) {
	privateKeyBytes, err := secrets.ExportPrivateKeyPEM(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to export private key")
	}
	privateKeyStr := string(privateKeyBytes)

	publicKey, err := secrets.ExportPublicKeyPEM(secrets.ExtractPublicKey(privateKey))
	if err != nil {
		return nil, errors.Wrap(err, "failed to extract public key")
	}
	publicKeyStr := string(publicKey)

	return requester.CreateUserConfig("", privateKeyStr, publicKeyStr)
}

// requireUrlArg returns an error if there is not url args.
func requireUrlArg() cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("The program required url as an argument")
		}
		return nil
	}
}

func DepositCreatorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "integrum <insolar_endpoint>",
		Short:   ApplicationShortDescription,
		Args:    requireUrlArg(),
		Example: "./integrum http://localhost:19001/admin-api/rpc -p http://localhost:19101/api/rpc -k 'Path to a key pair' -u postgresql://localhost:5432/postgres?sslmode=disable -v",
		RunE: func(_ *cobra.Command, args []string) error {
			// no need to check args size because of requireUrlArg
			apiURL = args[0]
			if len(apiURL) > 0 {
				ok, err := isUrl(apiURL)
				if !ok {
					return errors.Wrap(err, "URL parameter is incorrect.")
				}
			}

			e := verifyParams()
			if e != nil {
				return e
			}
			ctx := getContextWithLogger()
			requester.SetVerbose(verbose)

			userConfig, e := createUserConfig(memberPrivateKey)
			if e != nil {
				return e
			}

			reference, e := getMemberReference(ctx, userConfig)
			if e != nil {
				return errors.Wrap(e, "fail to get from member reference")
			}
			userConfig.Caller = reference

			opt, err := pg.ParseURL(dbURL)
			if err != nil {
				panic(err)
			}

			db := pg.Connect(opt)
			defer db.Close()
			var members []DepositMember
			rows, err := db.Query(&members, "SELECT d.deposit_ref, d.member_ref FROM deposits d "+
				"WHERE d.eth_hash != 'genesis_deposit' AND d.eth_hash != 'genesis_deposit2' "+
				"AND d.eth_hash NOT IN (SELECT d1.eth_hash FROM deposits d1 WHERE d1.eth_hash like '%_2');")
			if err != nil {
				panic(err)
			}

			if rows.RowsAffected() == 0 {
				return errors.New("Members without second deposit not found")
			}
			fmt.Println(rows.RowsAffected())

			request := &requester.ContractRequest{
				Request: requester.Request{
					Version: "2.0",
					Method:  "contract.call",
				},
				Params: requester.Params{
					CallSite:  "deposit.create",
					Reference: reference,
					PublicKey: userConfig.PublicKey,
				},
			}

			membersLen := len(members)
			for i, m := range members {
				request.ID = uint64(i)
				request.Params.CallParams = struct {
					DepositReference string `json:"depositReference"`
					MemberReference  string `json:"memberReference"`
				}{
					DepositReference: insolar.NewReferenceFromBytes(m.DepositReference).String(),
					MemberReference:  insolar.NewReferenceFromBytes(m.MemberReference).String(),
				}
				response, err := requester.Send(ctx, apiURL, userConfig, &request.Params)
				if err != nil {
					return err
				}
				requestInfo := fmt.Sprintf("[%d of %d] Request to member %s for deposit creation %s: ", i+1, membersLen, insolar.NewReferenceFromBytes(m.MemberReference).String(), insolar.NewReferenceFromBytes(m.MemberReference).String())
				_, _ = os.Stdout.Write([]byte(requestInfo))
				_, _ = os.Stdout.Write(response)
				time.Sleep(time.Millisecond * time.Duration(waitTime))
			}

			return nil
		},
	}
	parseInputParams(cmd)
	return cmd
}
