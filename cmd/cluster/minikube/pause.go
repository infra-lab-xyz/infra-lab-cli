package minikube

import (
	mksrc "infra-lab-cli/src/minikube"

	"github.com/spf13/cobra"
)

var PauseClusterCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause cluster if exist",
	RunE:  runPauseCluster,
}

func runPauseCluster(cmd *cobra.Command, args []string) error {
	return mksrc.PauseCluster(cluster)
}
