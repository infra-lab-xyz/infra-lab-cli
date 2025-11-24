package podman

import (
	"fmt"
	"infra-lab-cli/config"
)

var cfg *config.ILCConfig

func init() {
	cfg = config.GetConfig()
}

func isMachineExist(machineName string) (exist bool) {
	machines, err := GetMachineList()
	if err != nil {
		fmt.Printf("Error getting machines list: %s\n", err)
	}
	for idx := range machines {
		if machines[idx].Name == machineName {
			return true
		}
	}

	return false
}
