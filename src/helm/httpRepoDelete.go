package helm

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

func httpRepoDelete(repo HelmRepo) (err error) {
	args := fmt.Sprintf("repo rm %s", repo.Name)

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
func HTTPRepoDelete(repo HelmRepo) error {
	if !utils.IsBinaryInPath(cfg.Apps.Helm.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Helm.Binary))
		return nil
	}

	if isRepoNameExist(repo.Name) {
		err := httpRepoDelete(repo)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("Repo %s does not exists. Nothing to do\n", repo.Name)
	}

	return nil
}
