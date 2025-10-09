package podman

import (
	podmansrc "infra-lab-cli/src/podman"

	"github.com/spf13/cobra"
)

var RestartMachineCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart podman machine",
	RunE:  runRestartMachine,
}

func runRestartMachine(cmd *cobra.Command, args []string) error {
	return podmansrc.RestartMachine(machineName)
}
