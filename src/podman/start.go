package podman

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

func StartMachine(machineName string) (err error) {
	if !utils.IsBinaryInPath(cfg.Apps.Podman.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Podman.Binary))
		return nil
	}

	_, _, err = utils.ExecBinaryCommand(
		cfg.Apps.Podman.Binary,
		fmt.Sprintf("machine start %s", machineName),
		true,
		false,
		[]string{},
	)
	if err != nil {
		return err
	}

	return nil
}
