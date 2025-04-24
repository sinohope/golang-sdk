package features

import (
	"github.com/sinohope/golang-sdk/common"
)

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

type AccountAndAddressAPI interface {
	SinohopeWaaSAPI

	// CreateWallets 创建钱包
	// POST: /v1/waas/mpc/create_wallets
	CreateWallets(param *common.WaaSCreateBatchWalletParam) ([]*common.WaaSWalletInfoData, error)

	// ListAddress 查询链地址
	// POST: /v1/waas/mpc/wallet/list_addresses
	ListAddress(param *common.WaaSListAddressesParam) (*common.WaaSListAddressesResult, error)

	// GenerateChainAddresses 生成链地址
	// POST: /v1/waas/mpc/wallet/generate_chain_addresses
	GenerateChainAddresses(param *common.WaaSGenerateChainAddressParam) ([]*common.WaaSAddressInfoData, error)

	// ListWallets 查询钱包列表
	// POST: /v1/waas/mpc/wallet/list_wallets
	ListWallets(param *common.WaaSListWalletsParam) (*common.WaaSListWalletsResult, error)

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
	TransferAddressBook(param *common.WaaSTransferAddressBookParam) (*common.WaaSAddressBookResult, error)

	// TransferRiskControlSwitch 查询当前金库是否设置了风控开关
	// POST: /v1/waas/mpc/wallet/transfer_risk_control_switch
	TransferRiskControlSwitch(param *common.WaaSVaultIdDTO) (*common.WaaSTransferAddressSwitchDTOData, error)
}

type TransactionAPI interface {
	SinohopeWaaSAPI

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
	ListTransactions(param *common.WalletTransactionQueryWAASParam) (*common.TransferHistoryWAASDTO, error)

	// TransactionsByRequestIds 根据requestIds查询交易列表
	// POST: /v1/waas/mpc/transaction/transactions_by_request_ids
	TransactionsByRequestIds(param *common.WalletTransactionQueryWAASRequestIdParam) (*common.TransferHistoryWAASDTO, error)

	// TransactionsBySinoIds 根据sinoIds查询交易列表
	// POST: /v1/waas/mpc/transaction/transactions_by_sino_ids
	TransactionsBySinoIds(param *common.WalletTransactionQueryBySinoIdParam) (*common.TransferHistoryWAASDTO, error)

	// TransactionsByTxHash 根据txHash查询交易列表
	// POST: /v1/waas/mpc/transaction/transactions_by_tx_hash
	TransactionsByTxHash(param *common.WalletTransactionQueryWAASTxHashdParam) (*common.TransferHistoryWAASDTO, error)

	// Fee 估算交易所需费用
	// POST: /v1/waas/mpc/transaction/fee
	Fee(param *common.WalletTransactionFeeWAASParam) (*common.WalletTransactionFeeWAASResponse, error)

	// SignMessage 按已知的规范签名消息（EIP-191、 EIP-712）
	// POST: /v1/waas/mpc/web3/sign_message
	SignMessage(param *common.SignMessageParam) (*common.WaaSMessageHashResult, error)
	PageAvailableVouts(param *common.PageAvailableVoutsParam) (*common.PageAvailableVoutsRes, error)

	SetDelegateEnergy(param *common.SetDelegateEnergyReq) (*common.Response, error)
}

type AdvanceAPI interface {
	SinohopeWaaSAPI

	// SignRawData 原始数据签名
	// POST: /v1/waas/mpc/wallet/advance/sign_raw_data
	SignRawData(param *common.WaaSSignRawDataParam) (*common.WaaSSignRawDataRes, error)

	// GenAddressByPath 根据指定的路径创建地址
	// POST: /v1/waas/mpc/wallet/advance/gen_address_by_path
	GenAddressByPath(param *common.WaaSAddressPathParam) (*common.WaaSAddressInfoData, error)

	// UpdateWallet 更新钱包属性（高级功能开启、关闭）
	// POST: /v1/waas/mpc/wallet/advance/update_wallet
	UpdateWallet(param *common.WaaSUpdateWalletParam) (*common.WaaSUpdateWalletRes, error)
}

type MPCNodeAPI interface {
	SinohopeWaaSAPI

	// ListMPCRequests 查询mpc协议执行记录
	// POST: /v1/waas/mpc/mpcnode/list_mpc_requests
	ListMPCRequests(param *common.WaasMpcNodeExecRecordParam) (*common.WaaSMPCNodeRequestRes, error)

	// Status 查询MPC node状态
	// POST: /v1/waas/mpc/mpcnode/status
	Status() (*common.WaaSMpcNodeStatusDTOData, error)
}

