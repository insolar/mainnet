// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

// +build functest

package functest

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/insolar/insolar/api/requester"
	"github.com/insolar/insolar/applicationbase/testutils/launchnet"
	"github.com/insolar/insolar/applicationbase/testutils/testrequest"
	"github.com/stretchr/testify/require"
)

var IsDepositFundCreated bool

func TestDepositCreateFund(t *testing.T) {
	t.Run("happy_path", func(t *testing.T) {
		if IsDepositFundCreated {
			// run this test only once
			t.Skip("fund was already created in previous run")
		}

		IsDepositFundCreated = true

		lockupEndDate := time.Now().Unix()

		_, deposits := getBalanceAndDepositsNoErr(t, &MigrationAdmin, MigrationAdmin.Ref)
		require.Contains(t, deposits, "genesis_deposit2")

		// check double creation
		err := registerCreateFundCall(t, map[string]interface{}{
			"lockupEndDate": strconv.FormatInt(lockupEndDate, 10),
		})
		require.Error(t, err)
		requesterErr, ok := err.(*requester.Error)
		require.True(t, ok)
		trace := strings.Join(requesterErr.Data.Trace, ": ")
		require.Contains(t, trace, "fund already created")
	})

	t.Run("without_permissions", func(t *testing.T) {
		ordinaryMember := createMember(t)
		lockupEndDate := time.Now().Unix()

		// make call
		err := registerCreateFundMemberCall(t, map[string]interface{}{
			"lockupEndDate": strconv.FormatInt(lockupEndDate, 10),
		}, ordinaryMember)

		// check errors
		require.Error(t, err)
		requesterErr, ok := err.(*requester.Error)
		require.True(t, ok)
		trace := strings.Join(requesterErr.Data.Trace, ": ")
		require.Contains(t, trace, "only migration admin can call this method")
	})
}

func registerCreateFundCall(t *testing.T, params map[string]interface{}) error {
	return registerCreateFundMemberCall(t, params, &MigrationAdmin)
}

func registerCreateFundMemberCall(t *testing.T, params map[string]interface{}, member *AppUser) error {
	_, _, err := testrequest.MakeSignedRequest(launchnet.TestRPCUrl, member, "deposit.createFund", params)
	if err != nil {
		return err
	}
	return nil
}
