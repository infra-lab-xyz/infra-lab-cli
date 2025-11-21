package helm

import (
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
	"os"

	"github.com/spf13/viper"
)

var OCIHelmRepos OCIRepos
var OCIReposConfig *viper.Viper

func GetOCIRepos() *OCIRepos {
	return &OCIHelmRepos
}

func LoadOCIRepos() (err error) {
	configPath := utils.ExpandPath(fmt.Sprintf("%s/oci-repos.yaml", config.ProjectPath))

	OCIReposConfig.SetConfigFile(configPath)
	_ = OCIReposConfig.ReadInConfig()
	if err = OCIReposConfig.Unmarshal(&OCIHelmRepos); err != nil {
		return err
	}

	return nil
}

func SaveOCIRepos() (err error) {
	// TODO: check if dir exists
	// OCIHelmRepos.Repos = append(OCIHelmRepos.Repos, HelmRepo{
	// 	Name: "vault-secrets-webhook",
	// 	Url:  "oci://ghcr.io/bank-vaults/helm-chart",
	// })
	// fmt.Println(OCIReposConfig.ConfigFileUsed())
	projectDir := utils.ExpandPath(config.ProjectPath)
	if !utils.IsDirExist(projectDir) {
		fmt.Printf("%s does not exists", config.ProjectPath)
		// TODO: perms should not be hardcoded
		err := os.Mkdir(config.ProjectPath, 0700)
		if err != nil {
			fmt.Printf("Error creating project directory: %s\nWith error: %v\n", projectDir, err)
		}
	}
	OCIReposConfig.Set("repos", OCIHelmRepos.Repos)
	err = OCIReposConfig.WriteConfig()
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
