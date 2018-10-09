package cmd

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func hash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// Helper function execute commands on the commandline.
func shellOut(c []string) error {
	return exec.Command("sh", "-c", strings.Join(c, " ")).Run()
}

func respond(resp Response) error {
	// Format the output as JSON.
	output, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	fmt.Println(string(output))

	return nil
}
