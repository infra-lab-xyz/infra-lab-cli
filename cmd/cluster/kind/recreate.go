package kind

import (
	kindsrc "infra-lab-cli/src/kind"

	"github.com/spf13/cobra"
)

var RecreateClustersCmd = &cobra.Command{
	Use:     "recreate",
	Aliases: []string{"r"},
	Short:   "Recreate cluster",
	RunE:    runRecreateClusters,
}

func runRecreateClusters(cmd *cobra.Command, args []string) error {
	return kindsrc.RecreateCluster(cluster)
}
