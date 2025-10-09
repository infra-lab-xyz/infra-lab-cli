package kind

import (
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
)

// TODO: different style. Yet easier to implement. Should be standardized in the future

func ListClusters() (err error) {
	cfg := config.GetConfig()

	if !utils.IsBinaryInPath(cfg.Apps.Kind.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Kind.Binary))
		return nil
	}

	_, _, err = utils.ExecBinaryCommand(
		cfg.Apps.Kind.Binary,
		"get clusters",
		true,
		false,
		[]string{},
	)
	if err != nil {
		return err
	}

	return nil
}
