package sdk

import (
	"encoding/json"
	"fmt"

	"github.com/sinohope/sinohope-golang-sdk/common"
	"github.com/sinohope/sinohope-golang-sdk/core/gateway"
	"github.com/sinohope/sinohope-golang-sdk/features"
)

type mpcNodeAPI struct {
	gw features.Gateway
}

func NewMPCNodeAPI(baseUrl, private string) (features.MPCNodeAPI, error) {
	gw, err := gateway.NewGateway(baseUrl, private)
	if err != nil {
		return nil, fmt.Errorf("create new gateway failed, %v", err)
	}
	return &mpcNodeAPI{
		gw: gw,
	}, nil
}

// ListMPCRequests 查询mpc协议执行记录
// POST: /v1/waas/mpc/mpcnode/list_mpc_requests
func (m *mpcNodeAPI) ListMPCRequests(param *common.WaasMpcNodeExecRecordParam) (*common.TransferHistoryWAASDTO, error) {
	if response, err := m.gw.Post("/v1/waas/mpc/mpcnode/list_mpc_requests", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *common.TransferHistoryWAASDTO
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// Status 查询MPC node状态
// POST: /v1/waas/mpc/mpcnode/status
func (m *mpcNodeAPI) Status() (*common.WaaSMpcNodeStatusDTOData, error) {
	if response, err := m.gw.Post("/v1/waas/mpc/mpcnode/status", nil); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *common.WaaSMpcNodeStatusDTOData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}
