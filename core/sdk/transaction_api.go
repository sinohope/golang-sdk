package sdk

import (
	"encoding/json"
	"fmt"

	commonData "github.com/sinohope/sinohope-golang-sdk/common"

	"github.com/sinohope/sinohope-golang-sdk/core/gateway"
	"github.com/sinohope/sinohope-golang-sdk/features"
)

type transactionAPI struct {
	gw features.Gateway
}

func NewTransactionAPI(baseUrl, private string) (features.TransactionAPI, error) {
	gw, err := gateway.NewGateway(baseUrl, private)
	if err != nil {
		return nil, fmt.Errorf("create new gateway failed, %v", err)
	}
	return &transactionAPI{
		gw: gw,
	}, nil
}

// CreateTransfer 发起转账交易
// POST: /v1/waas/mpc/transaction/create_transfer
func (t *transactionAPI) CreateTransfer(param *commonData.WalletTransactionSendWAASParam) (*commonData.CreateSettlementTxResData, error) {
	if response, err := t.gw.Post("/v1/waas/mpc/transaction/create_transfer", param); err != nil {
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

// CreateTransaction 发起任意交易
// POST: /v1/waas/mpc/transaction/create_transaction
func (t *transactionAPI) CreateTransaction(param *commonData.WalletTransactionSendDataWAASParam) (*commonData.CreateSettlementTxResData, error) {
	if response, err := t.gw.Post("/v1/waas/mpc/transaction/create_transaction", param); err != nil {
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
// POST: /v1/waas/mpc/transaction/speedup_transaction
func (t *transactionAPI) SpeedupTransaction(param *commonData.WalletTransactionSpeedupWAASParam) (*commonData.CreateSettlementTxResData, error) {
	if response, err := t.gw.Post("/v1/waas/mpc/transaction/speedup_transaction", param); err != nil {
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
// POST: /v1/waas/mpc/transaction/cancel_transaction
func (t *transactionAPI) CancelTransaction(param *commonData.WalletTransactionCancelWAASParam) (*commonData.CreateSettlementTxResData, error) {
	if response, err := t.gw.Post("/v1/waas/mpc/transaction/cancel_transaction", param); err != nil {
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

// ListTransactions 交易列表
// POST: /v1/waas/mpc/transaction/list_transactions
func (t *transactionAPI) ListTransactions(param *commonData.WalletTransactionQueryWAASParam) (*commonData.TransferHistoryWAASDTO, error) {
	if response, err := t.gw.Post("/v1/waas/mpc/transaction/list_transactions", param); err != nil {
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

// TransactionsByRequestIds 根据requestIds查询交易列表
// POST: /v1/waas/mpc/transaction/transactions_by_request_ids
func (t *transactionAPI) TransactionsByRequestIds(param *commonData.WalletTransactionQueryWAASRequestIdParam) (*commonData.TransferHistoryWAASDTO, error) {
	if response, err := t.gw.Post("/v1/waas/mpc/transaction/transactions_by_request_ids", param); err != nil {
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
// POST: /v1/waas/mpc/transaction/transactions_by_sino_ids
func (t *transactionAPI) TransactionsBySinoIds(param *commonData.WalletTransactionQueryBySinoIdParam) (*commonData.TransferHistoryWAASDTO, error) {
	if response, err := t.gw.Post("/v1/waas/mpc/transaction/transactions_by_sino_ids", param); err != nil {
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
// POST: /v1/waas/mpc/transaction/transactions_by_tx_hash
func (t *transactionAPI) TransactionsByTxHash(param *commonData.WalletTransactionQueryWAASTxHashdParam) (*commonData.TransferHistoryWAASDTO, error) {
	if response, err := t.gw.Post("/v1/waas/mpc/transaction/transactions_by_tx_hash", param); err != nil {
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

// SignMessage 按已知的规范签名消息（EIP-191、 EIP-712）
// POST: /v1/waas/mpc/web3/sign_message
func (t *transactionAPI) SignMessage(param *commonData.SignMessageParam) (*commonData.WaaSSignatureData, error) {
	if response, err := t.gw.Post("/v1/waas/mpc/sign_message", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.WaaSSignatureData
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}
