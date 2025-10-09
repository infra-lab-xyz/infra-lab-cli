package podman

import (
	podmansrc "infra-lab-cli/src/podman"

	"github.com/spf13/cobra"
)

var ListMachinesCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List podman machines",
	RunE:    runListMachines,
}

func runListMachines(cmd *cobra.Command, args []string) error {
	return podmansrc.ListMachines()
}
