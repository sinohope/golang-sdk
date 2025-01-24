package sdk

import (
	"github.com/sinohope/golang-sdk/common"
	"github.com/sinohope/golang-sdk/log"
	"testing"
)

func TestOpenApiFee(t *testing.T) {
	a := log.App{}
	l := log.Log{}
	log.SetLogDetailsByConfig(a, l)

	m, err := NewOpenAPI(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatalf("create new mpc api failed, %v", err)
	}

	if fee, err := m.Fee(&common.WalletTransactionFeeWAASParam{
		OperationType: "TRANSFER",
		From:          "fc48dc25a8f162cfecc331f96de246591d1b09a5",
		To:            "fc48dc25a8f162cfecc331f96de246591d1b09a5",
		AssetId:       "ETH_SEPOLIA",
		ChainSymbol:   "SEPOLIA",
		Amount:        "1748812",
		InputData:     "0x095ea7b3000000000000000000000000387d311e47e80b498169e6fb51d3193167d89f7d0000000000000000000000000000000000000000000000000de0b6b3a7640000",
	}); err != nil {
		t.Fatalf("%v", err)
	} else {
		t.Logf("%v", fee)
	}
}
