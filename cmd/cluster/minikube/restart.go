package minikube

import (
	mksrc "infra-lab-cli/src/minikube"

	"github.com/spf13/cobra"
)

var RestartClusterCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart cluster",
	RunE:  runRestartCluster,
}

func runRestartCluster(cmd *cobra.Command, args []string) error {
	return mksrc.RestartCluster(cluster)
}
