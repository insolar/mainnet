// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

// +build functest

package functest

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/insolar/insolar/api/requester"
	"github.com/insolar/insolar/applicationbase/testutils/launchnet"
	"github.com/insolar/insolar/applicationbase/testutils/testrequest"
	"github.com/insolar/insolar/logicrunner/builtin/foundation"
	"github.com/insolar/insolar/testutils"
)

func TestDepositTransferToken(t *testing.T) {
	ethHash := testutils.RandomEthHash()

	member := fullMigration(t, ethHash)

	firstBalance := getBalanceNoErr(t, member, member.Ref)
	secondBalance := new(big.Int).Add(firstBalance, big.NewInt(1000))

	anon := func() *foundation.Error {
		_, _, err := testrequest.MakeSignedRequest(launchnet.TestRPCUrlPublic, member,
			"deposit.transfer", map[string]interface{}{"amount": "1000", "ethTxHash": ethHash})

		data := checkConvertRequesterError(t, err).Data
		for _, v := range data.Trace {
			if !strings.Contains(v, "hold period didn't end") {
				return nil
			}
		}
		return &foundation.Error{S: err.Error()}
	}

	err := waitUntilRequestProcessed(anon, time.Second*30, time.Second, 30)
	require.NoError(t, err)
	anon = func() *foundation.Error {
		_, _, err := testrequest.MakeSignedRequest(launchnet.TestRPCUrlPublic, member,
			"deposit.transfer", map[string]interface{}{"amount": "1000", "ethTxHash": ethHash})
		if err == nil {
			return nil
		}
		return &foundation.Error{S: err.Error()}
	}
	err = waitUntilRequestProcessed(anon, time.Second*30, time.Second, 30)
	require.NoError(t, err)
	checkBalanceFewTimes(t, member, member.Ref, secondBalance)
}

func TestDepositTransferBeforeUnhold(t *testing.T) {
	ethHash := testutils.RandomEthHash()

	member := fullMigration(t, ethHash)

	_, err := testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrlPublic, member,
		"deposit.transfer", map[string]interface{}{"amount": "100", "ethTxHash": ethHash})
	data := checkConvertRequesterError(t, err).Data
	require.Contains(t, data.Trace, "hold period didn't end", "check lockup_pulse_period param at bootstrap.yaml: it maybe too low")
}

func TestDepositTransferBiggerAmount(t *testing.T) {
	ethHash := testutils.RandomEthHash()

	member := fullMigration(t, ethHash)

	_, err := testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrlPublic, member,
		"deposit.transfer", map[string]interface{}{"amount": "10000000000000", "ethTxHash": ethHash})
	data := checkConvertRequesterError(t, err).Data
	require.Contains(t, data.Trace, "not enough balance for transfer")
}

func TestDepositTransferAnotherTx(t *testing.T) {
	ethHash := testutils.RandomEthHash()

	member := fullMigration(t, ethHash)

	_, err := testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrlPublic, member,
		"deposit.transfer", map[string]interface{}{"amount": "100", "ethTxHash": testutils.RandomEthHash()})
	data := checkConvertRequesterError(t, err).Data
	require.Contains(t, data.Trace, "can't find deposit")
}

func TestDepositTransferWrongValueAmount(t *testing.T) {
	ethHash := testutils.RandomEthHash()

	member := fullMigration(t, ethHash)

	_, err := testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrlPublic, member,
		"deposit.transfer", map[string]interface{}{"amount": "foo", "ethTxHash": ethHash})
	require.Error(t, err)
	data := checkConvertRequesterError(t, err).Data
	testrequest.ExpectedError(t, data.Trace, `Error at "/params/callParams/amount":JSON string doesn't match the regular expression '^[1-9][0-9]*$`)
}

func TestDepositTransferNotEnoughConfirms(t *testing.T) {
	ethHash := testutils.RandomEthHash()

	activateDaemons(t, countTwoActiveDaemon)
	member := createMigrationMemberForMA(t)
	_ = migrate(t, member.Ref, "1000", ethHash, member.MigrationAddress, 2)

	_, err := testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrlPublic, member,
		"deposit.transfer", map[string]interface{}{"amount": "100", "ethTxHash": ethHash})
	data := checkConvertRequesterError(t, err).Data
	require.Contains(t, data.Trace, "not enough balance for transfer")
}

func TestDepositTransfer(t *testing.T) {
	t.Run("from_deposit_2", func(t *testing.T) {
		ethHash := testutils.RandomEthHash()
		member := fullMigration(t, ethHash)

		deposit2 := fmt.Sprintf("%s_2", ethHash)

		err := registerDepositTransferCall(t, member, deposit2, "1000")

		// We expect error here
		// cause hold period for additional deposit equals hold period of base deposit
		// plus hardcoded 3 year's period.
		require.Error(t, err)
		requesterError, ok := err.(*requester.Error)
		require.True(t, ok)
		trace := strings.Join(requesterError.Data.Trace, ": ")
		require.Contains(t, trace, "hold period didn't end")

	})
}

func registerDepositTransferCall(t *testing.T, member *AppUser, ethHash, amount string) error {
	method := "deposit.transfer"
	_, _, err := testrequest.MakeSignedRequest(launchnet.TestRPCUrlPublic, member, method,
		map[string]interface{}{
			"amount":    amount,
			"ethTxHash": ethHash,
		},
	)

	if err != nil {
		var suffix string
		requesterError, ok := err.(*requester.Error)
		if ok {
			suffix = " [" + strings.Join(requesterError.Data.Trace, ": ") + "]"
		}
		t.Log("[" + method + "]" + err.Error() + suffix)
		return err
	}

	return nil
}
