package oraclebmc_sdk

import (
	"time"
)

type Vnic struct {
	AvailabilityDomain string
	CompartmentId      string
	DisplayName        string
	Id                 string
	InstanceId         string
	LifecycleState     string
	PrivateIp          string
	PublicIp           string
	SubnetId           string
	TimeCreated        time.Time
}

func (vnic *Vnic) getId() string {
	return vnic.Id
}
func (vnic *Vnic) getState() string {
	return vnic.LifecycleState
}

func (vnic *Vnic) retryCount() int {
	return 50
}

func (vnic *Vnic) endpoint() string {
	return "vnics"
}

func (vnic *Vnic) validStates() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"DISABLED",
		"DELETED"}
}
