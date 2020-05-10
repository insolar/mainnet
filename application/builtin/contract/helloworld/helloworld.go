// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

package helloworld

import (
"github.com/insolar/insolar/logicrunner/builtin/foundation"
)

type HelloWorld struct {
	foundation.BaseContract
	Message string
}

func New() (*HelloWorld, error) {
	return &HelloWorld{Message: "Hello, world"}, nil
}

func (hw *HelloWorld) ShowHelloWorldMessage() (string, error) {
	return hw.Message, nil
}
