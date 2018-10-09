package cmd

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

type cmdUnmountDevice struct {
	MountDir string
}

func (cmd *cmdUnmountDevice) run(c *kingpin.ParseContext) error {
	return respond(Response{
		Status: StatusNotSupported,
	})
}

// UnmountDevice declares the "unmountdevice" subcommand.
func UnmountDevice(app *kingpin.Application) {
	c := new(cmdUnmountDevice)

	cmd := app.Command("unmountdevice", "Unmounts a volume’s device from a directory").Action(c.run)

	cmd.Arg("mount-dir", "Directory to unmount").Required().StringVar(&c.MountDir)
}
