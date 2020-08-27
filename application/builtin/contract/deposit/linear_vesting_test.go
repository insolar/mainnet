// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

package deposit

import (
	"math"
	"math/big"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinearVestedByNow(t *testing.T) {
	t.Run("zero_amount", func(t *testing.T) {
		zero := big.NewInt(0)
		passedSteps := rand.Uint64()
		totalSteps := rand.Uint64()

		vested := LinearVestedByNow(zero, passedSteps, totalSteps)
		require.Equal(t, zero, vested)
	})

	t.Run("division_by_zero", func(t *testing.T) {
		amount := big.NewInt(int64(rand.Uint32()))
		passedSteps := rand.Uint64()
		totalSteps := uint64(0)

		vested := LinearVestedByNow(amount, passedSteps, totalSteps)
		require.Equal(t, amount, vested)
	})

	t.Run("zero_step", func(t *testing.T) {
		amount := big.NewInt(int64(rand.Uint32()))
		passedSteps := uint64(0)
		totalSteps := rand.Uint64()

		vested := LinearVestedByNow(amount, passedSteps, totalSteps)
		zero := big.NewInt(0)
		require.Equal(t, zero, vested)
	})

	t.Run("one_step", func(t *testing.T) {
		amount := big.NewInt(int64(rand.Uint32()))
		passedSteps := uint64(1)
		totalSteps := rand.Uint64()

		vested := LinearVestedByNow(amount, passedSteps, totalSteps)

		quo := new(big.Int).Div(
			amount,
			new(big.Int).SetUint64(totalSteps),
		)
		require.Equal(t, quo, vested)
	})

	t.Run("last_step", func(t *testing.T) {
		amount := big.NewInt(int64(rand.Uint32()))
		totalSteps := rand.Uint64()
		passedSteps := totalSteps

		vested := LinearVestedByNow(amount, passedSteps, totalSteps)
		require.Equal(t, amount, vested)
	})

	t.Run("passed", func(t *testing.T) {
		amount := big.NewInt(int64(rand.Uint32()))
		totalSteps := uint64(rand.Int63n(math.MaxInt64 - 1))
		passedSteps := totalSteps + 1

		vested := LinearVestedByNow(amount, passedSteps, totalSteps)
		require.Equal(t, amount, vested)
	})

	t.Run("big_amount", func(t *testing.T) {
		amount, ok := new(big.Int).SetString("10000000000000000000", 10)
		require.True(t, ok)
		totalSteps := uint64(math.MaxUint64)
		passedSteps := totalSteps - 1

		vested := LinearVestedByNow(amount, passedSteps, totalSteps)

		quo := new(big.Int).Div(
			new(big.Int).Mul(
				amount,
				new(big.Int).SetUint64(passedSteps),
			),
			new(big.Int).SetUint64(totalSteps),
		)

		require.Equal(t, quo, vested)
	})

	t.Run("by_table", func(t *testing.T) {
		table := []struct {
			amount   int64
			passed   uint64
			total    uint64
			expected int64
		}{
			{10, 0, 6, 0},
			{10, 1, 6, 1},
			{10, 2, 6, 3},
			{10, 3, 6, 5},
			{10, 4, 6, 6},
			{10, 5, 6, 8},
			{10, 6, 6, 10},
			{10, 7, 6, 10},
		}

		for _, row := range table {
			amount := big.NewInt(row.amount)
			vested := LinearVestedByNow(amount, row.passed, row.total)

			expected := big.NewInt(row.expected)
			require.Equal(t, expected, vested)
		}
	})
}
