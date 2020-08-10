// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

package functest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/insolar/insolar/insolar/defaults"
	yaml "gopkg.in/yaml.v2"

	"github.com/pkg/errors"

	"github.com/insolar/insolar/applicationbase/testutils/launchnet"

	"github.com/insolar/mainnet/application"
	"github.com/insolar/mainnet/application/builtin/contract/deposit"
	"github.com/insolar/mainnet/application/genesisrefs"
	"github.com/insolar/mainnet/application/sdk"
)

const insolarRootMemberKeys = "root_member_keys.json"
const insolarMigrationAdminMemberKeys = "migration_admin_member_keys.json"
const insolarFeeMemberKeys = "fee_member_keys.json"

var ApplicationIncentives [application.GenesisAmountApplicationIncentivesMembers]*AppUser
var NetworkIncentives [application.GenesisAmountNetworkIncentivesMembers]*AppUser
var Enterprise [application.GenesisAmountEnterpriseMembers]*AppUser
var Foundation [application.GenesisAmountFoundationMembers]*AppUser

var AppPath = []string{"insolar", "mainnet"}

var info *sdk.InfoResponse
var Root AppUser
var MigrationAdmin AppUser
var FeeMember AppUser
var MigrationDaemons [application.GenesisAmountMigrationDaemonMembers]*AppUser

type AppUser struct {
	Ref              string
	PrivKey          string
	PubKey           string
	MigrationAddress string
}

func (m *AppUser) GetReference() string {
	return m.Ref
}

func (m *AppUser) GetPrivateKey() string {
	return m.PrivKey
}

func (m *AppUser) GetPublicKey() string {
	return m.PubKey
}

func GetNumShards() (int, error) {
	type bootstrapConf struct {
		PKShardCount int `yaml:"ma_shard_count"`
	}

	var conf bootstrapConf

	path, err := launchnet.LaunchnetPath(AppPath, "bootstrap.yaml")
	if err != nil {
		return 0, err
	}
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, errors.Wrap(err, "[ GetNumShards ] Can't read bootstrap config")
	}

	err = yaml.Unmarshal(buff, &conf)
	if err != nil {
		return 0, errors.Wrap(err, "[ GetNumShards ] Can't parse bootstrap config")
	}

	return conf.PKShardCount, nil
}

func loadMemberKeys(keysPath string, member *AppUser) error {
	text, err := ioutil.ReadFile(keysPath)
	if err != nil {
		return errors.Wrapf(err, "[ loadMemberKeys ] could't load member keys")
	}
	var data map[string]string
	err = json.Unmarshal(text, &data)
	if err != nil {
		return errors.Wrapf(err, "[ loadMemberKeys ] could't unmarshal member keys")
	}
	if data["private_key"] == "" || data["public_key"] == "" {
		return errors.New("[ loadMemberKeys ] could't find any keys")
	}
	member.PrivKey = data["private_key"]
	member.PubKey = data["public_key"]

	return nil
}

func LoadAllMembersKeys() error {
	path, err := launchnet.LaunchnetPath(AppPath, "configs", insolarRootMemberKeys)
	if err != nil {
		return err
	}
	err = loadMemberKeys(path, &Root)
	if err != nil {
		return err
	}
	path, err = launchnet.LaunchnetPath(AppPath, "configs", insolarFeeMemberKeys)
	if err != nil {
		return err
	}
	err = loadMemberKeys(path, &FeeMember)
	if err != nil {
		return err
	}
	path, err = launchnet.LaunchnetPath(AppPath, "configs", insolarMigrationAdminMemberKeys)
	if err != nil {
		return err
	}
	err = loadMemberKeys(path, &MigrationAdmin)
	if err != nil {
		return err
	}
	for i := range MigrationDaemons {
		path, err := launchnet.LaunchnetPath(AppPath, "configs", "migration_daemon_"+strconv.Itoa(i)+"_member_keys.json")
		if err != nil {
			return err
		}
		var md AppUser
		err = loadMemberKeys(path, &md)
		if err != nil {
			return err
		}
		MigrationDaemons[i] = &md
	}

	for i := 0; i < application.GenesisAmountApplicationIncentivesMembers; i++ {
		path, err := launchnet.LaunchnetPath(AppPath, "configs", "application_incentives_"+strconv.Itoa(i)+"_member_keys.json")
		if err != nil {
			return err
		}
		var md AppUser
		err = loadMemberKeys(path, &md)
		if err != nil {
			return err
		}
		ApplicationIncentives[i] = &md
	}

	for i := 0; i < application.GenesisAmountNetworkIncentivesMembers; i++ {
		path, err := launchnet.LaunchnetPath(AppPath, "configs", "network_incentives_"+strconv.Itoa(i)+"_member_keys.json")
		if err != nil {
			return err
		}
		var md AppUser
		err = loadMemberKeys(path, &md)
		if err != nil {
			return err
		}
		NetworkIncentives[i] = &md
	}

	for i := 0; i < application.GenesisAmountFoundationMembers; i++ {
		path, err := launchnet.LaunchnetPath(AppPath, "configs", "foundation_"+strconv.Itoa(i)+"_member_keys.json")
		if err != nil {
			return err
		}
		var md AppUser
		err = loadMemberKeys(path, &md)
		if err != nil {
			return err
		}
		Foundation[i] = &md
	}

	for i := 0; i < application.GenesisAmountEnterpriseMembers; i++ {
		path, err := launchnet.LaunchnetPath(AppPath, "configs", "enterprise_"+strconv.Itoa(i)+"_member_keys.json")
		if err != nil {
			return err
		}
		var md AppUser
		err = loadMemberKeys(path, &md)
		if err != nil {
			return err
		}
		Enterprise[i] = &md
	}

	return nil
}

