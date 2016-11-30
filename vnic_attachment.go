package oraclebmc_sdk

import (
	"time"
)

type VnicAttachment struct {
	AvailabilityDomain string
	CompartmentId      string
	DisplayName        string
	Id                 string
	InstanceId         string
	LifecycleState     string
	SubnetId           string
	VnicId             string
	TimeCreated        time.Time
}

func (vnicAttachment *VnicAttachment) getId() string {
	return vnicAttachment.Id
}
func (vnicAttachment *VnicAttachment) getState() string {
	return vnicAttachment.LifecycleState
}

func (vnicAttachment *VnicAttachment) retryCount() int {
	return 50
}

func (vnicAttachment *VnicAttachment) endpoint() string {
	return "vnicAttachments"
}

func (vincAttachment *VnicAttachment) validStates() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"DISABLED",
		"DELETED"}
}
