package sdk

import (
	"encoding/json"
	"fmt"
	commonData "github.com/sinohope/golang-sdk/common"

	"github.com/sinohope/golang-sdk/common"
	"github.com/sinohope/golang-sdk/core/gateway"
	"github.com/sinohope/golang-sdk/features"
)

type runeApi struct {
	gw features.Gateway
}

func NewRuneApi(baseUrl, private string) (features.RuneApi, error) {
	gw, err := gateway.NewGateway(baseUrl, private)
	if err != nil {
		return nil, fmt.Errorf("create new gateway failed, %v", err)
	}
	return &runeApi{
		gw: gw,
	}, nil
}

func (t *runeApi) Transfer(param *common.RuneTransferParam) (*common.InscribeRes, error) {
	if response, err := t.gw.Post("/v1/waas/mpc/runes/transfer", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.InscribeRes
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}
func (t *runeApi) PageBalanceSummary(param *common.RunePageBalanceSummaryParam) (*common.RunePageBalanceSummaryRes, error) {
	if response, err := t.gw.Post("/v1/waas/mpc/runes/page_balance_summary", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.RunePageBalanceSummaryRes
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

func (t *runeApi) QueryBalance(param *common.RuneQueryBalanceParam) (*common.RuneQueryBalanceRes, error) {
	if response, err := t.gw.Post("/v1/waas/mpc/runes/balance", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.RuneQueryBalanceRes
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}
