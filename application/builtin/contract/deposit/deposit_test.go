package deposit

import (
	"math/big"
	"testing"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/gen"
	"github.com/insolar/insolar/logicrunner/builtin/foundation"
	"github.com/insolar/insolar/pulse"
	"github.com/stretchr/testify/require"

	"github.com/insolar/mainnet/application/appfoundation"
)

func TestDeposit_availableAmountZeroStep(t *testing.T) {
	amount := "123"
	expectedAmount, _ := new(big.Int).SetString(amount, 10)
	unholdPulse := pulse.OfNow()
	vesting := int64(0)
	vestingStep := int64(0)
	d := &Deposit{
		Balance:            amount,
		Amount:             amount,
		PulseDepositUnHold: unholdPulse,
		VestingType:        appfoundation.Vesting2,
		TxHash:             "tx_hash",
		Lockup:             int64(unholdPulse - pulse.MinTimePulse),
		Vesting:            vesting,
		VestingStep:        vestingStep,
		IsConfirmed:        true,
	}

	prepareContext(unholdPulse)
	require.NotPanics(t, func() {
		available, err := d.availableAmount()
		require.NoError(t, err)
		require.Equal(t, expectedAmount, available)
	})
}

func prepareContext(pulse pulse.Number) {
	request := gen.ReferenceWithPulse(pulse)
	ctx := &insolar.LogicCallContext{
		Request: &request,
	}
	foundation.SetLogicalContext(ctx)
}
