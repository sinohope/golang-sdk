package signer

import (
	"encoding/json"
	"github.com/sinohope/sinohope-golang-sdk/common"
	"github.com/sinohope/sinohope-golang-sdk/log"
	"testing"
)

func TestSign(t *testing.T) {
	a := log.App{
		Name: "sinohope-golang-sdk",
	}
	l := log.Log{
		Stdout: struct {
			Enable bool `toml:"enable"`
			Level  int  `toml:"level"`
		}{
			Enable: true,
			Level:  5,
		},
	}
	log.SetLogDetailsByConfig(a, l)

	s, err := NewSigner(common.FakePrivateKey)
	if err != nil {
		t.Fatalf("create new signer failed, %v", err)
	}
	request := common.WaasMpcNodeExecRecordParam{
		BusinessExecType:   1,
		BusinessExecStatus: 10,
		SinoId:             "fake sino id",
		PageIndex:          0,
		PageSize:           120,
	}
	payload, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal payload failed, %v", err)
	}
	if signature, err := s.Sign("/test", "123", string(payload)); err != nil {
		t.Fatalf("sign failed, %v", err)
	} else {
		t.Logf("signature: %s", signature)
	}
}

func TestPublicKey(t *testing.T) {
	s, err := NewSigner(common.FakePrivateKey)
	if err != nil {
		t.Fatalf("create new signer failed, %v", err)
	}
	if public := s.PublicKey(); public == "" {
		t.Fatalf("signer public key is empty")
	} else {
		t.Logf("signer public key: %s", public)
	}
}
