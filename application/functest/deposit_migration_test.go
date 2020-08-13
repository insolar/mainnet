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
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/insolar/insolar/applicationbase/testutils/launchnet"
	"github.com/insolar/insolar/applicationbase/testutils/testrequest"
	"github.com/insolar/insolar/testutils"

	depositContract "github.com/insolar/mainnet/application/builtin/contract/deposit"
	"github.com/insolar/mainnet/application/genesisrefs"
)

func TestDepositMigration(t *testing.T) {
	t.Run("happy_path", func(t *testing.T) {
		activeDaemons := activateDaemons(t, countTwoActiveDaemon)
		member := createMigrationMemberForMA(t)

		const insAmount = "1000"
		const xnsAmount = "10000"
		ethHash := testutils.RandomEthHash()
		additionalDepositHash := ethHash + "_2"
		initialMainFundBalance := getMainFundBalance(t)
		initialFund2Balance := getFund2Balance(t)

		// create new deposit and make first call deposit.Confirm
		deposit := migrate(t, member.Ref, insAmount, ethHash, member.MigrationAddress, 0)

		// check balances
		mainDepositBalance := getDepositBalanceNoErr(t, member, member.Ref, ethHash)
		require.Equal(t, big.NewInt(0), mainDepositBalance)
		require.Equal(t, initialMainFundBalance, getMainFundBalance(t))
		require.Equal(t, initialFund2Balance, getFund2Balance(t))

		// make rest deposit.Confirm calls
		for i := 1; i < len(activeDaemons); i++ {
			deposit = migrate(t, member.Ref, insAmount, ethHash, member.MigrationAddress, i)
		}

		// check confirmation amounts
		confirmations := deposit["confirmerReferences"].(map[string]interface{})
		for _, amount := range confirmations {
			require.Equal(t, xnsAmount, amount)
		}

		// check balances
		numericXNSAmount, ok := new(big.Int).SetString(xnsAmount, 10)
		require.True(t, ok)
		mainDepositBalance = getDepositBalanceNoErr(t, member, member.Ref, ethHash)
		additionalDepositBalance := getDepositBalanceNoErr(t, member, member.Ref, additionalDepositHash)
		require.Equal(t, numericXNSAmount, mainDepositBalance)
		require.Equal(t, numericXNSAmount, additionalDepositBalance)
		diff := new(big.Int).Sub(initialMainFundBalance, getMainFundBalance(t))
		require.Equal(t, numericXNSAmount, diff)
		diff = new(big.Int).Sub(initialFund2Balance, getFund2Balance(t))
		require.Equal(t, numericXNSAmount, diff)
	})
}

func getMainFundBalance(t *testing.T) *big.Int {
	return getDepositBalanceNoErr(t, &MigrationAdmin, MigrationAdmin.Ref, genesisrefs.FundsDepositName)
}

func getFund2Balance(t *testing.T) *big.Int {
	return getDepositBalanceNoErr(t, &MigrationAdmin, MigrationAdmin.Ref, depositContract.PublicAllocation2DepositName)
}

func TestMigrationTokenOneActiveDaemon(t *testing.T) {
	// one daemon confirmation can't change balance
	activateDaemons(t, countOneActiveDaemon)
	daemonIndex := 0
	member := createMigrationMemberForMA(t)

	ethHash := testutils.RandomEthHash()

	deposit := migrate(t, member.Ref, "1000", ethHash, member.MigrationAddress, daemonIndex)
	balance := deposit["balance"].(string)
	require.Equal(t, "0", balance)

	confirmations := deposit["confirmerReferences"].(map[string]interface{})
	require.Equal(t, "10000", confirmations[MigrationDaemons[daemonIndex].Ref])
}

