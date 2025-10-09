package minikube

import (
	mksrc "infra-lab-cli/src/minikube"

	"github.com/spf13/cobra"
)

var UnpauseClusterCmd = &cobra.Command{
	Use:   "unpause",
	Short: "Unpause cluster if exist",
	RunE:  runUnpauseCluster,
}

func runUnpauseCluster(cmd *cobra.Command, args []string) error {
	return mksrc.UnpauseCluster(cluster)
}
