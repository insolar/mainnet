package genesis

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/insolar/insolar/applicationbase/genesisrefs"
	"github.com/insolar/insolar/insolar"
)

func TestContractPublicKeyShards(t *testing.T) {
	for i, ref := range ContractPublicKeyShards(100) {
		require.Equal(t, genesisrefs.GenesisRef(GenesisNamePKShard+strconv.Itoa(i)), ref)
	}
}

func TestContractMigrationAddressShards(t *testing.T) {
	for i, ref := range ContractMigrationAddressShards(100) {
		require.Equal(t, genesisrefs.GenesisRef(GenesisNameMigrationShard+strconv.Itoa(i)), ref)
	}
}

func TestReferences(t *testing.T) {
	pairs := map[string]struct {
		got    insolar.Reference
		expect string
	}{
		GenesisNameRootDomain: {
			got:    ContractRootDomain,
			expect: "insolar:1AAEAAciWtcmPVgAcaIvICkgnSsJmp4Clp650xOHjYks",
		},
		GenesisNameRootMember: {
			got:    ContractRootMember,
			expect: "insolar:1AAEAAWeNhA_NwKaH6E36IJ-2PLvXnJRxiTTNWq1giOg",
		},
		GenesisNameRootWallet: {
			got:    ContractRootWallet,
			expect: "insolar:1AAEAAVEt_2mipoVG73cbK-v33ne0krJWXkZibayYKJc",
		},
		GenesisNameRootAccount: {
			got:    ContractRootAccount,
			expect: "insolar:1AAEAAYUzPb6A9YCwdhstSMjq8g4dV_059cFrscpHemo",
		},
		GenesisNameDeposit: {
			got:    ContractDeposit,
			expect: "insolar:1AAEAAVnfpSe6gLpJptcggYUNNGIu0-_kxjnef5G-nR0",
		},
		GenesisNameCostCenter: {
			got:    ContractCostCenter,
			expect: "insolar:1AAEAASCuYBHyztkBdO3b5lDXgDsrk12PKOTixEW6kvY",
		},
		GenesisNameFeeAccount: {
			got:    ContractFeeAccount,
			expect: "insolar:1AAEAAaN2AfiHUl4HxtCcMV-KhWirOx2MA69ndZVAIpM",
		},
		GenesisNameFeeWallet: {
			got:    ContractFeeWallet,
			expect: "insolar:1AAEAAVsLfNjPCXS5hsvt1WHuo0RZIYCs1H3oFC2jxIM",
		},
		GenesisNamePKShard: {
			got:    ContractPublicKeyShards(10)[0],
			expect: "insolar:1AAEAAVM1LnFXwPa92gplaRhMroFeWi-gxznLptCtPCc",
		},
		GenesisNameMigrationShard: {
			got:    ContractMigrationAddressShards(10)[0],
			expect: "insolar:1AAEAAVvqEYcaRInGWx75iJWNLUSuWXU5XA3L8Qh0fpU",
		},
		GenesisNameMigrationAdminAccount: {
			got:    ContractMigrationAccount,
			expect: "insolar:1AAEAAYHatrmIZwsoJ3Fy78F9Q1zZ9bEuxRrZasAcYYo",
		},
	}

	for n, p := range pairs {
		t.Run(n, func(t *testing.T) {
			require.Equal(t, p.expect, p.got.String(), "reference is stable")
		})
	}
}

func TestRootDomain(t *testing.T) {
	ref1 := ContractRootDomain
	ref2 := genesisrefs.GenesisRef(GenesisNameRootDomain)
	require.Equal(t, ref1.String(), ref2.String(), "reference is the same")
}
