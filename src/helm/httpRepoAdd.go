package helm

import (
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
)

func httpRepoAdd(repo config.HelmRepo) (err error) {
	args := fmt.Sprintf("repo add %s %s", repo.Name, repo.Url)

	_, _, err = utils.ExecBinaryCommand(
		cfg.Apps.Helm.Binary,
		args,
		true,
		false,
		[]string{},
	)

	return err
}

func HTTPRepoAdd(repo config.HelmRepo) error {
	if !utils.IsBinaryInPath(cfg.Apps.Helm.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Helm.Binary))
		return nil
	}

	repos, err := getHTTPRepos()
	if err != nil {
		return err
	}

	if isHTTPRepoNameExist(repo.Name, repos) {
		fmt.Printf("Repo %s already exists. Please use re-add command instead\n", repo.Name)
	} else {
		err = httpRepoAdd(repo)
		if err != nil {
			return err
		}
	}

	return nil
}
