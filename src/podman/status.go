package podman

import (
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
)

func GetMachineStatus(machineName string) error {
	cfg := config.GetConfig()

	if !utils.IsBinaryInPath(cfg.Apps.Podman.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Podman.Binary))
		return nil
	}

	machine, err := InspectMachine(machineName)
	if err != nil {
		return err
	}

	fmt.Printf("%s\t %s\t %d cpu\t %.1f GiB\t %d GiB\n",
		machine.Name, machine.State,
		machine.Resources.CPUs,
		utils.ConvertMiBToGiB(machine.Resources.Memory),
		machine.Resources.DiskSize)

	return nil
}
