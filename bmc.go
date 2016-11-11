package oraclebmc_sdk

import (
	"fmt"
)

type ComputeApi struct {
	Config *oracle_config
}

func (computeApi *ComputeApi) get(id string, resourceable Resourceable) error {
	suffix := fmt.Sprintf("/%s/%s", resourceable.endpoint(), id)

	orReq := oracleRequest{
		Url:          computeApi.Config.core_endpoint,
		Suffix:       suffix,
		Method:       "GET",
		OracleConfig: computeApi.Config,
		Body:         nil,
		QueryParams:  nil,
		Output:       resourceable}
	err := orReq.doReq()
	if err != nil {
		return err
	}
	return nil
}

func (computeApi *ComputeApi) createResource(resourceInput ResourceInput, resourceable Resourceable) error {
	suffix := resourceable.endpoint()
	var instance Instance
	output := &instance
	orReq := oracleRequest{
		Url:          computeApi.Config.core_endpoint,
		Suffix:       suffix,
		Method:       "POST",
		OracleConfig: computeApi.Config,
		Body:         resourceInput.asJSON(),
		QueryParams:  nil,
		Output:       output}

	err := orReq.doReq()
	if err != nil {
		return nil, err
	}
	return nil
}

func (computeApi *ComputeApi) GetImage(imageId string) (*Image, error) {
	var image Image
	output := &image
	err := computeApi.get(imageId, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (computeApi *ComputeApi) GetInstance(instanceId string) (*Instance, error) {
	var instance Instance
	output := &instance
	err := computeApi.get(instanceId, output)
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

func (computeApi *ComputeApi) CreateImage(createImageInput *CreateImageInput) (*Instance, error) {
	var image Image
	output := &image
	err := computeApi.createResource(createImageInput, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (computeApi *ComputeApi) CreateInstance(launchInstanceInput *LaunchInstanceInput) (*Instance, error) {
	var instance Instance
	output := &instance
	err := computeApi.createResource(launchInstanceInput, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
