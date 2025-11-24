package helm

import (
	"infra-lab-cli/config"
	helmsrc "infra-lab-cli/src/helm"

	"github.com/spf13/cobra"
)

var RepoDeleteCmd = &cobra.Command{
	Use:     "repo-delete",
	Aliases: []string{},
	Short:   "Delete HTTP or OCI repo by name",
	RunE:    runRepoDelete,
}

func runRepoDelete(cmd *cobra.Command, args []string) error {
	return helmsrc.RepoDelete(repoName)
}

func init() {
	cfg = *config.GetConfig()
	ociRepos = *helmsrc.GetOCIRepos()

	RepoDeleteCmd.Flags().StringVarP(&repoName, "name", "n", "", "Name of the repository")
}
