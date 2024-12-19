package main

import (
	"flag"
	"fmt"

	"github.com/sinohope/golang-sdk/core/sdk"
	"github.com/sinohope/golang-sdk/log"
	"github.com/sirupsen/logrus"

	"github.com/sinohope/golang-sdk/common"
)

var (
	private = flag.String("private", common.FakePrivateKey, "Your private key.")
	url     = flag.String("url", common.BaseUrl, "Base url to Sinohope WaaS service.")
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
		return fmt.Errorf("base url can not be empty")
	}
	return nil
}

func checkMPCNodeStatus() {
	logrus.
		WithField("private", private).
		Infof("after generate ECDSA keypair")
	client, err := sdk.NewApiClient(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		logrus.Errorf("create mpc node sdk failed, %v", err)
		return
	}
	request := &common.WaasMpcNodeExecRecordParam{
		BusinessExecType:   1,
		BusinessExecStatus: 10,
		SinoId:             "fake-sino-id",
		PageIndex:          0,
		PageSize:           40,
	}
	var result *common.WaaSMPCNodeRequestRes
	if result, err = client.MPCNode.ListMPCRequests(request); err != nil {
		logrus.Errorf("list mpc requests failed, %v", err)
	} else {
		// TODO: do something with result
		fmt.Printf("-----------> [%v]", result.List)
	}
	var status *common.WaaSMpcNodeStatusDTOData
	if status, err = client.MPCNode.Status(); err != nil {
		logrus.Errorf("get mpc node status failed, %v", err)
	} else {
		logrus.Infof("get mpc node status success, %v", status)
	}
}
