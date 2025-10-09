package minikube

import (
	"encoding/json"
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
	"strings"
)

var cfg *config.ILCConfig

func init() {
	cfg = config.GetConfig()
}

// TODO: these functions look very similar and duplicated

func GetSupportedKubeVersions() (versions []string, err error) {
	stdout, _, err := utils.ExecBinaryCommand(
		cfg.Apps.Minikube.Binary,
		"config defaults kubernetes-version -o json",
		false,
		false,
		[]string{},
	)
	if err != nil {
		return nil, err
	}

	data := strings.Join(stdout, "\n")
	jsonBytes := []byte(data)
	err = json.Unmarshal(jsonBytes, &versions)
	if err != nil {
		return nil, fmt.Errorf("failed to get the list of supported k8s versions: %v", err)
	}

	return versions, nil
}

// TODO: Should I keep this function?
func ListSupportedKubeVersions() (err error) {
	versions, err := GetSupportedKubeVersions()
	if err != nil {
		return err
	}
	for _, version := range versions {
		fmt.Println(version)
	}
	return nil
}

func GetSupportedDrivers() (versions []string, err error) {
	stdout, _, err := utils.ExecBinaryCommand(
		cfg.Apps.Minikube.Binary,
		"config defaults driver -o json",
		false,
		false,
		[]string{},
	)
	if err != nil {
		return nil, err
	}

	data := strings.Join(stdout, "\n")
	jsonBytes := []byte(data)
	err = json.Unmarshal(jsonBytes, &versions)
	if err != nil {
		return nil, fmt.Errorf("failed to get the list of supported k8s versions: %v", err)
	}

	return versions, nil
}

func getClusters() (clusters []Cluster, err error) {
	stdout, _, err := utils.ExecBinaryCommand(
		cfg.Apps.Minikube.Binary,
		"profile list -o json",
		false,
		false,
		[]string{},
	)
	if err != nil {
		return nil, err
	}

	var mkList MkList
	data := strings.Join(stdout, "\n")
	jsonBytes := []byte(data)
	err = json.Unmarshal(jsonBytes, &mkList)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}
	clusters = mkList.Valid

	return clusters, nil
}

func getClusterIfExists(newCluster Cluster, clusters []Cluster) *Cluster {
	for _, cluster := range clusters {
		if cluster.Name == newCluster.Name {
			return &cluster
		}
	}
	return nil
}
