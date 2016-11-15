package oraclebmc_sdk

type ComputeApi struct {
	Config *OracleConfig
}

func (ComputeApi *ComputeApi) WaitForInstance(instance *Instance, state string) error {
	return ComputeApi.waitForState(instance, state)
}

func (ComputeApi *ComputeApi) RefreshInstance(instance *Instance) error {
	return ComputeApi.refresh(instance)
}

func (ComputeApi *ComputeApi) TerminateInstance(instance *Instance) error {
	return ComputeApi.deleteResource(instance)
}

func (ComputeApi *ComputeApi) WaitForImage(image *Image, state string) error {
	return ComputeApi.waitForState(image, state)
}

func (ComputeApi *ComputeApi) RefreshImage(image *Image) error {
	return ComputeApi.refresh(image)
}

func (computeApi *ComputeApi) CreateImage(createImageInput *CreateImageInput) (*Image, error) {
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
