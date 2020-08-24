// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

package deposit

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/insolar/insolar/pulse"
	"github.com/pkg/errors"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/logicrunner/builtin/foundation"
	"github.com/insolar/insolar/logicrunner/builtin/foundation/safemath"

	"github.com/insolar/mainnet/application/appfoundation"
	"github.com/insolar/mainnet/application/builtin/proxy/deposit"
	"github.com/insolar/mainnet/application/builtin/proxy/member"
	"github.com/insolar/mainnet/application/builtin/proxy/migrationdaemon"
	"github.com/insolar/mainnet/application/builtin/proxy/wallet"
	"github.com/insolar/mainnet/application/genesisrefs"
)

const (
	numConfirmation              = 2
	PublicAllocation2DepositName = "genesis_deposit2"
)

// Deposit is like an account. But it holds migrated money.
type Deposit struct {
	foundation.BaseContract
	Balance                 string                    `json:"balance"`
	PulseDepositUnHold      insolar.PulseNumber       `json:"holdReleaseDate"`
	MigrationDaemonConfirms foundation.StableMap      `json:"confirmerReferences"`
	Amount                  string                    `json:"amount"`
	TxHash                  string                    `json:"ethTxHash"`
	VestingType             appfoundation.VestingType `json:"vestingType"`
	Lockup                  int64                     `json:"lockupInPulses"`
	Vesting                 int64                     `json:"vestingInPulses"`
	VestingStep             int64                     `json:"vestingStepInPulses"`
	IsConfirmed             bool                      `json:"isConfirmed"`
}

// New creates new deposit.
func New(txHash string, lockup int64, vesting int64, vestingStep int64) (*Deposit, error) {

	if vestingStep > 0 && vesting%vestingStep != 0 {
		return nil, errors.New("vesting is not multiple of vestingStep")
	}

	migrationDaemonConfirms := make(foundation.StableMap)

	return &Deposit{
		Balance:                 "0",
		MigrationDaemonConfirms: migrationDaemonConfirms,
		Amount:                  "0",
		TxHash:                  txHash,
		Lockup:                  lockup,
		Vesting:                 vesting,
		VestingStep:             vestingStep,
		VestingType:             appfoundation.DefaultVesting,
	}, nil
}

// NewFund creates new public allocation 2 deposit
func NewFund(lockupEndDate int64) (*Deposit, error) {
	unholdPulse := pulse.OfUnixTime(lockupEndDate)
	return &Deposit{
		Balance:            "0",
		Amount:             "0",
		PulseDepositUnHold: unholdPulse,
		VestingType:        appfoundation.Vesting2,
		TxHash:             PublicAllocation2DepositName,
		Lockup:             int64(unholdPulse - pulse.MinTimePulse),
		Vesting:            0,
		VestingStep:        0,
		IsConfirmed:        true,
	}, nil
}

// Form of Deposit that is applied in API
type DepositOut struct {
	Balance                 string                    `json:"balance"`
	HoldStartDate           int64                     `json:"holdStartDate"`
	PulseDepositUnHold      int64                     `json:"holdReleaseDate"`
	MigrationDaemonConfirms []DaemonConfirm           `json:"confirmerReferences"`
	Amount                  string                    `json:"amount"`
	TxHash                  string                    `json:"ethTxHash"`
	VestingType             appfoundation.VestingType `json:"vestingType"`
	Lockup                  int64                     `json:"lockup"`
	Vesting                 int64                     `json:"vesting"`
	VestingStep             int64                     `json:"vestingStep"`
}

type DaemonConfirm struct {
	Reference string `json:"reference"`
	Amount    string `json:"amount"`
}

// GetTxHash gets transaction hash.
// ins:immutable
func (d *Deposit) GetTxHash() (string, error) {
	return d.TxHash, nil
}

// GetAmount gets amount.
// ins:immutable
func (d *Deposit) GetAmount() (string, error) {
	return d.Amount, nil
}

