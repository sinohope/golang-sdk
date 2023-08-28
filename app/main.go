package main

import (
	"flag"
	"fmt"
	"github.com/sinohope/sinohope-golang-sdk/core/http"
	"github.com/sinohope/sinohope-golang-sdk/core/sdk"
	"github.com/sinohope/sinohope-golang-sdk/core/signer"
	"github.com/sinohope/sinohope-golang-sdk/log"
	"github.com/sirupsen/logrus"

	"github.com/sinohope/sinohope-golang-sdk/common"
)

var (
	private = flag.String("private", common.FakePrivateKey, "Your private key.")
	url     = flag.String("url", common.BaseUrl, "URL to mpc-proxy.")
)

func main() {
	flag.Parse()

	if err := check(); err != nil {
		panic(err)
	}

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
		File: struct {
			Enable bool   `toml:"enable"`
			Level  int    `toml:"level"`
			Path   string `toml:"path"`
			MaxAge int    `toml:"max-age"`
		}{
			Enable: true,
			Level:  5,
			Path:   "./logs/sinohope-golang-sdk.log",
			MaxAge: 7,
		},
	}
	log.SetLogDetailsByConfig(a, l)

	checkMPCNodeStatus()
}

func check() error {
	if *private == "" {
		return fmt.Errorf("private key can not be empty")
	}
	if *url == "" {
		return fmt.Errorf("mpc-proxy url can not be empty")
	}
	return nil
}

func checkMPCNodeStatus() {
	s, err := signer.NewSigner(common.FakePrivateKey)
	if err != nil {
		logrus.Errorf("create new signer failed, %v", err)
		return
	}
	p, err := http.NewHTTP(common.BaseUrl, s)
	if err != nil {
		logrus.Errorf("create http failed, %v", err)
		return
	}
	m, err := sdk.NewMPCNodeAPI(p)
	if err != nil {
		logrus.Errorf("create mpc node sdk failed, %v", err)
		return
	}
	request := &common.WaasMpcNodeExecRecordParam{
		BusinessExecType:   1,
		BusinessExecStatus: 10,
		SinoId:             "fake sino id",
		PageIndex:          0,
		PageSize:           40,
	}
	var result *common.PageData
	if result, err = m.ListMPCRequests(request); err != nil {
		logrus.Errorf("list mpc requests failed, %v", err)
	} else {
		// TODO: do something with result
		fmt.Printf("-----------> [%s]", result.List)
	}
	var status *common.WaaSMpcNodeStatusDTOData
	if status, err = m.Status(); err != nil {
		logrus.Errorf("get mpc node status failed, %v", err)
	} else {
		logrus.Infof("get mpc node status success, %v", status)
	}
}
