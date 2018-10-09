package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

type cmdMount struct {
	MountDir string
	Options  string
}

func (cmd *cmdMount) run(c *kingpin.ParseContext) error {
	var options map[string]string

	if err := json.Unmarshal([]byte(cmd.Options), &options); err != nil {
		return respond(Response{
			Status:  StatusFailure,
			Message: fmt.Sprintf("Failed unmarshal options: %s", err),
		})
	}

	if _, ok := options[OptionBucket]; !ok {
		return respond(Response{
			Status:  StatusFailure,
			Message: fmt.Sprintf("Not found: %s", OptionBucket),
		})
	}

	if _, ok := options[OptionUID]; !ok {
		return respond(Response{
			Status:  StatusFailure,
			Message: fmt.Sprintf("Not found: %s", OptionUID),
		})
	}

	if _, ok := options[OptionGID]; !ok {
		return respond(Response{
			Status:  StatusFailure,
			Message: fmt.Sprintf("Not found: %s", OptionGID),
		})
	}

	// Before we perform the mount we need to make sure the directory exists.
	err := os.MkdirAll(cmd.MountDir, os.ModePerm)
	if err != nil {
		return respond(Response{
			Status:  StatusFailure,
			Message: fmt.Sprintf("Failed create directory for Goofys volume: %s", err),
		})
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return respond(Response{
			Status:  StatusFailure,
			Message: fmt.Sprintf("Failed find directory of Flexvolume: %s", err),
		})
	}

	mount := []string{
		path.Join(dir, "mount"),
		"--uid", options[OptionUID],
		"--gid", options[OptionGID],
		"-o", "allow_other",
		options[OptionBucket],
		cmd.MountDir,
	}

	err = shellOut(mount)
	if err != nil {
		return respond(Response{
			Status:  StatusFailure,
			Message: fmt.Sprintf("Failed to mount Goofys volume: %s: %s", strings.Join(mount, " "), err),
		})
	}

	return respond(Response{
		Status:  StatusSuccess,
		Message: "Mounted Goofys volume",
	})
}

// Mount declares the "mount" subcommand.
func Mount(app *kingpin.Application) {
	c := new(cmdMount)

	cmd := app.Command("mount", "Mounts a volume to directory").Action(c.run)

	cmd.Arg("mount-dir", "Directory to mount the device").Required().StringVar(&c.MountDir)
	cmd.Arg("options", "Mount options for the device").Required().StringVar(&c.Options)
}
