// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/insolar/blob/master/LICENSE.md.

package contracts

import (
	"encoding/json"
	"fmt"
	"github.com/insolar/insolar/applicationbase/genesis"
	"github.com/insolar/mainnet/application/builtin/contract/account"
	"github.com/insolar/mainnet/application/builtin/contract/costcenter"
	"github.com/insolar/mainnet/application/builtin/contract/deposit"
	"github.com/insolar/mainnet/application/builtin/contract/migrationadmin"
	"github.com/insolar/mainnet/application/builtin/contract/migrationdaemon"
	"github.com/insolar/mainnet/application/builtin/contract/migrationshard"
	"github.com/insolar/mainnet/application/builtin/contract/pkshard"
	"github.com/insolar/mainnet/application/builtin/contract/rootdomain"
	"github.com/insolar/mainnet/application/builtin/contract/wallet"
	"github.com/pkg/errors"
	"io/ioutil"
	"time"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/logicrunner/builtin/foundation"
	"github.com/insolar/insolar/pulse"
	"github.com/insolar/mainnet/application/appfoundation"

	"github.com/insolar/insolar/log"
	"github.com/insolar/mainnet/application/builtin/contract/member"
	. "github.com/insolar/mainnet/application/genesis"
)

const (
	XNS                        = "XNS"
	MigrationDaemonUnholdDate  = 1596456000 // 03.08.2020 12-00-00
	MigrationDaemonVesting     = 0
	MigrationDaemonVestingStep = 0

	NetworkIncentivesUnholdStartDate = 1583020800 // 01.03.2020 00-00-00
	NetworkIncentivesVesting         = 0
	NetworkIncentivesVestingStep     = 0

	ApplicationIncentivesUnholdStartDate = 1609459200 // 01.01.2021 00-00-00
	ApplicationIncentivesVesting         = 0
	ApplicationIncentivesVestingStep     = 0

	FoundationUnholdStartDate = 1609459200 // 01.01.2021 00-00-00
	FoundationVestingPeriod   = 0
	FoundationVestingStep     = 0
)

