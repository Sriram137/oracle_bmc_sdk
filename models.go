package oraclebmc_sdk

import (
	"bytes"
	"encoding/json"
	"io"
	"time"
)

type Image struct {
	BaseImageId            string
	CompartmentId          string
	CreateImageAllowed     string
	DisplayName            string
	Id                     string
	LifeCycleState         string
	OperatingSystemVersion string
	TimeCreated            time.Time
}

type Instance struct {
	AvailabilityDomain string
	CompartmentId      string
	DisplayName        string
	Id                 string
	ImageId            string
	LifecycleState     string
	Region             string
	Shape              string
	TimeCreated        time.Time
	Metadata           map[string]string
}

type LaunchInstanceInput struct {
	AvailabilityDomain string
	CompartmentId      string
	DisplayName        string
	ImageId            string
	Shape              string
	SubnetId           string
	Metadata           map[string]string
}

func (launchInstanceInput *LaunchInstanceInput) asJSON() (io.Reader) {
	body, _ := json.Marshal(launchInstanceInput)
	return bytes.NewBuffer(body)
}

type oracle_config struct {
	user                         string
	tenancy                      string
	fingerprint                  string
	private_key                  string
	core_endpoint                string
	endpoint_blockstorage_api    string
	endpoint_identity_api        string
	endpoint_compute_api         string
	endpoint_virtual_network_api string
	endpoint_object_storage_api  string
	verify_ssl                   bool
	log_requests                 bool
	additional_user_agent        string
}
