package cmd

import (
	"infra-lab-cli/cmd/cluster"
	"infra-lab-cli/cmd/config"
	"infra-lab-cli/cmd/helm"
	"infra-lab-cli/cmd/vm"
	"infra-lab-cli/cmd/webhook"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "infra-lab-cli",
	Short: "Local sandbox automation CLI",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(vm.RootCmd)
	rootCmd.AddCommand(cluster.RootCmd)
	rootCmd.AddCommand(webhook.RootCmd)
	rootCmd.AddCommand(helm.RootCmd)
	rootCmd.AddCommand(config.RootCmd)
	rootCmd.AddCommand(statusCmd)
}
