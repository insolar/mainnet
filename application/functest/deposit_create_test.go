// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

// +build functest

package functest

import (
	"strings"
	"testing"

	"github.com/insolar/insolar/api/requester"
	"github.com/insolar/insolar/applicationbase/testutils/launchnet"
	"github.com/insolar/insolar/applicationbase/testutils/testrequest"
	"github.com/insolar/insolar/insolar/gen"
	"github.com/insolar/insolar/testutils"
	"github.com/stretchr/testify/require"
)

func TestDepositCreate(t *testing.T) {
	t.Run("happy_path", func(t *testing.T) {
		ethHash := testutils.RandomEthHash()
		targetMember := fullMigration(t, ethHash)
		targetDeposit := getDepositReference(t, targetMember, ethHash)

		// make call
		err := registerDepositCreateAdminCall(t, map[string]interface{}{
			"depositReference": targetDeposit,
			"memberReference":  targetMember.Ref,
		})
		require.NoError(t, err)

		// check second deposit
		firstDepositBalance := getDepositBalanceNoErr(t, targetMember, targetMember.Ref, ethHash)
		secondHash := ethHash + "_2"
		secondDepositBalance := getDepositBalanceNoErr(t, targetMember, targetMember.Ref, secondHash)
		require.Equal(t, firstDepositBalance, secondDepositBalance)
	})

	t.Run("without_permissions", func(t *testing.T) {
		ethHash := testutils.RandomEthHash()
		targetMember := fullMigration(t, ethHash)
		targetDeposit := getDepositReference(t, targetMember, ethHash)

		// make call
		err := registerDepositCreateUserCall(t, map[string]interface{}{
			"depositReference": targetDeposit,
			"memberReference":  targetMember.Ref,
		}, targetMember)
		require.Error(t, err)
		requesterError, ok := err.(*requester.Error)
		require.True(t, ok)
		trace := strings.Join(requesterError.Data.Trace, ": ")
		require.Contains(t, trace, "only migration admin can call this method")
	})

	t.Run("invalid_target_member", func(t *testing.T) {
		firstMember := createMember(t)
		ethHash := testutils.RandomEthHash()
		targetMember := fullMigration(t, ethHash)
		targetDeposit := getDepositReference(t, targetMember, ethHash)

		// make call
		err := registerDepositCreateAdminCall(t, map[string]interface{}{
			"depositReference": targetDeposit,
			"memberReference":  firstMember.Ref,
		})
		require.Error(t, err)
		requesterError, ok := err.(*requester.Error)
		require.True(t, ok)
		trace := strings.Join(requesterError.Data.Trace, ": ")
		require.Contains(t, trace, "actual deposit ref is nil or deposit doesn't exist")
	})

	t.Run("invalid_target_deposit", func(t *testing.T) {
		ethHash := testutils.RandomEthHash()
		targetMember := fullMigration(t, ethHash)
		targetDeposit := gen.Reference().String()

		// make call
		err := registerDepositCreateAdminCall(t, map[string]interface{}{
			"depositReference": targetDeposit,
			"memberReference":  targetMember.Ref,
		})
		require.Error(t, err)
		requesterError, ok := err.(*requester.Error)
		require.True(t, ok)
		trace := strings.Join(requesterError.Data.Trace, ": ")
		require.Contains(t, trace, "failed to get deposit itself")
	})
}

func registerDepositCreateAdminCall(t *testing.T, params map[string]interface{}) error {
	return registerDepositCreateUserCall(t, params, &MigrationAdmin)
}

func registerDepositCreateUserCall(t *testing.T, params map[string]interface{}, user *AppUser) error {
	method := "deposit.create"
	_, _, err := testrequest.MakeSignedRequest(launchnet.TestRPCUrl, user, method, params)

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

func getDepositReference(t *testing.T, user *AppUser, ethHash string) string {
	_, deposits := getBalanceAndDepositsNoErr(t, user, user.Ref)
	depFace, ok := deposits[ethHash]
	require.True(t, ok)
	dep, ok := depFace.(map[string]interface{})
	require.True(t, ok)
	depositReferenceFace, ok := dep["reference"]
	require.True(t, ok)
	depositRef, ok := depositReferenceFace.(string)
	require.True(t, ok)
	return depositRef
}
