#!/usr/bin/make -f

export CGO_ENABLED=0

PROJECT=github.com/previousnext/k8s-aws-goofys

# Builds the project.
build:
	gox -os='linux darwin' -arch='amd64' -output='bin/flexvolume_{{.OS}}_{{.Arch}}' $(PROJECT)/flexvolume
	gox -os='linux darwin' -arch='amd64' -output='bin/controller_{{.OS}}_{{.Arch}}' $(PROJECT)/controller

# Run all lint checking with exit codes for CI.
lint:
	golint -set_exit_status `go list ./... | grep -v /vendor/`

.PHONY: build lint
