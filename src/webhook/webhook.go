package webhook

import (
	"infra-lab-cli/config"
)

var cfg *config.ILCConfig

func init() {
	cfg = config.GetConfig()
}
