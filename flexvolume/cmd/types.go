package cmd

type Status string

const (
	StatusSuccess      Status = "Success"
	StatusFailure      Status = "Failure"
	StatusNotSupported Status = "Not Supported"
)

type Response struct {
	Status       Status        `json:"status"`
	Message      string        `json:"message,omitempty"`
	DevicePath   string        `json:"device,omitempty"`
	VolumeName   string        `json:"volumeName,omitempty"`
	Attached     bool          `json:"attached,omitempty"`
	Capabilities *Capabilities `json:",omitempty"`
}

type Capabilities struct {
	Attach         bool `json:"attach"`
	SELinuxRelabel bool `json:"selinuxRelabel"`
}
