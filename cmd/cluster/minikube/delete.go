package minikube

import (
	mksrc "infra-lab-cli/src/minikube"

	"github.com/spf13/cobra"
)

var DeleteClusterCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete cluster",
	Aliases: []string{"d", "del"},
	RunE:    runDeleteCluster,
}

func runDeleteCluster(cmd *cobra.Command, args []string) error {
	return mksrc.DeleteCluster(cluster)
}
