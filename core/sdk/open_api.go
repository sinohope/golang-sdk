package sdk

import (
	"encoding/json"
	"fmt"
	commonData "github.com/sinohope/golang-sdk/common"
	"github.com/sinohope/golang-sdk/core/gateway"
	"github.com/sinohope/golang-sdk/features"
)

type openApi struct {
	gw features.Gateway
}

func NewOpenAPI(baseUrl, private string) (features.OpenAPI, error) {
	gw, err := gateway.NewGateway(baseUrl, private)
	if err != nil {
		return nil, fmt.Errorf("create new gateway failed, %v", err)
	}
	return &openApi{
		gw: gw,
	}, nil
}

// CreateWallets 创建钱包
func (m *openApi) CreateWallets(param *commonData.WaaSCreateBatchWalletParam) ([]*commonData.WaaSWalletInfoData, error) {
	if response, err := m.gw.Post("/v1/api/address/create_wallets", param); err != nil {
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

// ListWallets 查询钱包列表
func (m *openApi) ListWallets(param *commonData.WaaSListWalletsParam) (*commonData.WaaSListWalletsResult, error) {
	if response, err := m.gw.Post("/v1/api/address/list_wallets", param); err != nil {
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

// GenerateChainAddresses 生成链地址
// POST: /v1/waas/mpc/wallet/generate_chain_addresses
func (m *openApi) GenerateChainAddresses(param *commonData.WaaSGenerateChainAddressParam) ([]*commonData.WaaSAddressInfoData, error) {
	if response, err := m.gw.Post("/v1/api/address/generate_chain_addresses", param); err != nil {
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

// ListAddress 查询链地址
func (m *openApi) ListAddress(param *commonData.WaaSListAddressesParam) (*commonData.WaaSListAddressesResult, error) {
	if response, err := m.gw.Post("/v1/api/address/list_addresses", param); err != nil {
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

// ListAddedChains 查询指定钱包下已添加地址的链及其首个地址信息
func (m *openApi) ListAddedChains(param *commonData.WaaSListAddedChainsParam) ([]*commonData.WaaSListAddedChainsDTOData, error) {
	if response, err := m.gw.Post("/v1/api/address/list_added_chains", param); err != nil {
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
func (m *openApi) GetAddressBalance(param *commonData.WaaSGetAddressBalanceParam) (*commonData.WaaSGetWalletBalanceDTOData, error) {
	if response, err := m.gw.Post("/v1/api/address/get_address_balance", param); err != nil {
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

// GetSupportedChains 查询系统支持链
func (c *openApi) GetSupportedChains() ([]*commonData.WaasChainData, error) {
	if response, err := c.gw.Post("/v1/api/common/get_supported_chains", nil); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result []*commonData.WaasChainData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// GetSupportedCoins 查询链支持的币种列表
func (c *openApi) GetSupportedCoins(param *commonData.WaasChainParam) ([]*commonData.WaaSCoinDTOData, error) {
	if response, err := c.gw.Post("/v1/api/common/get_supported_coins", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result []*commonData.WaaSCoinDTOData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// GetVaults 金库列表
func (c *openApi) GetVaults() ([]*commonData.WaaSVaultInfoData, error) {
	if response, err := c.gw.Post("/v1/api/common/get_vaults", nil); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result []*commonData.WaaSVaultInfoData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// Fee 估算交易所需费用
func (t *openApi) Fee(param *commonData.WalletTransactionFeeWAASParam) (*commonData.WalletTransactionFeeWAASResponse, error) {
	if response, err := t.gw.Post("/v1/api/transaction/fee", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.WalletTransactionFeeWAASResponse
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// CreateTransfer 发起转账交易
func (t *openApi) CreateTransfer(param *commonData.WalletTransactionSendWAASParam) (*commonData.CreateSettlementTxResData, error) {
	if response, err := t.gw.Post("/v1/api/transaction/create_transfer", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.CreateSettlementTxResData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// SpeedupTransaction 交易加速
func (t *openApi) SpeedupTransaction(param *commonData.WalletTransactionSpeedupWAASParam) (*commonData.CreateSettlementTxResData, error) {
	if response, err := t.gw.Post("/v1/api/transaction/speedup_transaction", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.CreateSettlementTxResData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// CancelTransaction 交易取消，该方法只支持类eth系列交易取消
func (t *openApi) CancelTransaction(param *commonData.WalletTransactionCancelWAASParam) (*commonData.CreateSettlementTxResData, error) {
	if response, err := t.gw.Post("/v1/api/transaction/cancel_transaction", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.CreateSettlementTxResData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// TransactionsByRequestIds 根据requestIds查询交易列表
func (t *openApi) TransactionsByRequestIds(param *commonData.WalletTransactionQueryWAASRequestIdParam) (*commonData.TransferHistoryWAASDTO, error) {
	if response, err := t.gw.Post("/v1/api/transaction/transactions_by_request_ids", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.TransferHistoryWAASDTO
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// TransactionsBySinoIds 根据sinoIds查询交易列表
func (t *openApi) TransactionsBySinoIds(param *commonData.WalletTransactionQueryBySinoIdParam) (*commonData.TransferHistoryWAASDTO, error) {
	if response, err := t.gw.Post("/v1/api/transaction/transactions_by_sino_ids", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.TransferHistoryWAASDTO
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// TransactionsByTxHash 根据txHash查询交易列表
func (t *openApi) TransactionsByTxHash(param *commonData.WalletTransactionQueryWAASTxHashdParam) (*commonData.TransferHistoryWAASDTO, error) {
	if response, err := t.gw.Post("/v1/api/transaction/transactions_by_tx_hash", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.TransferHistoryWAASDTO
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// ListTransactions 交易列表
func (t *openApi) ListTransactions(param *commonData.WalletTransactionQueryWAASParam) (*commonData.TransferHistoryWAASDTO, error) {
	if response, err := t.gw.Post("/v1/api/transaction/list_transactions", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.TransferHistoryWAASDTO
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// IsValidAddress 检查币种地址是否正确
func (m *openApi) IsValidAddress(param *commonData.WaaSAddressCheckParam) (*commonData.WaaSAddressCheckDTOData, error) {
	if response, err := m.gw.Post("/v1/api/common/is_valid_address", param); err != nil {
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
