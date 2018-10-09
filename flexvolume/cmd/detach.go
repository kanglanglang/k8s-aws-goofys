package cmd

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

type cmdDetach struct {
	VolumeName string
	NodeName   string
}

func (cmd *cmdDetach) run(c *kingpin.ParseContext) error {
	return respond(Response{
		Status: StatusNotSupported,
	})
}

// Detach declares the "detach" subcommand.
func Detach(app *kingpin.Application) {
	c := new(cmdDetach)

	cmd := app.Command("detach", "Detach the device from the node").Action(c.run)

	cmd.Arg("volume-name", "Name of the volume to detach").Required().StringVar(&c.VolumeName)
	cmd.Arg("node-name", "Name of the node which to attach the volume").Required().StringVar(&c.NodeName)
}
