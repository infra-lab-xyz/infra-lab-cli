package minikube

import (
	mksrc "infra-lab-cli/src/minikube"

	"github.com/spf13/cobra"
)

var StopClusterCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop cluster",
	RunE:  runStopCluster,
}

func runStopCluster(cmd *cobra.Command, args []string) error {
	return mksrc.StopCluster(cluster)
}
