package helm

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

func httpRepoAddExec(repoName, repoUrl string, forceUpdate bool) (err error) {
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

func httpRepoAdd(repoName, repoUrl string, forceUpdate bool) (err error) {
	if !utils.IsBinaryInPath(cfg.Apps.Helm.Binary) {
		return fmt.Errorf("%s", utils.BinaryNotFoundError(cfg.Apps.Helm.Binary))
	}

	repo := getRepoByName(repoName)

	if repo.Name != "" {
		fmt.Printf("The repo %s\t%s already exists\n", repo.Name, repo.Url)
		if forceUpdate {
			fmt.Printf("The repo will be updated from: %s to: %s\n", repo.Url, repoUrl)

			schema, err := parseSchema(repo.Url)
			if err != nil {
				return err
			}

			switch schema {
			case "oci":
				err = ociRepoDelete(repo.Name)
				if err != nil {
					return err
				}
			case "http", "https":
				break
			default:
				return fmt.Errorf("unknown schema: %s", schema)
			}

		} else {
			fmt.Printf("The repo exists and forceUpdate flag was not provided\n")
			return nil
		}
	}

	err = httpRepoAddExec(repoName, repoUrl, forceUpdate)
	if err != nil {
		return err
	}

	return nil
}
