package helm

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

// TODO: probably not needed anymore, since the flag would be provided and this is implemented in HELM feature
func HTTPRepoReAdd(repo HelmRepo) (err error) {
	if !utils.IsBinaryInPath(cfg.Apps.Helm.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Helm.Binary))
		return nil
	}

	if isRepoNameExist(repo.Name) {
		err = httpRepoDelete(repo)
		if err != nil {
			return err
		}
	}

	err = httpRepoAdd(repo.Name, repo.Url, false)
	if err != nil {
		return err
	}

	return nil
}
