package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	"github.com/kubernetes-incubator/external-storage/lib/controller"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/previousnext/k8s-aws-goofys/controller/provisioner"
)

const apiVersion = "storage.skpr.io/goofys"

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// The controller needs to know what the server version is because out-of-tree
	// provisioners aren't officially supported until 1.5
	serverVersion, err := clientset.Discovery().ServerVersion()
	if err != nil {
		panic(err)
	}

	// Create the provisioner: it implements the Provisioner interface expected by the controller.
	provisioner, err := provisioner.New()
	if err != nil {
		panic(err)
	}

	glog.Infof("Running provisioner: %s", apiVersion)

	pc := controller.NewProvisionController(
		clientset,
		apiVersion,
		provisioner,
		serverVersion.GitVersion,
		controller.CreateProvisionedPVInterval(time.Minute*10),
		controller.LeaseDuration(time.Minute*10),
	)

	pc.Run(wait.NeverStop)
}
