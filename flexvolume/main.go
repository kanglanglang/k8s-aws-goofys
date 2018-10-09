package main

import (
	"os"

	"github.com/previousnext/k8s-aws-goofys/flexvolume/cmd"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New("goofys", "Kubernetes Flexvolume for Goofys (S3 Fuse Filesystem)")

	cmd.Attach(app)
	cmd.Detach(app)
	cmd.GetVolumeName(app)
	cmd.Init(app)
	cmd.IsAttached(app)
	cmd.Mount(app)
	cmd.MountDevice(app)
	cmd.Unmount(app)
	cmd.UnmountDevice(app)
	cmd.WaitForAttach(app)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
