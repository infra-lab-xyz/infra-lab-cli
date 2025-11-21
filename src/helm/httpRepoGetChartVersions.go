package helm

import (
	"encoding/json"
	"fmt"
	"infra-lab-cli/src/utils"
	"strings"
)

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
