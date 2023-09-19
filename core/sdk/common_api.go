package sdk

import (
	"encoding/json"
	"fmt"

	"github.com/sinohope/sinohope-golang-sdk/common"
	"github.com/sinohope/sinohope-golang-sdk/core/gateway"
	"github.com/sinohope/sinohope-golang-sdk/features"
)

type commonAPI struct {
	gw features.Gateway
}

func NewCommonAPI(baseUrl, private string) (features.CommonAPI, error) {
	gw, err := gateway.NewGateway(baseUrl, private)
	if err != nil {
		return nil, fmt.Errorf("create new gateway failed, %v", err)
	}
	return &commonAPI{
		gw: gw,
	}, nil
}

// GetSupportedChains 查询系统支持链
// POST: /v1/waas/commonAPI/get_supported_chains
func (c *commonAPI) GetSupportedChains() ([]*common.WaasChainData, error) {
	if response, err := c.gw.Post("/v1/waas/commonAPI/get_supported_chains", nil); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result []*common.WaasChainData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// GetSupportedCoins 查询链支持的币种列表
// POST: /v1/waas/commonAPI/get_supported_coins
func (c *commonAPI) GetSupportedCoins(param *common.WaasChainParam) ([]*common.WaaSCoinDTOData, error) {
	if response, err := c.gw.Post("/v1/waas/commonAPI/get_supported_coins", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result []*common.WaaSCoinDTOData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// GetVaults 金库列表
// POST: /v1/waas/commonAPI/get_vaults
func (c *commonAPI) GetVaults() ([]*common.WaaSVaultInfoData, error) {
	if response, err := c.gw.Post("/v1/waas/commonAPI/get_vaults", nil); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result []*common.WaaSVaultInfoData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}
