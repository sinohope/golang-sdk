package sdk

import (
	"github.com/sinohope/golang-sdk/common"
	"testing"
)

func TestCollectionStrategy(t *testing.T) {
	s, err := newSweepApi(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	p := &common.CollectionStrategyReq{
		VaultId:                 "534605276521221",
		ToAddress:               "tb1q7a0jm6rlv8umckjjxqqnwdpze30w4rwafjj4hj",
		AssetId:                 "BTC_BTC_TEST",
		MinimumCollectionAmount: 10000,
	}
	res, err := s.CollectionStrategy(p)
	if err != nil {
		t.Fatal(res)
	}
	t.Log(res)
}

func TestCollectionStrategyLists(t *testing.T) {
	s, err := newSweepApi(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	p := &common.CollectionStrategyListsReq{
		VaultId:   "534605276521221",
		PageIndex: 1,
		PageSize:  10,
		Offset:    0,
	}
	res, err := s.CollectionStrategyLists(p)
	if err != nil {
		t.Fatal(res)
	}
	t.Log(res)
}

func TestSetGasStation(t *testing.T) {
	s, err := newSweepApi(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	p := &common.SetGasStationReq{
		ChainSymbol: "BTC_TEST",
		Address:     "tb1q7a0jm6rlv8umckjjxqqnwdpze30w4rwafjj4hj",
	}
	res, err := s.SetGasStation(p)
	if err != nil {
		t.Fatal(res)
	}
	t.Log(res)
}

func TestGasStationLists(t *testing.T) {
	s, err := newSweepApi(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	p := &common.GasStationListsReq{
		ChainSymbol: "BTC_TEST",
	}
	res, err := s.GasStationLists(p)
	if err != nil {
		t.Fatal(res)
	}
	t.Log(res)
}
