package main

import (
	"context"

	basebootstrap "github.com/insolar/insolar/applicationbase/bootstrap"
	"github.com/insolar/mainnet/application/genesis/contracts"
	"github.com/spf13/cobra"
)

func bootstrapCommand() *cobra.Command {
	var (
		configPath         string
		properName         bool
		certificatesOutDir string
	)
	c := &cobra.Command{
		Use:   "bootstrap",
		Short: "creates files required for new network (keys, genesis config)",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			contractsConfig, err := contracts.CreateGenesisContractsConfig(ctx, configPath)
			check("failed to create genesis contracts config", err)

			config, err := basebootstrap.ParseConfig(configPath)
			check("bootstrap config error", err)
			if certificatesOutDir != "" {
				config.CertificatesOutDir = certificatesOutDir
			}

			err = basebootstrap.NewGeneratorWithConfig(config, contractsConfig).Run(ctx, properName)
			check("base bootstrap failed to start", err)
		},
	}
	c.Flags().StringVarP(
		&configPath, "config", "c", "bootstrap.yaml", "path to bootstrap config")

	c.Flags().BoolVarP(
		&properName, "propername", "p", false, "whenever to use proper cert names")
	c.Flags().StringVarP(
		&certificatesOutDir, "certificates-out-dir", "o", "", "dir with certificate files")
	c.Flags().MarkDeprecated("certificates-out-dir", "please switch to 'certificates_out_dir:' in config")

	return c
}
