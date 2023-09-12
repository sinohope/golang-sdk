package features

import "github.com/sinohope/sinohope-golang-sdk/common"

type SinohopeWaaSAPI interface{}

type CommonAPI interface {
	SinohopeWaaSAPI

	// GetSupportedChains 查询系统支持链
	// POST: /v1/waas/common/get_supported_chains
	GetSupportedChains() ([]*common.WaasChainData, error)

	// GetSupportedCoins 查询链支持的币种列表
	// POST: /v1/waas/common/get_supported_coins
	GetSupportedCoins(param *common.WaasChainParam) ([]*common.WaaSCoinDTOData, error)

	// GetVaults 金库列表
	// POST: /v1/waas/common/get_vaults
	GetVaults() ([]*common.WaaSVaultInfoData, error)
}

type MPCAPI interface {
	SinohopeWaaSAPI

	// CreateWallets 创建钱包
	// POST: /v1/waas/mpc/create_wallets
	CreateWallets(param *common.WaaSCreateBatchWalletParam) ([]*common.WaasWalletInfoData, error)

	// ListAddress 查询链地址
	// POST: /v1/waas/mpc/wallet/list_addresses
	ListAddress(param *common.WaaSListAddressesParam) ([]*common.WaaSAddressInfoData, error)

	// GenerateChainAddresses 生成链地址
	// POST: /v1/waas/mpc/wallet/generate_chain_addresses
	GenerateChainAddresses(param *common.WaaSGenerateChainAddressParam) ([]*common.WaaSAddressData, error)

	// ListWallets 查询钱包列表
	// POST: /v1/waas/mpc/wallet/list_wallets
	ListWallets(param *common.WaaSListWalletsParam) ([]*common.WaaSWalletInfoData, error)

	// ListAddedChains 查询指定钱包下已添加地址的链及其首个地址信息
	// POST: /v1/waas/mpc/wallet/list_added_chains
	ListAddedChains(param *common.WaaSListAddedChainsParam) ([]*common.WaaSListAddedChainsDTOData, error)

	// GetAddressBalance 查询地址余额信息
	// POST: /v1/waas/mpc/wallet/get_address_balance
	GetAddressBalance(param *common.WaaSGetAddressBalanceParam) (*common.WaaSGetWalletBalanceDTOData, error)

	// IsValidAddress 检查币种地址是否正确
	// POST: /v1/waas/mpc/is_valid_address
	IsValidAddress(param *common.WaaSAddressCheckParam) (*common.WaaSAddressCheckDTOData, error)

	// TransferAddressBook 当前金库设置开关以后,支持转账的地址簿
	// POST: /v1/waas/mpc/wallet/transfer_address_book
	TransferAddressBook(param *common.WaaSTransferAddressBookParam) ([]*common.WaaSTransferAddressBookDTOData, error)

	// TransferRiskControlSwitch 查询当前金库是否设置了风控开关
	// POST: /v1/waas/mpc/wallet/transfer_risk_control_switch
	TransferRiskControlSwitch(param *common.WaaSVaultIdDTO) (*common.WaaSTransferAddressSwitchDTOData, error)

	// CreateTransfer 发起转账交易
	// POST: /v1/waas/mpc/transaction/create_transfer
	CreateTransfer(param *common.WalletTransactionSendWAASParam) (*common.CreateSettlementTxResData, error)

	// CreateTransaction 发起任意交易
	// POST: /v1/waas/mpc/transaction/create_transaction
	CreateTransaction(param *common.WalletTransactionSendDataWAASParam) (*common.CreateSettlementTxResData, error)

	// SpeedupTransaction 交易加速
	// POST: /v1/waas/mpc/transaction/speedup_transaction
	SpeedupTransaction(param *common.WalletTransactionSpeedupWAASParam) (*common.CreateSettlementTxResData, error)

	// CancelTransaction 交易取消，该方法只支持类eth系列交易取消
	// POST: /v1/waas/mpc/transaction/cancel_transaction
	CancelTransaction(param *common.WalletTransactionCancelWAASParam) (*common.CreateSettlementTxResData, error)

	// ListTransactions 交易列表
	// POST: /v1/waas/mpc/transaction/list_transactions
	ListTransactions(param *common.WalletTransactionQueryWAASParam) ([]*common.TransferHistoryWAASDTO, error)

	// TransactionsByRequestIds 根据requestIds查询交易列表
	// POST: /v1/waas/mpc/transaction/transactions_by_request_ids
	TransactionsByRequestIds(param *common.WalletTransactionQueryWAASRequestIdParam) ([]*common.TransferHistoryWAASDTO, error)

	// TransactionsBySinoIds 根据sinoIds查询交易列表
	// POST: /v1/waas/mpc/transaction/transactions_by_sino_ids
	TransactionsBySinoIds(param *common.WalletTransactionQueryBySinoIdParam) ([]*common.TransferHistoryWAASDTO, error)

	// TransactionsByTxHash 根据txHash查询交易列表
	// POST: /v1/waas/mpc/transaction/transactions_by_tx_hash
	TransactionsByTxHash(param *common.WalletTransactionQueryWAASTxHashdParam) ([]*common.TransferHistoryWAASDTO, error)

	// SignMessage 按已知的规范签名消息（EIP-191、 EIP-712）
	// POST: /v1/waas/mpc/web3/sign_message
	SignMessage(param *common.SignMessageParam) (*common.WaaSSignatureData, error)
}

type AdvanceAPI interface {
	SinohopeWaaSAPI

	// SignRawData 原始数据签名
	// POST: /v1/waas/mpc/wallet/advance/sign_raw_data
	SignRawData(param *common.WaasSignRawDataParam) (*common.Response, error)

	// GenAddressByPath 根据指定的路径创建地址
	// POST: /v1/waas/mpc/wallet/advance/gen_address_by_path
	GenAddressByPath(param *common.WaasAddressPathParam) (*common.Response, error)

	// UpdateWallet 更新钱包属性（高级功能开启、关闭）
	// POST: /v1/waas/mpc/wallet/advance/update_wallet
	UpdateWallet(param *common.WaasUpdateWalletParam) (*common.Response, error)
}

type MPCNodeAPI interface {
	SinohopeWaaSAPI

	// ListMPCRequests 查询mpc协议执行记录
	// POST: /v1/waas/mpc/mpcnode/list_mpc_requests
	ListMPCRequests(param *common.WaasMpcNodeExecRecordParam) (*common.PageData, error)

	// Status 查询MPC node状态
	// POST: /v1/waas/mpc/mpcnode/status
	Status() (*common.WaaSMpcNodeStatusDTOData, error)
}
