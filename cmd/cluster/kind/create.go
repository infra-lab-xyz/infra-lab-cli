package kind

import (
	kindsrc "infra-lab-cli/src/kind"

	"github.com/spf13/cobra"
)

var CreateClustersCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Create cluster",
	RunE:    runCreateClusters,
}

func runCreateClusters(cmd *cobra.Command, args []string) error {
	return kindsrc.CreateCluster(cluster)
}
