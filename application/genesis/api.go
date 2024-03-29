package genesis

import (
	"github.com/insolar/insolar/api"
	"github.com/pkg/errors"
)

// initAPIInfoResponse creates application-specific data,
// that will be included in response from /admin-api/rpc#network.getInfo
func initAPIInfoResponse() (map[string]interface{}, error) {
	rootDomain := ContractRootDomain
	if rootDomain.IsEmpty() {
		return nil, errors.New("rootDomain ref is nil")
	}

	rootMember := ContractRootMember
	if rootMember.IsEmpty() {
		return nil, errors.New("rootMember ref is nil")
	}

	migrationDaemonMembers := ContractMigrationDaemonMembers
	migrationDaemonMembersStrs := make([]string, 0)
	for _, r := range migrationDaemonMembers {
		if r.IsEmpty() {
			return nil, errors.New("migration daemon members refs are nil")
		}
		migrationDaemonMembersStrs = append(migrationDaemonMembersStrs, r.String())
	}

	migrationAdminMember := ContractMigrationAdminMember
	if migrationAdminMember.IsEmpty() {
		return nil, errors.New("migration admin member ref is nil")
	}
	feeMember := ContractFeeMember
	if feeMember.IsEmpty() {
		return nil, errors.New("feeMember ref is nil")
	}
	return map[string]interface{}{
		"rootDomain":             rootDomain.String(),
		"rootMember":             rootMember.String(),
		"migrationAdminMember":   migrationAdminMember.String(),
		"feeMember":              feeMember.String(),
		"migrationDaemonMembers": migrationDaemonMembersStrs,
	}, nil
}

// initAPIOptions creates options object, that contains application-specific settings for api component.
func InitAPIOptions() (api.Options, error) {
	apiInfoResponse, err := initAPIInfoResponse()
	if err != nil {
		return api.Options{}, err
	}
	adminContractMethods := map[string]bool{
		"migration.deactivateDaemon": true,
		"migration.activateDaemon":   true,
		"migration.checkDaemon":      true,
		"migration.addAddresses":     true,
		"migration.getAddressCount":  true,
		"deposit.migration":          true,
		"deposit.createFund":         true,
		"member.getBalance":          true,
		"account.transferToDeposit":  true,
		"deposit.transferToDeposit":  true,
		"coin.burn":                  true,
	}
	contractMethods := map[string]bool{
		"member.create":          true,
		"member.get":             true,
		"member.transfer":        true,
		"member.migrationCreate": true,
		"deposit.transfer":       true,
	}
	proxyToRootMethods := []string{"member.create", "member.migrationCreate", "member.get"}

	return api.Options{
		AdminContractMethods: adminContractMethods,
		ContractMethods:      contractMethods,
		InfoResponse:         apiInfoResponse,
		RootReference:        ContractRootMember,
		ProxyToRootMethods:   proxyToRootMethods,
	}, nil
}
