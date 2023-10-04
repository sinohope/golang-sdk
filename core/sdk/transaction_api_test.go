package sdk

import (
	"testing"

	"github.com/sinohope/sinohope-golang-sdk/common"
	"github.com/sinohope/sinohope-golang-sdk/log"
)

func TestCreateTransfer(t *testing.T) {
	a := log.App{}
	l := log.Log{}
	log.SetLogDetailsByConfig(a, l)

	m, err := NewTransactionAPI(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatalf("create new mpc api failed, %v", err)
	}

	var walletInfo *common.CreateSettlementTxResData
	if walletInfo, err = m.CreateTransaction(&common.WalletTransactionSendDataWAASParam{
		VaultId:   "456517693645317",
		RequestId: "fake-request-id",
	}); err != nil {
		t.Fatalf("create wallets failed, %v", err)
	} else {
		t.Logf("create wallets success, %v", walletInfo)
	}
}
