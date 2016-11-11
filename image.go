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
	Id                     string
	LifecycleState         string
	DisplayName            string
	TimeCreated            time.Time
	BaseImageId            string
	CompartmentId          string
	CreateImageAllowed     string
	OperatingSystemVersion string
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
