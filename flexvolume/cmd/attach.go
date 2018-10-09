package cmd

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

type cmdAttach struct {
	Options  string
	NodeName string
}

func (cmd *cmdAttach) run(c *kingpin.ParseContext) error {
	return respond(Response{
		Status: StatusNotSupported,
	})
}

// Attach declares the "attach" subcommand.
func Attach(app *kingpin.Application) {
	c := new(cmdAttach)

	cmd := app.Command("attach", "Attach a device to a node").Action(c.run)

	cmd.Arg("options", "Attach options for the device").Required().StringVar(&c.Options)
	cmd.Arg("node-name", "Name of the node which to attach the device").Required().StringVar(&c.NodeName)
}
