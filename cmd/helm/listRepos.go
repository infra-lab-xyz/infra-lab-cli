package helm

import (
	helmsrc "infra-lab-cli/src/helm"

	"github.com/spf13/cobra"
)

var ListReposCmd = &cobra.Command{
	Use:     "repos-list",
	Aliases: []string{},
	Short:   "List HTTP and OCI repos",
	RunE:    runListRepos,
}

func runListRepos(cmd *cobra.Command, args []string) error {
	return helmsrc.ListRepos()
}
