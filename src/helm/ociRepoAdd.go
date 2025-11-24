package helm

import (
	"fmt"
)

func ociRepoAdd(repoName, repoUrl string, forceUpdate bool) (err error) {
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
				// TODO: probably I can just update the URL, this would lead to a single config update instead for double
				err = ociRepoDelete(repo.Name)
				if err != nil {
					return err
				}
				fmt.Println("OCI repo deleted")
			case "http", "https":
				err = httpRepoDelete(repo)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("unknown schema: %s", schema)
			}

		} else {
			fmt.Printf("The repo exists and forceUpdate flag was not provided\n")
			return nil
		}
	}

	ociRepos := GetOCIRepos()

	ociRepos.Repos = append(ociRepos.Repos, HelmRepo{
		Name: repoName,
		Url:  repoUrl,
	})

	err = SaveOCIRepos(ociRepos)
	fmt.Println("OCI repo added")

	return err
}
