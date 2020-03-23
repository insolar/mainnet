//  Copyright 2020 Insolar Network Ltd.
//  All rights reserved.
//  This material is licensed under the Insolar License version 1.0,
//  available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

package main

import (
	"github.com/insolar/insolar/log"
	"github.com/insolar/mainnet/application/cmd/requester/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal("requester execution failed:", err)
	}
}
