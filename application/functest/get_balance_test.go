//go:build functest
// +build functest

package functest

import (
	"math/big"
	"testing"

	"github.com/insolar/insolar/applicationbase/testutils/launchnet"
	"github.com/insolar/insolar/applicationbase/testutils/testrequest"
	"github.com/insolar/insolar/insolar/gen"

	"github.com/stretchr/testify/require"
)

func TestGetBalance(t *testing.T) {
	firstMember := createMember(t)
	firstBalance := getBalanceNoErr(t, firstMember, firstMember.Ref)
	r := big.NewInt(0)
	require.Equal(t, r, firstBalance)
}

func TestGetBalanceWrongRef(t *testing.T) {
	_, err := testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrl, &Root, "member.getBalance",
		map[string]interface{}{"reference": gen.Reference().String()})
	data := checkConvertRequesterError(t, err).Data
	require.Contains(t, data.Trace, "failed to fetch index from heavy")
}
