package helm

import (
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
)

func httpRepoUpdate(repo HelmRepo) (err error) {
	cfg := config.GetConfig()

	args := fmt.Sprintf("repo update %s", repo.Name)

	_, _, err = utils.ExecBinaryCommand(
		cfg.Apps.Helm.Binary,
		args,
		true,
		false,
		[]string{},
	)

	return err
}

// TODO: probably not needed anymore, since I'll have another entrance point
func HTTPRepoUpdate(repo HelmRepo) error {
	if !utils.IsBinaryInPath(cfg.Apps.Helm.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Helm.Binary))
		return nil
	}

	if isRepoNameExist(repo.Name) {
		err := httpRepoUpdate(repo)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("Repo %s does not exists. Nothing to do\n", repo.Name)
	}

	return nil
}
