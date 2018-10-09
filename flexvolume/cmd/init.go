package cmd

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

type cmdInit struct{}

func (cmd *cmdInit) run(c *kingpin.ParseContext) error {
	return respond(Response{
		Status: StatusSuccess,
		Capabilities: &Capabilities{
			Attach:         false,
			SELinuxRelabel: false,
		},
	})
}

// Init declares the "init" subcommand.
func Init(app *kingpin.Application) {
	c := new(cmdInit)

	app.Command("init", "Initializes the driver").Action(c.run)
}
