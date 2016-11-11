package oraclebmc_sdk

import (
	"fmt"
	"io"
	"time"
)

type ResourceInput interface {
	asJSON() io.Reader
}

type Resourceable interface {
	endpoint() string
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

func NewConfig(user string, tenancy string, fingerprint string, signing_key string) *oracle_config {
	core_endpoint := "https://iaas.us-phoenix-1.oraclecloud.com/20160918"
	obj_endpoint := "https://objectstorage.us-phoenix-1.oraclecloud.com"
	endpoint_identity_api := "https://identity.us-phoenix-1.oraclecloud.com/20160918"
	return &oracle_config{
		user:                         user,
		tenancy:                      tenancy,
		fingerprint:                  fingerprint,
		private_key:                  signing_key,
		core_endpoint:                core_endpoint,
		endpoint_blockstorage_api:    core_endpoint,
		endpoint_compute_api:         core_endpoint,
		endpoint_virtual_network_api: core_endpoint,
		endpoint_object_storage_api:  obj_endpoint,
		endpoint_identity_api:        endpoint_identity_api}
}

func (config *oracle_config) getKey() string {
	return fmt.Sprintf("%s/%s/%s", config.tenancy, config.user, config.fingerprint)
}

type OracleResource struct {
	Id             string
	LifecycleState string
	DisplayName    string
	TimeCreated    time.Time
}
