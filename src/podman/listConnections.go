package podman

import (
	"encoding/json"
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
	"strings"
)

// TODO: this function has wrong and confusing name

func getConnections(connections *[]Connection) (err error) {
	cfg := config.GetConfig()

	stdout, _, err := utils.ExecBinaryCommand(
		cfg.Apps.Podman.Binary,
		"system connection list --format json",
		false,
		false,
		[]string{},
	)
	if err != nil {
		return err
	}

	data := strings.Join(stdout, "\n")
	jsonBytes := []byte(data)
	err = json.Unmarshal(jsonBytes, &connections)
	if err != nil {
		return err
	}

	return nil
}

func GetConnections(connections *[]Connection) (err error) {
	err = getConnections(connections)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}

func GetDefaultMachineName(connections *[]Connection, machineName *string) (err error) {
	// TODO: what if list of connections is empty?
	// TODO: probably this function need to execute another to update connections
	prefix := "-root"
	for _, conn := range *connections {
		if conn.IsDefault {
			*machineName = strings.ReplaceAll(conn.Name, prefix, "")
			break
		}
	}

	return nil
}

func GetMachineNames(connections *[]Connection) (machineNames []string, err error) {
	// TODO: what if list of connections is empty?
	// TODO: probably this function need to execute another to update connections
	postfix := "-root"
	for _, conn := range *connections {
		if strings.Contains(conn.Name, postfix) {
			machineNames = append(machineNames, strings.ReplaceAll(conn.Name, postfix, ""))
		}
	}

	return machineNames, nil
}
