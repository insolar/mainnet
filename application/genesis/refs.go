package genesis

import (
	"github.com/insolar/insolar/applicationbase/genesisrefs"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/logicrunner/preprocessor"
	"strconv"
)

const (
	FundsDepositName = "genesis_deposit"
)

var applicationPrototypes = map[string]insolar.Reference{
	GenesisNameRootDomain + genesisrefs.PrototypeSuffix:            *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameRootDomain, 0),
	GenesisNameRootMember + genesisrefs.PrototypeSuffix:            *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameMember, 0),
	GenesisNameRootWallet + genesisrefs.PrototypeSuffix:            *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameWallet, 0),
	GenesisNameRootAccount + genesisrefs.PrototypeSuffix:           *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameAccount, 0),
	GenesisNameCostCenter + genesisrefs.PrototypeSuffix:            *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameCostCenter, 0),
	GenesisNameFeeMember + genesisrefs.PrototypeSuffix:             *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameMember, 0),
	GenesisNameFeeWallet + genesisrefs.PrototypeSuffix:             *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameWallet, 0),
	GenesisNameFeeAccount + genesisrefs.PrototypeSuffix:            *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameAccount, 0),
	GenesisNameDeposit + genesisrefs.PrototypeSuffix:               *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameDeposit, 0),
	GenesisNameMember + genesisrefs.PrototypeSuffix:                *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameMember, 0),
	GenesisNameMigrationAdminMember + genesisrefs.PrototypeSuffix:  *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameMember, 0),
	GenesisNameMigrationAdmin + genesisrefs.PrototypeSuffix:        *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameMigrationAdmin, 0),
	GenesisNameMigrationAdminWallet + genesisrefs.PrototypeSuffix:  *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameWallet, 0),
	GenesisNameMigrationAdminAccount + genesisrefs.PrototypeSuffix: *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameAccount, 0),
	GenesisNameMigrationAdminDeposit + genesisrefs.PrototypeSuffix: *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameDeposit, 0),
	GenesisNameWallet + genesisrefs.PrototypeSuffix:                *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameWallet, 0),
}

func init() {
	for i, val := range applicationPrototypes {
		genesisrefs.PredefinedPrototypes[i] = val
	}

	for _, el := range GenesisNameMigrationDaemonMembers {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameMember, 0)
	}

	for _, el := range GenesisNameMigrationDaemons {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameMigrationDaemon, 0)
	}

	// Incentives Application
	for _, el := range GenesisNameApplicationIncentivesMembers {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameMember, 0)
	}
	for _, el := range GenesisNameApplicationIncentivesWallets {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameWallet, 0)
	}
	for _, el := range GenesisNameApplicationIncentivesAccounts {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameAccount, 0)
	}
	for _, el := range GenesisNameApplicationIncentivesDeposits {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameDeposit, 0)
	}

	// Network Incentives
	for _, el := range GenesisNameNetworkIncentivesMembers {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameMember, 0)
	}
	for _, el := range GenesisNameNetworkIncentivesWallets {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameWallet, 0)
	}
	for _, el := range GenesisNameNetworkIncentivesAccounts {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameAccount, 0)
	}
	for _, el := range GenesisNameNetworkIncentivesDeposits {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameDeposit, 0)
	}

	// Foundation
	for _, el := range GenesisNameFoundationMembers {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameMember, 0)
	}
	for _, el := range GenesisNameFoundationWallets {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameWallet, 0)
	}
	for _, el := range GenesisNameFoundationAccounts {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameAccount, 0)
	}
	for _, el := range GenesisNameFoundationDeposits {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameDeposit, 0)
	}

	// Enterprise
	for _, el := range GenesisNameEnterpriseMembers {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameMember, 0)
	}
	for _, el := range GenesisNameEnterpriseWallets {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameWallet, 0)
	}
	for _, el := range GenesisNameEnterpriseAccounts {
		genesisrefs.PredefinedPrototypes[el+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(genesisrefs.PrototypeType, GenesisNameAccount, 0)
	}
}