// GetBalance gets balance.
// ins:immutable
func (d *Deposit) GetBalance() (string, error) {
	return d.Balance, nil
}

// Return pulse of unhold deposit.
// ins:immutable
func (d *Deposit) GetPulseUnHold() (insolar.PulseNumber, error) {
	return d.PulseDepositUnHold, nil
}

// Itself gets deposit information.
// ins:immutable
func (d *Deposit) Itself() (interface{}, error) {
	var daemonConfirms = make([]DaemonConfirm, 0, len(d.MigrationDaemonConfirms))
	var pulseDepositUnHold int64
	for k, v := range d.MigrationDaemonConfirms {
		daemonConfirms = append(daemonConfirms, DaemonConfirm{Reference: k, Amount: v})
	}
	t, err := d.PulseDepositUnHold.AsApproximateTime()
	if err == nil {
		pulseDepositUnHold = t.Unix()
	}
	holdStartDate := pulseDepositUnHold - d.Lockup
	if holdStartDate < 0 {
		holdStartDate = 0
	}
	return &DepositOut{
		Balance:                 d.Balance,
		HoldStartDate:           holdStartDate,
		PulseDepositUnHold:      pulseDepositUnHold,
		MigrationDaemonConfirms: daemonConfirms,
		Amount:                  d.Amount,
		TxHash:                  d.TxHash,
		VestingType:             d.VestingType,
		Lockup:                  d.Lockup,
		Vesting:                 d.Vesting,
		VestingStep:             d.VestingStep,
	}, nil
}

// Confirm adds confirm for deposit by migration daemon.
func (d *Deposit) Confirm(
	txHash string, amountStr string, fromMember insolar.Reference, request insolar.Reference, toMember insolar.Reference,
) error {

	migrationDaemonRef := fromMember.String()
	if txHash != d.TxHash {
		return errors.New("transaction hash is incorrect")
	}
	if confirmedAmount, ok := d.MigrationDaemonConfirms[migrationDaemonRef]; ok {
		if amountStr != confirmedAmount {
			return fmt.Errorf(
				"confirm from this migration daemon %s already exists with different amount: was %s, now %s",
				migrationDaemonRef,
				confirmedAmount,
				amountStr,
			)
		}
		return nil
	}

	if d.IsConfirmed {
		d.MigrationDaemonConfirms[migrationDaemonRef] = amountStr
		if amountStr != d.Amount {
			return fmt.Errorf(
				"migration is done for this deposit %s, but with different amount: confirmed is %s, from this daemon %s",
				txHash,
				d.Amount,
				amountStr,
			)
		}
		return nil
	}

	if len(d.MigrationDaemonConfirms) > 0 {
		canConfirm, errFromConfirm := d.checkConfirm(migrationDaemonRef, amountStr)
		if canConfirm {
			currentPulse, err := foundation.GetPulseNumber()
			if err != nil {
				return errors.Wrap(err, "failed to get current pulse")
			}
			d.Amount = amountStr
			d.PulseDepositUnHold = currentPulse + insolar.PulseNumber(d.Lockup)

			ma := member.GetObject(appfoundation.GetMigrationAdminMember())
			walletRef, err := ma.GetWallet()
			if err != nil {
				return errors.Wrap(err, "failed to get wallet")
			}
			ok, maDeposit, _ := wallet.GetObject(*walletRef).FindDeposit(genesisrefs.FundsDepositName)
			if !ok {
				return fmt.Errorf("failed to find source deposit - %s", walletRef.String())
			}

			err = deposit.GetObject(*maDeposit).TransferToDeposit(
				amountStr, d.GetReference(), appfoundation.GetMigrationAdminMember(), request, toMember, string(appfoundation.TTypeMigration))
			if err != nil {
				return errors.Wrap(err, "failed to transfer from migration deposit to deposit")
			}
			d.IsConfirmed = true
		}
		if errFromConfirm != nil {
			return errFromConfirm
		}
		return nil
	}
	d.MigrationDaemonConfirms[migrationDaemonRef] = amountStr
	return nil
}

