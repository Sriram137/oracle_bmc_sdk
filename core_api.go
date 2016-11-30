package oraclebmc_sdk

import (
	"errors"
	"fmt"
	"time"
)

func (ComputeApi *ComputeApi) refresh(resourceable Resourceable) error {
	return ComputeApi.get(resourceable.getId(), resourceable)
}

func (ComputeApi *ComputeApi) waitForState(resourceable Resourceable, state string) error {
	err := errors.New("State is invalid")
	for _, valid_state := range resourceable.validStates() {
		if valid_state == state {
			err = nil
			break
		}
	}
	if err != nil {
		return err
	}

	interval := time.Duration(60)
	retries := resourceable.retryCount()
	for ; retries > 0; retries-- {
		ComputeApi.refresh(resourceable)
		if resourceable.getState() == state {
			return nil
		}
		time.Sleep(interval * time.Second)

	}
	return errors.New("Time Out expired")
}
func (computeApi *ComputeApi) get(id string, resourceable Resourceable) error {
	if id == "" {
		err := errors.New("Id cannot be empty")
		return err
	}
	suffix := fmt.Sprintf("%s/%s", resourceable.endpoint(), id)

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
	orReq := oracleRequest{
		Url:          computeApi.Config.core_endpoint,
		Suffix:       suffix,
		Method:       "POST",
		OracleConfig: computeApi.Config,
		Body:         resourceInput.asJSON(),
		QueryParams:  nil,
		Output:       resourceable}

	err := orReq.doReq()
	if err != nil {
		return err
	}
	return nil
}

func (computeApi *ComputeApi) deleteResource(resourceable Resourceable) error {
	suffix := fmt.Sprintf("%s/%s", resourceable.endpoint(), resourceable.getId())
	orReq := oracleRequest{
		Url:          computeApi.Config.core_endpoint,
		Suffix:       suffix,
		Method:       "DELETE",
		OracleConfig: computeApi.Config,
		Body:         nil,
		QueryParams:  nil,
		Output:       nil}

	err := orReq.doReq()
	if err != nil {
		return err
	}
	return nil
}

func (computeApi *ComputeApi) ListVnicAttachments(compartment_id string, instanceId string) (*[]*VnicAttachment, error) {
	suffix := "/vnicAttachments"

	var vnicAttachments []*VnicAttachment
	output := &vnicAttachments

	params := make(map[string]string)
	params["compartmentId"] = compartment_id
	params["instanceId"] = instanceId

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
