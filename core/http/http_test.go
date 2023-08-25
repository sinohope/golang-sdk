package http

import (
	"github.com/sinohope/sinohope-golang-sdk/common"
	"github.com/sinohope/sinohope-golang-sdk/core/signer"
	"github.com/sinohope/sinohope-golang-sdk/log"
	"testing"
)

func TestPost(t *testing.T) {
	a := log.App{
		Name: "sinohope-golang-sdk",
	}
	l := log.Log{
		Stdout: struct {
			Enable bool `toml:"enable"`
			Level  int  `toml:"level"`
		}{
			Enable: true,
			Level:  5,
		},
	}
	log.SetLogDetailsByConfig(a, l)

	s, err := signer.NewSigner(common.FakePrivateKey)
	if err != nil {
		t.Fatalf("create new signer failed, %v", err)
	}
	p, err := NewHTTP(common.BaseUrl, s)
	if err != nil {
		t.Fatalf("create new http failed, %v", err)
	}
	path := "/v1/waas/mpc/mpcnode/list_mpc_requests"
	request := &common.ListMPCRequestsParam{
		BusinessExecType:   1,
		BusinessExecStatus: 10,
		SinoID:             "fake sino id",
		PageIndex:          0,
		PageSize:           40,
	}
	if response, err := p.Post(path, request); err != nil {
		t.Fatalf("post failed, %v", err)
	} else {
		t.Logf("response, code: %v, message: %v, success: %v, data: %v",
			response.Code, response.Message, response.Success,
			string(response.Data))
	}
}
