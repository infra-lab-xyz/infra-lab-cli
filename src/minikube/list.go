package minikube

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

// TODO: different style. Yet easier to implement

func ListProfiles() (err error) {
	if !utils.IsBinaryInPath(cfg.Apps.Minikube.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Minikube.Binary))
		return nil
	}

	_, _, err = utils.ExecBinaryCommand(
		cfg.Apps.Minikube.Binary,
		"profile list",
		true,
		false,
		[]string{},
	)
	if err != nil {
		return err
	}

	return nil
}
