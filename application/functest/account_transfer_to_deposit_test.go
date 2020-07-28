// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

// +build functest

package functest

import (
	"testing"

	"github.com/insolar/insolar/applicationbase/testutils/launchnet"
	"github.com/insolar/insolar/applicationbase/testutils/testrequest"
	"github.com/insolar/insolar/testutils"
	"github.com/stretchr/testify/require"
)

func TestAccountTransferToDeposit(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		ethHash := testutils.RandomEthHash()

		firstMember := createMember(t)
		secondMember := fullMigration(t, ethHash)

		// init money on member
		_, err := testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, &Root, "member.transfer",
			map[string]interface{}{"amount": "2000000000", "toMemberReference": firstMember.Ref})
		require.NoError(t, err)

		_, err = testrequest.SignedRequest(t, launchnet.TestRPCUrl, firstMember,
			"account.transferToDeposit", map[string]interface{}{"amount": "1000", "toDepositName": ethHash, "toMemberReference": secondMember.GetReference()})

		require.NoError(t, err)
	})

	t.Run("NotEnoughBalance", func(t *testing.T) {
		ethHash := testutils.RandomEthHash()

		member := createMember(t)
		member2 := fullMigration(t, ethHash)

		_, err := testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrl, member,
			"account.transferToDeposit", map[string]interface{}{"amount": "1000", "toDepositName": ethHash, "toMemberReference": member2.GetReference()})
		data := checkConvertRequesterError(t, err).Data
		require.Contains(t, data.Trace, "not enough balance for transfer")
	})
}
