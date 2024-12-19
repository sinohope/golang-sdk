package sdk

import (
	"encoding/json"
	"github.com/sinohope/golang-sdk/common"
	"testing"
)

func TestInscribeDeploy(t *testing.T) {
	c, _ := NewBrc20API(common.BaseUrl, common.FakePrivateKey)
	p := &common.InscribeDeployParam{
		RequestId:    "1",
		Ticker:       "joy1",
		TotalSupply:  "10000",
		LimitPerMint: "100",
		Decimals:     "6",
		ChainSymbol:  "BTC_TEST",
		From:         "tb1q7a0jm6rlv8umckjjxqqnwdpze30w4rwafjj4hj",
	}

	if res, err := c.InscribeDeploy(p); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}

func TestInscribeMint(t *testing.T) {
	c, _ := NewBrc20API(common.BaseUrl, common.FakePrivateKey)
	p := &common.InscribeMintParam{
		RequestId:   "26",
		Ticker:      "joy1",
		Amount:      "100",
		ChainSymbol: "BTC_TEST",
		From:        "tb1q7a0jm6rlv8umckjjxqqnwdpze30w4rwafjj4hj",
	}
	if res, err := c.InscribeMint(p); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}
func TestInscribeTransfer(t *testing.T) {
	c, _ := NewBrc20API(common.BaseUrl, common.FakePrivateKey)
	p := &common.InscribeTransferParam{
		RequestId:   "3",
		Ticker:      "joy1",
		Amount:      "10",
		ChainSymbol: "BTC_TEST",
		From:        "tb1q7a0jm6rlv8umckjjxqqnwdpze30w4rwafjj4hj",
	}
	if res, err := c.InscribeTransfer(p); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}

func TestInscribeTransferById(t *testing.T) {
	c, _ := NewBrc20API(common.BaseUrl, common.FakePrivateKey)
	p := &common.TransferByIdParam{
		RequestId:     "4",
		ChainSymbol:   "BTC_TEST",
		From:          "tb1q7a0jm6rlv8umckjjxqqnwdpze30w4rwafjj4hj",
		To:            "2N9Fi5qKJTQDtq8CTGcqAFHRpyGDW5isbc1",
		InscriptionId: "762972c705b604b64a6665faf5ef84ca0a4e3ebb46ba8ad495b8fb569d79187ei0",
		Ticker:        "joy1",
	}
	if res, err := c.InscribeTransferById(p); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}

func TestOneStopTransfer(t *testing.T) {
	c, _ := NewBrc20API(common.BaseUrl, common.FakePrivateKey)
	p := &common.OneStopTransferParam{
		RequestId:   "5",
		Ticker:      "joy1",
		Amount:      "10",
		ChainSymbol: "BTC_TEST",
		From:        "tb1q7a0jm6rlv8umckjjxqqnwdpze30w4rwafjj4hj",
		To:          "2N9Fi5qKJTQDtq8CTGcqAFHRpyGDW5isbc1",
	}
	if res, err := c.OneStopTransfer(p); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}

func TestQueryInscribeTransfers(t *testing.T) {
	c, _ := NewBrc20API(common.BaseUrl, common.FakePrivateKey)
	p := &common.WaasBrc20QueryInscribeTransferReq{
		Ticker:      "joy1",
		ChainSymbol: "BTC_TEST",
		Address:     "tb1q7a0jm6rlv8umckjjxqqnwdpze30w4rwafjj4hj",
	}
	if res, err := c.QueryInscribeTransfers(p); err != nil {
		t.Error(err)
	} else {
		b, _ := json.Marshal(res)
		t.Log(string(b))
	}
}

func TestQueryPageBalanceSummary(t *testing.T) {
	c, _ := NewBrc20API(common.BaseUrl, common.FakePrivateKey)
	p := &common.WaasBrc20PageQueryBalanceSummaryReq{
		ChainSymbol: "BTC_TEST",
		Address:     "tb1q7a0jm6rlv8umckjjxqqnwdpze30w4rwafjj4hj",
		Start:       0,
		Limit:       10,
	}
	if res, err := c.QueryPageBalanceSummary(p); err != nil {
		t.Error(err)
	} else {
		b, _ := json.Marshal(res)
		t.Log(string(b))
	}
}

func TestAddressTickerInfo(t *testing.T) {
	c, _ := NewBrc20API(common.BaseUrl, common.FakePrivateKey)
	p := &common.WaasBrc20QueryAddressTickerInfoReq{
		Ticker:      "joy1",
		ChainSymbol: "BTC_TEST",
		Address:     "tb1q7a0jm6rlv8umckjjxqqnwdpze30w4rwafjj4hj",
	}
	if res, err := c.AddressTickerInfo(p); err != nil {
		t.Error(err)
	} else {
		b, _ := json.Marshal(res)
		t.Log(string(b))
	}
}

func TestQueryAddressBalance(t *testing.T) {
	c, _ := NewBrc20API(common.BaseUrl, common.FakePrivateKey)
	p := &common.WaasBrc20QueryAddressBalanceReq{
		ChainSymbol: "BTC_TEST",
		Address:     "tb1q7a0jm6rlv8umckjjxqqnwdpze30w4rwafjj4hj",
	}
	if res, err := c.AddressBalance(p); err != nil {
		t.Error(err)
	} else {
		b, _ := json.Marshal(res)
		t.Log(string(b))
	}
}
