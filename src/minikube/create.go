package minikube

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

func createCluster(cluster Cluster) (err error) {
	fmt.Printf("Creating cluster: \n")
	fmt.Printf("\tName: %s\n", cluster.Name)
	fmt.Printf("\tNodes: %d\n", cluster.NodesCount)
	fmt.Printf("\tCPUs: %s\n", cluster.Config.CPUsFlag)
	fmt.Printf("\tMemory: %s\n", cluster.Config.MemoryFlag)
	fmt.Printf("\tDiskSize: %s\n", cluster.Config.DiskSizeFlag)
	fmt.Printf("\tK8S version: %s\n", cluster.Config.KubeConfig.KubeVersion)
	fmt.Printf("\tCNI: %s\n", cluster.CNI)
	fmt.Printf("\tDriver: %s\n", cluster.Config.Driver)
	fmt.Printf("\tCIDR: %s\n", cluster.CIDR)

	cmdArgs := utils.MapToString(map[string]any{
		"--cpus":               cluster.Config.CPUsFlag,
		"--memory":             cluster.Config.MemoryFlag,
		"--disk-size":          cluster.Config.DiskSizeFlag,
		"--nodes":              cluster.NodesCount,
		"--kubernetes-version": cluster.Config.KubeConfig.KubeVersion,
		"--extra-config=kubeadm.pod-network-cidr": cluster.CIDR,
		"--driver": cluster.Config.Driver,
	},
		"=",
	)
	_, _, err = utils.ExecBinaryCommand(
		cfg.Apps.Minikube.Binary,
		fmt.Sprintf("-p %s start %s %s",
			cluster.Name,
			cmdArgs,
			cluster.ExtraArgs,
		),
		true,
		false,
		[]string{},
	)

	return err
}

func CreateCluster(cluster Cluster) error {
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
		fmt.Printf("Cluster %s already exists. Please use recreate command instead\n", cluster.Name)
	} else {
		err = createCluster(cluster)
		if err != nil {
			return err
		}
	}

	return nil
}
