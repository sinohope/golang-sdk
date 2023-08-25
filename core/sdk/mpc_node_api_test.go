package sdk

import (
	"github.com/sinohope/sinohope-golang-sdk/common"
	"github.com/sinohope/sinohope-golang-sdk/core/http"
	"github.com/sinohope/sinohope-golang-sdk/core/signer"
	"github.com/sinohope/sinohope-golang-sdk/log"
	"testing"
)

func TestMPCNodeAPI(t *testing.T) {
	a := log.App{}
	l := log.Log{}
	log.SetLogDetailsByConfig(a, l)

	s, err := signer.NewSigner(common.FakePrivateKey)
	if err != nil {
		t.Fatalf("create new signer failed, %v", err)
	}
	p, err := http.NewHTTP(common.BaseUrl, s)
	if err != nil {
		t.Fatalf("create new http failed, %v", err)
	}
	m, err := NewMPCNodeAPI(p)
	if err != nil {
		t.Fatalf("create new mpc node api failed, %v", err)
	}
	request := &common.WaasMpcNodeExecRecordParam{
		BusinessExecType:   1,
		BusinessExecStatus: 10,
		SinoId:             "fake sino id",
		PageIndex:          0,
		PageSize:           40,
	}
	var result *common.PageData
	if result, err = m.ListMPCRequests(request); err != nil {
		t.Fatalf("list mpc requests failed, %v", err)
	} else {
		t.Logf("list mpc requests success, %v", result)
	}

	var status *common.WaaSMpcNodeStatusDTOData
	if status, err = m.Status(); err != nil {
		t.Fatalf("get mpc node status failed, %v", err)
	} else {
		t.Logf("get mpc node status success, %v", status)
	}
}
