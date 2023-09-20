package appfoundation

import (
	"regexp"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/mainnet/application/genesis"
)

const AllowedVersionSmartContract = 2

// Get reference CostCenter contract.
func GetCostCenter() insolar.Reference {
	return genesis.ContractCostCenter
}

// Get reference MigrationAdminMember contract.
func GetMigrationAdminMember() insolar.Reference {
	return genesis.ContractMigrationAdminMember
}

// Get reference on MigrationAdmin contract.
func GetMigrationAdmin() insolar.Reference {
	return genesis.ContractMigrationAdmin
}

// Get reference on  migrationdaemon contract by  migration member.
func GetMigrationDaemon(migrationMember insolar.Reference) (insolar.Reference, error) {
	return genesis.ContractMigrationMap[migrationMember], nil
}

// Check member is migration daemon member or not
func IsMigrationDaemonMember(member insolar.Reference) bool {
	for _, mDaemonMember := range genesis.ContractMigrationDaemonMembers {
		if mDaemonMember.Equal(member) {
			return true
		}
	}
	return false
}

var etheriumAddressRegex = regexp.MustCompile(`^(0x)?[\dA-Fa-f]{40}$`)

// IsEthereumAddress Ethereum address format verifier
func IsEthereumAddress(s string) bool {
	return etheriumAddressRegex.MatchString(s)
}
