package burnedaccount

import (
	"math/big"

	"github.com/insolar/insolar/logicrunner/builtin/foundation"
	"github.com/insolar/insolar/logicrunner/builtin/foundation/safemath"
	"github.com/pkg/errors"
)

type BurnedAccount struct {
	foundation.BaseContract
	Balance string
}

func New(balance string) (*BurnedAccount, error) {
	return &BurnedAccount{Balance: balance}, nil
}

// GetBalance gets total balance.
// ins:immutable
func (a *BurnedAccount) GetBalance() (string, error) {
	return a.Balance, nil
}

// IncreaseBalance increases the current balance by the amount.
func (a *BurnedAccount) IncreaseBalance(amountStr string) error {
	amount, ok := new(big.Int).SetString(amountStr, 10)
	if !ok {
		return errors.New("can't parse input amount")
	}
	if amount.Sign() <= 0 {
		return errors.New("amount should be greater then zero")
	}
	balance, ok := new(big.Int).SetString(a.Balance, 10)
	if !ok {
		return errors.New("can't parse account balance")
	}
	newBalance, err := safemath.Add(balance, amount)
	if err != nil {
		return errors.Wrap(err, "failed to add amount to balance")
	}
	a.Balance = newBalance.String()
	return nil
}
