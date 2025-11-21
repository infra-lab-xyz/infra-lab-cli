package config

import (
	"infra-lab-cli/config"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{},
	Short:   "config management",
}

var cfg config.ILCConfig

func init() {
	cfg = *config.GetConfig()
}
