package podman

import (
	podmansrc "infra-lab-cli/src/podman"

	"github.com/spf13/cobra"
)

var StartMachineCmd = &cobra.Command{
	Use:   "start",
	Short: "Start podman machine",
	RunE:  runStartMachine,
}

func runStartMachine(cmd *cobra.Command, args []string) error {
	return podmansrc.StartMachine(machineName)
}
