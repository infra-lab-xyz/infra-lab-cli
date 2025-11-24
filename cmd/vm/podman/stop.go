package podman

import (
	podmansrc "infra-lab-cli/src/podman"

	"github.com/spf13/cobra"
)

var StopMachineCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop podman machine",
	RunE:  runStopMachine,
}

func runStopMachine(cmd *cobra.Command, args []string) error {
	return podmansrc.StopMachine(machineName)
}
