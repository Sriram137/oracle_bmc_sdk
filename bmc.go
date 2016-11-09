package oraclebmc_sdk

import "fmt"

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

type ComputeApi struct {
	Config *oracle_config
}

func (computeApi *ComputeApi) GetInstance(instanceId string) (*Instance, error) {
	suffix := fmt.Sprintf("/instances/%s", instanceId)

	var instance Instance
	output := &instance
	orReq := oracleRequest{
		Url:          computeApi.Config.core_endpoint,
		Suffix:       suffix,
		Method:       "GET",
		OracleConfig: computeApi.Config,
		Body:         nil,
		QueryParams:  nil,
		Output:       output}
	err := orReq.doReq()
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (computeApi *ComputeApi) ListImages(compartment_id string) (*[]*Image, error) {
	suffix := "/images"

	var images []*Image
	output := &images

	params := make(map[string]string)
	params["compartmentId"] = compartment_id

	orReq := oracleRequest{
		Url:          computeApi.Config.core_endpoint,
		Suffix:       suffix,
		Method:       "GET",
		OracleConfig: computeApi.Config,
		QueryParams:  params,
		Body:         nil,
		Output:       output}

	err := orReq.doReq()

	if err != nil {
		return nil, err
	}
	return output, nil
}

func (computeApi *ComputeApi) CreateInstance(launchInstanceInput *LaunchInstanceInput) (Instance, error) {
	suffix := "/instances"
	var instance Instance
	output := &instance
	orReq := oracleRequest{
		Url:          computeApi.Config.core_endpoint,
		Suffix:       suffix,
		Method:       "POST",
		OracleConfig: computeApi.Config,
		Body:         launchInstanceInput,
		QueryParams:  nil,
		Output:       output}

	err := orReq.doReq()
	if err != nil {
		return nil, err
	}
	return output, nil
}
