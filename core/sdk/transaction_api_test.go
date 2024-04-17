package sdk

import (
	"testing"

	"github.com/sinohope/sinohope-golang-sdk/common"
	"github.com/sinohope/sinohope-golang-sdk/log"
)

func TestFee(t *testing.T) {
	a := log.App{}
	l := log.Log{}
	log.SetLogDetailsByConfig(a, l)

	m, err := NewTransactionAPI(common.BaseUrl, common.FakePrivateKey)
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

// encoding 3 -> 2
func TestNested2NativeTransfer(t *testing.T) {
	api, err := NewTransactionAPI(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatalf("create new mpc api failed, %v", err)
	}
	param := &common.WalletTransactionSendWAASParam{

		RequestId:   "100005",
		ChainSymbol: "BTC_TEST",
		AssetId:     "BTC_BTC_TEST",
		From:        "2N9Fi5qKJTQDtq8CTGcqAFHRpyGDW5isbc1",
		To:          "tb1q7a0jm6rlv8umckjjxqqnwdpze30w4rwafjj4hj",
		Amount:      "100000",
		VaultId:     "534605276521221",
		WalletId:    "534606724211461",
		ToTag:       "32143",
		Note:        "用户交易信息备注",
		Fee:         "50000",
	}
	res, err := api.CreateTransfer(param)
	if err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(res)
	}
}

// encode 2 -> 4
func TestNative2Taproot(t *testing.T) {
	api, err := NewTransactionAPI(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatalf("create new mpc api failed, %v", err)
	}
	param := &common.WalletTransactionSendWAASParam{

		RequestId:   "100006",
		ChainSymbol: "BTC_TEST",
		AssetId:     "BTC_BTC_TEST",
		From:        "tb1q7a0jm6rlv8umckjjxqqnwdpze30w4rwafjj4hj",
		To:          "tb1pv443quwd6n5mssmauukf0t37r3muxc0hmy45092n6yjxt2afd2hqqy8rze",
		Amount:      "90000",
		VaultId:     "534605276521221",
		WalletId:    "534606724211461",
		ToTag:       "32143",
		Note:        "用户交易信息备注",
		Fee:         "10000",
	}
	res, err := api.CreateTransfer(param)
	if err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(res)
	}
}

func TestLegacy2Native(t *testing.T) {
	api, err := NewTransactionAPI(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatalf("create new mpc api failed, %v", err)
	}
	param := &common.WalletTransactionSendWAASParam{

		RequestId:   "100007",
		ChainSymbol: "BTC_TEST",
		AssetId:     "BTC_BTC_TEST",
		From:        "n1QwsxKD3uZH9mmoCQKdsMbqi81tuD9QR3",
		To:          "tb1q7a0jm6rlv8umckjjxqqnwdpze30w4rwafjj4hj",
		Amount:      "90000",
		VaultId:     "534605276521221",
		WalletId:    "534606724211461",
		ToTag:       "32143",
		Note:        "用户交易信息备注",
	}
	res, err := api.CreateTransfer(param)
	if err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(res)
	}
}
