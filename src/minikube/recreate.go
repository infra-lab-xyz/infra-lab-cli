package minikube

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

func RecreateCluster(cluster Cluster) (err error) {
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
		err = deleteCluster(cluster.Name)
		if err != nil {
			return err
		}
	}

	err = createCluster(cluster)
	if err != nil {
		return err
	}

	return nil
}
