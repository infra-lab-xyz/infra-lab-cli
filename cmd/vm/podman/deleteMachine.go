package podman

import (
	podmansrc "infra-lab-cli/src/podman"

	"github.com/spf13/cobra"
)

var DeleteMachineCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm", "del"},
	Short:   "Delete podman machine",
	RunE:    runDeleteMachine,
}

func runDeleteMachine(cmd *cobra.Command, args []string) error {
	return podmansrc.DeleteMachine(machineName)
}

func init() {
}
