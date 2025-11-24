package minikube

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

func Tunnel(cluster Cluster) (err error) {
	if !utils.IsBinaryInPath(cfg.Apps.Minikube.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Minikube.Binary))
		return nil
	}

	_, _, err = utils.ExecBinaryCommand(
		cfg.Apps.Minikube.Binary,
		fmt.Sprintf("-p %s tunnel", cluster.Name),
		true,
		true,
		[]string{},
	)
	if err != nil {
		return err
	}

	return nil
}
