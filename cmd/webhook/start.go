package webhook

import (
	webhooksrc "infra-lab-cli/src/webhook"

	"github.com/spf13/cobra"
)

var StartWebhookCmd = &cobra.Command{
	Use:   "start",
	Short: "Start webhook",
	RunE:  runStartWebhook,
}

func runStartWebhook(cmd *cobra.Command, args []string) error {
	return webhooksrc.StartWebhook(webhook)
}
