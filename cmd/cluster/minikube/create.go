package minikube

import (
	"infra-lab-cli/config"
	mksrc "infra-lab-cli/src/minikube"

	"github.com/spf13/cobra"
)

var CreateClusterCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new cluster or recreate existing by using --recreate",
	Aliases: []string{"c"},
	RunE:    runCreateCluster,
}

var recreate bool

func runCreateCluster(cmd *cobra.Command, args []string) error {
	if recreate {
		return mksrc.RecreateCluster(cluster)
	} else {
		return mksrc.CreateCluster(cluster)
	}
}

func init() {
	cfg = *config.GetConfig()

	// TODO: somehow I could not figure out how to use `cluster.KubeVersions, _ = mksrc.GetSupportedKubeVersions(binaryName)` because it adds a list to completion as plain text
	CreateClusterCmd.Flags().StringVarP(&cluster.Config.KubeConfig.KubeVersion, "kubeVersion", "v", cfg.Apps.Minikube.KubeVersion, "Version of Kubernetes cluster")
	_ = CreateClusterCmd.RegisterFlagCompletionFunc("kubeVersion", getSupportedKubeVersions)

	// TODO: Should the default value be gathered from config in HOME dir?
	CreateClusterCmd.Flags().StringVarP(&cluster.Config.Driver, "driver", "", cfg.Apps.Minikube.Driver, "Which driver to use")
	_ = CreateClusterCmd.RegisterFlagCompletionFunc("driver", getSupportedDrivers)

	CreateClusterCmd.Flags().StringVarP(&cluster.CNI, "cni", "", cfg.Apps.Minikube.CNI, "The CNI binary to use")
	_ = CreateClusterCmd.RegisterFlagCompletionFunc("cni", getSupportedCNIs)

	CreateClusterCmd.Flags().StringVarP(&cluster.Config.CPUsFlag, "cpus", "c", cfg.Apps.Minikube.CPUs, "The amount of CPUs to use")
	CreateClusterCmd.Flags().StringVarP(&cluster.Config.MemoryFlag, "memory", "m", cfg.Apps.Minikube.Memory, "The amount of memory to use")
	CreateClusterCmd.Flags().StringVarP(&cluster.Config.DiskSizeFlag, "disk-size", "d", cfg.Apps.Minikube.DiskSize, "Disk size allocated to the minikube VM")
	CreateClusterCmd.Flags().StringVarP(&cluster.CIDR, "cidr", "", cfg.Apps.Minikube.CIDR, "The CIDR to use")
	CreateClusterCmd.Flags().IntVarP(&cluster.NodesCount, "nodes", "n", cfg.Apps.Minikube.NodesCount, "Number of nodes")
	CreateClusterCmd.Flags().BoolVarP(&recreate, "recreate", "", false, "Recreate existing cluster")
}
