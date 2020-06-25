// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

package main

import (
	"fmt"
	"math/big"

	"github.com/insolar/mainnet/application/sdk"
)

type walletToWalletSimpleTransferScenario struct {
	insSDK  *sdk.SDK
	members []sdk.Member

	balanceCheckMembers []sdk.Member
}

func (s *walletToWalletSimpleTransferScenario) canBeStarted() error {
	if len(s.members) < concurrent*2 {
		return fmt.Errorf("not enough members for start")
	}
	return nil
}

func (s *walletToWalletSimpleTransferScenario) prepare(repetition int) {
	members, err := getMembers(s.insSDK, concurrent*2, false)
	check("Error while loading members: ", err)

	if useMembersFromFile {
		members = members[:len(members)-2]
	}

	s.members = members

	s.balanceCheckMembers = make([]sdk.Member, len(s.members), len(s.members)+2)
	copy(s.balanceCheckMembers, s.members)
	s.balanceCheckMembers = append(s.balanceCheckMembers, s.insSDK.GetFeeMember())
	s.balanceCheckMembers = append(s.balanceCheckMembers, s.insSDK.GetMigrationAdminMember())
}

func (s *walletToWalletSimpleTransferScenario) start(concurrentIndex int, repetitionIndex int) (string, error) {
	from := s.members[concurrentIndex*2]
	to := s.members[concurrentIndex*2+1]

	return s.insSDK.TransferSimple(big.NewInt(transferAmount).String(), from, to)
}

func (s *walletToWalletSimpleTransferScenario) getBalanceCheckMembers() []sdk.Member {
	return s.balanceCheckMembers
}
