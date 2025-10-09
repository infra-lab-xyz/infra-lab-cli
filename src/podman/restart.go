package podman

import (
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
)

func RestartMachine(machineName string) (err error) {
	cfg := config.GetConfig()

	if !utils.IsBinaryInPath(cfg.Apps.Podman.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Podman.Binary))
		return nil
	}

	machine, err := InspectMachine(machineName)
	if err != nil {
		return err
	}

	if machine.State == "running" {
		err = StopMachine(machineName)
		if err != nil {
			return err
		}
	}

	err = StartMachine(machineName)
	if err != nil {
		return err
	}

	return nil
}
