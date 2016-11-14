package oraclebmc_sdk

import (
	"bytes"
	"encoding/json"
	"io"
	"time"
)

type CreateImageInput struct {
	CompartmentId string `json:"compartmentId"`
	DisplayName   string `json:"displayName"`
	InstanceId    string `json:"instanceId"`
}

type Image struct {
	BaseImageId            string
	CompartmentId          string
	CreateImageAllowed     bool
	DisplayName            string
	Id                     string
	LifecycleState         string
	operatingSystem        string
	OperatingSystemVersion string
	TimeCreated            time.Time
}

func (createImageInput *CreateImageInput) asJSON() io.Reader {
	body, _ := json.Marshal(createImageInput)
	return bytes.NewBuffer(body)
}

func (image *Image) getId() string {
	return image.Id
}
func (image *Image) getState() string {
	return image.LifecycleState
}

func (image *Image) retryCount() int {
	return 50
}

func (image *Image) endpoint() string {
	return "images"
}

func (image *Image) validStates() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"DISABLED",
		"DELETED"}
}
