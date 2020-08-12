// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

// +build functest

package functest

import (
	"strconv"
	"testing"
	"time"

	"github.com/insolar/insolar/applicationbase/testutils/launchnet"
	"github.com/insolar/insolar/applicationbase/testutils/testrequest"
	"github.com/stretchr/testify/require"
)

func TestDepositCreateFund(t *testing.T) {
	lockupEndDate := time.Now().Unix()
	err := registerCreateFundCall(t, map[string]interface{}{
		"lockupEndDate": strconv.FormatInt(lockupEndDate, 10),
	})
	require.NoError(t, err)

	_, deposits := getBalanceAndDepositsNoErr(t, &MigrationAdmin, MigrationAdmin.Ref)
	require.Contains(t, deposits, "genesis_deposit2")
}

func registerCreateFundCall(t *testing.T, params map[string]interface{}) error {
	_, err := testrequest.SignedRequest(t, launchnet.TestRPCUrl, &MigrationAdmin, "deposit.createFund", params)
	if err != nil {
		return err
	}
	return nil
}
