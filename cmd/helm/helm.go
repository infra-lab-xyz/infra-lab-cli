package helm

import (
	"infra-lab-cli/config"
	helmSrc "infra-lab-cli/src/helm"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "helm",
	Aliases: []string{},
	Short:   "Helm commands",
}

var cfg config.ILCConfig
var ociRepos helmSrc.OCIRepos
var helmBinaryName string
var skopeoBinaryName string

func init() {
	cfg = *config.GetConfig()
	ociRepos = *helmSrc.GetOCIRepos()

	// TODO: Well, I need 2 binaries for helm: helm and skopeo
	RootCmd.PersistentFlags().StringVarP(&helmBinaryName, "binary", "", cfg.Apps.Helm.Binary, "Helm binary to use")
	RootCmd.PersistentFlags().StringVarP(&skopeoBinaryName, "skopeo-binary", "", cfg.Apps.Skopeo.Binary, "Skopeo binary to use. Used for OCI repos")

	RootCmd.AddCommand(ListReposCmd)
}
