package kind

import (
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
)

var cfg *config.ILCConfig

func init() {
	cfg = config.GetConfig()
}

func getClusters() (stdout []string, err error) {
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