func TestMigrationTokenThreeActiveDaemons(t *testing.T) {
	activeDaemons := activateDaemons(t, countThreeActiveDaemon)
	member := createMigrationMemberForMA(t)

	ethHash := testutils.RandomEthHash()

	for i := 0; i < len(activeDaemons)-1; i++ {
		_ = migrate(t, member.Ref, "1000", ethHash, member.MigrationAddress, i)
	}

	_, err := testrequest.SignedRequest(t,
		launchnet.TestRPCUrl,
		MigrationDaemons[countThreeActiveDaemon-1],
		"deposit.migration",
		map[string]interface{}{"amount": "1000", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
	require.NoError(t, err)
}

func TestMigrationTokenOnDifferentDeposits(t *testing.T) {
	activateDaemons(t, countTwoActiveDaemon)
	member := createMigrationMemberForMA(t)

	ethHash := testutils.RandomEthHash()

	_ = migrate(t, member.Ref, "1000", ethHash, member.MigrationAddress, 0)
	deposit := migrate(t, member.Ref, "1000", ethHash, member.MigrationAddress, 1)

	confirmations := deposit["confirmerReferences"].(map[string]interface{})
	require.Equal(t, "10000", confirmations[MigrationDaemons[0].Ref])
	require.Equal(t, "10000", confirmations[MigrationDaemons[1].Ref])
}

func TestMigrationTokenNotInTheList(t *testing.T) {
	migrationAddress := testutils.RandomEthMigrationAddress()
	_, err := testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrl,
		&MigrationAdmin,
		"deposit.migration",
		map[string]interface{}{"amount": "1000", "ethTxHash": testutils.RandomEthHash(), "migrationAddress": migrationAddress})
	data := checkConvertRequesterError(t, err).Data
	require.Contains(t, data.Trace, "the member is not migration daemon")
}

func TestMigrationTokenZeroAmount(t *testing.T) {
	member := createMigrationMemberForMA(t)

	result, err := testrequest.SignedRequestWithEmptyRequestRef(t,
		launchnet.TestRPCUrl,
		MigrationDaemons[0],
		"deposit.migration",
		map[string]interface{}{"amount": "0", "ethTxHash": testutils.RandomEthHash(), "migrationAddress": member.MigrationAddress})

	data := checkConvertRequesterError(t, err).Data
	testrequest.ExpectedError(t, data.Trace, `Error at "/params/callParams/amount":JSON string doesn't match the regular expression '^[1-9][0-9]*$`)
	require.Nil(t, result)
}

func TestMigrationTokenMistakeField(t *testing.T) {
	member := createMigrationMemberForMA(t)

	result, err := testrequest.SignedRequestWithEmptyRequestRef(t,
		launchnet.TestRPCUrl,
		MigrationDaemons[0],
		"deposit.migration",
		map[string]interface{}{"amount1": "0", "ethTxHash": testutils.RandomEthHash(), "migrationAddress": member.MigrationAddress})
	data := checkConvertRequesterError(t, err).Data
	testrequest.ExpectedError(t, data.Trace, "Property 'amount' is missing")
	require.Nil(t, result)
}

func TestMigrationTokenNilValue(t *testing.T) {
	member := createMigrationMemberForMA(t)

	result, err := testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrl, MigrationDaemons[0],
		"deposit.migration", map[string]interface{}{"amount": "20", "ethTxHash": nil, "migrationAddress": member.MigrationAddress})
	data := checkConvertRequesterError(t, err).Data
	testrequest.ExpectedError(t, data.Trace, `Error at "/params/callParams/ethTxHash":Value is not nullable`)
	require.Nil(t, result)

}

func TestMigrationTokenMaxAmount(t *testing.T) {
	activateDaemons(t, countTwoActiveDaemon)
	member := createMigrationMemberForMA(t)

	result, err := testrequest.SignedRequest(t,
		launchnet.TestRPCUrl,
		MigrationDaemons[0],
		"deposit.migration",
		map[string]interface{}{"amount": "500000000000000000", "ethTxHash": testutils.RandomEthHash(), "migrationAddress": member.MigrationAddress})
	require.NoError(t, err)
	require.Equal(t, result.(map[string]interface{})["memberReference"].(string), member.Ref)
}

