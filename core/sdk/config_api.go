package sdk

import (
	"fmt"

	"github.com/sinohope/golang-sdk/common"
	"github.com/sinohope/golang-sdk/core/gateway"
	"github.com/sinohope/golang-sdk/features"
)

type configAPI struct {
	gw features.Gateway
}

func NewConfigAPI(baseUrl, private string) (features.ConfigAPI, error) {
	gw, err := gateway.NewGateway(baseUrl, private)
	if err != nil {
		return nil, fmt.Errorf("create new gateway failed, %v", err)
	}
	return &configAPI{
		gw: gw,
	}, nil
}

func (api *configAPI) SetTransferStrategy(param *common.SetTransferStrategyReq) (*common.Response, error) {
	if response, err := api.gw.Post("/v1/api/config/set_transfer_strategy", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		return response, nil
	}
}
