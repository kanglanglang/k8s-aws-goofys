package cmd

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

type cmdMountDevice struct {
	MountDir   string
	DeviceName string
	Options    string
}

func (cmd *cmdMountDevice) run(c *kingpin.ParseContext) error {
	return respond(Response{
		Status: StatusNotSupported,
	})
}

// MountDevice declares the "mountdevice" subcommand.
func MountDevice(app *kingpin.Application) {
	c := new(cmdMountDevice)

	cmd := app.Command("mountdevice", "Mounts a volume’s device to a directory").Action(c.run)

	cmd.Arg("mount-dir", "Directory to mount the device").Required().StringVar(&c.MountDir)
	cmd.Arg("device-name", "Name of the device to mount").Required().StringVar(&c.DeviceName)
	cmd.Arg("options", "Mount options for the device").Required().StringVar(&c.Options)
}
