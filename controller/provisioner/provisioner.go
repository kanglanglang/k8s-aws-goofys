package provisioner

import (
	"bytes"
	"net/http"
	"os"
	"text/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/glog"
	"github.com/kubernetes-incubator/external-storage/lib/controller"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// DriverName for installing a Flexvolume.
	DriverName = "skpr/goofys"
	// DriverType for providing developers with more context around this storage option.
	DriverType = "fuse"
	// OptionBucket passed to the Flexvolume for mounting.
	OptionBucket = "Bucket"
	// EnvRegion for operators to declare which region to provision buckets.
	EnvRegion = "AWS_REGION"
	// EnvFormat for provisioning buckets.
	EnvFormat = "GOOFYS_BUCKET_FORMAT"
	// DefaultEnvFormat is a fallback for when a format is not provided.
	DefaultEnvFormat = "goofys-{{ .PVC.ObjectMeta.Namespace }}-{{ .PVC.ObjectMeta.Name }}"
	// DefaultEnvRegion is a fallback for when a region is not provided.
	DefaultEnvRegion = "ap-southeast-2"
)

var _ controller.Provisioner = &goofys{}

type goofys struct {
	region string
	format string
}

// New return a Goofys provisioner.
func New() (controller.Provisioner, error) {
	region := os.Getenv(EnvRegion)
	if region == "" {
		region = DefaultEnvRegion
	}

	format := os.Getenv(EnvFormat)
	if format == "" {
		format = DefaultEnvFormat
	}

	provisioner := &goofys{
		region: region,
		format: format,
	}

	return provisioner, nil
}

// Provision creates a storage asset and returns a PV object representing it.
func (p *goofys) Provision(options controller.VolumeOptions) (*v1.PersistentVolume, error) {
	// This is a consistent naming pattern for provisioning our S3 bucket.
	name, err := formatName(p.format, options)
	if err != nil {
		return nil, err
	}

	glog.Infof("Provisioning bucket: %s", name)

	// Check if the bucket exists, if it does not, create the bucket if it does not exist.
	svc := s3.New(session.New(&aws.Config{Region: aws.String(p.region)}))
	params := &s3.CreateBucketInput{
		Bucket: aws.String(name),
	}
	_, err = svc.CreateBucket(params)
	if err != nil {
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			if reqErr.StatusCode() != http.StatusConflict {
				return nil, err
			}
		}
	}

	glog.Infof("Responding with persistent volume spec: %s", name)

	pv := &v1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: v1.PersistentVolumeSpec{
			PersistentVolumeReclaimPolicy: v1.PersistentVolumeReclaimRetain,
			AccessModes:                   options.PVC.Spec.AccessModes,
			Capacity: v1.ResourceList{
				v1.ResourceName(v1.ResourceStorage): resource.MustParse("8.0E"),
			},
			PersistentVolumeSource: v1.PersistentVolumeSource{
				FlexVolume: &v1.FlexPersistentVolumeSource{
					Driver: DriverName,
					FSType: DriverType,
					Options: map[string]string{
						OptionBucket: name,
					},
				},
			},
		},
	}

	return pv, nil
}

// Delete removes the storage asset that was created by Provision represented
// by the given PV.
// @todo, Tag Bucket as "ready for removal"
// @todo, Tag Bucket with a date to show how old it is.
func (p *goofys) Delete(volume *v1.PersistentVolume) error {
	return nil
}

// Helper function for building hostname.
func formatName(format string, options controller.VolumeOptions) (string, error) {
	var formatted bytes.Buffer

	t := template.Must(template.New("name").Parse(format))

	err := t.Execute(&formatted, options)
	if err != nil {
		return "", err
	}

	return formatted.String(), nil
}
