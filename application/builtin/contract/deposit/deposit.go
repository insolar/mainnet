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
	"github.com/insolar/mainnet/application/builtin/proxy/migrationadmin"
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
func New(txHash string, lockup int64, vesting int64, vestingStep int64, balance string, pulseDepositUnHold pulse.Number, confirms []appfoundation.DaemonConfirm, amount string, vestingType appfoundation.VestingType, isConfirmed bool) (*Deposit, error) {

	if vestingStep > 0 && vesting%vestingStep != 0 {
		return nil, errors.New("vesting is not multiple of vestingStep")
	}

	migrationDaemonConfirms := make(foundation.StableMap)
	for _, confirm := range confirms {
		migrationDaemonConfirms[confirm.Reference] = confirm.Amount
	}

	return &Deposit{
		Balance:                 balance,
		PulseDepositUnHold:      pulseDepositUnHold,
		MigrationDaemonConfirms: migrationDaemonConfirms,
		Amount:                  amount,
		TxHash:                  txHash,
		Lockup:                  lockup,
		Vesting:                 vesting,
		VestingStep:             vestingStep,
		VestingType:             vestingType,
		IsConfirmed:             isConfirmed,
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
	Ref                     string                        `json:"reference"`
	Balance                 string                        `json:"balance"`
	HoldStartDate           int64                         `json:"holdStartDate"`
	PulseDepositUnHold      int64                         `json:"holdReleaseDate"`
	MigrationDaemonConfirms []appfoundation.DaemonConfirm `json:"confirmerReferences"`
	Amount                  string                        `json:"amount"`
	TxHash                  string                        `json:"ethTxHash"`
	VestingType             appfoundation.VestingType     `json:"vestingType"`
	Lockup                  int64                         `json:"lockup"`
	Vesting                 int64                         `json:"vesting"`
	VestingStep             int64                         `json:"vestingStep"`
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

// Return pulse of unhold deposit.
// ins:immutable
func (d *Deposit) GetPulseUnHold() (insolar.PulseNumber, error) {
	return d.PulseDepositUnHold, nil
}

// Itself gets deposit information.
// ins:immutable
func (d *Deposit) Itself() (interface{}, error) {
	var daemonConfirms = make([]appfoundation.DaemonConfirm, 0, len(d.MigrationDaemonConfirms))
	var pulseDepositUnHold int64
	for k, v := range d.MigrationDaemonConfirms {
		daemonConfirms = append(daemonConfirms, appfoundation.DaemonConfirm{Reference: k, Amount: v})
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
		Ref:                     d.GetReference().String(),
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

// Confirm saves confirmation of deposit by migration daemon,
// checks confirmation sufficiency,
// if all confirmations are collected makes transfer from fund to deposit
// and eventually activates deposit.
func (d *Deposit) Confirm(txHash string, proposedAmount string, migrationDaemonRef insolar.Reference, requestRef insolar.Reference, toMember insolar.Reference) (err error) {
	// check args
	if txHash != d.TxHash {
		return errors.New("transaction hash is incorrect")
	}
	if numericValue, ok := assertBigInt(proposedAmount); !ok {
		return errors.New("invalid amount")
	} else if lessOrEqualZero(numericValue) { // amount <= 0
		return errors.New("amount must be greater than zero")
	}
	if migrationDaemonRef.IsEmpty() {
		return errors.New("empty migrationDaemonRef")
	}
	if requestRef.IsEmpty() {
		return errors.New("empty requestRef reference")
	}
	if toMember.IsEmpty() {
		return errors.New("empty toMember reference")
	}

	// check confirmation existence
	migrationDaemon := migrationDaemonRef.String()
	if previousProposal, ok := d.MigrationDaemonConfirms[migrationDaemon]; ok {
		if proposedAmount != previousProposal {
			return fmt.Errorf(
				"confirm from this migration daemon %s already exists with different amount: was %s, now %s",
				migrationDaemonRef,
				previousProposal,
				proposedAmount,
			)
		}
		return nil
	}

	// save confirmation data
	d.MigrationDaemonConfirms[migrationDaemon] = proposedAmount

	if d.IsConfirmed {
		return nil
	}

	// check confirmations sufficiency
	if len(d.MigrationDaemonConfirms) < numConfirmation {
		return nil
	}
	amounts := d.amounts()
	if !d.allAmountsEqual(amounts) {
		return fmt.Errorf("some of confirmation amounts aren't equal others confirms=%v",
			d.MigrationDaemonConfirms)
	}

	// transfer to deposit
	err = d.acquireFundAssets(proposedAmount, requestRef, toMember)
	if err != nil {
		return errors.Wrap(err, "failed to acquire assets from migration admin fund")
	}

	// activate deposit
	d.Amount = proposedAmount
	currentPulse, err := foundation.GetPulseNumber()
	if err != nil {
		return errors.Wrap(err, "failed to get current pulse")
	}
	d.PulseDepositUnHold = currentPulse + insolar.PulseNumber(d.Lockup)
	d.IsConfirmed = true

	// create additional deposit with linear vesting process
	err = d.createAdditionalDeposit(requestRef, toMember)
	if err != nil {
		return errors.Wrap(err, "failed to create additional linear deposit")
	}
	return nil
}

// TransferToDeposit transfers funds to deposit.
func (d *Deposit) TransferToDeposit(
	amountStr string,
	toDeposit insolar.Reference,
	fromMember insolar.Reference,
	request insolar.Reference,
	toMember insolar.Reference,
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

func (d *Deposit) amounts() []string {
	var amounts []string
	for _, amount := range d.MigrationDaemonConfirms {
		amounts = append(amounts, amount)
	}
	return amounts
}

func (d *Deposit) allAmountsEqual(amounts []string) bool {
	if len(amounts) < 1 {
		return false
	}
	amount := amounts[0]
	for i := 1; i < len(amounts); i++ {
		if amounts[i] != amount {
			return false
		}
	}
	return true
}

func (d *Deposit) acquireFundAssets(amount string, requestRef insolar.Reference, toMember insolar.Reference) error {
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
		amount, d.GetReference(), appfoundation.GetMigrationAdminMember(), requestRef, toMember,
	)
	if err != nil {
		return errors.Wrap(err, "failed to transfer from migration deposit to deposit")
	}
	return nil
}

func (d *Deposit) createAdditionalDeposit(requestRef insolar.Reference, targetMember insolar.Reference) error {
	// get target wallet object
	targetWallet, err := member.GetObject(targetMember).GetWallet()
	if err != nil {
		if strings.Contains(err.Error(), "index not found") {
			return fmt.Errorf("target member does not exist")
		}
		return errors.Wrap(err, "failed to get target wallet")
	}
	targetWalletObj := wallet.GetObject(*targetWallet)

	// get target deposit info
	depositInfoFace, err := d.Itself()
	if err != nil {
		return errors.Wrap(err, "failed to get deposit itself")
	}

	depositInfo, ok := depositInfoFace.(*DepositOut)
	if !ok {
		return fmt.Errorf("failed to assert deposit.Itseft() result to *deposit.DepositOut actualType=%T",
			depositInfoFace)
	}

	// calc new vesting parameters
	vestingParams, err := migrationadmin.GetObject(appfoundation.GetMigrationAdmin()).GetLinearDepositParameters()
	if err != nil {
		return errors.Wrap(err, "failed to get linear deposit parameters")
	}
	pulseDepositUnHold := pulse.Number(depositInfo.HoldStartDate + vestingParams.Lockup)

	// try to create new deposit
	newTxHash := depositInfo.TxHash + "_2"
	const (
		ZeroBalance    = "0"
		FullyConfirmed = true
	)
	newDeposit, err := targetWalletObj.FindOrCreateDeposit(
		newTxHash,
		vestingParams.Lockup,
		vestingParams.Vesting,
		vestingParams.VestingStep,
		ZeroBalance,
		pulseDepositUnHold,
		depositInfo.MigrationDaemonConfirms,
		depositInfo.Amount,
		appfoundation.LinearVesting,
		FullyConfirmed,
	)
	if err != nil {
		return errors.Wrap(err, "failed to find or create deposit")
	}

	// migrate money to new deposit
	maRef := appfoundation.GetMigrationAdminMember()
	ma := member.GetObject(maRef)
	adminWalletRef, err := ma.GetWallet()
	if err != nil {
		return errors.Wrap(err, "failed to get wallet")
	}
	ok, fund, _ := wallet.GetObject(*adminWalletRef).FindDeposit(PublicAllocation2DepositName)
	if !ok {
		return fmt.Errorf("failed to find source deposit - %s", adminWalletRef.String())
	}
	return deposit.GetObject(*fund).TransferToDeposit(
		depositInfo.Amount,
		*newDeposit,
		maRef,
		requestRef,
		targetMember,
	)
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

func assertBigInt(amount string) (*big.Int, bool) {
	return new(big.Int).SetString(amount, 10)
}

func lessOrEqualZero(val *big.Int) bool {
	return val.Cmp(big.NewInt(0)) <= 0
}
