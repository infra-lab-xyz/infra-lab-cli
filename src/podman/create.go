package podman

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

func CreateMachine(machineName string, params ConfigParams, startMachine bool) (err error) {
	if !utils.IsBinaryInPath(cfg.Apps.Podman.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Podman.Binary))
		return nil
	}

	if isMachineExist(machineName) {
		return fmt.Errorf("machine %s already exists", machineName)
	}

	cmdArgs := utils.MapToString(map[string]any{
		"--cpus":      params.CPUs.ValueFlag,
		"--memory":    params.Memory.ValueFlag,
		"--disk-size": params.DiskSize.ValueFlag,
	},
		" ",
	)

	if startMachine {
		cmdArgs += " --now"
	}
	_, _, err = utils.ExecBinaryCommand(
		cfg.Apps.Podman.Binary,
		fmt.Sprintf(
			"machine init %s --rootful %s",
			cmdArgs,
			machineName,
		),
		true,
		false,
		[]string{},
	)
	if err != nil {
		return err
	}

	return nil
}
