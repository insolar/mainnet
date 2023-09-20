//go:build functest
// +build functest

package functest

import (
	"math/big"
	"testing"

	"github.com/insolar/insolar/applicationbase/testutils/launchnet"
	"github.com/insolar/insolar/applicationbase/testutils/testrequest"
	"github.com/insolar/insolar/testutils"
	"github.com/stretchr/testify/require"
)

func TestAccountTransferToDeposit(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		ethHash := testutils.RandomEthHash()

		firstMember := createMember(t)            // initial zero balance
		secondMember := fullMigration(t, ethHash) // deposit with 3600000

		// init money on member
		_, err := testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, &Root, "member.transfer",
			map[string]interface{}{"amount": "2000000000", "toMemberReference": firstMember.Ref})
		require.NoError(t, err)

		_, err = testrequest.SignedRequest(t, launchnet.TestRPCUrl, firstMember,
			"account.transferToDeposit", map[string]interface{}{"amount": "1000", "toDepositName": ethHash, "toMemberReference": secondMember.GetReference()})

		require.NoError(t, err)

		firstBalance := getBalanceNoErr(t, firstMember, firstMember.Ref)
		expectedFirst, _ := new(big.Int).SetString("1999999000", 10)
		require.Equal(t, expectedFirst, firstBalance)

		secondBalance, err := getDepositBalance(t, secondMember, secondMember.Ref, ethHash)
		require.NoError(t, err)
		expectedSecond, _ := new(big.Int).SetString("3601000", 10)
		require.Equal(t, expectedSecond, secondBalance)
	})

	t.Run("NotEnoughBalance", func(t *testing.T) {
		ethHash := testutils.RandomEthHash()

		member := createMember(t)
		member2 := fullMigration(t, ethHash)

		_, err := testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrl, member,
			"account.transferToDeposit", map[string]interface{}{"amount": "1000", "toDepositName": ethHash, "toMemberReference": member2.GetReference()})
		data := checkConvertRequesterError(t, err).Data
		require.Contains(t, data.Trace, "not enough balance for transfer")
	})
}
