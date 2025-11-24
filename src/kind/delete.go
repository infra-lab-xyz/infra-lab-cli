package kind

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

func deleteCluster(clusterName string) (err error) {
	_, _, err = utils.ExecBinaryCommand(
		cfg.Apps.Kind.Binary,
		fmt.Sprintf("delete cluster --name %s", clusterName),
		true,
		false,
		[]string{},
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCluster(cluster Cluster) (err error) {
	if !utils.IsBinaryInPath(cfg.Apps.Kind.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Kind.Binary))
		return nil
	}

	clusters, err := getClusters()
	if err != nil {
		return err
	}

	if utils.IfStringInSlice(cluster.Name, clusters) {
		err = deleteCluster(cluster.Name)
		if err != nil {
			return err
		}
	}

	return nil
}
