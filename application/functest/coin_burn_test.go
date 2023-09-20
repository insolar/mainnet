//go:build functest
// +build functest

package functest

import (
	"encoding/json"
	"math/big"
	"strings"
	"testing"

	"github.com/insolar/insolar/applicationbase/testutils/launchnet"
	"github.com/insolar/insolar/applicationbase/testutils/testrequest"
	"github.com/stretchr/testify/require"

	"github.com/insolar/mainnet/application/builtin/contract/member"
)

func TestCoinBurn(t *testing.T) {
	t.Run("happy_path", func(t *testing.T) {
		donor := createMember(t) // initial zero balance

		const amount = "12345"
		grantCoinsFromRoot(t, amount, donor) // now we have some nonzero balance

		// Let's check how many coins have already been burned.
		wasBurned := getBurnedBalance(t)

		// burning amount
		err := coinBurn(amount, donor)
		require.NoError(t, err)

		// We expect the donor's balance to be zero.
		donorBalance := getBalanceNoErr(t, donor, donor.Ref)
		zeroBalance := big.NewInt(0)
		require.Equal(t, zeroBalance, donorBalance)

		// BurnedBalance should be increased by the amount.
		becameBurned := getBurnedBalance(t)
		expectedBurned, _ := new(big.Int).SetString(amount, 10)
		actualBurned := new(big.Int).Sub(becameBurned, wasBurned)
		require.Equal(t, expectedBurned, actualBurned)
	})

	t.Run("not_enough_balance", func(t *testing.T) {
		donor := createMember(t) // initial zero balance
		wasBalance := getBalanceNoErr(t, donor, donor.Ref)

		const amount = "12345"

		// Let's check how many coins have already been burned.
		wasBurned := getBurnedBalance(t)

		// burning amount
		err := coinBurn(amount, donor)
		require.Error(t, err)
		trace := checkConvertRequesterError(t, err).Data.Trace
		require.Contains(t, strings.Join(trace, ":"), "not enough balance to burn")

		// We expect that the balance has not changed.
		becameBalance := getBalanceNoErr(t, donor, donor.Ref)
		require.Equal(t, wasBalance, becameBalance)

		// BurnedBalance should stay the same.
		becameBurned := getBurnedBalance(t)
		require.Equal(t, wasBurned, becameBurned)
	})

	t.Run("invalid_amount", func(t *testing.T) {
		donor := createMember(t) // initial zero balance
		wasBalance := getBalanceNoErr(t, donor, donor.Ref)

		const amount = "invalid_amount"

		// Let's check how many coins have already been burned.
		wasBurned := getBurnedBalance(t)

		// burning amount
		err := coinBurn(amount, donor)
		require.Error(t, err)
		trace := checkConvertRequesterError(t, err).Data.Trace
		require.Contains(t, strings.Join(trace, ":"),
			"request don't pass OpenAPI schema validation")

		// We expect that the balance has not changed.
		becameBalance := getBalanceNoErr(t, donor, donor.Ref)
		require.Equal(t, wasBalance, becameBalance)

		// BurnedBalance should stay the same.
		becameBurned := getBurnedBalance(t)
		require.Equal(t, wasBurned, becameBurned)
	})
}

func coinBurn(amount string, member *AppUser) error {
	_, _, err := testrequest.MakeSignedRequest(
		launchnet.TestRPCUrl,
		member,
		"coin.burn",
		map[string]interface{}{
			"amount": amount,
		})
	return err
}

func grantCoinsFromRoot(t *testing.T, amount string, toMember *AppUser) {
	_, _, err := testrequest.MakeSignedRequest(
		launchnet.TestRPCUrlPublic,
		&Root,
		"member.transfer",
		map[string]interface{}{
			"amount":            amount,
			"toMemberReference": toMember.Ref,
		})
	require.NoError(t, err)
}

func getBurnedBalance(t *testing.T) *big.Int {
	res, _, err := testrequest.MakeSignedRequest(
		launchnet.TestRPCUrl,
		&MigrationAdmin,
		"member.getBalance",
		map[string]interface{}{
			"reference": MigrationAdmin.Ref,
		})
	require.NoError(t, err)

	jsonBuf, err := json.Marshal(res)
	require.NoError(t, err)
	getBalanceResp := member.GetBalanceResponse{}
	err = json.Unmarshal(jsonBuf, &getBalanceResp)
	require.NoError(t, err)

	burnedBalance, ok := new(big.Int).SetString(getBalanceResp.BurnedBalance, 10)
	if !ok {
		return big.NewInt(0)
	}

	return burnedBalance
}
