package podman

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

func DeleteMachine(machineName string) (err error) {
	if !utils.IsBinaryInPath(cfg.Apps.Podman.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Podman.Binary))
		return nil
	}

	if !isMachineExist(machineName) {
		return fmt.Errorf("machine %s does not exist", machineName)
	}

	_, _, err = utils.ExecBinaryCommand(
		cfg.Apps.Podman.Binary,
		fmt.Sprintf(
			"machine rm --force %s",
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
