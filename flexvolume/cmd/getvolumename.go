package cmd

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

type cmdGetVolumeName struct {
	Options string
}

func (cmd *cmdGetVolumeName) run(c *kingpin.ParseContext) error {
	return respond(Response{
		Status: StatusNotSupported,
	})
}

// GetVolumeName declares the "getvolumename" subcommand.
func GetVolumeName(app *kingpin.Application) {
	c := new(cmdGetVolumeName)

	cmd := app.Command("getvolumename", "Get the unique name of the volume").Action(c.run)

	cmd.Arg("options", "Volume name options").Required().StringVar(&c.Options)
}
