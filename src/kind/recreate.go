package kind

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

func RecreateCluster(cluster Cluster) (err error) {
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

	err = createCluster(cluster)
	if err != nil {
		return err
	}

	return nil
}
