package minikube

func RestartCluster(cluster Cluster) (err error) {
	err = StopCluster(cluster)
	if err != nil {
		return err
	}

	err = StartCluster(cluster)
	if err != nil {
		return err
	}

	return nil
}
