package oraclebmc_sdk

import "time"

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
