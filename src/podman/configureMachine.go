package podman

import (
	"fmt"
	"infra-lab-cli/src/utils"
	"strconv"
)

func ConfigureMachine(machineName string, params ConfigParams) error {
	// TODO: is it wise to move this check to a function, or this action would not help with code duplication?
	if !utils.IsBinaryInPath(cfg.Apps.Podman.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cfg.Apps.Podman.Binary))
		return nil
	}

	machine, err := InspectMachine(machineName)
	if err != nil {
		return err
	}

	// TODO: don't know how to test it
	err = checkIfParamsWereChanged(&params, machine)
	if err != nil {
		fmt.Printf("An error occurred while checking if params were changed: %s\n", err)
		return nil
	}
	if !params.CPUs.IsChanged && !params.Memory.IsChanged && !params.DiskSize.IsChanged {
		fmt.Println("No changes detected in configuration.")
		return nil
	}

	isRunning := machine.State == "running"
	if isRunning {
		err := StopMachine(machineName)
		if err != nil {
			return err
		}
	}

	if params.CPUs.IsChanged {
		_, _, err := utils.ExecBinaryCommand(
			cfg.Apps.Podman.Binary,
			fmt.Sprintf("machine set --cpus %s %s", strconv.Itoa(params.CPUs.Value), machineName),
			false,
			false,
			[]string{},
		)
		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Printf("CPU was updated from %d to %d\n", machine.Resources.CPUs, params.CPUs.Value)
	}

	if params.Memory.IsChanged {
		_, _, err := utils.ExecBinaryCommand(
			cfg.Apps.Podman.Binary,
			fmt.Sprintf("machine set --memory %s %s", strconv.Itoa(params.Memory.Value), machineName),
			false,
			false,
			[]string{},
		)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("Memory was updated from %s to %s\n",
			utils.ConvertToDesiredUnit(fmt.Sprintf("%d%s", machine.Resources.Memory, "M"), "G").FloatStr,
			utils.ConvertToDesiredUnit(fmt.Sprintf("%d%s", params.Memory.Value, "M"), "G").FloatStr,
		)
	}

	if params.DiskSize.IsChanged {
		if params.DiskSize.Value > machine.Resources.DiskSize {
			_, _, err := utils.ExecBinaryCommand(
				cfg.Apps.Podman.Binary,
				fmt.Sprintf("machine set --disk-size %s %s", strconv.Itoa(params.DiskSize.Value), machineName),
				false,
				false,
				[]string{},
			)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Printf("Disk size was updated from %s to %s\n",
				utils.ConvertToDesiredUnit(fmt.Sprintf("%d%s", machine.Resources.DiskSize, "G"), "G").IntStr,
				utils.ConvertToDesiredUnit(fmt.Sprintf("%d%s", params.DiskSize.Value, "G"), "G").IntStr,
			)
		} else {
			fmt.Println("Disk size must be greater than the current one.")
		}
	}

	if isRunning {
		err = StartMachine(machineName)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkIfParamChanged(param *ConfigParam, currentValue int) (err error) {
	if param.ValueFlag != currentValue {
		param.IsChanged = true
	}
	return nil
}

func checkIfParamsWereChanged(params *ConfigParams, machine *InspectedMachine) (err error) {
	if params.CPUs.IsProvided {
		err := checkIfParamChanged(&params.CPUs, machine.Resources.CPUs)
		if err != nil {
			return err
		}
	}

	if params.Memory.IsProvided {
		err := checkIfParamChanged(&params.Memory, machine.Resources.Memory)
		if err != nil {
			return err
		}
	}

	if params.DiskSize.IsProvided {
		err := checkIfParamChanged(&params.DiskSize, machine.Resources.DiskSize)
		if err != nil {
			return err
		}
	}

	return nil
}
