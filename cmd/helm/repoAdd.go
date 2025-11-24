package helm

import (
	"infra-lab-cli/config"
	helmsrc "infra-lab-cli/src/helm"

	"github.com/spf13/cobra"
)

var RepoAddCmd = &cobra.Command{
	Use:     "repo-add",
	Aliases: []string{},
	Short:   "Add HTTP or OCI repo",
	RunE:    runAddRepo,
}

func runAddRepo(cmd *cobra.Command, args []string) error {
	return helmsrc.RepoAdd(repoName, repoUrl, forceUpdate)
}

func init() {
	cfg = *config.GetConfig()
	ociRepos = *helmsrc.GetOCIRepos()

	RepoAddCmd.Flags().StringVarP(&repoName, "name", "n", "", "Name of the repository")
	RepoAddCmd.Flags().StringVarP(&repoUrl, "url", "u", "", "URL of the repository")
	RepoAddCmd.Flags().BoolVarP(&forceUpdate, "forceUpdate", "f", false, "Force update")
}