func InitStates(genesisConfigPath string) ([]genesis.ContractState, error) {
	b, err := ioutil.ReadFile(genesisConfigPath)
	if err != nil {
		log.Fatalf("failed to load genesis configuration from file: %v", genesisConfigPath)
	}
	var config struct {
		ContractsConfig GenesisContractsConfig
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Fatalf("failed to parse genesis configuration from file: %v", genesisConfigPath)
	}

	contractsConfig := config.ContractsConfig

	migrationAccounts := make(foundation.StableMap)
	migrationAccounts[XNS] = ContractMigrationAccount.String()

	migrationDeposits := make(foundation.StableMap)
	migrationDeposits[FundsDepositName] = ContractMigrationDeposit.String()

	ContractMigrationAddressShardRefs(contractsConfig.MAShardCount)
	ContractPublicKeyShardRefs(contractsConfig.PKShardCount)

	// Hint: order matters, because of dependency contracts on each other.
	states := []genesis.ContractState{
		rootDomain(contractsConfig.PKShardCount),
		GetMemberGenesisContractState(contractsConfig.RootPublicKey, GenesisNameRootMember, GenesisNameRootDomain, ContractRootWallet),
		GetMemberGenesisContractState(contractsConfig.MigrationAdminPublicKey, GenesisNameMigrationAdminMember, GenesisNameRootDomain, ContractMigrationWallet),
		GetMemberGenesisContractState(contractsConfig.FeePublicKey, GenesisNameFeeMember, GenesisNameRootDomain, ContractFeeWallet),

		GetWalletGenesisContractState(GenesisNameRootWallet, GenesisNameRootDomain, ContractRootAccount),
		GetPreWalletGenesisContractState(GenesisNameMigrationAdminWallet, GenesisNameRootDomain, migrationAccounts, migrationDeposits),
		GetWalletGenesisContractState(GenesisNameFeeWallet, GenesisNameRootDomain, ContractFeeAccount),

		GetAccountGenesisContractState(contractsConfig.RootBalance, GenesisNameRootAccount, GenesisNameRootDomain),
		GetAccountGenesisContractState("0", GenesisNameMigrationAdminAccount, GenesisNameRootDomain),
		GetAccountGenesisContractState("0", GenesisNameFeeAccount, GenesisNameRootDomain),

		GetDepositGenesisContractState(
			contractsConfig.MDBalance,
			MigrationDaemonVesting,
			MigrationDaemonVestingStep,
			appfoundation.Vesting2,
			pulse.OfUnixTime(MigrationDaemonUnholdDate), // Unhold date
			GenesisNameMigrationAdminDeposit,
			GenesisNameRootDomain,
		),
		GetMigrationAdminGenesisContractState(contractsConfig.LockupPeriodInPulses, contractsConfig.VestingPeriodInPulses, contractsConfig.VestingStepInPulses, contractsConfig.MAShardCount),
		GetCostCenterGenesisContractState(),
	}

	for i, key := range contractsConfig.MigrationDaemonPublicKeys {
		states = append(states, GetMemberGenesisContractState(key, GenesisNameMigrationDaemonMembers[i], GenesisNameRootDomain, *insolar.NewEmptyReference()))
		states = append(states, GetMigrationDaemonGenesisContractState(i))
	}

	for i, key := range contractsConfig.ApplicationIncentivesPublicKeys {
		states = append(states, GetMemberGenesisContractState(key, GenesisNameApplicationIncentivesMembers[i], GenesisNameRootDomain, ContractApplicationIncentivesWallets[i]))

		states = append(states, GetAccountGenesisContractState("0", GenesisNameApplicationIncentivesAccounts[i], GenesisNameRootDomain))

		unholdWithMonth := time.Unix(ApplicationIncentivesUnholdStartDate, 0).AddDate(0, i, 0).Unix()

		states = append(states, GetDepositGenesisContractState(
			AppIncentivesDistributionAmount,
			ApplicationIncentivesVesting,
			ApplicationIncentivesVestingStep,
			appfoundation.Vesting2,
			pulse.OfUnixTime(unholdWithMonth),
			GenesisNameApplicationIncentivesDeposits[i],
			GenesisNameRootDomain,
		))

		membersAccounts := make(foundation.StableMap)
		membersAccounts[XNS] = ContractApplicationIncentivesAccounts[i].String()

		membersDeposits := make(foundation.StableMap)
		membersDeposits[FundsDepositName] = ContractApplicationIncentivesDeposits[i].String()

		states = append(states, GetPreWalletGenesisContractState(
			GenesisNameApplicationIncentivesWallets[i],
			GenesisNameRootDomain,
			membersAccounts,
			membersDeposits,
		))
	}

	for i, key := range contractsConfig.NetworkIncentivesPublicKeys {
		states = append(states, GetMemberGenesisContractState(key, GenesisNameNetworkIncentivesMembers[i], GenesisNameRootDomain, ContractNetworkIncentivesWallets[i]))
		states = append(states, GetAccountGenesisContractState("0", GenesisNameNetworkIncentivesAccounts[i], GenesisNameRootDomain))

		unholdWithMonth := time.Unix(NetworkIncentivesUnholdStartDate, 0).AddDate(0, i, 0).Unix()

		states = append(states, GetDepositGenesisContractState(
			NetworkIncentivesDistributionAmount,
			NetworkIncentivesVesting,
			NetworkIncentivesVestingStep,
			appfoundation.Vesting2,
			pulse.OfUnixTime(unholdWithMonth),
			GenesisNameNetworkIncentivesDeposits[i],
			GenesisNameRootDomain,
		))

		membersAccounts := make(foundation.StableMap)
		membersAccounts[XNS] = ContractNetworkIncentivesAccounts[i].String()

		membersDeposits := make(foundation.StableMap)
		membersDeposits[FundsDepositName] = ContractNetworkIncentivesDeposits[i].String()

		states = append(states, GetPreWalletGenesisContractState(
			GenesisNameNetworkIncentivesWallets[i],
			GenesisNameRootDomain,
			membersAccounts,
			membersDeposits,
		))
	}

	for i, key := range contractsConfig.FoundationPublicKeys {
		states = append(states, GetMemberGenesisContractState(key, GenesisNameFoundationMembers[i], GenesisNameRootDomain, ContractFoundationWallets[i]))
		states = append(states, GetAccountGenesisContractState("0", GenesisNameFoundationAccounts[i], GenesisNameRootDomain))

		unholdWithMonth := time.Unix(FoundationUnholdStartDate, 0).AddDate(0, i, 0).Unix()

		states = append(states, GetDepositGenesisContractState(
			FoundationDistributionAmount,
			FoundationVestingPeriod,
			FoundationVestingStep,
			appfoundation.Vesting2,
			pulse.OfUnixTime(unholdWithMonth),
			GenesisNameFoundationDeposits[i],
			GenesisNameRootDomain,
		))

		membersAccounts := make(foundation.StableMap)
		membersAccounts[XNS] = ContractFoundationAccounts[i].String()

		membersDeposits := make(foundation.StableMap)
		membersDeposits[FundsDepositName] = ContractFoundationDeposits[i].String()

		states = append(states, GetPreWalletGenesisContractState(
			GenesisNameFoundationWallets[i],
			GenesisNameRootDomain,
			membersAccounts,
			membersDeposits,
		))
	}

	for i, key := range contractsConfig.EnterprisePublicKeys {
		states = append(states, GetMemberGenesisContractState(key, GenesisNameEnterpriseMembers[i], GenesisNameRootDomain, ContractEnterpriseWallets[i]))
		states = append(states, GetAccountGenesisContractState(
			EnterpriseDistributionAmount,
			GenesisNameEnterpriseAccounts[i],
			GenesisNameRootDomain,
		))

		membersAccounts := make(foundation.StableMap)
		membersAccounts[XNS] = ContractEnterpriseAccounts[i].String()

		membersDeposits := make(foundation.StableMap)

		states = append(states, GetPreWalletGenesisContractState(
			GenesisNameEnterpriseWallets[i],
			GenesisNameRootDomain,
			membersAccounts,
			membersDeposits,
		))
	}

	if contractsConfig.PKShardCount <= 0 {
		panic(fmt.Sprintf("[genesis] store contracts failed: setup pk_shard_count parameter, current value %v", contractsConfig.PKShardCount))
	}
	if contractsConfig.VestingStepInPulses > 0 && contractsConfig.VestingPeriodInPulses%contractsConfig.VestingStepInPulses != 0 {
		panic(fmt.Sprintf("[genesis] store contracts failed: vesting_pulse_period (%d) is not a multiple of vesting_pulse_step (%d)", contractsConfig.VestingPeriodInPulses, contractsConfig.VestingStepInPulses))
	}

	// Split genesis members by PK shards
	var membersByPKShards []foundation.StableMap
	for i := 0; i < contractsConfig.PKShardCount; i++ {
		membersByPKShards = append(membersByPKShards, make(foundation.StableMap))
	}
	trimmedRootPublicKey, err := foundation.ExtractCanonicalPublicKey(contractsConfig.RootPublicKey)
	if err != nil {
		panic(errors.Wrapf(err, "[genesis] extracting canonical pk failed, current value %v", contractsConfig.RootPublicKey))
	}
	index := foundation.GetShardIndex(trimmedRootPublicKey, contractsConfig.PKShardCount)
	membersByPKShards[index][trimmedRootPublicKey] = ContractRootMember.String()

	trimmedMigrationAdminPublicKey, err := foundation.ExtractCanonicalPublicKey(contractsConfig.MigrationAdminPublicKey)
	if err != nil {
		panic(errors.Wrapf(err, "[genesis] extracting canonical pk failed, current value %v", contractsConfig.MigrationAdminPublicKey))
	}
	index = foundation.GetShardIndex(trimmedMigrationAdminPublicKey, contractsConfig.PKShardCount)
	membersByPKShards[index][trimmedMigrationAdminPublicKey] = ContractMigrationAdminMember.String()

	trimmedFeeMemberPublicKey, err := foundation.ExtractCanonicalPublicKey(contractsConfig.FeePublicKey)
	if err != nil {
		panic(errors.Wrapf(err, "[genesis] extracting canonical pk failed, current value %v", contractsConfig.FeePublicKey))
	}
	index = foundation.GetShardIndex(trimmedFeeMemberPublicKey, contractsConfig.PKShardCount)
	membersByPKShards[index][trimmedFeeMemberPublicKey] = ContractFeeMember.String()

	for i, key := range contractsConfig.MigrationDaemonPublicKeys {
		trimmedMigrationDaemonPublicKey, err := foundation.ExtractCanonicalPublicKey(key)
		if err != nil {
			panic(errors.Wrapf(err, "[genesis] extracting canonical pk failed, current value %v", key))
		}
		index := foundation.GetShardIndex(trimmedMigrationDaemonPublicKey, contractsConfig.PKShardCount)
		membersByPKShards[index][trimmedMigrationDaemonPublicKey] = ContractMigrationDaemonMembers[i].String()
	}

	for i, key := range contractsConfig.NetworkIncentivesPublicKeys {
		trimmedNetworkIncentivesPublicKey, err := foundation.ExtractCanonicalPublicKey(key)
		if err != nil {
			panic(errors.Wrapf(err, "[genesis] extracting canonical pk failed, current value %v", key))
		}
		index := foundation.GetShardIndex(trimmedNetworkIncentivesPublicKey, contractsConfig.PKShardCount)
		membersByPKShards[index][trimmedNetworkIncentivesPublicKey] = ContractNetworkIncentivesMembers[i].String()
	}

	for i, key := range contractsConfig.ApplicationIncentivesPublicKeys {
		trimmedApplicationIncentivesPublicKey, err := foundation.ExtractCanonicalPublicKey(key)
		if err != nil {
			panic(errors.Wrapf(err, "[genesis] extracting canonical pk failed, current value %v", key))
		}
		index := foundation.GetShardIndex(trimmedApplicationIncentivesPublicKey, contractsConfig.PKShardCount)
		membersByPKShards[index][trimmedApplicationIncentivesPublicKey] = ContractApplicationIncentivesMembers[i].String()
	}

	for i, key := range contractsConfig.FoundationPublicKeys {
		trimmedFoundationPublicKey, err := foundation.ExtractCanonicalPublicKey(key)
		if err != nil {
			panic(errors.Wrapf(err, "[genesis] extracting canonical pk failed, current value %v", key))
		}
		index := foundation.GetShardIndex(trimmedFoundationPublicKey, contractsConfig.PKShardCount)
		membersByPKShards[index][trimmedFoundationPublicKey] = ContractFoundationMembers[i].String()
	}

	for i, key := range contractsConfig.EnterprisePublicKeys {
		trimmedEnterprisePublicKey, err := foundation.ExtractCanonicalPublicKey(key)
		if err != nil {
			panic(errors.Wrapf(err, "[genesis] extracting canonical pk failed, current value %v", key))
		}
		index := foundation.GetShardIndex(trimmedEnterprisePublicKey, contractsConfig.PKShardCount)
		membersByPKShards[index][trimmedEnterprisePublicKey] = ContractEnterpriseMembers[i].String()
	}

	// Append states for shards
	for i, name := range ContractPublicKeyNameShards(contractsConfig.PKShardCount) {
		states = append(states, GetPKShardGenesisContractState(name, membersByPKShards[i]))
	}
	for i, name := range ContractMigrationAddressNameShards(contractsConfig.MAShardCount) {
		states = append(states, GetMigrationShardGenesisContractState(name, contractsConfig.MigrationAddresses[i]))
	}

	return states, nil
}

