// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

package appfoundation

// DaemonConfirm holds info (daemon reference and confirmed amount) about success daemon's confirmation requests.
type DaemonConfirm struct {
	Reference string `json:"reference"`
	Amount    string `json:"amount"`
}
