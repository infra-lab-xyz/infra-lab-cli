package helm

import (
	"encoding/json"
	"fmt"
	"infra-lab-cli/src/utils"
	"slices"
	"strings"
)

func getHTTPRepos() (repos []HelmRepo, err error) {
	args := "repo list --output json"
	stdout, _, err := utils.ExecBinaryCommand(
		cfg.Apps.Helm.Binary,
		args,
		false,
		false,
		[]string{},
	)
	if err != nil {
		return nil, err
	}

	data := strings.Join(stdout, "\n")
	jsonBytes := []byte(data)
	err = json.Unmarshal(jsonBytes, &repos)
	if err != nil {
		return nil, fmt.Errorf("failed to get the list repos: %v", err)
	}

	return repos, nil
}

func GetRepos() (repos []HelmRepo, err error) {
	httpRepos, _ := getHTTPRepos()

	_ = LoadOCIRepos()
	ociRepos := *GetOCIRepos()

	repos = slices.Concat(httpRepos, ociRepos.Repos)
	slices.SortFunc(repos, func(a, b HelmRepo) int {
		return strings.Compare(a.Name, b.Name)
	})

	return repos, nil
}
