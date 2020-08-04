// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

package wallet

import (
	"fmt"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/logicrunner/builtin/foundation"
	"github.com/insolar/insolar/pulse"

	"github.com/insolar/mainnet/application/appfoundation"
	depositContract "github.com/insolar/mainnet/application/builtin/contract/deposit"
	"github.com/insolar/mainnet/application/builtin/proxy/account"
	"github.com/insolar/mainnet/application/builtin/proxy/deposit"
)

const (
	XNS = "XNS"
)

// Wallet - basic wallet contract.
type Wallet struct {
	foundation.BaseContract
	Accounts foundation.StableMap
	Deposits foundation.StableMap
}

// New creates new wallet.
func New(accountReference insolar.Reference) (*Wallet, error) {
	if accountReference.IsEmpty() {
		return nil, fmt.Errorf("reference is empty")
	}
	accounts := make(foundation.StableMap)
	// TODO: Think about creating of new types of assets and initial balance
	accounts[XNS] = accountReference.String()

	return &Wallet{
		Accounts: accounts,
		Deposits: make(foundation.StableMap),
	}, nil
}

// GetAccount returns account ref
// ins:immutable
func (w *Wallet) GetAccount(assetName string) (*insolar.Reference, error) {
	accountReference, ok := w.Accounts[assetName]
	if !ok {
		return nil, fmt.Errorf("asset not found: %s", assetName)
	}
	return insolar.NewObjectReferenceFromString(accountReference)
}

// Transfer transfers money to given wallet.
// ins:immutable
func (w *Wallet) Transfer(
	assetName string, amountStr string, toMember *insolar.Reference,
	fromMember insolar.Reference, request insolar.Reference,
) (interface{}, error) {
	accRef, err := w.GetAccount(assetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get account by asset name: %s", err.Error())
	}
	acc := account.GetObject(*accRef)
	return acc.Transfer(amountStr, toMember, fromMember, request)
}

// GetBalance gets balance by asset name.
// ins:immutable
func (w *Wallet) GetBalance(assetName string) (string, error) {
	accRef, err := w.GetAccount(assetName)
	if err != nil {
		return "", fmt.Errorf("failed to get account by asset: %s", err.Error())
	}
	acc := account.GetObject(*accRef)
	return acc.GetBalance()
}

// GetDeposits get all deposits for this wallet
// ins:immutable
func (w *Wallet) GetDeposits() ([]interface{}, error) {
	result := make([]interface{}, 0)
	for _, dRef := range w.Deposits {

		reference, err := insolar.NewObjectReferenceFromString(dRef)
		if err != nil {
			return nil, err
		}
		d := deposit.GetObject(*reference)

		depositInfo, err := d.Itself()
		if err != nil {
			return nil, fmt.Errorf("failed to get deposit itself: %s", err.Error())
		}

		result = append(result, depositInfo)
	}
	return result, nil
}

// FindDeposit finds deposit for this wallet with this transaction hash.
// ins:immutable
func (w *Wallet) FindDeposit(transactionHash string) (bool, *insolar.Reference, error) {
	if depositReferenceStr, ok := w.Deposits[transactionHash]; ok {
		depositReference, _ := insolar.NewObjectReferenceFromString(depositReferenceStr)
		return true, depositReference, nil
	}
	return false, nil, nil
}

// FindOrCreateDeposit finds deposit for this wallet with this transaction hash or creates new one with link in this wallet.
func (w *Wallet) FindOrCreateDeposit(
	balance string,
	pulseDepositUnHold pulse.Number,
	confirms []appfoundation.DaemonConfirm,
	amount string,
	txHash string,
	vestingType appfoundation.VestingType,
	lockup int64,
	vesting int64,
	vestingStep int64,
	isConfirmed bool,
) (*insolar.Reference, error) {
	found, dRef, err := w.FindDeposit(txHash)
	if err != nil {
		return nil, fmt.Errorf("failed to find deposit: %s", err.Error())
	}

	if found {
		return dRef, nil
	}

	dHolder := deposit.New(
		balance,
		pulseDepositUnHold,
		confirms,
		amount,
		txHash,
		vestingType,
		lockup,
		vesting,
		vestingStep,
		isConfirmed,
	)
	txDeposit, err := dHolder.AsChild(w.GetReference())
	if err != nil {
		return nil, fmt.Errorf("failed to save deposit as child: %s", err.Error())
	}

	ref := txDeposit.GetReference()
	w.Deposits[txHash] = ref.String()

	return &ref, err
}

// CreateFund creates new one public allocation 2 deposit with specified lockup end date.
func (w *Wallet) CreateFund(lockupEndDate int64) (*insolar.Reference, error) {
	depositHolder := deposit.NewFund(lockupEndDate)
	fund, err := depositHolder.AsChild(w.GetReference())
	if err != nil {
		return nil, fmt.Errorf("failed to save deposit as child: %s", err.Error())
	}

	ref := fund.GetReference()
	w.Deposits[depositContract.PublicAllocation2DepositName] = ref.String()
	return &ref, nil
}
