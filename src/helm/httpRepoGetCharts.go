package helm

import (
	"encoding/json"
	"fmt"
	"infra-lab-cli/src/utils"
	"strings"
)

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
