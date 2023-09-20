//go:build functest
// +build functest

package functest

import (
	"fmt"
	"testing"

	"github.com/insolar/insolar/applicationbase/testutils/launchnet"
	"github.com/insolar/insolar/applicationbase/testutils/testrequest"

	"github.com/stretchr/testify/require"
)

func TestMemberGet(t *testing.T) {
	member1 := *createMember(t)
	member2, _ := newUserWithKeys()
	member2.PubKey = member1.PubKey
	member2.PrivKey = member1.PrivKey
	res, err := testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, member2, "member.get", nil)
	require.Nil(t, err)
	require.Equal(t, member1.Ref, res.(map[string]interface{})["reference"].(string))
}

func TestMigrationMemberGet(t *testing.T) {
	member1, _ := newUserWithKeys()

	res1, err := testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, member1, "member.migrationCreate", nil)
	require.Nil(t, err)

	decodedRes1, ok := res1.(map[string]interface{})
	require.True(t, ok, fmt.Sprintf("failed to decode: expected map[string]interface{}, got %T", res1))

	res2, err := testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, member1, "member.get", nil)
	require.Nil(t, err)

	decodedRes2, ok := res2.(map[string]interface{})
	require.True(t, ok, fmt.Sprintf("failed to decode: expected map[string]interface{}, got %T", res2))

	require.Equal(t, decodedRes1["reference"].(string), decodedRes2["reference"].(string))
	require.Equal(t, decodedRes1["migrationAddress"], res2.(map[string]interface{})["migrationAddress"].(string))
}

func TestMemberGetWrongPublicKey(t *testing.T) {
	member1, _ := newUserWithKeys()
	_, err := testrequest.SignedRequestWithEmptyRequestRef(t, launchnet.TestRPCUrlPublic, member1, "member.get", nil)
	data := checkConvertRequesterError(t, err).Data
	require.Contains(t, data.Trace, "failed to find reference by key")
}

func TestMemberGetGenesisMember(t *testing.T) {
	res, err := testrequest.SignedRequest(t, launchnet.TestRPCUrlPublic, &MigrationAdmin, "member.get", nil)
	require.Nil(t, err)
	require.Equal(t, MigrationAdmin.Ref, res.(map[string]interface{})["reference"].(string))
}
