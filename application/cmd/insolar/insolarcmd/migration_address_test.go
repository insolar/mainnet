package insolarcmd

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/insolar/mainnet/application/appfoundation"
)

func TestGenerateMigrationAddresses(t *testing.T) {
	genAddrsCount := 12
	var b bytes.Buffer
	err := GenerateMigrationAddresses(&b, genAddrsCount)
	require.NoError(t, err, "generete error check")

	var addresses []string
	err = json.Unmarshal(b.Bytes(), &addresses)
	require.NoError(t, err, "json unmarshal error check")

	require.Equal(t, genAddrsCount, len(addresses), "generates expected addresses count")

	for _, addr := range addresses {
		assert.Truef(t, appfoundation.IsEthereumAddress(addr), "validate ethereum address %v", addr)
	}
}
