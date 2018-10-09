package cmd

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

type cmdWaitForAttach struct {
	DeviceName string
	Options    string
}

func (cmd *cmdWaitForAttach) run(c *kingpin.ParseContext) error {
	return respond(Response{
		Status: StatusNotSupported,
	})
}

// WaitForAttach declares the "waitforattach" subcommand.
func WaitForAttach(app *kingpin.Application) {
	c := new(cmdWaitForAttach)

	cmd := app.Command("waitforattach", "Waits until a volume is fully attached to a node and its device emerges").Action(c.run)

	cmd.Arg("device-name", "Name of the device to wait for").Required().StringVar(&c.DeviceName)
	cmd.Arg("options", "Attach options for the device").Required().StringVar(&c.Options)
}
