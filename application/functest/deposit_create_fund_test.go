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
	"github.com/insolar/insolar/testutils"
	"github.com/stretchr/testify/require"
)

func TestDepositCreateFund(t *testing.T) {
	ethHash := testutils.RandomEthHash()
	member := fullMigration(t, ethHash)
	lockupEndDate := time.Now().Unix()
	ref, err := registerCreateFundCall(t, member, map[string]interface{}{
		"lockupEndDate": strconv.FormatInt(lockupEndDate, 10),
	})
	require.NoError(t, err)
	require.NotNil(t, ref)

	_, deposits := getBalanceAndDepositsNoErr(t, member, member.Ref)
	require.Contains(t, deposits, "genesis_deposit2")
}

func registerCreateFundCall(t *testing.T, member *AppUser, params map[string]interface{}) (string, error) {
	res, err := testrequest.SignedRequest(t, launchnet.TestRPCUrl, member, "deposit.createFund", params)
	if err != nil {
		return "", err
	}
	return res.(string), nil
}
