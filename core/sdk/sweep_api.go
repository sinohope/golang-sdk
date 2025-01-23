package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/sinohope/golang-sdk/common"
	"github.com/sinohope/golang-sdk/core/gateway"
	"github.com/sinohope/golang-sdk/features"
)

type sweepApi struct {
	gw features.Gateway
}

func newSweepApi(baseUrl, private string) (features.SweepApi, error) {
	gw, err := gateway.NewGateway(baseUrl, private)
	if err != nil {
		return nil, fmt.Errorf("create new gateway failed, %v", err)
	}
	return &sweepApi{
		gw: gw,
	}, nil
}

func (api *sweepApi) CollectionStrategy(param *common.CollectionStrategyReq) (*common.Response, error) {
	if response, err := api.gw.Post("/v1/api/sweep/set_strategy", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		return response, nil
	}
}

func (api *sweepApi) CollectionStrategyLists(param *common.CollectionStrategyListsReq) (*common.PageCollectStrategyRes, error) {
	if response, err := api.gw.Post("/v1/api/sweep/strategy_lists", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *common.PageCollectStrategyRes
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

func (api *sweepApi) SetGasStation(param *common.SetGasStationReq) (*common.Response, error) {
	if response, err := api.gw.Post("/v1/api/sweep/set_gas_station", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		return response, nil
	}
}

func (api *sweepApi) GasStationLists(param *common.GasStationListsReq) ([]*common.GasStationRes, error) {
	if response, err := api.gw.Post("/v1/api/sweep/gas_station_lists", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result []*common.GasStationRes
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}
