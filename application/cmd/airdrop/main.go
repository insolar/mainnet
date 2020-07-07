//  Copyright 2020 Insolar Network Ltd.
//  All rights reserved.
//  This material is licensed under the Insolar License version 1.0,
//  available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
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
	xecdsa "github.com/insolar/x-crypto/ecdsa"
	xelliptic "github.com/insolar/x-crypto/elliptic"
	"github.com/insolar/x-crypto/x509"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const ApplicationShortDescription string = "airdrop is tool for airdrop distribution at Insolar Platform"

const FractionsInXNS int64 = 10000000000
var (
	hexPrivate       string
	membersPath      string
	apiURL           string
	verbose          bool
	memberPrivateKey crypto.PrivateKey
	// request          *requester.ContractRequest
	airMembers       []AirMember
)

func parseInputParams(cmd *cobra.Command) {
	flags := cmd.Flags()
	flags.StringVarP(&hexPrivate, "hexPrivate", "p", "", "Private key pair")
	flags.StringVarP(&membersPath, "membersPath", "m", "", "Path to a file with members")
	flags.BoolVarP(&verbose, "verbose", "v", false, "Print request information")
	flags.BoolP("help", "h", false, "Help for airdrop")
}

type AirMember struct {
	Wallet string
	Tokens int64
}

func main() {
	err := AirdropCommand().Execute()
	if err != nil {
		log.Fatal("airdrop execution failed:", err)
	}
}

func getMemberPrivateKey(hexPrivate string) (crypto.PrivateKey, error)  {
	// Declare a new big int variable, specify the key as its value,
	// and set its format to base 16:
	i := new(big.Int)
	i.SetString(hexPrivate, 16)

	// Create a new elliptic curve and feed the value to it
	// to get the X and Y values of the public key:
	privateKey := new(xecdsa.PrivateKey)
	privateKey.PublicKey.Curve = xelliptic.P256K()
	privateKey.D = i
	privateKey.PublicKey.X, privateKey.PublicKey.Y = xelliptic.P256K().ScalarBaseMult(i.Bytes())

	// Convert the private key to PEM:
	x509Encoded, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, errors.Wrapf(err, "Fail convert private key to PEM")
	}
	pemPrivateKey := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	memberPrivateKey, err = secrets.ImportPrivateKeyPEM([]byte(string(pemPrivateKey)))
	if err != nil {
		return nil, errors.Wrapf(err, "Fail import private key")
	}
	return memberPrivateKey, nil
}

func getMemberReference(ctx context.Context, userConfig *requester.UserConfigJSON) (string, error)  {
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
	getmemberResponse, err := requester.Send(ctx, apiURL, userConfig, &getMemberRequest.Params)
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
	var err error
	memberPrivateKey, err = getMemberPrivateKey(hexPrivate)
	if err != nil {
		return errors.Wrapf(err, "Fail get from member private key")
	}

	// try to read airMembers
	// verify that the member airMembers paramsFile is exist
	if !isFileExists(membersPath) {
		return errors.New("File with members does not exists")
	}
	b, err := ioutil.ReadFile(membersPath)
	if err != nil {
		return errors.Wrapf(err, " couldn't read file %v", membersPath)
	}

	err = json.NewDecoder(bytes.NewReader(b)).Decode(&airMembers)
	if err != nil {
		return errors.Wrapf(err, "fail unmarshal data")
	}

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

func AirdropCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "airdrop <insolar_endpoint>",
		Short:   ApplicationShortDescription,
		Args:    requireUrlArg(),
		Example: "./airdrop http://localhost:19101/api/rpc -p '<from member hex private key>' -m members_test.json -v",
		RunE: func(_ *cobra.Command, args []string) error {
			// no need to check args size because of requireUrlArg
			apiURL = args[0]
			if len(apiURL) > 0 {
				ok, err := isUrl(apiURL)
				if !ok {
					return errors.Wrap(err, "URL parameter is incorrect. ")
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

			request := &requester.ContractRequest{
				Request: requester.Request{
					Version: "2.0",
					ID:      1,
					Method:  "contract.call",
				},
				Params: requester.Params{
					CallSite:  "member.transfer",
					Reference: reference,
					PublicKey: userConfig.PublicKey,
				},
			}

			for _, m := range airMembers {
				request.Params.CallParams = struct {
					Amount            string `json:"amount"`
					ToMemberReference string `json:"toMemberReference"`
				}{
					Amount:            big.NewInt(m.Tokens*FractionsInXNS).String(), // do we need multiply?
					ToMemberReference: m.Wallet,
				}
				response, err := requester.Send(ctx, apiURL, userConfig, &request.Params)
				if err != nil {
					return err
				}
				requestInfo := fmt.Sprintf("Request to member %s for transferring %s tokens: ", m.Wallet, big.NewInt(m.Tokens*FractionsInXNS).String())
				_, _ = os.Stdout.Write([]byte(requestInfo))
				_, _ = os.Stdout.Write(response)
				time.Sleep(time.Millisecond*500)
			}

			return nil
		},
	}
	parseInputParams(cmd)
	return cmd
}
