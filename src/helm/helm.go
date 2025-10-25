package helm

import (
	"encoding/json"
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
	"slices"
	"strings"
)

var cfg *config.ILCConfig

func init() {
	cfg = config.GetConfig()
}

func getHTTPRepos() (repos []config.HelmRepo, err error) {
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

	for idx := range repos {
		repos[idx].Type = "http"
	}

	return repos, nil
}

func isHTTPRepoNameExist(repoName string, repos []config.HelmRepo) (exist bool) {
	for _, repo := range repos {
		if repo.Name == repoName {
			return true
		}
	}
	return false
}

func GetHTTPRepoCharts(repoName string) (charts []Chart, err error) {
	args := fmt.Sprintf("repo search list %s/ --output json", repoName)
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
	err = json.Unmarshal(jsonBytes, &charts)
	if err != nil {
		return nil, fmt.Errorf("failed to get the list repos: %v", err)
	}

	return charts, nil
}

func GetHTTPRepoChartsSlice(repoName string) (chartsSlice []string, err error) {
	charts, err := GetHTTPRepoCharts(repoName)
	if err != nil {
		return nil, err
	}

	for _, chart := range charts {
		chartsSlice = append(chartsSlice, chart.Name)
	}

	return chartsSlice, nil
}

func GetHTTPRepoChartVersions(chartName string) (chartVersions []Chart, err error) {
	args := fmt.Sprintf("repo search list %s --output json --verbose", chartName)
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
	err = json.Unmarshal(jsonBytes, &chartVersions)
	if err != nil {
		return nil, fmt.Errorf("failed to get the list repos: %v", err)
	}

	return chartVersions, nil
}

func GetHTTPRepoChartVersionsSlice(chartName string) (chartVersionsSlice []string, err error) {
	chartVersions, err := GetHTTPRepoChartVersions(chartName)
	if err != nil {
		return nil, err
	}

	for _, chartVersion := range chartVersions {
		chartVersionsSlice = append(chartVersionsSlice, chartVersion.Version)
	}

	return chartVersionsSlice, nil
}

func GetRepos() (repos []config.HelmRepo, err error) {
	httpRepos, _ := getHTTPRepos()

	_ = config.LoadOCIRepos()
	ociRepos := *config.GetOCIRepos()

	repos = slices.Concat(httpRepos, ociRepos.Repos)
	slices.SortFunc(repos, func(a, b config.HelmRepo) int {
		return strings.Compare(a.Name, b.Name)
	})

	return repos, nil
}
