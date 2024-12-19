package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/sinohope/golang-sdk/common"
	commonData "github.com/sinohope/golang-sdk/common"
	"github.com/sinohope/golang-sdk/core/gateway"
	"github.com/sinohope/golang-sdk/features"
)

type stakingApi struct {
	gw features.Gateway
}

func NewStakingApi(baseUrl, private string) (features.StakingApi, error) {
	gw, err := gateway.NewGateway(baseUrl, private)
	if err != nil {
		return nil, fmt.Errorf("create new gateway failed, %v", err)
	}
	return &stakingApi{
		gw: gw,
	}, nil
}

func (s *stakingApi) Create(param *common.StakingCreateParam) (*common.StakingRes, error) {
	if response, err := s.gw.Post("/v1/waas/mpc/staking/create", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.StakingRes
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

func (s *stakingApi) Unbond(param *common.UnbondCreateParam) (*common.StakingRes, error) {
	if response, err := s.gw.Post("/v1/waas/mpc/staking/unbond", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.StakingRes
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

func (s *stakingApi) SpendingTimeLockPathTx(param *common.SpendingTimeLockPathTxParam) (*common.StakingRes, error) {
	if response, err := s.gw.Post("/v1/waas/mpc/staking/spending_time_lock_path_tx", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.StakingRes
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}
