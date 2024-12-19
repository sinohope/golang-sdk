package sdk

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/sinohope/golang-sdk/common"
	"github.com/sinohope/golang-sdk/log"
	"testing"
)

func TestGenAddress(t *testing.T) {
	a := log.App{}
	l := log.Log{}
	log.SetLogDetailsByConfig(a, l)

	m, err := NewAdvanceAPI(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatalf("create new mpc api failed, %v", err)
	}
	param := &common.WaaSAddressPathParam{
		VaultId:  "534605276521221",
		WalletId: "534606724211461",
		CoinType: 130}
	res, _ := m.GenAddressByPath(param)
	d, _ := json.Marshal(res)
	t.Log(string(d))
	//    advance_api_test.go:25: {"id":"8","address":"0xe7b4d75721cdbbb9775715631864719244c6a5bd","hdPath":"m/1/1/130/0","pubkey":"02482125d4306ec0d6b60a8161b5f3ce6051c7cb7d5bf69e11b2897d1997964863"}
}

func TestSignRawData(t *testing.T) {
	a := log.App{}
	l := log.Log{}
	log.SetLogDetailsByConfig(a, l)

	m, err := NewAdvanceAPI(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatalf("create new mpc api failed, %v", err)
	}
	messageToSign := []byte("hello")
	if err != nil {
		t.Fatal(err)
	}
	tapScriptRoot := sha256.Sum256(messageToSign)
	param := &common.WaaSSignRawDataParam{
		RequestId:     "113",
		VaultId:       "534605276521221",
		WalletId:      "534606724211461",
		HdPath:        "m/1/1/130/0",
		RawData:       "4f2de62c7f2e3393a5b20b85435a76de87dfc1088102d795efb573d3db8ba52a",
		AlgorithmType: 2,
		TapScriptRoot: hex.EncodeToString(tapScriptRoot[:]),
	}
	res, _ := m.SignRawData(param)
	d, _ := json.Marshal(res)
	t.Log(string(d))
	//    advance_api_test.go:25: {"id":"8","address":"0xe7b4d75721cdbbb9775715631864719244c6a5bd","hdPath":"m/1/1/130/0","pubkey":"02482125d4306ec0d6b60a8161b5f3ce6051c7cb7d5bf69e11b2897d1997964863"}
}
