package sdk

import (
	"github.com/sinohope/sinohope-golang-sdk/common"
	"testing"
)

func TestCreate(t *testing.T) {
	s, err := NewStakingApi(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	p := &common.StakingCreateParam{
		RequestId:     "2012",
		ChainSymbol:   "SIGNET",
		From:          "tb1q5gddgnx7umasha6vs7xd4784mfv858nzrt0l23",
		FeeRate:       "80",
		StakingAmount: "30000",
		StakingTime:   "150",
	}
	res, err := s.Create(p)
	if err != nil {
		t.Fatal(res)
	}
	t.Log(res)
}

func TestUnbond(t *testing.T) {
	s, err := NewStakingApi(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	p := &common.UnbondCreateParam{
		RequestId:    "1114",
		OriRequestId: "1012",
	}
	res, err := s.Unbond(p)
	if err != nil {
		t.Fatal(res)
	}
	t.Log(res)
}

func TestSpend(t *testing.T) {
	s, err := NewStakingApi(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	p := &common.SpendingTimeLockPathTxParam{
		RequestId:    "1116",
		OriRequestId: "1012",
		FeeRate:      "50",
	}
	res, err := s.SpendingTimeLockPathTx(p)
	if err != nil {
		t.Fatal(res)
	}
	t.Log(res)
}
