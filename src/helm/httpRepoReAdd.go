package helm

import (
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
)

func HTTPRepoReAdd(repo config.HelmRepo) (err error) {
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
	}

	err = httpRepoAdd(repo)
	if err != nil {
		return err
	}

	return nil
}
