//go:build tools
// +build tools

package tools

import (
	_ "github.com/insolar/insolar/cmd/insgocc"
	_ "github.com/insolar/insolar/cmd/keeperd"
	_ "github.com/insolar/insolar/cmd/pulsard"
	_ "github.com/insolar/insolar/cmd/pulsewatcher"
	_ "golang.org/x/tools/cmd/stringer"
)
