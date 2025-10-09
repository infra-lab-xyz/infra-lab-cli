package helm

import (
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
)

func httpRepoDelete(repo config.HelmRepo) (err error) {
	cfg := config.GetConfig()

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

func HTTPRepoDelete(repo config.HelmRepo) error {
	cfg := config.GetConfig()

	if !utils.IsBinaryInPath(cfg.Apps.Helm.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Helm.Binary))
		return nil
	}

	repos, err := getHTTPRepos()
	if err != nil {
		return err
	}

	if isHTTPRepoNameExist(repo.Name, repos) {
		err = httpRepoDelete(repo)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("Repo %s does not exists. Nothing to do\n", repo.Name)
	}

	return nil
}
