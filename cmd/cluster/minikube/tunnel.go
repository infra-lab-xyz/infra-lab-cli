package minikube

import (
	mksrc "infra-lab-cli/src/minikube"

	"github.com/spf13/cobra"
)

var TunnelCmd = &cobra.Command{
	Use:     "tunnel",
	Aliases: []string{"t"},
	Short:   "Tunnel creates a route to services deployed with type LoadBalancer and sets their Ingress to their ClusterIP. for a\ndetailed example see https://minikube.sigs.k8s.io/docs/tasks/loadbalancer",
	RunE:    runTunnel,
}

func runTunnel(cmd *cobra.Command, args []string) error {
	return mksrc.Tunnel(cluster)
}