// TransferToDeposit transfers funds to deposit.
func (d *Deposit) TransferToDeposit(
	amountStr string,
	toDeposit insolar.Reference,
	fromMember insolar.Reference,
	request insolar.Reference,
	toMember insolar.Reference,
	txType string,
) error {
	amount, ok := new(big.Int).SetString(amountStr, 10)
	if !ok {
		return errors.New("can't parse input amount")
	}
	balance, ok := new(big.Int).SetString(d.Balance, 10)
	if !ok {
		return errors.New("can't parse deposit balance")
	}
	if balance.Sign() <= 0 {
		return errors.New("not enough balance for transfer")
	}
	newBalance, err := safemath.Sub(balance, amount)
	if err != nil {
		return errors.Wrap(err, "not enough balance for transfer")
	}
	d.Balance = newBalance.String()
	destination := deposit.GetObject(toDeposit)
	acceptDepositErr := destination.Accept(appfoundation.SagaAcceptInfo{
		Amount:     amountStr,
		FromMember: fromMember,
		Request:    request,
	})
	if acceptDepositErr == nil {
		return nil
	}
	d.Balance = balance.String()
	return errors.Wrap(err, "failed to transfer amount")

}

// Check amount field in confirmation from migration daemons.
func (d *Deposit) checkAmount(activeDaemons map[string]string, migrationDaemonRef string, amountStr string) (bool, error) {
	confirmed := false
	if activeDaemons == nil || len(activeDaemons) == 0 {
		return false, errors.New("list with migration daemons member is empty")
	}
	amountConfirms := make(map[string]int) // amount: num of confirms
	var errDaemon []string
	for migrationRef, amount := range activeDaemons {
		if amount != amountStr {
			errDaemon = append(errDaemon, fmt.Sprintf("%s send amount %s", migrationRef, amount))
		}
		amountConfirms[amount] = amountConfirms[amount] + 1
	}
	amountConfirms[amountStr] = amountConfirms[amountStr] + 1
	if amountConfirms[amountStr] >= numConfirmation {
		confirmed = true
	}
	var err error
	if len(errDaemon) > 0 {
		if !confirmed {
			errDaemon = append(errDaemon, fmt.Sprintf("%s send amount %s", migrationDaemonRef, amountStr))
		}
		err = fmt.Errorf("several migration daemons send different amount: %s", strings.Join(errDaemon, ": "))
	}
	return confirmed, err
}

func (d *Deposit) checkConfirm(migrationDaemonRef string, amountStr string) (bool, error) {
	activateDaemons := make(map[string]string)

	for ref, a := range d.MigrationDaemonConfirms {
		migrationDaemonMemberRef, err := insolar.NewObjectReferenceFromString(ref)
		if err != nil {
			return false, errors.New("failed to parse params.Reference")
		}

		migrationDaemonContractRef, err := appfoundation.GetMigrationDaemon(*migrationDaemonMemberRef)
		if err != nil || migrationDaemonContractRef.IsEmpty() {
			return false, errors.Wrap(err, "get migration daemon contract from foundation failed")
		}

		migrationDaemonContract := migrationdaemon.GetObject(migrationDaemonContractRef)
		result, err := migrationDaemonContract.GetActivationStatus()

		if err != nil {
			return false, err
		}
		if result {
			activateDaemons[ref] = a
		}
	}
	d.MigrationDaemonConfirms[migrationDaemonRef] = amountStr
	if len(activateDaemons) > 0 {
		canConfirm, err := d.checkAmount(activateDaemons, migrationDaemonRef, amountStr)
		if err != nil {
			return canConfirm, errors.Wrap(err, "failed to check amount in confirmation from migration daemon")
		}
		return canConfirm, nil
	}
	return false, nil
}

