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
)