func rootDomain(pkShardCount int) genesis.ContractState {

	return genesis.ContractState{
		Name:       GenesisNameRootDomain,
		Prototype:  GenesisNameRootDomain,
		ParentName: "",

		Memory: genesis.MustGenMemory(&rootdomain.RootDomain{
			PublicKeyShards: ContractPublicKeyShards(pkShardCount),
		}),
	}
}

func GetMemberGenesisContractState(publicKey string, name string, parent string, walletRef insolar.Reference) genesis.ContractState {
	m, err := member.New(publicKey, "", *insolar.NewEmptyReference())
	if err != nil {
		panic(fmt.Sprintf("'%s' member constructor failed", name))
	}

	m.Wallet = walletRef

	return genesis.ContractState{
		Name:       name,
		Prototype:  GenesisNameMember,
		ParentName: parent,
		Memory:     genesis.MustGenMemory(m),
	}
}

func GetWalletGenesisContractState(name string, parent string, accountRef insolar.Reference) genesis.ContractState {
	w, err := wallet.New(accountRef)
	if err != nil {
		panic("failed to create ` " + name + "` wallet instance")
	}

	return genesis.ContractState{
		Name:       name,
		Prototype:  GenesisNameWallet,
		ParentName: parent,
		Memory:     genesis.MustGenMemory(w),
	}
}

