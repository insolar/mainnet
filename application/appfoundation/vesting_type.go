// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

package appfoundation

// VestingType type of vesting process
type VestingType int

//go:generate ../../bin/stringer -type=VestingType
const (
	// Non-linear process for regular users
	DefaultVesting VestingType = iota
	// Deprecated: Never used
	Vesting1
	// Vesting type for funds (usually with zero vesting period and zero-step)
	Vesting2
	// Deprecated: Never used
	Vesting3
	// Deprecated: Never used
	Vesting4
	// Linear process for regular users
	LinearVesting
)