func TestMigrationDoubleMigrationFromSameDaemon(t *testing.T) {
	activateDaemons(t, countTwoActiveDaemon)
	member := createMigrationMemberForMA(t)

	ethHash := testutils.RandomEthHash()

	resultMigr1, err := testrequest.SignedRequest(t,
		launchnet.TestRPCUrl,
		MigrationDaemons[0], "deposit.migration",
		map[string]interface{}{"amount": "20", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
	require.NoError(t, err)
	require.Equal(t, member.Ref, resultMigr1.(map[string]interface{})["memberReference"].(string))

	resultMigr2, err := testrequest.SignedRequest(t,
		launchnet.TestRPCUrl,
		MigrationDaemons[0], "deposit.migration",
		map[string]interface{}{"amount": "20", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
	require.NoError(t, err)
	require.Equal(t, member.Ref, resultMigr2.(map[string]interface{})["memberReference"].(string))
}

func TestMigrationDoubleMigrationFromSameDaemon_WithDifferentAmount(t *testing.T) {
	activateDaemons(t, countTwoActiveDaemon)
	member := createMigrationMemberForMA(t)

	ethHash := testutils.RandomEthHash()

	resultMigr1, err := testrequest.SignedRequest(t,
		launchnet.TestRPCUrl,
		MigrationDaemons[0], "deposit.migration",
		map[string]interface{}{"amount": "20", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
	require.NoError(t, err)
	require.Equal(t, member.Ref, resultMigr1.(map[string]interface{})["memberReference"].(string))

	_, err = testrequest.SignedRequestWithEmptyRequestRef(t,
		launchnet.TestRPCUrl,
		MigrationDaemons[0],
		"deposit.migration",
		map[string]interface{}{"amount": "30", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
	data := checkConvertRequesterError(t, err).Data
	require.Contains(
		t,
		data.Trace,
		fmt.Sprintf("confirm from this migration daemon %s already exists with different amount", MigrationDaemons[0].Ref),
	)
}

func TestMigrationAnotherAmountSameTx(t *testing.T) {
	activateDaemons(t, countThreeActiveDaemon)

	member := createMigrationMemberForMA(t)

	ethHash := testutils.RandomEthHash()

	_, err := testrequest.SignedRequest(t,
		launchnet.TestRPCUrl,
		MigrationDaemons[0], "deposit.migration",
		map[string]interface{}{"amount": "20", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
	require.NoError(t, err)

	_, _, err = testrequest.MakeSignedRequest(
		launchnet.TestRPCUrl,
		MigrationDaemons[2],
		"deposit.migration",
		map[string]interface{}{"amount": "30", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
	data := checkConvertRequesterError(t, err).Data
	trace := strings.Join(data.Trace, ": ")
	require.Contains(t, trace, "some of confirmation amounts aren't equal others")
	require.Contains(t, trace, fmt.Sprintf("%s:200", MigrationDaemons[0].Ref))
	require.Contains(t, trace, fmt.Sprintf("%s:300", MigrationDaemons[2].Ref))
}

func TestMigration_WrongSecondAmount(t *testing.T) {
	activateDaemons(t, countThreeActiveDaemon)

	member := createMigrationMemberForMA(t)

	ethHash := testutils.RandomEthHash()

	_, err := testrequest.SignedRequest(t,
		launchnet.TestRPCUrl,
		MigrationDaemons[0], "deposit.migration",
		map[string]interface{}{"amount": "100", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
	require.NoError(t, err)

	_, _, err = testrequest.MakeSignedRequest(
		launchnet.TestRPCUrl,
		MigrationDaemons[1],
		"deposit.migration",
		map[string]interface{}{"amount": "200", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
	data := checkConvertRequesterError(t, err).Data
	trace := strings.Join(data.Trace, ": ")
	require.Contains(t, trace, "some of confirmation amounts aren't equal others")
	require.Contains(t, trace, fmt.Sprintf("%s:1000", MigrationDaemons[0].Ref))
	require.Contains(t, trace, fmt.Sprintf("%s:2000", MigrationDaemons[1].Ref))

	_, _, err = testrequest.MakeSignedRequest(
		launchnet.TestRPCUrl,
		MigrationDaemons[2],
		"deposit.migration",
		map[string]interface{}{"amount": "100", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
	data = checkConvertRequesterError(t, err).Data
	trace = strings.Join(data.Trace, ": ")
	require.Contains(t, trace, "some of confirmation amounts aren't equal others")
	require.Contains(t, trace, fmt.Sprintf("%s:2000", MigrationDaemons[1].Ref))

	_, deposits := getBalanceAndDepositsNoErr(t, member, member.Ref)
	deposit, ok := deposits[ethHash].(map[string]interface{})
	require.True(t, ok)
	require.Equal(t, ethHash, deposit["ethTxHash"])
	require.Equal(t, "0", deposit["amount"])
	memberBalance := deposit["balance"].(string)
	require.Equal(t, "0", memberBalance)
	confirmations := deposit["confirmerReferences"].(map[string]interface{})
	require.Equal(t, "1000", confirmations[MigrationDaemons[0].Ref])
	require.Equal(t, "2000", confirmations[MigrationDaemons[1].Ref])
	require.Equal(t, "1000", confirmations[MigrationDaemons[2].Ref])
}

func TestMigration_WrongFirstAmount(t *testing.T) {
	activateDaemons(t, countThreeActiveDaemon)

	member := createMigrationMemberForMA(t)

	ethHash := testutils.RandomEthHash()

	_, err := testrequest.SignedRequest(t,
		launchnet.TestRPCUrl,
		MigrationDaemons[0], "deposit.migration",
		map[string]interface{}{"amount": "200", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
	require.NoError(t, err)

	_, _, err = testrequest.MakeSignedRequest(
		launchnet.TestRPCUrl,
		MigrationDaemons[1],
		"deposit.migration",
		map[string]interface{}{"amount": "100", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
	data := checkConvertRequesterError(t, err).Data
	trace := strings.Join(data.Trace, ": ")
	require.Contains(t, trace, "some of confirmation amounts aren't equal others")
	require.Contains(t, trace, fmt.Sprintf("%s:2000", MigrationDaemons[0].Ref))
	require.Contains(t, trace, fmt.Sprintf("%s:1000", MigrationDaemons[1].Ref))

	_, _, err = testrequest.MakeSignedRequest(
		launchnet.TestRPCUrl,
		MigrationDaemons[2],
		"deposit.migration",
		map[string]interface{}{"amount": "100", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
	data = checkConvertRequesterError(t, err).Data
	trace = strings.Join(data.Trace, ": ")
	require.Contains(t, trace, "some of confirmation amounts aren't equal others")
	require.Contains(t, trace, fmt.Sprintf("%s:2000", MigrationDaemons[0].Ref))

	_, deposits := getBalanceAndDepositsNoErr(t, member, member.Ref)
	deposit, ok := deposits[ethHash].(map[string]interface{})
	require.True(t, ok)
	require.Equal(t, ethHash, deposit["ethTxHash"])
	require.Equal(t, "0", deposit["amount"])
	memberBalance := deposit["balance"].(string)
	require.Equal(t, "0", memberBalance)
	confirmations := deposit["confirmerReferences"].(map[string]interface{})
	require.Equal(t, "2000", confirmations[MigrationDaemons[0].Ref])
	require.Equal(t, "1000", confirmations[MigrationDaemons[1].Ref])
	require.Equal(t, "1000", confirmations[MigrationDaemons[2].Ref])
}

func TestMigration_WrongAllAmount(t *testing.T) {
	activateDaemons(t, countThreeActiveDaemon)

	member := createMigrationMemberForMA(t)

	ethHash := testutils.RandomEthHash()

	_, err := testrequest.SignedRequest(t,
		launchnet.TestRPCUrl,
		MigrationDaemons[0], "deposit.migration",
		map[string]interface{}{"amount": "100", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
	require.NoError(t, err)

	_, _, err = testrequest.MakeSignedRequest(
		launchnet.TestRPCUrl,
		MigrationDaemons[1],
		"deposit.migration",
		map[string]interface{}{"amount": "200", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
	data := checkConvertRequesterError(t, err).Data
	trace := strings.Join(data.Trace, ": ")
	require.Contains(t, trace, "some of confirmation amounts aren't equal others")
	require.Contains(t, trace, fmt.Sprintf("%s:1000", MigrationDaemons[0].Ref))
	require.Contains(t, trace, fmt.Sprintf("%s:2000", MigrationDaemons[1].Ref))

	_, _, err = testrequest.MakeSignedRequest(
		launchnet.TestRPCUrl,
		MigrationDaemons[2],
		"deposit.migration",
		map[string]interface{}{"amount": "300", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
	data = checkConvertRequesterError(t, err).Data
	trace = strings.Join(data.Trace, ": ")
	require.Contains(t, trace, "some of confirmation amounts aren't equal others")
	require.Contains(t, trace, fmt.Sprintf("%s:1000", MigrationDaemons[0].Ref))
	require.Contains(t, trace, fmt.Sprintf("%s:2000", MigrationDaemons[1].Ref))
	require.Contains(t, trace, fmt.Sprintf("%s:3000", MigrationDaemons[2].Ref))

	_, deposits := getBalanceAndDepositsNoErr(t, member, member.Ref)
	deposit, ok := deposits[ethHash].(map[string]interface{})
	require.True(t, ok)
	require.Equal(t, ethHash, deposit["ethTxHash"])
	require.Equal(t, "0", deposit["amount"])
	memberBalance := deposit["balance"].(string)
	require.Equal(t, "0", memberBalance)
	confirmations := deposit["confirmerReferences"].(map[string]interface{})
	require.Equal(t, "1000", confirmations[MigrationDaemons[0].Ref])
	require.Equal(t, "2000", confirmations[MigrationDaemons[1].Ref])
	require.Equal(t, "3000", confirmations[MigrationDaemons[2].Ref])
}

func TestMigrationTokenDoubleSpend(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	_ = activateDaemons(t, countThreeActiveDaemon)
	member := createMigrationMemberForMA(t)
	anotherMember := createMember(t)

	ethHash := testutils.RandomEthHash()

	deposit := migrate(t, member.Ref, "1000", ethHash, member.MigrationAddress, 0)
	firstMemberBalance := deposit["balance"].(string)

	require.Equal(t, "0", firstMemberBalance)
	firstMABalance, err := getAdminDepositBalance(t, &MigrationAdmin, MigrationAdmin.Ref)
	require.NoError(t, err)

	for i := 1; i < countThreeActiveDaemon; i++ {
		go func(i int) {
			res, _, err := testrequest.MakeSignedRequest(
				launchnet.TestRPCUrl,
				MigrationDaemons[i],
				"deposit.migration",
				map[string]interface{}{"amount": "1000", "ethTxHash": ethHash, "migrationAddress": member.MigrationAddress})
			if err != nil {
				requestErrorData := checkConvertRequesterError(t, err).Data
				t.Log(requestErrorData)
			} else {
				t.Log(res)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	_, deposits := getBalanceAndDepositsNoErr(t, anotherMember, member.Ref)
	deposit, ok := deposits[ethHash].(map[string]interface{})
	require.True(t, ok)

	require.Equal(t, ethHash, deposit["ethTxHash"])
	require.Equal(t, "10000", deposit["amount"])
	secondMemberBalance := deposit["balance"].(string)
	require.Equal(t, "10000", secondMemberBalance)
	secondMABalance, err := getAdminDepositBalance(t, &MigrationAdmin, MigrationAdmin.Ref)
	require.NoError(t, err)
	dif := new(big.Int).Sub(firstMABalance, secondMABalance)
	require.Equal(t, "10000", dif.String())
}
