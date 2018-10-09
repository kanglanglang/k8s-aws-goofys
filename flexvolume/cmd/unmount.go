package cmd

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

type cmdUnmount struct {
	MountDir string
}

func (cmd *cmdUnmount) run(c *kingpin.ParseContext) error {
	// Unmount the filesystem.
	err := shellOut([]string{"fusermount", "-u", cmd.MountDir})
	if err != nil {
		return respond(Response{
			Status:  StatusFailure,
			Message: fmt.Sprintf("Failed to unmount Goofys volume: %s", err),
		})
	}

	return respond(Response{
		Status:  StatusSuccess,
		Message: "Unmounted Goofys volume",
	})
}

// Unmount declares the "unmount" subcommand.
func Unmount(app *kingpin.Application) {
	c := new(cmdUnmount)

	cmd := app.Command("unmount", "Unmounts a volume from a directory").Action(c.run)

	cmd.Arg("mount-dir", "Directory to unmount").Required().StringVar(&c.MountDir)
}
