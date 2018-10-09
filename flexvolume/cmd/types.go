package cmd

// Status returned action each action.
type Status string

const (
	// StatusSuccess is returned if the action was a success.
	StatusSuccess Status = "Success"
	// StatusFailure is returned if the action was a failure.
	StatusFailure Status = "Failure"
	// StatusNotSupported is returned if the action is not supported.
	StatusNotSupported Status = "Not Supported"
)

// Response returned to the Kubelet after each action.
type Response struct {
	Status       Status        `json:"status"`
	Message      string        `json:"message,omitempty"`
	DevicePath   string        `json:"device,omitempty"`
	VolumeName   string        `json:"volumeName,omitempty"`
	Attached     bool          `json:"attached,omitempty"`
	Capabilities *Capabilities `json:",omitempty"`
}

// Capabilities of this flexvolume.
type Capabilities struct {
	Attach         bool `json:"attach"`
	SELinuxRelabel bool `json:"selinuxRelabel"`
}
