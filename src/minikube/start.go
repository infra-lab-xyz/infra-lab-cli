package minikube

import (
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
)

func startCluster(cluster Cluster) (err error) {
	cfg := config.GetConfig()

	_, _, err = utils.ExecBinaryCommand(
		cfg.Apps.Minikube.Binary,
		fmt.Sprintf("-p %s start", cluster.Name),
		true,
		false,
		[]string{},
	)
	if err != nil {
		return err
	}

	return nil
}

func StartCluster(cluster Cluster) error {
	cfg := config.GetConfig()

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
		fmt.Printf("Cluster %s already exists.\n", cluster.Name)
		// TODO: not sure which else statuses are possible and this is for sure a hardcode
		if existingCluster.Status == "Stopped" {
			fmt.Printf("Cluster %s is stopped. Trying to start it\n", cluster.Name)
			err = startCluster(cluster)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Cluster %s is not stopped. Do nothing\n", cluster.Name)
		}
	} else {
		fmt.Printf("Cluster %s does not exist.\n", cluster.Name)
	}

	return nil
}