var (
	// ContractRootDomain is the root domain contract reference.
	ContractRootDomain = genesisrefs.GenesisRef(GenesisNameRootDomain)
	// ContractRootMember is the root member contract reference.
	ContractRootMember = genesisrefs.GenesisRef(GenesisNameRootMember)
	// ContractRootWallet is the root wallet contract reference.
	ContractRootWallet = genesisrefs.GenesisRef(GenesisNameRootWallet)
	// ContractRootAccount is the root account contract reference.
	ContractRootAccount = genesisrefs.GenesisRef(GenesisNameRootAccount)
	// ContractMigrationAdminMember is the migration admin member contract reference.
	ContractMigrationAdminMember = genesisrefs.GenesisRef(GenesisNameMigrationAdminMember)
	// ContractMigrationAdmin is the migration wallet contract reference.
	ContractMigrationAdmin = genesisrefs.GenesisRef(GenesisNameMigrationAdmin)
	// ContractMigrationWallet is the migration wallet contract reference.
	ContractMigrationWallet = genesisrefs.GenesisRef(GenesisNameMigrationAdminWallet)
	// ContractMigrationAccount is the migration account contract reference.
	ContractMigrationAccount = genesisrefs.GenesisRef(GenesisNameMigrationAdminAccount)
	// ContractMigrationDeposit is the migration deposit contract reference.
	ContractMigrationDeposit = genesisrefs.GenesisRef(GenesisNameMigrationAdminDeposit)
	// ContractDeposit is the deposit contract reference.
	ContractDeposit = genesisrefs.GenesisRef(GenesisNameDeposit)
	// ContractCostCenter is the cost center contract reference.
	ContractCostCenter = genesisrefs.GenesisRef(GenesisNameCostCenter)
	// ContractFeeMember is the fee member contract reference.
	ContractFeeMember = genesisrefs.GenesisRef(GenesisNameFeeMember)
	// ContractFeeWallet is the fee wallet contract reference.
	ContractFeeWallet = genesisrefs.GenesisRef(GenesisNameFeeWallet)
	// ContractFeeAccount is the fee account contract reference.
	ContractFeeAccount = genesisrefs.GenesisRef(GenesisNameFeeAccount)

	// ContractMigrationDaemonMembers is the migration daemon members contracts references.
	ContractMigrationDaemonMembers = func() (result [GenesisAmountMigrationDaemonMembers]insolar.Reference) {
		for i, name := range GenesisNameMigrationDaemonMembers {
			result[i] = genesisrefs.GenesisRef(name)
		}
		return
	}()

	// ContractMigrationMap where key is migration daemon member  references and value related migration daemon contract
	ContractMigrationMap = func() (result map[insolar.Reference]insolar.Reference) {
		result = make(map[insolar.Reference]insolar.Reference)
		for i := 0; i < GenesisAmountMigrationDaemonMembers; i++ {
			result[genesisrefs.GenesisRef(GenesisNameMigrationDaemonMembers[i])] = genesisrefs.GenesisRef(GenesisNameMigrationDaemons[i])
		}
		return
	}()

	// ContractNetworkIncentivesMembers is the network incentives members contracts references.
	ContractNetworkIncentivesMembers = func() (result [GenesisAmountNetworkIncentivesMembers]insolar.Reference) {
		for i, name := range GenesisNameNetworkIncentivesMembers {
			result[i] = genesisrefs.GenesisRef(name)
		}
		return
	}()

	// ContractApplicationIncentivesMembers is the application incentives members contracts references.
	ContractApplicationIncentivesMembers = func() (result [GenesisAmountApplicationIncentivesMembers]insolar.Reference) {
		for i, name := range GenesisNameApplicationIncentivesMembers {
			result[i] = genesisrefs.GenesisRef(name)
		}
		return
	}()

	// ContractFoundationMembers is the foundation members contracts references.
	ContractFoundationMembers = func() (result [GenesisAmountFoundationMembers]insolar.Reference) {
		for i, name := range GenesisNameFoundationMembers {
			result[i] = genesisrefs.GenesisRef(name)
		}
		return
	}()

	// ContractEnterpriseMembers is the enterprise members contracts references.
	ContractEnterpriseMembers = func() (result [GenesisAmountEnterpriseMembers]insolar.Reference) {
		for i, name := range GenesisNameEnterpriseMembers {
			result[i] = genesisrefs.GenesisRef(name)
		}
		return
	}()

	// ContractNetworkIncentivesWallets is the network incentives members contracts references.
	ContractNetworkIncentivesWallets = func() (result [GenesisAmountNetworkIncentivesMembers]insolar.Reference) {
		for i, name := range GenesisNameNetworkIncentivesWallets {
			result[i] = genesisrefs.GenesisRef(name)
		}
		return
	}()

	// ContractApplicationIncentivesWallets is the application incentives members contracts references.
	ContractApplicationIncentivesWallets = func() (result [GenesisAmountApplicationIncentivesMembers]insolar.Reference) {
		for i, name := range GenesisNameApplicationIncentivesWallets {
			result[i] = genesisrefs.GenesisRef(name)
		}
		return
	}()

	// ContractFoundationWallets is the foundation members contracts references.
	ContractFoundationWallets = func() (result [GenesisAmountFoundationMembers]insolar.Reference) {
		for i, name := range GenesisNameFoundationWallets {
			result[i] = genesisrefs.GenesisRef(name)
		}
		return
	}()

	// ContractEnterpriseWallets is the enterprise members contracts references.
	ContractEnterpriseWallets = func() (result [GenesisAmountEnterpriseMembers]insolar.Reference) {
		for i, name := range GenesisNameEnterpriseWallets {
			result[i] = genesisrefs.GenesisRef(name)
		}
		return
	}()

	// ContractNetworkIncentivesDeposits is the network incentives deposits contracts references.
	ContractNetworkIncentivesDeposits = func() (result [GenesisAmountNetworkIncentivesMembers]insolar.Reference) {
		for i, name := range GenesisNameNetworkIncentivesDeposits {
			result[i] = genesisrefs.GenesisRef(name)
		}
		return
	}()

	// ContractApplicationIncentivesDeposits is the application incentives deposits contracts references.
	ContractApplicationIncentivesDeposits = func() (result [GenesisAmountApplicationIncentivesMembers]insolar.Reference) {
		for i, name := range GenesisNameApplicationIncentivesDeposits {
			result[i] = genesisrefs.GenesisRef(name)
		}
		return
	}()

	// ContractFoundationDeposits is the foundation deposits contracts references.
	ContractFoundationDeposits = func() (result [GenesisAmountFoundationMembers]insolar.Reference) {
		for i, name := range GenesisNameFoundationDeposits {
			result[i] = genesisrefs.GenesisRef(name)
		}
		return
	}()

	// ContractNetworkIncentivesAccounts is the network incentives accounts contracts references.
	ContractNetworkIncentivesAccounts = func() (result [GenesisAmountNetworkIncentivesMembers]insolar.Reference) {
		for i, name := range GenesisNameNetworkIncentivesAccounts {
			result[i] = genesisrefs.GenesisRef(name)
		}
		return
	}()

	// ContractApplicationIncentivesAccounts is the application incentives accounts contracts references.
	ContractApplicationIncentivesAccounts = func() (result [GenesisAmountApplicationIncentivesMembers]insolar.Reference) {
		for i, name := range GenesisNameApplicationIncentivesAccounts {
			result[i] = genesisrefs.GenesisRef(name)
		}
		return
	}()

	// ContractFoundationAccounts is the foundation accounts contracts references.
	ContractFoundationAccounts = func() (result [GenesisAmountFoundationMembers]insolar.Reference) {
		for i, name := range GenesisNameFoundationAccounts {
			result[i] = genesisrefs.GenesisRef(name)
		}
		return
	}()

	// ContractEnterpriseAccounts is the enterprise accounts contracts references.
	ContractEnterpriseAccounts = func() (result [GenesisAmountFoundationMembers]insolar.Reference) {
		for i, name := range GenesisNameEnterpriseAccounts {
			result[i] = genesisrefs.GenesisRef(name)
		}
		return
	}()
)

