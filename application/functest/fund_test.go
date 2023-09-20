//go:build functest
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
	"github.com/insolar/mainnet/application/genesis"
)

func TestFoundationMemberCreate(t *testing.T) {
	for _, m := range Foundation {
		err := verifyFundsMembersAndDeposits(t, m, genesis.FoundationDistributionAmount)
		if err != nil {
			require.NoError(t, err)
		}
	}
}

func TestEnterpriseMemberCreate(t *testing.T) {
	for _, m := range Enterprise {
		err := verifyFundsMembersExist(t, m, genesis.EnterpriseDistributionAmount)
		if err != nil {
			require.NoError(t, err)
		}
	}
}

func TestNetworkIncentivesMemberCreate(t *testing.T) {
	// for speed up test check only last member
	m := NetworkIncentives[genesis.GenesisAmountNetworkIncentivesMembers-1]

	err := verifyFundsMembersAndDeposits(t, m, genesis.NetworkIncentivesDistributionAmount)
	if err != nil {
		require.NoError(t, err)
	}
}

func TestApplicationIncentivesMemberCreate(t *testing.T) {
	for _, m := range ApplicationIncentives {
		err := verifyFundsMembersAndDeposits(t, m, genesis.AppIncentivesDistributionAmount)
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
		depositStr = deposits[genesis.FundsDepositName].(map[string]interface{})["balance"].(string)
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
	lastIdx := genesis.GenesisAmountNetworkIncentivesMembers - 1
	m := NetworkIncentives[lastIdx]

	res2, err := testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, m, "member.get", nil)
	require.NoError(t, err)
	decodedRes2, ok := res2.(map[string]interface{})
	m.Ref = decodedRes2["reference"].(string)
	require.True(t, ok, fmt.Sprintf("failed to decode: expected map[string]interface{}, got %T", res2))

	_, err = testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, m,
		"deposit.transfer", map[string]interface{}{"amount": "100", "ethTxHash": genesis.FundsDepositName},
	)
	require.NoError(t, err)
	depositAmount, ok := new(big.Int).SetString(genesis.NetworkIncentivesDistributionAmount, 10)
	require.True(t, ok, "can't parse NetworkIncentivesDistributionAmount")
	checkBalanceAndDepositFewTimes(t, m, "100", depositAmount.Sub(depositAmount, big.NewInt(100)).String())
}

func TestApplicationIncentivesTransferDeposit(t *testing.T) {
	for _, m := range ApplicationIncentives {
		res2, err := testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, m, "member.get", nil)
		require.NoError(t, err)
		decodedRes2, ok := res2.(map[string]interface{})
		m.Ref = decodedRes2["reference"].(string)
		require.True(t, ok, fmt.Sprintf("failed to decode: expected map[string]interface{}, got %T", res2))

		_, err = testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, m,
			"deposit.transfer", map[string]interface{}{"amount": "100", "ethTxHash": genesis.FundsDepositName},
		)
		require.NoError(t, err)
		depositAmount, ok := new(big.Int).SetString(genesis.AppIncentivesDistributionAmount, 10)
		require.True(t, ok, "can't parse AppIncentivesDistributionAmount")
		checkBalanceAndDepositFewTimes(t, m, "100", depositAmount.Sub(depositAmount, big.NewInt(100)).String())
	}
}

func TestFoundationTransferDeposit(t *testing.T) {
	for _, m := range Foundation {
		res2, err := testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, m, "member.get", nil)
		require.NoError(t, err)
		decodedRes2, ok := res2.(map[string]interface{})
		m.Ref = decodedRes2["reference"].(string)
		require.True(t, ok, fmt.Sprintf("failed to decode: expected map[string]interface{}, got %T", res2))

		_, err = testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, m,
			"deposit.transfer", map[string]interface{}{"amount": "100", "ethTxHash": genesis.FundsDepositName},
		)
		require.NoError(t, err)
		depositAmount, ok := new(big.Int).SetString(genesis.FoundationDistributionAmount, 10)
		require.True(t, ok, "can't parse FoundationDistributionAmount")
		checkBalanceAndDepositFewTimes(t, m, "100", depositAmount.Sub(depositAmount, big.NewInt(100)).String())
	}
}

func TestMigrationDaemonTransferDeposit(t *testing.T) {
	m := &MigrationAdmin

	res, err := testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, m, "member.get", nil)
	require.NoError(t, err)
	decodedRes2, ok := res.(map[string]interface{})
	require.True(t, ok, fmt.Sprintf("failed to decode: expected map[string]interface{}, got %T", res))
	m.Ref = decodedRes2["reference"].(string)

	oldBalance, deposits := getBalanceAndDepositsNoErr(t, m, m.Ref)
	oldDepositStr := deposits[genesis.FundsDepositName].(map[string]interface{})["balance"].(string)

	_, err = testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, m,
		"deposit.transfer", map[string]interface{}{"amount": "100", "ethTxHash": genesis.FundsDepositName},
	)
	require.NoError(t, err)
	newBalance, newDeposits := getBalanceAndDepositsNoErr(t, m, m.Ref)
	newDepositStr := newDeposits[genesis.FundsDepositName].(map[string]interface{})["balance"].(string)
	amount := int64(100)
	require.Equal(t, oldBalance.Add(oldBalance, big.NewInt(amount)).String(), newBalance.String())
	oldDeposit, ok := new(big.Int).SetString(oldDepositStr, 10)
	require.True(t, ok, "can't parse oldDepositStr")
	require.Equal(t, oldDeposit.Sub(oldDeposit, big.NewInt(amount)).String(), newDepositStr)
}
