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

func isHTTPRepoNameExist(repoName string, repos []HelmRepo) (exist bool) {
	for _, repo := range repos {
		if repo.Name == repoName {
			return true
		}
	}
	return false
}
