//go:build functest
// +build functest

package functest

import (
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/insolar/insolar/applicationbase/testutils/launchnet"
	"github.com/insolar/insolar/insolar/defaults"
	"github.com/insolar/mainnet/application/cmd/insolar/insolarcmd"
	"github.com/insolar/mainnet/application/genesis/contracts"
)

func TestAddMigrationAddresses(t *testing.T) {
	extraAddrsDir := filepath.Join(defaults.LaunchnetConfigDir(), "extra_addrs")
	if _, err := os.Stat(extraAddrsDir); err == nil {
		// run this test only once
		t.Skip(extraAddrsDir, "extra addresses dir already exists")
	}

	bootCfg, err := contracts.ParseContractsConfig(filepath.Join(defaults.LaunchnetDir(), "bootstrap.yaml"))
	require.NoError(t, err, "bootstrap config parse")

	shardsCount := bootCfg.MAShardCount
	// one query gets 10 shards according to migrationadmin code: const maxNumberOfElements = 10
	startWithIndex := rand.Intn(shardsCount - 10)

	migrationShardsBefore := getAddressCount(t, startWithIndex)

	dirErr := os.Mkdir(extraAddrsDir, 0755)
	require.NoError(t, dirErr, "directory for additonal addresses creation error")

	addrsByShard := insolarcmd.NRandomMigrationAddressesSplitByShard(40000, shardsCount)
	genErr := insolarcmd.WritesShardedMigrationsAddressesToDir(extraAddrsDir, addrsByShard)
	require.NoError(t, genErr, "extra migration address files creation error")

	addErr := insolarcmd.AddMigrationAddresses(
		[]string{launchnet.TestRPCUrl},
		[]string{launchnet.TestRPCUrlPublic},
		defaults.LaunchnetConfigDir(),
		extraAddrsDir,
	)
	require.NoError(t, addErr, "extra migration address adding error")

	migrationShardsAfter := getAddressCount(t, startWithIndex)
	for n, addrsCount := range migrationShardsAfter {
		expectCount := migrationShardsBefore[n] + len(addrsByShard[n])
		assert.Equalf(t, addrsCount, expectCount, "%v addresses added to shard", len(addrsByShard[n]), n)
	}
}
