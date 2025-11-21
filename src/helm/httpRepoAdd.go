package helm

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

func httpRepoAdd(repoName, repoUrl string, forceUpdate bool) (err error) {
	args := fmt.Sprintf("repo add %s %s", repoName, repoUrl)
	if forceUpdate {
		args = fmt.Sprintf("%s --force-update", args)
	}

	_, _, err = utils.ExecBinaryCommand(
		cfg.Apps.Helm.Binary,
		args,
		true,
		false,
		[]string{},
	)

	return err
}

// TODO: probably not needed anymore, since I have another entrance point
func HTTPRepoAdd(repo HelmRepo) error {
	if !utils.IsBinaryInPath(cfg.Apps.Helm.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Helm.Binary))
		return nil
	}

	if isRepoNameExist(repo.Name) {
		fmt.Printf("Repo %s already exists. Please use re-add command instead\n", repo.Name)
	} else {
		err := httpRepoAdd(repo.Name, repo.Url, false)
		if err != nil {
			return err
		}
	}

	return nil
}