func (d *Deposit) availableAmount() (*big.Int, error) {
	if d.VestingType == appfoundation.DefaultVesting && !d.IsConfirmed {
		return nil, errors.New("number of confirms is less then 2")
	}

	currentPulse, err := foundation.GetPulseNumber()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get pulse number")
	}
	if currentPulse < d.PulseDepositUnHold {
		return nil, errors.New("hold period didn't end")
	}

	amount, ok := new(big.Int).SetString(d.Amount, 10)
	if !ok {
		return nil, errors.New("can't parse deposit amount")
	}
	balance, ok := new(big.Int).SetString(d.Balance, 10)
	if !ok {
		return nil, errors.New("can't parse deposit balance")
	}

	// Allow to transfer whole balance if vesting period has already finished
	if currentPulse >= d.PulseDepositUnHold+insolar.PulseNumber(d.Vesting) {
		return balance, nil
	}

	// Total number of vesting steps in vesting period
	totalSteps := uint64(d.Vesting / d.VestingStep)
	// Vesting steps already passed by now
	passedSteps := uint64(int64(currentPulse-d.PulseDepositUnHold) / d.VestingStep)
	// Amount that has been vested by now
	vestedByNow := VestedByNow(amount, passedSteps, totalSteps)
	// Amount that is still locked on deposit
	onHold := new(big.Int).Sub(amount, vestedByNow)
	// Amount that is now available for withdrawal
	availableNow := new(big.Int).Sub(balance, onHold)

	// availableNow can become negative when balance is 0 and vesting has already started
	if availableNow.Cmp(big.NewInt(0)) == -1 {
		return big.NewInt(0), nil
	}

	return availableNow, nil
}

func (d *Deposit) canTransfer(transferAmount *big.Int) error {
	availableAmount, err := d.availableAmount()
	if err != nil {
		return err
	}
	if transferAmount.Cmp(availableAmount) == 1 {
		return errors.New("not enough unholded balance for transfer")
	}
	return nil
}

// Transfer transfers money from deposit to wallet. It can be called only after deposit hold period.
func (d *Deposit) Transfer(
	amountStr string, memberRef insolar.Reference, request insolar.Reference,
) (interface{}, error) {

	amount, ok := new(big.Int).SetString(amountStr, 10)
	if !ok {
		return nil, errors.New("can't parse input amount")
	}

	balance, ok := new(big.Int).SetString(d.Balance, 10)
	if !ok {
		return nil, errors.New("can't parse deposit balance")
	}
	if balance.Sign() <= 0 {
		return nil, errors.New("not enough balance for transfer")
	}
	newBalance, err := safemath.Sub(balance, amount)
	if err != nil {
		return nil, errors.Wrap(err, "not enough balance for transfer")
	}
	err = d.canTransfer(amount)
	if err != nil {
		return nil, errors.Wrap(err, "can't start transfer")
	}
	d.Balance = newBalance.String()

	m := member.GetObject(memberRef)
	acceptMemberErr := m.Accept(appfoundation.SagaAcceptInfo{
		Amount:     amountStr,
		FromMember: memberRef,
		Request:    request,
	})
	if acceptMemberErr == nil {
		return nil, nil
	}
	d.Balance = balance.String()
	return nil, errors.Wrap(acceptMemberErr, "failed to transfer amount")
}

// Accept accepts transfer to balance.
// ins:saga(INS_FLAG_NO_ROLLBACK_METHOD)
func (d *Deposit) Accept(arg appfoundation.SagaAcceptInfo) error {

	amount := new(big.Int)
	amount, ok := amount.SetString(arg.Amount, 10)
	if !ok {
		return errors.New("can't parse input amount")
	}

	balance := new(big.Int)
	balance, ok = balance.SetString(d.Balance, 10)
	if !ok {
		return errors.New("can't parse deposit balance")
	}

	b, err := safemath.Add(balance, amount)
	if err != nil {
		return errors.Wrap(err, "failed to add amount to balance")
	}
	d.Balance = b.String()

	return nil
}
