package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/sinohope/sinohope-golang-sdk/common"
	"github.com/sinohope/sinohope-golang-sdk/features"
)

type advanceAPI struct {
	h features.HTTP
}

func NewAdvanceAPI(h features.HTTP) (features.AdvanceAPI, error) {
	return &advanceAPI{
		h: h,
	}, nil
}

// SignRawData 原始数据签名
// POST: /v1/waas/mpc/wallet/advance/sign_raw_data
func (a *advanceAPI) SignRawData(param *common.WaasSignRawDataParam) (*common.Response, error) {
	if response, err := a.h.Post("/v1/waas/mpc/wallet/advance/sign_raw_data", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *common.Response
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// GenAddressByPath 根据指定的路径创建地址
// POST: /v1/waas/mpc/wallet/advance/gen_address_by_path
func (a *advanceAPI) GenAddressByPath(param *common.WaasAddressPathParam) (*common.Response, error) {
	if response, err := a.h.Post("/v1/waas/mpc/wallet/advance/gen_address_by_path", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *common.Response
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// UpdateWallet 更新钱包属性（高级功能开启、关闭）
// POST: /v1/waas/mpc/wallet/advance/update_wallet
func (a *advanceAPI) UpdateWallet(param *common.WaasUpdateWalletParam) (*common.Response, error) {
	if response, err := a.h.Post("/v1/waas/mpc/wallet/advance/update_wallet", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *common.Response
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}
