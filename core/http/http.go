package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	http2 "net/http"

	"github.com/sinohope/sinohope-golang-sdk/common"
	"github.com/sinohope/sinohope-golang-sdk/features"
	"github.com/sirupsen/logrus"
)

type http struct {
	baseUrl string
	signer  features.Signer
}

func NewHTTP(baseUrl string, signer features.Signer) (features.HTTP, error) {
	return &http{
		baseUrl: baseUrl,
		signer:  signer,
	}, nil
}

func (h *http) Post(path string, request interface{}) (*common.Response, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("marshal payload failed, %v", err)
	}
	logrus.
		WithField("path", path).
		WithField("payload", string(payload)).
		Debugf("prepare to post")
	t := timestamp()
	header := make(map[string]string, 4)
	header["Content-Type"] = "application/json"
	// ---------------------------------------------
	// TODO:
	header["publicKey"] = h.signer.PublicKey()
	header["privateKey"] = h.signer.PrivateKey()
	// ---------------------------------------------
	header[common.BizApiKey] = h.signer.PublicKey()
	header[common.BizApiNone] = t
	if signature, err := h.signer.Sign(path, t, string(payload)); err != nil {
		return nil, fmt.Errorf("sign request failed, %v", err)
	} else {
		header[common.BizApiSignature] = signature
	}
	url := fmt.Sprintf("%s%s", h.baseUrl, path)
	if result, err := h.post(url, header, request); err != nil {
		return nil, fmt.Errorf("post request failed, %v", err)
	} else {
		logrus.
			WithField("path", path).
			WithField("payload", string(payload)).
			WithField("response", string(result)).
			Debugf("post success")
		response := &common.Response{}
		if err := json.Unmarshal(result, response); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		} else {
			return response, nil
		}
	}
}

func (h *http) post(url string, headers map[string]string, body interface{}) ([]byte, error) {
	payloadBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	client := &http2.Client{}
	req, err := http2.NewRequest("POST", url, bytes.NewReader(payloadBytes))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
