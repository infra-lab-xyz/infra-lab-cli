package podman

import (
	podmansrc "infra-lab-cli/src/podman"

	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status of podman machine",
	RunE:  runStatus,
}

func runStatus(cmd *cobra.Command, args []string) error {
	return podmansrc.GetMachineStatus(machineName)
}
