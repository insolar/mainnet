package functest

import (
	"github.com/insolar/insolar/applicationbase/testutils/launchnet"
	"github.com/insolar/insolar/applicationbase/testutils/testrequest"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestCreateHelloWorldCall (t *testing.T) {
	res, err := testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, &Root, "helloworld.create", nil)
	require.Nil(t, err)
	require.NotEqual(t, "", res.(map[string]interface{})["helloWorldRef"].(string))
}
