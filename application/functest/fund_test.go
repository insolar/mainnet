// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

// +build functest

package functest

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/insolar/insolar/applicationbase/testutils/launchnet"
	"github.com/insolar/insolar/applicationbase/testutils/testrequest"
	"github.com/insolar/insolar/insolar"

	"github.com/insolar/mainnet/application"
	"github.com/insolar/mainnet/application/genesisrefs"
	"github.com/insolar/mainnet/cmd/insolard/genesisstate"
)

const (
	ZeroBalance = "0"
)

func TestFoundationMemberCreate(t *testing.T) {
	for _, m := range Foundation {
		err := verifyFundsMembersAndDeposits(t, m, application.FoundationDistributionAmount, ZeroBalance)
		if err != nil {
			require.NoError(t, err)
		}
	}
}

func TestEnterpriseMemberCreate(t *testing.T) {
	for _, m := range Enterprise {
		err := verifyFundsMembersExist(t, m, ZeroBalance)
		if err != nil {
			require.NoError(t, err)
		}
	}
}

func TestNetworkIncentivesMemberCreate(t *testing.T) {
	// for speed up test check only last member
	m := NetworkIncentives[application.GenesisAmountNetworkIncentivesMembers-1]

	err := verifyFundsMembersAndDeposits(t, m, application.NetworkIncentivesDistributionAmount, ZeroBalance)
	if err != nil {
		require.NoError(t, err)
	}
}

func TestApplicationIncentivesMemberCreate(t *testing.T) {
	for _, m := range ApplicationIncentives {
		err := verifyFundsMembersAndDeposits(t, m, application.AppIncentivesDistributionAmount, ZeroBalance)
		if err != nil {
			require.NoError(t, err)
		}
	}
}

func checkBalanceAndDepositFewTimes(t *testing.T, m *AppUser, expectedBalance string, expectedDeposit string) {
	var balance *big.Int
	var depositStr string
	for i := 0; i < times; i++ {
		balance, deposits := getBalanceAndDepositsNoErr(t, m, m.Ref)
		depositStr = deposits[genesisrefs.FundsDepositName].(map[string]interface{})["balance"].(string)
		if balance.String() == expectedBalance && depositStr == expectedDeposit {
			return
		}
		time.Sleep(time.Second)
	}
	t.Errorf("Received balance or deposite is not equal expected: current balance %s, expected %s;"+
		" current deposite %s, expected %s",
		balance, expectedBalance,
		depositStr, expectedDeposit)
}

func TestNetworkIncentivesTransferDeposit(t *testing.T) {
	// for speed up test check only last member
	lastIdx := application.GenesisAmountNetworkIncentivesMembers - 1
	m := NetworkIncentives[lastIdx]

	res2, err := testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, m, "member.get", nil)
	require.NoError(t, err)
	decodedRes2, ok := res2.(map[string]interface{})
	m.Ref = decodedRes2["reference"].(string)
	require.True(t, ok, fmt.Sprintf("failed to decode: expected map[string]interface{}, got %T", res2))

	_, err = testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrlPublic, m,
		"deposit.transfer", map[string]interface{}{"amount": "100", "ethTxHash": genesisrefs.FundsDepositName},
	)
	data := checkConvertRequesterError(t, err).Data
	require.Contains(t, data.Trace, "not enough balance for transfer")
}

func TestApplicationIncentivesTransferDeposit(t *testing.T) {
	for _, m := range ApplicationIncentives {
		res2, err := testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, m, "member.get", nil)
		require.NoError(t, err)
		decodedRes2, ok := res2.(map[string]interface{})
		m.Ref = decodedRes2["reference"].(string)
		require.True(t, ok, fmt.Sprintf("failed to decode: expected map[string]interface{}, got %T", res2))

		_, err = testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrlPublic, m,
			"deposit.transfer", map[string]interface{}{"amount": "100", "ethTxHash": genesisrefs.FundsDepositName},
		)
		data := checkConvertRequesterError(t, err).Data
		require.Contains(t, data.Trace, "not enough balance for transfer")
	}
}

func TestFoundationTransferDeposit(t *testing.T) {
	for _, m := range Foundation {
		res2, err := testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, m, "member.get", nil)
		require.NoError(t, err)
		decodedRes2, ok := res2.(map[string]interface{})
		m.Ref = decodedRes2["reference"].(string)
		require.True(t, ok, fmt.Sprintf("failed to decode: expected map[string]interface{}, got %T", res2))

		_, err = testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrlPublic, m,
			"deposit.transfer", map[string]interface{}{"amount": "100", "ethTxHash": genesisrefs.FundsDepositName},
		)
		data := checkConvertRequesterError(t, err).Data
		require.Contains(t, data.Trace, "not enough balance for transfer")
	}
}

func TestMigrationDaemonTransferDeposit(t *testing.T) {
	m := &MigrationAdmin

	oldBalance, deposits := getBalanceAndDepositsNoErr(t, m, m.Ref)
	oldDepositStr := deposits[genesisrefs.FundsDepositName].(map[string]interface{})["balance"].(string)

	_, reqRefStr, err := testrequest.MakeSignedRequest(launchnet.TestRPCUrlPublic, m, "member.get", nil)
	require.NoError(t, err)
	reqRef, err := insolar.NewReferenceFromString(reqRefStr)
	require.NoError(t, err)
	currentTime, err := reqRef.GetLocal().Pulse().AsApproximateTime()
	require.NoError(t, err)

	if currentTime.Unix() < genesisstate.MigrationDaemonUnholdDate {
		_, err = testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrlPublic, m,
			"deposit.transfer", map[string]interface{}{"amount": "100", "ethTxHash": genesisrefs.FundsDepositName},
		)
		data := checkConvertRequesterError(t, err).Data
		require.Contains(t, data.Trace, "hold period didn't end")

		newBalance, newDeposits := getBalanceAndDepositsNoErr(t, m, m.Ref)
		newDepositStr := newDeposits[genesisrefs.FundsDepositName].(map[string]interface{})["balance"].(string)
		require.Equal(t, oldBalance.String(), newBalance.String())
		require.Equal(t, oldDepositStr, newDepositStr)
	} else {
		_, err = testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, m,
			"deposit.transfer", map[string]interface{}{"amount": "100", "ethTxHash": genesisrefs.FundsDepositName},
		)
		require.NoError(t, err)
		newBalance, newDeposits := getBalanceAndDepositsNoErr(t, m, m.Ref)
		newDepositStr := newDeposits[genesisrefs.FundsDepositName].(map[string]interface{})["balance"].(string)
		amount := int64(100)
		require.Equal(t, oldBalance.Add(oldBalance, big.NewInt(amount)).String(), newBalance.String())
		oldDeposit, ok := new(big.Int).SetString(oldDepositStr, 10)
		require.True(t, ok, "can't parse oldDepositStr")
		require.Equal(t, oldDeposit.Sub(oldDeposit, big.NewInt(amount)).String(), newDepositStr)
	}
}
