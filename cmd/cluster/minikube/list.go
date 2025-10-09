package minikube

import (
	mksrc "infra-lab-cli/src/minikube"

	"github.com/spf13/cobra"
)

var ListProfilesCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List profiles",
	RunE:    runListProfiles,
}

func runListProfiles(cmd *cobra.Command, args []string) error {
	return mksrc.ListProfiles()
}
