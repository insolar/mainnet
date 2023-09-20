//go:build functest
// +build functest

package functest

import (
	"github.com/insolar/insolar/applicationbase/testutils/launchnet"
	"github.com/insolar/insolar/applicationbase/testutils/testrequest"
	"github.com/insolar/insolar/testutils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDepositTransferToDepositToken(t *testing.T) {
	ethHash := testutils.RandomEthHash()

	member := fullMigration(t, ethHash)
	member2 := fullMigration(t, ethHash)

	_, err := testrequest.SignedRequest(t, launchnet.TestRPCUrl, member,
		"deposit.transferToDeposit", map[string]interface{}{"fromDepositName": ethHash, "toDepositName": ethHash, "toMemberReference": member2.GetReference()})

	require.NoError(t, err)
}

func TestDepositTransferToDepositAnotherTx(t *testing.T) {
	ethHash := testutils.RandomEthHash()

	member := fullMigration(t, ethHash)
	member2 := fullMigration(t, ethHash)

	_, err := testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrl, member,
		"deposit.transferToDeposit", map[string]interface{}{
			"fromDepositName": ethHash, "toDepositName": testutils.RandomEthHash(),
			"toMemberReference": member2.GetReference()})
	data := checkConvertRequesterError(t, err).Data
	require.Contains(t, data.Trace, "failed to find toDeposit object")

	_, err = testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrl, member,
		"deposit.transferToDeposit", map[string]interface{}{
			"fromDepositName": testutils.RandomEthHash(), "toDepositName": ethHash,
			"toMemberReference": member2.GetReference()})
	data = checkConvertRequesterError(t, err).Data
	require.Contains(t, data.Trace, "failed to find fromDeposit object")
}

func TestDepositTransferToDepositNotEnoughConfirms(t *testing.T) {
	ethHash := testutils.RandomEthHash()

	activateDaemons(t, countTwoActiveDaemon)
	member := createMigrationMemberForMA(t)
	member2 := fullMigration(t, ethHash)

	_ = migrate(t, member.Ref, "1000", ethHash, member.MigrationAddress, 2)

	_, err := testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrl, member,
		"deposit.transferToDeposit", map[string]interface{}{"fromDepositName": ethHash,
			"toDepositName": ethHash, "toMemberReference": member2.GetReference()})
	data := checkConvertRequesterError(t, err).Data
	require.Contains(t, data.Trace, "not enough balance for transfer")
}

func TestDepositTransferToDepositYouOwnDeposit(t *testing.T) {
	ethHash := testutils.RandomEthHash()

	member := fullMigration(t, ethHash)

	_, err := testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrl, member,
		"deposit.transferToDeposit", map[string]interface{}{
			"fromDepositName": ethHash, "toDepositName": ethHash,
			"toMemberReference": member.GetReference()})
	data := checkConvertRequesterError(t, err).Data
	require.Contains(t, data.Trace, "it is impossible to make a transfer and accrual to the same deposit")
}
