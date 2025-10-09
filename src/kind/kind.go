package kind

import (
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
)

func getClusters() (stdout []string, err error) {
	cfg := config.GetConfig()

	stdout, _, err = utils.ExecBinaryCommand(
		cfg.Apps.Kind.Binary,
		"get clusters",
		false,
		false,
		[]string{},
	)
	if err != nil {
		return nil, err
	}

	return stdout, nil
}
