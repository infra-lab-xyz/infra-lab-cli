package kind

import (
	kindsrc "infra-lab-cli/src/kind"

	"github.com/spf13/cobra"
)

var DeleteClustersCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Delete cluster",
	RunE:    runDeleteClusters,
}

func runDeleteClusters(cmd *cobra.Command, args []string) error {
	return kindsrc.DeleteCluster(cluster)
}
