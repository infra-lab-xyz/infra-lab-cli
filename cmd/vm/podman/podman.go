package podman

import (
	"infra-lab-cli/config"
	podmansrc "infra-lab-cli/src/podman"
	"infra-lab-cli/src/utils"
	"strconv"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "podman",
	Short: "Manage podman machines",
}

var machineName string
var binaryName = "podman"
var defaultMachineName string
var connections []podmansrc.Connection
var cfg config.ILCConfig

func machineNameCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	machineNames, err := podmansrc.GetMachineNames(&connections)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}
	return machineNames, cobra.ShellCompDirectiveNoFileComp
}

func commonFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&cpus, "cpus", "c", cfg.Apps.Podman.CPUs, "Number of CPUs to allocate to the podman machine")
	cmd.Flags().StringVarP(&memory, "memory", "m", cfg.Apps.Podman.Memory, "Memory in GiB or in MiB to allocate to the podman machine. E.g. 2G, 2048, 2048M")
	cmd.Flags().StringVarP(&diskSize, "disk-size", "d", cfg.Apps.Podman.DiskSize, "Disk size in GiB for the podman machine")
}

func commonParams(cmd *cobra.Command, cpus, memory, diskSize string) podmansrc.ConfigParams {
	cpuVal, _ := strconv.Atoi(cpus)
	memVal := utils.ConvertToDesiredUnit(memory, "M")
	diskSizeVal := utils.ConvertToDesiredUnit(diskSize, "G")
	return podmansrc.ConfigParams{
		CPUs:     podmansrc.ConfigParam{ValueFlag: cpuVal, IsProvided: cmd.Flags().Changed("cpus")},
		Memory:   podmansrc.ConfigParam{ValueFlag: memVal.IntValue, IsProvided: cmd.Flags().Changed("memory")},
		DiskSize: podmansrc.ConfigParam{ValueFlag: diskSizeVal.IntValue, IsProvided: cmd.Flags().Changed("disk-size")},
	}
}

func init() {
	cfg = *config.GetConfig()
	// TODO: Select the default machine name based on the default system connection
	// TODO: Add possibility to autocomplete machine name when using the `--name` flag. Correlated with the previous TODO.
	RootCmd.PersistentFlags().StringVarP(&binaryName, "binary", "b", cfg.Apps.Podman.Binary, "Binary to use")

	_ = podmansrc.GetConnections(&connections)
	_ = podmansrc.GetDefaultMachineName(&connections, &defaultMachineName)

	RootCmd.PersistentFlags().StringVarP(&machineName, "name", "n", defaultMachineName, "Name of the podman machine")
	_ = RootCmd.RegisterFlagCompletionFunc("name", machineNameCompletion)

	RootCmd.AddCommand(ListMachinesCmd)
	RootCmd.AddCommand(StartMachineCmd)
	RootCmd.AddCommand(StopMachineCmd)
	RootCmd.AddCommand(RestartMachineCmd)
	RootCmd.AddCommand(ConfigMachineCmd)
	RootCmd.AddCommand(CreateMachineCmd)
	RootCmd.AddCommand(DeleteMachineCmd)
	RootCmd.AddCommand(StatusCmd)
}
