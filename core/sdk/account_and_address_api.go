package sdk

import (
	"encoding/json"
	"fmt"

	commonData "github.com/sinohope/sinohope-golang-sdk/common"
	"github.com/sinohope/sinohope-golang-sdk/core/gateway"
	"github.com/sinohope/sinohope-golang-sdk/features"
)

type accountAndAddressAPI struct {
	gw features.Gateway
}

func NewAccountAndAddressAPI(baseUrl, private string) (features.AccountAndAddressAPI, error) {
	gw, err := gateway.NewGateway(baseUrl, private)
	if err != nil {
		return nil, fmt.Errorf("create new gateway failed, %v", err)
	}
	return &accountAndAddressAPI{
		gw: gw,
	}, nil
}

// CreateWallets 创建钱包
// POST: /v1/waas/mpc/create_wallets
func (m *accountAndAddressAPI) CreateWallets(param *commonData.WaaSCreateBatchWalletParam) ([]*commonData.WaaSWalletInfoData, error) {
	if response, err := m.gw.Post("/v1/waas/mpc/create_wallets", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result []*commonData.WaaSWalletInfoData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// ListAddress 查询链地址
// POST: /v1/waas/mpc/wallet/list_addresses
func (m *accountAndAddressAPI) ListAddress(param *commonData.WaaSListAddressesParam) (*commonData.WaaSListAddressesResult, error) {
	if response, err := m.gw.Post("/v1/waas/mpc/wallet/list_addresses", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.WaaSListAddressesResult
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// GenerateChainAddresses 生成链地址
// POST: /v1/waas/mpc/wallet/generate_chain_addresses
func (m *accountAndAddressAPI) GenerateChainAddresses(param *commonData.WaaSGenerateChainAddressParam) ([]*commonData.WaaSAddressInfoData, error) {
	if response, err := m.gw.Post("/v1/waas/mpc/wallet/generate_chain_addresses", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result []*commonData.WaaSAddressInfoData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// ListWallets 查询钱包列表
// POST: /v1/waas/mpc/wallet/list_wallets
func (m *accountAndAddressAPI) ListWallets(param *commonData.WaaSListWalletsParam) (*commonData.WaaSListWalletsResult, error) {
	if response, err := m.gw.Post("/v1/waas/mpc/wallet/list_wallets", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.WaaSListWalletsResult
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// ListAddedChains 查询指定钱包下已添加地址的链及其首个地址信息
// POST: /v1/waas/mpc/wallet/list_added_chains
func (m *accountAndAddressAPI) ListAddedChains(param *commonData.WaaSListAddedChainsParam) ([]*commonData.WaaSListAddedChainsDTOData, error) {
	if response, err := m.gw.Post("/v1/waas/mpc/wallet/list_added_chains", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result []*commonData.WaaSListAddedChainsDTOData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// GetAddressBalance 查询地址余额信息
// POST: /v1/waas/mpc/wallet/get_address_balance
func (m *accountAndAddressAPI) GetAddressBalance(param *commonData.WaaSGetAddressBalanceParam) (*commonData.WaaSGetWalletBalanceDTOData, error) {
	if response, err := m.gw.Post("/v1/waas/mpc/wallet/get_address_balance", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.WaaSGetWalletBalanceDTOData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// IsValidAddress 检查币种地址是否正确
// POST: /v1/waas/mpc/is_valid_address
func (m *accountAndAddressAPI) IsValidAddress(param *commonData.WaaSAddressCheckParam) (*commonData.WaaSAddressCheckDTOData, error) {
	if response, err := m.gw.Post("/v1/waas/mpc/wallet/is_valid_address", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.WaaSAddressCheckDTOData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// TransferAddressBook 当前金库设置开关以后,支持转账的地址簿
// POST: /v1/waas/mpc/wallet/transfer_address_book
func (m *accountAndAddressAPI) TransferAddressBook(param *commonData.WaaSTransferAddressBookParam) (*commonData.WaaSAddressBookResult, error) {
	if response, err := m.gw.Post("/v1/waas/mpc/wallet/transfer_address_book", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.WaaSAddressBookResult
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// TransferRiskControlSwitch 查询当前金库是否设置了风控开关
// POST: /v1/waas/mpc/wallet/transfer_risk_control_switch
func (m *accountAndAddressAPI) TransferRiskControlSwitch(param *commonData.WaaSVaultIdDTO) (*commonData.WaaSTransferAddressSwitchDTOData, error) {
	if response, err := m.gw.Post("/v1/waas/mpc/wallet/transfer_risk_control_switch", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.WaaSTransferAddressSwitchDTOData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}
