package minikube

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

func unpauseCluster(clusterName string) (err error) {
	_, _, err = utils.ExecBinaryCommand(
		cfg.Apps.Minikube.Binary,
		fmt.Sprintf("-p %s unpause", clusterName),
		true,
		false,
		[]string{},
	)
	if err != nil {
		return err
	}

	return nil
}

func UnpauseCluster(cluster Cluster) (err error) {
	if !utils.IsBinaryInPath(cfg.Apps.Minikube.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Minikube.Binary))
		return nil
	}

	clusters, err := getClusters()
	if err != nil {
		return err
	}
	existingCluster := getClusterIfExists(cluster, clusters)

	if existingCluster != nil {
		err = unpauseCluster(cluster.Name)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("Cluster %s does not exist.\n", cluster.Name)
	}

	return nil
}