// Get reference RootMember contract.
func GetRootMember() insolar.Reference {
	return ContractRootMember
}

// Get reference on RootDomain contract.
func GetRootDomain() insolar.Reference {
	return ContractRootDomain
}

// ContractPublicKeyNameShards is the public key shards contracts names.
func ContractPublicKeyNameShards(pkShardCount int) []string {
	result := make([]string, pkShardCount)
	for i := 0; i < pkShardCount; i++ {
		name := GenesisNamePKShard + strconv.Itoa(i)
		result[i] = name
	}
	return result
}

// ContractPublicKeyShards is the public key shards contracts references.
func ContractPublicKeyShards(pkShardCount int) []insolar.Reference {
	result := make([]insolar.Reference, pkShardCount)
	for i, name := range ContractPublicKeyNameShards(pkShardCount) {
		result[i] = genesisrefs.GenesisRef(name)
	}
	return result
}

// ContractMigrationAddressNameShards is the migration address shards contracts names.
func ContractMigrationAddressNameShards(maShardCount int) []string {
	result := make([]string, maShardCount)
	for i := 0; i < maShardCount; i++ {
		name := GenesisNameMigrationShard + strconv.Itoa(i)
		result[i] = name
	}
	return result
}

// ContractMigrationAddressShards is the migration address shards contracts references.
func ContractMigrationAddressShards(maShardCount int) []insolar.Reference {
	result := make([]insolar.Reference, maShardCount)
	for i, name := range ContractMigrationAddressNameShards(maShardCount) {
		result[i] = genesisrefs.GenesisRef(name)
	}
	return result
}

func ContractPublicKeyShardRefs(pkShardCount int) {
	for _, name := range ContractPublicKeyNameShards(pkShardCount) {
		genesisrefs.PredefinedPrototypes[name+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(preprocessor.PrototypeType, GenesisNamePKShard, 0)
	}
}

func ContractMigrationAddressShardRefs(maShardCount int) {
	for _, name := range ContractMigrationAddressNameShards(maShardCount) {
		genesisrefs.PredefinedPrototypes[name+genesisrefs.PrototypeSuffix] = *genesisrefs.GenerateProtoReferenceFromContractID(preprocessor.PrototypeType, GenesisNameMigrationShard, 0)
	}
}
