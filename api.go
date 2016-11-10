package oraclebmc_sdk

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/99designs/httpsignatures-go"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type oracleRequest struct {
	Url          string
	Suffix       string
	Method       string
	Body         io.Reader
	Output       interface{}
	OracleConfig *oracle_config
	QueryParams  map[string]string
}

func (orReq *oracleRequest) doReq() error {
	//orReq.Body = bytes.NewBufferString(`{"displayName": "qbol-acc99-cl410-master-1", "metadata": {}, "availabilityDomain": "ncSu:PHX-AD-1", "shape": "BM.DenseIO1.36", "compartmentId": "ocid1.compartment.oc1..aaaaaaaa754zm7nutveaoklas34rorcznxvshimdnrwekqgx5dc6wk5b5tgq", "subnetId": "ocid1.subnet.oc1.phx.aaaaaaaaxkefohaf242whnerssetx2lis5g5paikorm4xb4zr4cxt2m6iiaq", "imageId": "ocid1.image.oc1.phx.aaaaaaaapkuvkljm6ku254ykuaufb36hyz7oqohvryjq4dhiodrhfi5ve6xq"}`)
	req, err := http.NewRequest(orReq.Method, orReq.Url+orReq.Suffix, orReq.Body)
	if err != nil {
		return err
	}
	if orReq.QueryParams != nil {
		url := req.URL
		q := url.Query()
		for key, val := range orReq.QueryParams {
			q.Set(key, val)
		}
		url.RawQuery = q.Encode()

	}
	orReq.inject_headers(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(orReq.Output)
	return nil
}

func (orReq *oracleRequest) inject_headers(request *http.Request) {
	oracleConfig := orReq.OracleConfig
	var required_headers []string
	request.Header.Set("host", request.URL.Host)
	if request.Method == "POST" || request.Method == "PUT" {

		required_headers = []string{
			"date",
			httpsignatures.RequestTarget,
			"host",
			"content-length",
			"content-type",
			"x-content-sha256"}

		body, _ := ioutil.ReadAll(request.Body)
		hash := sha256.New()
		hash.Write(body)
		content_body := hash.Sum(nil)

		request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		request.Header.Set("x-content-sha256", base64.StdEncoding.EncodeToString([]byte((content_body))))
		request.Header.Set("content-type", "application/json")
		request.Header.Set("content-length", strconv.Itoa(len(string(body))))
	} else {
		required_headers = []string{httpsignatures.RequestTarget, "date", "host"}
	}
	signer := httpsignatures.NewSigner(httpsignatures.AlgorithmRsaSha256, required_headers...)
	signer.AuthRequest(oracleConfig.getKey(), oracleConfig.private_key, request)
}