func SetInfo() error {
	var err error
	info, err = sdk.Info(launchnet.TestRPCUrl)
	if err != nil {
		return errors.Wrap(err, "[ setInfo ] error sending request")
	}

	err = PostGenesis()
	if err != nil {
		return errors.Wrap(err, "[ setInfo ] failed to execute post genesis")
	}
	return nil
}

func AfterSetup() {
	Root.Ref = info.RootMember
	MigrationAdmin.Ref = info.MigrationAdminMember
}

func PostGenesis() error {
	fmt.Println("[ PostGenesis ] starting...")
	err := preparePublicAllocation2()
	if err != nil {
		return errors.Wrap(err, "failed to create fund public allocation 2")
	}
	return nil
}

func preparePublicAllocation2() error {
	insSDK, err := sdk.NewSDK(
		[]string{launchnet.TestRPCUrl},
		[]string{launchnet.TestRPCUrlPublic},
		defaults.LaunchnetConfigDir(),
		sdk.DefaultOptions)
	if err != nil {
		return errors.Wrap(err, "SDK is not initialized")
	}

	lockupEndDate := time.Now().Unix()
	_, err = insSDK.CreateFund(strconv.FormatInt(lockupEndDate, 10))
	if err != nil {
		return errors.Wrap(err, "failed to call deposit.createFund")
	}

	migrationAdmin := insSDK.GetMigrationAdminMember()

	for i, donor := range insSDK.GetEnterpriseMembers() {
		err := transferFromAccountToDeposit(
			insSDK,
			donor,
			deposit.PublicAllocation2DepositName,
			migrationAdmin.GetReference(),
			application.EnterpriseDistributionAmount,
		)
		if err != nil {
			return errors.Wrapf(err, "failed to transfer money from enterprise member #%d", i)
		}
	}

	for i, donor := range insSDK.GetApplicationIncentivesMembers() {
		err = transferFromDepositToDeposit(
			insSDK,
			donor,
			genesisrefs.FundsDepositName,
			deposit.PublicAllocation2DepositName,
			migrationAdmin.GetReference(),
		)
		if err != nil {
			return errors.Wrapf(err, "failed to transfer money from application incentives member #%d", i)
		}
	}

	for i, donor := range insSDK.GetFoundationMembers() {
		err = transferFromDepositToDeposit(
			insSDK,
			donor,
			genesisrefs.FundsDepositName,
			deposit.PublicAllocation2DepositName,
			migrationAdmin.GetReference(),
		)
		if err != nil {
			return errors.Wrapf(err, "failed to transfer money from foundation member #%d", i)
		}
	}

	for i, donor := range insSDK.GetNetworkIncentivesMembers() {
		err = transferFromDepositToDeposit(
			insSDK,
			donor,
			genesisrefs.FundsDepositName,
			deposit.PublicAllocation2DepositName,
			migrationAdmin.GetReference(),
		)
		if err != nil {
			return errors.Wrapf(err, "failed to transfer money from network incentives member #%d", i)
		}
	}

	return nil
}

func transferFromDepositToDeposit(insSDK *sdk.SDK,
	from sdk.Member,
	fromDepositName string,
	toDepositName string,
	toMemberRef string) error {

	_, err := insSDK.TransferFromDepositToDeposit(from, fromDepositName, toDepositName, toMemberRef)
	if err != nil {
		return errors.Wrap(err, "failed to call deposit.transferToDeposit")
	}
	return nil
}

func transferFromAccountToDeposit(insSDK *sdk.SDK,
	from sdk.Member,
	toDepositName string,
	toMemberRef string,
	amount string,
) error {

	_, err := insSDK.TransferFromAccountToDeposit(from, toDepositName, toMemberRef, amount)
	if err != nil {
		return errors.Wrap(err, "failed to call account.transferToDeposit")
	}
	return nil
}
