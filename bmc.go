package oraclebmc_sdk

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/99designs/httpsignatures-go"
	"io"
	"io/ioutil"
	"net/http"
)

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

type ComputeApi struct {
	Config *oracle_config
}

type oracleRequest struct {
	Url          string
	Suffix       string
	Method       string
	Output       interface{}
	OracleConfig *oracle_config
	Params       map[string]string
}

func (orReq *oracleRequest) doReq() error {
	var body io.Reader
	if orReq.Params != nil && len(orReq.Params) > 0 {
		val, err := json.Marshal(orReq.Params)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(val)
	} else {
		body = nil
	}
	req, err := http.NewRequest(orReq.Method, orReq.Url+orReq.Suffix, body)
	if body == nil {
		url := req.URL
		q := url.Query()
		for key, val := range orReq.Params {
			q.Set(key, val)
		}
		url.RawQuery = q.Encode()

	}
	if err != nil {
		return err
	}
	inject_headers(orReq.OracleConfig, req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(orReq.Output)
	return nil
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
		Params:       nil,
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
		Params:       nil,
		Output:       output}

	err := orReq.doReq()

	if err != nil {
		return nil, err
	}
	return output, nil
}

func inject_headers(oracleConfig *oracle_config, request *http.Request) {
	var required_headers []string
	request.Header.Set("host", request.URL.Host)
	if request.Method == "POST" || request.Method == "PUT" {

		required_headers = []string{
			httpsignatures.RequestTarget,
			"date",
			"host",
			"x-content-sha256",
			"content-type",
			"content-length"}

		body, _ := ioutil.ReadAll(request.Body)
		hash := sha256.New()
		hash.Write([]byte(body))
		content_body := hash.Sum(nil)

		request.Header.Set("x-content-sha256", string(content_body))
		request.Header.Set("content-length", string(len(body)))
		request.Header.Set("content-type", "application/json")
	} else {
		required_headers = []string{httpsignatures.RequestTarget, "date", "host"}
	}
	signer := httpsignatures.NewSigner(httpsignatures.AlgorithmRsaSha256, required_headers...)
	signer.AuthRequest(oracleConfig.getKey(), oracleConfig.private_key, request)
}
