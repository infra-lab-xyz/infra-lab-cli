package minikube

import (
	mksrc "infra-lab-cli/src/minikube"

	"github.com/spf13/cobra"
)

var StartClusterCmd = &cobra.Command{
	Use:   "start",
	Short: "Start cluster if exist",
	RunE:  runStartCluster,
}

func runStartCluster(cmd *cobra.Command, args []string) error {
	return mksrc.StartCluster(cluster)
}