func GetPreWalletGenesisContractState(name string, parent string, accounts foundation.StableMap, deposits foundation.StableMap) genesis.ContractState {
	return genesis.ContractState{
		Name:       name,
		Prototype:  GenesisNameWallet,
		ParentName: parent,
		Memory: genesis.MustGenMemory(&wallet.Wallet{
			Accounts: accounts,
			Deposits: deposits,
		}),
	}
}

func GetAccountGenesisContractState(balance string, name string, parent string) genesis.ContractState {
	w, err := account.New(balance)
	if err != nil {
		panic("failed to create ` " + name + "` account instance")
	}

	return genesis.ContractState{
		Name:       name,
		Prototype:  GenesisNameAccount,
		ParentName: parent,
		Memory:     genesis.MustGenMemory(w),
	}
}

func GetCostCenterGenesisContractState() genesis.ContractState {
	cc, err := costcenter.New(&ContractFeeMember)
	if err != nil {
		panic("failed to create cost center instance")
	}

	return genesis.ContractState{
		Name:       GenesisNameCostCenter,
		Prototype:  GenesisNameCostCenter,
		ParentName: GenesisNameRootDomain,
		Memory:     genesis.MustGenMemory(cc),
	}
}

func GetPKShardGenesisContractState(name string, members foundation.StableMap) genesis.ContractState {
	s, err := pkshard.New(members)
	if err != nil {
		panic(fmt.Sprintf("'%s' shard constructor failed", name))
	}

	return genesis.ContractState{
		Name:       name,
		Prototype:  GenesisNamePKShard,
		ParentName: GenesisNameRootDomain,
		Memory:     genesis.MustGenMemory(s),
	}
}

