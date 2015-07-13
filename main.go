package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	gems, err := getGems()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, gem := range gems {
		cmd := exec.Command("gem", "uninstall", "-axI", gem)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}

func getGems() ([]string, error) {
	out, err := exec.Command("gem", "list", "--no-versions").Output()

	if err != nil {
		return nil, err
	}

	result := string(out)
	tokens := strings.Split(result, "\n")

	gems := []string{}
	for _, token := range tokens {
		if token == "" || strings.HasPrefix(token, "*") {
			continue
		}
		gems = append(gems, token)
	}

	return gems, nil
}
