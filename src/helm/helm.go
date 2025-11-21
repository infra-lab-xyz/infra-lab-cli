package helm

import (
	"infra-lab-cli/config"

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
