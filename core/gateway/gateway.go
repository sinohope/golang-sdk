package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	http2 "net/http"
	"regexp"
	"time"

	"github.com/sinohope/golang-sdk/common"
	"github.com/sinohope/golang-sdk/features"
	"github.com/sirupsen/logrus"
)

type gateway struct {
	baseUrl string
	s       *Signer
}

func NewGateway(baseUrl, private string) (features.Gateway, error) {
	s, err := NewSigner(private)
	if err != nil {
		return nil, fmt.Errorf("create new signer failed, %v", err)
	}
	return &gateway{
		baseUrl: baseUrl,
		s:       s,
	}, nil
}

func (g *gateway) Post(path string, request interface{}) (*common.Response, error) {
	var err error
	var payload []byte
	if request != nil {
		payload, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("marshal payload failed, %v", err)
		}
	}

	replaceReg := regexp.MustCompile("\\s*|\n|\r|\t")
	signPayload := replaceReg.ReplaceAllLiteralString(string(payload), "")

	logrus.
		WithField("path", path).
		WithField("payload", signPayload).
		Debugf("prepare to post")
	t := timestamp()
	header := make(map[string]string, 4)
	header["Content-Type"] = "application/json"
	header[common.BizApiKey] = g.s.PublicKey()
	header[common.BizApiNone] = t
	if signature, err := g.s.Sign(path, t, signPayload); err != nil {
		return nil, fmt.Errorf("sign request failed, %v", err)
	} else {
		header[common.BizApiSignature] = signature
	}
	url := fmt.Sprintf("%s%s", g.baseUrl, path)
	if result, err := doPost(url, header, request); err != nil {
		return nil, fmt.Errorf("post request failed, %v", err)
	} else {
		logrus.
			WithField("path", path).
			WithField("payload", signPayload).
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

func doPost(url string, headers map[string]string, body interface{}) ([]byte, error) {
	var err error
	var payloadBytes []byte
	if body != nil {
		payloadBytes, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
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

func timestamp() string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(loc)
	return fmt.Sprintf("%d", now.UnixNano()/int64(time.Millisecond))
}
