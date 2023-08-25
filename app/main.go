package main

import (
	"flag"
	"fmt"
	"github.com/sinohope/sinohope-golang-sdk/log"

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
