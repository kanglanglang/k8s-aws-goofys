package cmd

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

type cmdIsAttached struct {
	Options  string
	NodeName string
}

func (cmd *cmdIsAttached) run(c *kingpin.ParseContext) error {
	return respond(Response{
		Status: StatusNotSupported,
	})
}

// IsAttached declares the "isattached" subcommand.
func IsAttached(app *kingpin.Application) {
	c := new(cmdIsAttached)

	cmd := app.Command("isattached", "Checks that a volume is attached to a node").Action(c.run)

	cmd.Arg("options", "Is attached options for the device").Required().StringVar(&c.Options)
	cmd.Arg("node-name", "Name of the node which the device is attached").Required().StringVar(&c.NodeName)
}
