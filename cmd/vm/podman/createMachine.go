package podman

import (
	podmansrc "infra-lab-cli/src/podman"

	"github.com/spf13/cobra"
)

var startMachine bool

var CreateMachineCmd = &cobra.Command{
	Use:   "create",
	Short: "Create podman machine",
	RunE:  runCreateMachine,
}

func runCreateMachine(cmd *cobra.Command, args []string) error {

	return podmansrc.CreateMachine(
		machineName,
		commonParams(cmd, cpus, memory, diskSize),
		startMachine,
	)
}

func init() {
	commonFlags(CreateMachineCmd)
	CreateMachineCmd.Flags().BoolVar(&startMachine, "start", true, "Start the podman machine after creation")
}
