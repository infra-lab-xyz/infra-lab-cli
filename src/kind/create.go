package kind

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

func createCluster(cluster Cluster) (err error) {
	args := fmt.Sprintf("create cluster --name %s", cluster.Name)
	if cluster.ConfigPath != "" {
		args += fmt.Sprintf(" --config=%s", cluster.ConfigPath)
	}

	_, _, err = utils.ExecBinaryCommand(
		cfg.Apps.Kind.Binary,
		args,
		true,
		false,
		[]string{},
	)

	return err
}

func CreateCluster(cluster Cluster) (err error) {
	if !utils.IsBinaryInPath(cfg.Apps.Kind.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Kind.Binary))
		return nil
	}

	// TODO: check if VM is running, but which podman, docker, colima, what if we have multiple online or want to run kind only in specific env?

	clusters, err := getClusters()
	if err != nil {
		return err
	}

	if utils.IfStringInSlice(cluster.Name, clusters) {
		fmt.Printf("Cluster %s already exists. Please use recreate command instead\n", cluster.Name)
	} else {
		err = createCluster(cluster)
		if err != nil {
			return err
		}
	}

	return nil
}
