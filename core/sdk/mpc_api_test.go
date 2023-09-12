package sdk

import (
	"github.com/sinohope/sinohope-golang-sdk/common"
	"github.com/sinohope/sinohope-golang-sdk/core/http"
	"github.com/sinohope/sinohope-golang-sdk/core/signer"
	"github.com/sinohope/sinohope-golang-sdk/log"
	"testing"
)

func TestMPCAPI(t *testing.T) {
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

	m, err := NewMPCAPI(p)
	if err != nil {
		t.Fatalf("create new mpc api failed, %v", err)
	}

	var walletInfo []*common.WaasWalletInfoData
	if walletInfo, err = m.CreateWallets(&common.WaaSCreateBatchWalletParam{
		VaultId:   "456517693645317",
		RequestId: "fake-request-id",
	}); err != nil {
		t.Fatalf("create wallets failed, %v", err)
	} else {
		t.Logf("create wallets success, %v", walletInfo)
	}
}
