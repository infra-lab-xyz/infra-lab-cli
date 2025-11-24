package helm

import (
	"errors"
	"infra-lab-cli/config"
	"net/url"

	"github.com/spf13/viper"
)

var cfg *config.ILCConfig

func init() {
	cfg = config.GetConfig()
	OCIReposConfig = viper.New()
}

func isRepoNameExist(repoName string) (exist bool) {
	repos, _ := GetRepos()

	for _, repo := range repos {
		if repo.Name == repoName {
			return true
		}
	}
	return false
}

func getRepoByName(repoName string) (repo HelmRepo) {
	repos, _ := GetRepos()

	for _, repo := range repos {
		if repo.Name == repoName {
			return repo
		}
	}
	return repo
}

func parseSchema(repoUrl string) (schema string, err error) {
	parsedUrl, err := url.Parse(repoUrl)
	if err != nil {
		return "", err
	}

	schema = parsedUrl.Scheme
	if schema == "" {
		return "", errors.New("no schema provided or wrong Url")
	}

	return schema, nil
}
