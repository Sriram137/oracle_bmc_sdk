package oraclebmc_sdk

import (
	"bytes"
	"encoding/json"
	"io"
)

type Instance struct {
	OracleResource
	AvailabilityDomain string
	CompartmentId      string
	ImageId            string
	Region             string
	Shape              string
	Metadata           map[string]string
}

type LaunchInstanceInput struct {
	AvailabilityDomain string            `json:"availabilityDomain"`
	CompartmentId      string            `json:"compartmentId"`
	DisplayName        string            `json:"displayName"`
	ImageId            string            `json:"imageId"`
	Shape              string            `json:"shape"`
	SubnetId           string            `json:"subnetId"`
	Metadata           map[string]string `json:"metadata"`
}

func (launchInstanceInput *LaunchInstanceInput) asJSON() io.Reader {
	body, _ := json.Marshal(launchInstanceInput)
	return bytes.NewBuffer(body)
}

func (instance *Instance) endpoint() string {
	return "instances"
}

func (instance *Instance) validStates() []string {
	return []string{
		"PROVISIONING",
		"RUNNING",
		"STARTING",
		"STOPPING",
		"STOPPED",
		"CREATING_IMAGE",
		"TERMINATING",
		"TERMINATED"}
}
