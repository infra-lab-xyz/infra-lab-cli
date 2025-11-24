package podman

import (
	podmansrc "infra-lab-cli/src/podman"

	"github.com/spf13/cobra"
)

var (
	cpus     string
	memory   string
	diskSize string
)

var ConfigMachineCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure podman machine",
	RunE:  runConfigMachine,
}

// TODO: Add recreate flag?
func runConfigMachine(cmd *cobra.Command, args []string) error {
	if !cmd.Flags().Changed("cpus") &&
		!cmd.Flags().Changed("memory") &&
		!cmd.Flags().Changed("disk-size") {
		return cmd.Help()
	}

	return podmansrc.ConfigureMachine(
		machineName,
		commonParams(cmd, cpus, memory, diskSize),
	)
}

func init() {
	commonFlags(ConfigMachineCmd)
}
