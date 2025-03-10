package sdk

import (
	"testing"

	"github.com/sinohope/golang-sdk/common"
	"github.com/sinohope/golang-sdk/log"
)

func TestMPCAPI(t *testing.T) {
	a := log.App{}
	l := log.Log{}
	log.SetLogDetailsByConfig(a, l)

	m, err := NewAccountAndAddressAPI(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatalf("create new mpc api failed, %v", err)
	}

	var walletInfo []*common.WaaSWalletInfoData
	if walletInfo, err = m.CreateWallets(&common.WaaSCreateBatchWalletParam{
		VaultId:   "456517693645317",
		RequestId: "fake-request-id",
	}); err != nil {
		t.Fatalf("create wallets failed, %v", err)
	} else {
		t.Logf("create wallets success, %v", walletInfo)
	}
}

func TestGenChainAddress(t *testing.T) {
	a := log.App{}
	l := log.Log{}
	log.SetLogDetailsByConfig(a, l)

	m, err := NewAccountAndAddressAPI(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatalf("create new mpc api failed, %v", err)
	}
	param := &common.WaaSGenerateChainAddressParam{
		RequestId:   "1005",
		VaultId:     "534605276521221",
		WalletId:    "534606724211461",
		Count:       1,
		ChainSymbol: "SIGNET",
		Encoding:    2,
	}
	res, err := m.GenerateChainAddresses(param)
	if err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(res)
	}

}
