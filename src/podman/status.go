package podman

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

func GetMachineStatus(machineName string) error {
	if !utils.IsBinaryInPath(cfg.Apps.Podman.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Podman.Binary))
		return nil
	}

	machine, err := InspectMachine(machineName)
	if err != nil {
		return err
	}

	fmt.Printf("%s\t%s\t%d cpu\t%s\t%s\n",
		machine.Name, machine.State,
		machine.Resources.CPUs,
		utils.ConvertToDesiredUnit(fmt.Sprintf("%d%s", machine.Resources.Memory, "M"), "G").FloatStr,
		utils.ConvertToDesiredUnit(fmt.Sprintf("%d%s", machine.Resources.DiskSize, "G"), "G").IntStr,
	)

	return nil
}
