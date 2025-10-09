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

func runConfigMachine(cmd *cobra.Command, args []string) error {
	if !cmd.Flags().Changed("cpus") &&
		!cmd.Flags().Changed("memory") &&
		!cmd.Flags().Changed("disk-size") {
		return cmd.Help()
	}

	params := podmansrc.ConfigParams{
		CPUs:     podmansrc.ConfigParam{ValueFlag: cpus, IsProvided: cmd.Flags().Changed("cpus")},
		Memory:   podmansrc.ConfigParam{ValueFlag: memory, IsProvided: cmd.Flags().Changed("memory")},
		DiskSize: podmansrc.ConfigParam{ValueFlag: diskSize, IsProvided: cmd.Flags().Changed("disk-size")},
	}

	return podmansrc.ConfigureMachine(machineName, params)
}

func init() {
	commonFlags(ConfigMachineCmd)
}
