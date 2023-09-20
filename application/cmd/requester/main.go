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
