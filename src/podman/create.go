package podman

import (
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
	"strconv"
)

func CreateMachine(machineName string, params ConfigParams, startMachine bool) (err error) {
	cfg := config.GetConfig()

	if !utils.IsBinaryInPath(cfg.Apps.Podman.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Podman.Binary))
		return nil
	}

	if isMachineExist(machineName) {
		return fmt.Errorf("machine %s already exists", machineName)
	}

	cmdArgs := utils.MapToString(map[string]string{
		"--cpus":      strconv.Itoa(params.CPUs.Value),
		"--memory":    strconv.Itoa(params.Memory.Value),
		"--disk-size": strconv.Itoa(params.DiskSize.Value),
	},
		" ",
	)
	if startMachine {
		cmdArgs += " --now"
	}
	_, _, err = utils.ExecBinaryCommand(
		cfg.Apps.Podman.Binary,
		fmt.Sprintf(
			"machine init %s %s",
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