type Brc20Api interface {
	SinohopeWaaSAPI

	InscribeDeploy(param *common.InscribeDeployParam) (*common.InscribeRes, error)
	InscribeMint(param *common.InscribeMintParam) (*common.InscribeRes, error)
	InscribeTransfer(param *common.InscribeTransferParam) (*common.InscribeRes, error)
	InscribeTransferById(param *common.TransferByIdParam) (*common.InscribeRes, error)
	OneStopTransfer(param *common.OneStopTransferParam) (*common.InscribeRes, error)
	QueryInscribeTransfers(param *common.WaasBrc20QueryInscribeTransferReq) (*common.WaasBrc20QueryInscribeTransfersRes, error)
	QueryPageBalanceSummary(param *common.WaasBrc20PageQueryBalanceSummaryReq) (*common.WaasBrc20PageQueryBalanceSummaryRes, error)
	AddressTickerInfo(param *common.WaasBrc20QueryAddressTickerInfoReq) (*common.WaasBrc20QueryAddressTickerInfoRes, error)
	AddressBalance(param *common.WaasBrc20QueryAddressBalanceReq) (*common.WaasBrc20QueryAddressBalanceRes, error)
}

type RuneApi interface {
	Transfer(param *common.RuneTransferParam) (*common.InscribeRes, error)
	PageBalanceSummary(param *common.RunePageBalanceSummaryParam) (*common.RunePageBalanceSummaryRes, error)
	QueryBalance(param *common.RuneQueryBalanceParam) (*common.RuneQueryBalanceRes, error)
}
type StakingApi interface {
	Create(param *common.StakingCreateParam) (*common.StakingRes, error)
	Unbond(param *common.UnbondCreateParam) (*common.StakingRes, error)
	SpendingTimeLockPathTx(param *common.SpendingTimeLockPathTxParam) (*common.StakingRes, error)
}

type SweepApi interface {
	CollectionStrategy(param *common.CollectionStrategyReq) (*common.Response, error)
	CollectionStrategyLists(param *common.CollectionStrategyListsReq) (*common.PageCollectStrategyRes, error)
	SetGasStation(param *common.SetGasStationReq) (*common.Response, error)
	GasStationLists(param *common.GasStationListsReq) ([]*common.GasStationRes, error)
}

type OpenAPI interface {
	CreateWallets(param *common.WaaSCreateBatchWalletParam) ([]*common.WaaSWalletInfoData, error)
	ListWallets(param *common.WaaSListWalletsParam) (*common.WaaSListWalletsResult, error)
	GenerateChainAddresses(param *common.WaaSGenerateChainAddressParam) ([]*common.WaaSAddressInfoData, error)
	ListAddress(param *common.WaaSListAddressesParam) (*common.WaaSListAddressesResult, error)
	ListAddedChains(param *common.WaaSListAddedChainsParam) ([]*common.WaaSListAddedChainsDTOData, error)
	GetAddressBalance(param *common.WaaSGetAddressBalanceParam) (*common.WaaSGetWalletBalanceDTOData, error)
	GetSupportedChains() ([]*common.WaasChainData, error)
	GetSupportedCoins(param *common.WaasChainParam) ([]*common.WaaSCoinDTOData, error)
	GetVaults() ([]*common.WaaSVaultInfoData, error)
	Fee(param *common.WalletTransactionFeeWAASParam) (*common.WalletTransactionFeeWAASResponse, error)
	CreateTransfer(param *common.WalletTransactionSendWAASParam) (*common.CreateSettlementTxResData, error)
	SpeedupTransaction(param *common.WalletTransactionSpeedupWAASParam) (*common.CreateSettlementTxResData, error)
	CancelTransaction(param *common.WalletTransactionCancelWAASParam) (*common.CreateSettlementTxResData, error)
	TransactionsByRequestIds(param *common.WalletTransactionQueryWAASRequestIdParam) (*common.TransferHistoryWAASDTO, error)
	TransactionsBySinoIds(param *common.WalletTransactionQueryBySinoIdParam) (*common.TransferHistoryWAASDTO, error)
	TransactionsByTxHash(param *common.WalletTransactionQueryWAASTxHashdParam) (*common.TransferHistoryWAASDTO, error)
	ListTransactions(param *common.WalletTransactionQueryWAASParam) (*common.TransferHistoryWAASDTO, error)
	IsValidAddress(param *common.WaaSAddressCheckParam) (*common.WaaSAddressCheckDTOData, error)
}

type ConfigAPI interface {
	SetTransferStrategy(param *common.SetTransferStrategyReq) (*common.Response, error)
}