func GetMigrationShardGenesisContractState(name string, migrationAddresses []string) genesis.ContractState {
	s, err := migrationshard.New(migrationAddresses)
	if err != nil {
		panic(fmt.Sprintf("'%s' shard constructor failed", name))
	}

	return genesis.ContractState{
		Name:       name,
		Prototype:  GenesisNameMigrationShard,
		ParentName: GenesisNameRootDomain,
		Memory:     genesis.MustGenMemory(s),
	}
}

func GetMigrationAdminGenesisContractState(lockup int64, vesting int64, vestingStep int64, maShardCount int) genesis.ContractState {
	return genesis.ContractState{
		Name:       GenesisNameMigrationAdmin,
		Prototype:  GenesisNameMigrationAdmin,
		ParentName: GenesisNameRootDomain,
		Memory: genesis.MustGenMemory(&migrationadmin.MigrationAdmin{
			MigrationAddressShards: ContractMigrationAddressShards(maShardCount),
			MigrationAdminMember:   ContractMigrationAdminMember,
			VestingParams: &migrationadmin.VestingParams{
				Lockup:      lockup,
				Vesting:     vesting,
				VestingStep: vestingStep,
			},
		}),
	}
}

func GetDepositGenesisContractState(
	amount string,
	vesting int64,
	vestingStep int64,
	vestingType appfoundation.VestingType,
	pulseDepositUnHold insolar.PulseNumber,
	name string, parent string,
) genesis.ContractState {
	return genesis.ContractState{
		Name:       name,
		Prototype:  GenesisNameDeposit,
		ParentName: parent,
		Memory: genesis.MustGenMemory(&deposit.Deposit{
			Balance:            amount,
			Amount:             amount,
			PulseDepositUnHold: pulseDepositUnHold,
			VestingType:        vestingType,
			TxHash:             FundsDepositName,
			Lockup:             int64(pulseDepositUnHold - pulse.MinTimePulse),
			Vesting:            vesting,
			VestingStep:        vestingStep,
			IsConfirmed:        true,
		}),
	}
}

func GetMigrationDaemonGenesisContractState(numberMigrationDaemon int) genesis.ContractState {

	return genesis.ContractState{
		Name:       GenesisNameMigrationDaemons[numberMigrationDaemon],
		Prototype:  GenesisNameMigrationDaemon,
		ParentName: GenesisNameRootDomain,
		Memory: genesis.MustGenMemory(&migrationdaemon.MigrationDaemon{
			IsActive:              false,
			MigrationDaemonMember: ContractMigrationDaemonMembers[numberMigrationDaemon],
		}),
	}
}
