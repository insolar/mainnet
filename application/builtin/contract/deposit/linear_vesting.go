// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

package deposit

import (
	"math/big"
)

func LinearVestedByNow(amount *big.Int, passedSteps, totalSteps uint64) *big.Int {
	// Passed and total steps have a uint64 type
	// so check equality total steps to zero (to prevent division by zero) is unnecessary.
	if passedSteps >= totalSteps {
		return amount
	}

	passed := new(big.Int).SetUint64(passedSteps)
	total := new(big.Int).SetUint64(totalSteps)

	coeff := new(big.Int).Mul(amount, passed)
	return new(big.Int).Div(coeff, total)
}
