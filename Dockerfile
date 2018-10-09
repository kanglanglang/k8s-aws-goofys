FROM golang:1.11 as build
RUN go get github.com/mitchellh/gox
ADD . /go/src/github.com/previousnext/k8s-aws-goofys
WORKDIR /go/src/github.com/previousnext/k8s-aws-goofys
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
# Default executable for provisioning PersistentVolumeClaims.
COPY --from=build /go/src/github.com/previousnext/k8s-aws-goofys/bin/controller_linux_amd64 /controller
# Components for installing on a Node:
# * deploy.sh = Will sync
# * flexvolume = Kubernetes Flexvolume plugin
# * goofys = https://github.com/kahing/goofys
ADD deploy.sh /deploy.sh
RUN chmod +x /deploy.sh
COPY --from=build /go/src/github.com/previousnext/k8s-aws-goofys/bin/flexvolume_linux_amd64 /flexvolume
ADD https://github.com/kahing/goofys/releases/download/v0.19.0/goofys /goofys
CMD ["/controller"]
