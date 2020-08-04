// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

package appfoundation

type TransactionType string

const (
	TTypeMigration  TransactionType = "migration"
	TTypeAllocation TransactionType = "allocation"
)
