package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/sinohope/sinohope-golang-sdk/common"
	"github.com/sinohope/sinohope-golang-sdk/features"
)

type mpcAPI struct {
	h features.HTTP
}

func NewMPCAPI(h features.HTTP) (features.MPCAPI, error) {
	return &mpcAPI{
		h: h,
	}, nil
}

// CreateWallets 创建钱包
// POST: /v1/waas/mpc/create_wallets
func (m *mpcAPI) CreateWallets(param *common.WaaSCreateBatchWalletParam) ([]*common.WaasWalletInfoData, error) {
	if response, err := m.h.Post("/v1/waas/mpc/create_wallets", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result []*common.WaasWalletInfoData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// ListAddress 查询链地址
// POST: /v1/waas/mpc/wallet/list_addresses
func (m *mpcAPI) ListAddress(param *common.WaaSListAddressesParam) ([]*common.WaaSAddressInfoData, error) {
	if response, err := m.h.Post("/v1/waas/mpc/wallet/list_addresses", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *common.PageData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		var addressInfo []*common.WaaSAddressInfoData
		if err := json.Unmarshal(result.List, &addressInfo); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return addressInfo, nil
	}
}

// GenerateChainAddresses 生成链地址
// POST: /v1/waas/mpc/wallet/generate_chain_addresses
func (m *mpcAPI) GenerateChainAddresses(param *common.WaaSGenerateChainAddressParam) ([]*common.WaaSAddressData, error) {
	if response, err := m.h.Post("/v1/waas/mpc/wallet/generate_chain_addresses", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result []*common.WaaSAddressData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// ListWallets 查询钱包列表
// POST: /v1/waas/mpc/wallet/list_wallets
func (m *mpcAPI) ListWallets(param *common.WaaSListWalletsParam) (*common.TransferHistoryWAASDTO, error) {
	if response, err := m.h.Post("/v1/waas/mpc/wallet/list_wallets", param); err != nil {
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

// ListAddedChains 查询指定钱包下已添加地址的链及其首个地址信息
// POST: /v1/waas/mpc/wallet/list_added_chains
func (m *mpcAPI) ListAddedChains(param *common.WaaSListAddedChainsParam) ([]*common.WaaSListAddedChainsDTOData, error) {
	if response, err := m.h.Post("/v1/waas/mpc/wallet/list_added_chains", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result []*common.WaaSListAddedChainsDTOData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// GetAddressBalance 查询地址余额信息
// POST: /v1/waas/mpc/wallet/get_address_balance
func (m *mpcAPI) GetAddressBalance(param *common.WaaSGetAddressBalanceParam) (*common.WaaSGetWalletBalanceDTOData, error) {
	if response, err := m.h.Post("/v1/waas/mpc/wallet/get_address_balance", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *common.WaaSGetWalletBalanceDTOData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// IsValidAddress 检查币种地址是否正确
// POST: /v1/waas/mpc/is_valid_address
func (m *mpcAPI) IsValidAddress(param *common.WaaSAddressCheckParam) (*common.WaaSAddressCheckDTOData, error) {
	if response, err := m.h.Post("/v1/waas/mpc/wallet/is_valid_address", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *common.WaaSAddressCheckDTOData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// TransferAddressBook 当前金库设置开关以后,支持转账的地址簿
// POST: /v1/waas/mpc/wallet/transfer_address_book
func (m *mpcAPI) TransferAddressBook(param *common.WaaSTransferAddressBookParam) (*common.TransferHistoryWAASDTO, error) {
	if response, err := m.h.Post("/v1/waas/mpc/wallet/transfer_address_book", param); err != nil {
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

// TransferRiskControlSwitch 查询当前金库是否设置了风控开关
// POST: /v1/waas/mpc/wallet/transfer_risk_control_switch
func (m *mpcAPI) TransferRiskControlSwitch(param *common.WaaSVaultIdDTO) (*common.WaaSTransferAddressSwitchDTOData, error) {
	if response, err := m.h.Post("/v1/waas/mpc/wallet/transfer_risk_control_switch", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *common.WaaSTransferAddressSwitchDTOData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// CreateTransfer 发起转账交易
// POST: /v1/waas/mpc/transaction/create_transfer
func (m *mpcAPI) CreateTransfer(param *common.WalletTransactionSendWAASParam) (*common.CreateSettlementTxResData, error) {
	if response, err := m.h.Post("/v1/waas/mpc/wallet/create_transfer", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *common.CreateSettlementTxResData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// CreateTransaction 发起任意交易
// POST: /v1/waas/mpc/transaction/create_transaction
func (m *mpcAPI) CreateTransaction(param *common.WalletTransactionSendDataWAASParam) (*common.CreateSettlementTxResData, error) {
	if response, err := m.h.Post("/v1/waas/mpc/wallet/create_transaction", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *common.CreateSettlementTxResData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// SpeedupTransaction 交易加速
// POST: /v1/waas/mpc/transaction/speedup_transaction
func (m *mpcAPI) SpeedupTransaction(param *common.WalletTransactionSpeedupWAASParam) (*common.CreateSettlementTxResData, error) {
	if response, err := m.h.Post("/v1/waas/mpc/wallet/speedup_transaction", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *common.CreateSettlementTxResData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// CancelTransaction 交易取消，该方法只支持类eth系列交易取消
// POST: /v1/waas/mpc/transaction/cancel_transaction
func (m *mpcAPI) CancelTransaction(param *common.WalletTransactionCancelWAASParam) (*common.CreateSettlementTxResData, error) {
	if response, err := m.h.Post("/v1/waas/mpc/wallet/cancel_transaction", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *common.CreateSettlementTxResData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// ListTransactions 交易列表
// POST: /v1/waas/mpc/transaction/list_transactions
func (m *mpcAPI) ListTransactions(param *common.WalletTransactionQueryWAASParam) (*common.TransferHistoryWAASDTO, error) {
	if response, err := m.h.Post("/v1/waas/mpc/wallet/list_transactions", param); err != nil {
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

// TransactionsByRequestIds 根据requestIds查询交易列表
// POST: /v1/waas/mpc/transaction/transactions_by_request_ids
func (m *mpcAPI) TransactionsByRequestIds(param *common.WalletTransactionQueryWAASRequestIdParam) (*common.TransferHistoryWAASDTO, error) {
	if response, err := m.h.Post("/v1/waas/mpc/wallet/transactions_by_request_ids", param); err != nil {
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

// TransactionsBySinoIds 根据sinoIds查询交易列表
// POST: /v1/waas/mpc/transaction/transactions_by_sino_ids
func (m *mpcAPI) TransactionsBySinoIds(param *common.WalletTransactionQueryBySinoIdParam) (*common.TransferHistoryWAASDTO, error) {
	if response, err := m.h.Post("/v1/waas/mpc/wallet/transactions_by_sino_ids", param); err != nil {
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

// TransactionsByTxHash 根据txHash查询交易列表
// POST: /v1/waas/mpc/transaction/transactions_by_tx_hash
func (m *mpcAPI) TransactionsByTxHash(param *common.WalletTransactionQueryWAASTxHashdParam) (*common.TransferHistoryWAASDTO, error) {
	if response, err := m.h.Post("/v1/waas/mpc/wallet/transactions_by_tx_hash", param); err != nil {
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

// SignMessage 按已知的规范签名消息（EIP-191、 EIP-712）
// POST: /v1/waas/mpc/web3/sign_message
func (m *mpcAPI) SignMessage(param *common.SignMessageParam) (*common.WaaSSignatureData, error) {
	if response, err := m.h.Post("/v1/waas/mpc/web3/sign_message", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != common.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *common.WaaSSignatureData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}
