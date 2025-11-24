package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"sort"
	"strings"
)

func IsBinaryInPath(binary string) bool {
	_, err := exec.LookPath(binary)
	return err == nil
}

func ExecBinaryCommand(binaryName, args string, showOutput, inputRequired bool, envs []string) (stdout, stderr []string, err error) {
	// TODO: is it wise to split this function to 2 different: exec and interactiveExec?
	cmd := exec.Command(binaryName, strings.Split(args, " ")...)
	cmd.Env = append(os.Environ(), envs...)

	if inputRequired {
		cmd.Stdin = os.Stdin
	}

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get stdout pipe: %w", err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get stderr pipe: %w", err)
	}

	err = cmd.Start()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start command: %w", err)
	}

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			stdout = append(stdout, scanner.Text())
			if showOutput {
				fmt.Println(scanner.Text())
			}
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			stderr = append(stderr, scanner.Text())
			if showOutput {
				fmt.Println(scanner.Text())
			}
		}
	}()

	return stdout, stderr, cmd.Wait()
}

func IfStringInSlice(str string, list []string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}

	return false
}

func ExpandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		usr, err := user.Current()
		if err != nil {
			fmt.Println(err)
			return path
		}
		homeDir := usr.HomeDir

		path = strings.Replace(path, "~/", homeDir+"/", 1)
	}
	path = os.ExpandEnv(path)

	return path
}

func IsDirExist(path string) bool {
	path = ExpandPath(path)
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}

	return stat.IsDir()
}

func MapToString(data map[string]any, separator string) (result string) {
	parts := make([]string, 0, len(data))
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := data[key]
		parts = append(parts, fmt.Sprintf("%s%s%v", key, separator, value))
	}

	return strings.Join(parts, " ")
}
