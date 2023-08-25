package common

import "encoding/json"

const (
	FakePublicKey  = "3059301306072a8648ce3d020106082a8648ce3d0301070342000421627f551985a685e00188665a6a2f4fa80d5569fd53b31471785aad62fc9b2e23e8a9ba8a3d03b90d2c8fa5fad26fb2a05a9e7477a1ee5e4228bd85bd660309"
	FakePrivateKey = "308193020100301306072a8648ce3d020106082a8648ce3d030107047930770201010420d59deed8651f9dc130ce12c7ce9ddbda1129a2dc0d57c6e42596188c041a9aa8a00a06082a8648ce3d030107a1440342000421627f551985a685e00188665a6a2f4fa80d5569fd53b31471785aad62fc9b2e23e8a9ba8a3d03b90d2c8fa5fad26fb2a05a9e7477a1ee5e4228bd85bd660309"
	BaseUrl        = "http://mpc-exchange-demo.mpc.qa1.newhuoapps.com"

	BizApiKey       = "BIZ-API-KEY"
	BizApiNone      = "BIZ-API-NONCE"
	BizApiSignature = "BIZ-API-SIGNATURE"

	MPCProxyStatusOk = 200
	MPCProxySuccess  = true
)

type Response struct {
	Code      int             `json:"code,omitempty"`
	Message   string          `json:"msg,omitempty"`
	Success   bool            `json:"success,omitempty"`
	RequestID string          `json:"requestId,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
}

// Response data

type WaasChainData struct {
	ChainName   string `json:"chainName,omitempty"`   // 链全名
	ChainSymbol string `json:"chainSymbol,omitempty"` // 链名称 简称 链标识 具有唯一性
}

type WaaSCoinDTOData struct {
	AssetName    string `json:"assetName,omitempty"`    // 币种全称
	AssetId      string `json:"assetId,omitempty"`      // 币标识 具有唯一性
	AssetDecimal int    `json:"assetDecimal,omitempty"` // 币种精度
}

type WaaSVaultInfoData struct {
	BusinessType           int                  `json:"businessType,omitempty"`
	VaultInfoOfOpenApiList []VaultInfoOfOpenApi `json:"vaultInfoOfOpenApiList,omitempty"`
}

type WaasWalletInfoData struct {
	WalletId        string `json:"walletId,omitempty"`        // 钱包id
	WalletName      string `json:"walletName,omitempty"`      // 钱包名称
	AdvancedEnabled int    `json:"advancedEnabled,omitempty"` // 高级功能开关 (关：0，开：1)
}

type PageData struct {
	PageIndex int             `json:"pageIndex,omitempty"`
	PageSize  int             `json:"pageSize,omitempty"`
	TotalSize int             `json:"totalSize,omitempty"`
	List      json.RawMessage `json:"list,omitempty"`
}

type WaaSAddressInfoData struct {
	ID       string `json:"id,omitempty"`       // 地址id
	Address  string `json:"address,omitempty"`  // 链地址
	HdPath   string `json:"hdPath,omitempty"`   // 地址对应的path
	Encoding int    `json:"encoding,omitempty"` // 地址类型（例：BTC链：Legacy（P2PKH）格式：0，Native SegWit（Bech32）格式：2）
	Pubkey   string `json:"pubkey,omitempty"`   // 公钥
}

type WaaSAddressData struct {
	ID       string `json:"id,omitempty"`       // 地址id
	Address  string `json:"address,omitempty"`  // 链地址
	HdPath   string `json:"hdPath,omitempty"`   // 地址对应的path
	Encoding int    `json:"encoding,omitempty"` // 地址类型（例：BTC链：Legacy（P2PKH）格式：0，Native SegWit（Bech32）格式：2）
	Pubkey   string `json:"pubkey,omitempty"`   // 公钥
}

type WaaSWalletInfoData struct {
	WalletId        string `json:"walletId,omitempty"`        // 钱包id
	WalletName      string `json:"walletName,omitempty"`      // 钱包名称
	AdvancedEnabled int    `json:"advancedEnabled,omitempty"` // 高级功能开关 (关：0，开：1)
}

type WaaSListAddedChainsDTOData struct {
	ChainSymbol      string          `json:"chainSymbol,omitempty"`      // 链名称 简称 链标识
	FirstAddressInfo WaaSAddressData `json:"firstAddressInfo,omitempty"` // 地址
}

type WaaSGetWalletBalanceDTOData struct {
	ChainSymbol  string `json:"chainSymbol,omitempty"`  // 链名称 简称 链标识
	AssetId      string `json:"assetId,omitempty"`      // 币名称 简称 币标识
	Balance      string `json:"balance,omitempty"`      // 金额
	AssetDecimal int    `json:"assetDecimal,omitempty"` // 币种精度
}

type WaaSAddressCheckDTOData struct {
	IsValid bool `json:"isValid,omitempty"` // 检查结果
}

type WaaSTransferAddressBookDTOData struct {
	Address string `json:"address,omitempty"` // 地址
}

type WaaSTransferAddressSwitchDTOData struct {
	RiskControlEnabled bool `json:"riskControlEnabled,omitempty"` // 风控开关 (true 开启, false 关闭)
}

type CreateSettlementTxResData struct {
	SinoId        int64  `json:"sinoId,omitempty"`
	ChainSymbol   string `json:"chainSymbol,omitempty"`
	TxType        string `json:"txType,omitempty"`
	HasToBeSigned string `json:"hasToBeSigned,omitempty"`
	Data          TxInfo `json:"data,omitempty"`
	Hdpath        string `json:"hdpath,omitempty"`
	Cryptography  string `json:"cryptography,omitempty"`
}

type TxInfo struct {
	Asset             string     `json:"asset,omitempty"`
	From              string     `json:"from,omitempty"`
	FromTag           string     `json:"fromTag,omitempty"`
	To                string     `json:"to,omitempty"`
	ToTag             string     `json:"toTag,omitempty"`
	Amount            string     `json:"amount,omitempty"`
	Fee               string     `json:"fee,omitempty"`
	TxType            int        `json:"txType,omitempty"`
	FeeAsset          string     `json:"feeAsset,omitempty"`
	FeePrice          string     `json:"feePrice,omitempty"`
	FeeStep           string     `json:"feeStep,omitempty"`
	ChainId           string     `json:"chainId,omitempty"`
	Nonce             string     `json:"nonce,omitempty"`
	Decimal           int        `json:"decimal,omitempty"`
	CurrentTime       int64      `json:"currentTime,omitempty"`
	ExpireTime        int64      `json:"expireTime,omitempty"`
	ExpireBlockHeight int64      `json:"expireBlockHeight,omitempty"`
	ReferBlockHeight  int64      `json:"referBlockHeight,omitempty"`
	Vin               []VinItem  `json:"vin,omitempty"`
	Vout              []VoutItem `json:"vout,omitempty"`
	FunName           string     `json:"funName,omitempty"`
	Params            []string   `json:"params,omitempty"`
}

type VinItem struct {
	Id        int64  `json:"id,omitempty"`
	Hash      string `json:"hash,omitempty"`
	VoutIndex string `json:"voutIndex,omitempty"`
	Address   string `json:"address,omitempty"`
	Amount    string `json:"amount,omitempty"`
	Asset     string `json:"asset,omitempty"`
}

type VoutItem struct {
	Asset   string `json:"asset,omitempty"`
	Address string `json:"address,omitempty"`
	Amount  string `json:"amount,omitempty"`
}

type BaseWaasParam struct {
	PageIndex int `json:"pageIndex,omitempty"`
	PageSize  int `json:"pageSize,omitempty"`
	Offset    int `json:"offset,omitempty"`
}

type TransferHistoryWAASDTO struct {
	BaseWaasParam

	SinoId       int64           `json:"sinoId,omitempty"`
	AssetId      string          `json:"assetId,omitempty"`
	AssetName    string          `json:"assetName,omitempty"`
	ChainSymbol  string          `json:"chainSymbol,omitempty"`
	ChainName    string          `json:"chainName,omitempty"`
	TxDirection  string          `json:"txDirection,omitempty"`
	TxType       string          `json:"txType,omitempty"`
	Decimal      int             `json:"decimal,omitempty"`
	CoinDecimals int             `json:"coinDecimals,omitempty"`
	Transaction  TransactionWAAS `json:"transaction,omitempty"`
	UpdateTime   string          `json:"updateTime,omitempty"`
}

type TransactionWAAS struct {
	TxHash            string                          `json:"txHash,omitempty"`
	From              string                          `json:"from,omitempty"`
	FromTag           string                          `json:"fromTag,omitempty"`
	To                string                          `json:"to,omitempty"`
	ToTag             string                          `json:"toTag,omitempty"`
	Amount            string                          `json:"amount,omitempty"`
	FeeAsset          string                          `json:"feeAsset,omitempty"`
	FeePrice          string                          `json:"feePrice,omitempty"`
	FeeStep           string                          `json:"feeStep,omitempty"`
	State             int                             `json:"state,omitempty"`
	Nonce             string                          `json:"nonce,omitempty"`
	Timestamp         int64                           `json:"timestamp,omitempty"`
	Note              string                          `json:"note,omitempty"`
	Fee               float64                         `json:"fee,omitempty"`
	TransactionFee    float64                         `json:"transactionFee,omitempty"`
	Asset             string                          `json:"asset,omitempty"`
	Decimal           int                             `json:"decimal,omitempty"`
	CurrentTime       int64                           `json:"currentTime,omitempty"`
	ExpireTime        int64                           `json:"expireTime,omitempty"`
	ExpireBlockHeight int64                           `json:"expireBlockHeight,omitempty"`
	ReferBlockHeight  int64                           `json:"referBlockHeight,omitempty"`
	ChainId           string                          `json:"chainId,omitempty"`
	Vin               []CreateSettlementTxResVinItem  `json:"vin,omitempty"`
	Vout              []CreateSettlementTxResVoutItem `json:"vout,omitempty"`
}

type CreateSettlementTxResVinItem struct {
	ID        int64  `json:"id,omitempty"`
	Hash      string `json:"hash,omitempty"`
	VoutIndex string `json:"voutIndex,omitempty"`
	Address   string `json:"address,omitempty"`
	Amount    string `json:"amount,omitempty"`
	Asset     string `json:"asset,omitempty"`
}

type CreateSettlementTxResVoutItem struct {
	Asset   string `json:"asset,omitempty"`
	Address string `json:"address,omitempty"`
	Amount  string `json:"amount,omitempty"`
}

type WaaSMpcNodeExecRecordDTO struct {
	RequestId          string `json:"requestId,omitempty"`          // 用户发起业务自己生成的唯一请求id
	SinoId             string `json:"sinoId,omitempty"`             // sinoId sinohope 根据当前业务生成的唯一id
	RequestTime        string `json:"requestTime,omitempty"`        // 请求时间 格式 "2022-02-02 00:00:00"
	Param              string `json:"param,omitempty"`              // 参数
	Result             string `json:"result,omitempty"`             // 失败原因
	BusinessExecType   int    `json:"businessExecType,omitempty"`   // 业务执行类型（KeyGen 类型请求：1，signTx 类型请求：2，signMessage 类型请求：3，signRawData 类型请求：4）
	BusinessExecStatus int    `json:"businessExecStatus,omitempty"` // 业务执行状态 (进行中：0，成功：1，失败：2）
	FailedReason       string `json:"failedReason,omitempty"`       // 失败原因
}

type WaaSMpcNodeStatusDTOData struct {
	NodeId       string `json:"nodeId,omitempty"`       // 节点id
	OnlineStatus int    `json:"onlineStatus,omitempty"` // 在线状态 (offline：0，online：1) MpcNode 和 sinohope的连接状态
}

type WaaSSignatureData struct {
	SinoId      string `json:"sinoId,omitempty"`
	MessageHash string `json:"messageHash,omitempty"`
}

// Request params

type WaasChainParam struct {
	ChainName   string `json:"chainName,omitempty"`   // 链全名
	ChainSymbol string `json:"chainSymbol,omitempty"` // 链名称 简称 链标识 具有唯一性
}

type WaaSCoinDTOParam struct {
	AssetName    string `json:"assetName,omitempty"`    // 币种全称
	AssetId      string `json:"assetId,omitempty"`      // 币标识 具有唯一性
	AssetDecimal int    `json:"assetDecimal,omitempty"` // 币种精度
}

type WaaSVaultInfoParam struct {
	BusinessType           int                  `json:"businessType,omitempty"`
	VaultInfoOfOpenApiList []VaultInfoOfOpenApi `json:"vaultInfoOfOpenApiList,omitempty"`
}

type VaultInfoOfOpenApi struct {
	VaultId       string `json:"vaultId,omitempty"`       // 金库id
	VaultName     string `json:"vaultName,omitempty"`     // 金库名称
	CreateTime    string `json:"createTime,omitempty"`    // 创建时间
	AuthorityType string `json:"authorityType,omitempty"` // 权限 查询 1、操作 2、提现 3、任意交易0
}

type WaaSCreateBatchWalletParam struct {
	Count       int                    `json:"count,omitempty"` // 创建多少个 默认值为 1
	VaultId     string                 `json:"vaultId,omitempty"`
	WalletInfos []WaaSCreateWalletInfo `json:"walletInfos,omitempty"` // 钱包信息 如果不为空 count值必须 walletInfos的数量一致
	RequestId   string                 `json:"requestId,omitempty"`
}

type WaaSCreateWalletInfo struct {
	WalletName      string `json:"walletName,omitempty"`      // 钱包名称 为空的话用默认值: wallet+随机值 不为空时：部门下钱包名称不能给重复
	AdvancedEnabled int    `json:"advancedEnabled,omitempty"` // 高级功能开关 (关：0，开：1) 开了以后支持的功能: 根据指定的路径创建地址 ,原始数据签名
}

type WaaSListAddressesParam struct {
	PageIndex   int    `json:"pageIndex,omitempty"`   // 当前页码，首页为0,默认0
	PageSize    int    `json:"pageSize,omitempty"`    // 每页数据条数（不得小于1,不得大于50）
	WalletId    string `json:"walletId,omitempty"`    // 钱包id
	ChainSymbol string `json:"chainSymbol,omitempty"` // 链名称 简称 链标识
}

type WaaSGenerateChainAddressParam struct {
	RequestId   string `json:"requestId,omitempty"`   // 请求ID
	WalletId    string `json:"walletId,omitempty"`    // 钱包id
	Count       int    `json:"count,omitempty"`       // 创建多少个,不传默认为1
	ChainSymbol string `json:"chainSymbol,omitempty"` // 链名称 简称 链标识 具有唯一性
}

type WaaSListWalletsParam struct {
	VaultId   string `json:"vaultId,omitempty"`   // 金库ID
	PageIndex int    `json:"pageIndex,omitempty"` // 当前页码，首页为0,默认0
	PageSize  int    `json:"pageSize,omitempty"`  // 每页数据条数（不得小于1,不得大于50）
}

type WaaSListAddedChainsParam struct {
	WalletId string `json:"walletId,omitempty"` // 钱包id
}

type WaaSGetAddressBalanceParam struct {
	AssetId string `json:"assetId,omitempty"` // 币名称 简称 币标识 具有唯一性
	Address string `json:"address,omitempty"` // 地址
}

type WaaSGetWalletBalanceParam struct {
	WalletId  string `json:"walletId,omitempty"`  // 钱包id
	PageIndex int    `json:"pageIndex,omitempty"` // 当前页码，首页为0,默认0
	PageSize  int    `json:"pageSize,omitempty"`  // 每页数据条数（不得小于1,不得大于50）
}

type WaaSAddressCheckParam struct {
	AssetId string `json:"assetId,omitempty"` // 币种代号 币标识 具有唯一性
	Address string `json:"address,omitempty"` // 地址
}

type WaaSTransferAddressBookParam struct {
	ChainSymbol string `json:"chainSymbol,omitempty"` // 链名称 简称 链标识 具有唯一性
	PageIndex   int    `json:"pageIndex,omitempty"`   // 当前页码，首页为0,默认0
	PageSize    int    `json:"pageSize,omitempty"`    // 每页数据条数（不得小于1,不得大于50）
}

type WaaSVaultIdDTO struct {
	VaultId string `json:"vaultId,omitempty"` // 金库id
}

type WalletTransactionSendWAASParam struct {
	WalletId    string `json:"walletId,omitempty"`    // 钱包id
	RequestId   string `json:"requestId,omitempty"`   // 请求方交易的requestId
	ChainSymbol string `json:"chainSymbol,omitempty"` // 链标识
	AssetId     string `json:"assetId,omitempty"`     // 资产id
	From        string `json:"from,omitempty"`        // from 地址
	To          string `json:"to,omitempty"`          // to地址
	ToTag       string `json:"toTag,omitempty"`       // memo
	Amount      string `json:"amount,omitempty"`      // 金额
	Fee         string `json:"fee,omitempty"`         // 手续费 对于 UTXO 类的非EVM兼容链的交易,自设置fee, 如参数为 UTXO 资产转账提供，表示每字节的手续费
	GasPrice    string `json:"gasPrice,omitempty"`    // 交易gasPrice，燃料价格，ETH 账号模型适用，单位为 wei
	GasLimit    string `json:"gasLimit,omitempty"`    // 交易gasLimit，燃料上限，ETH 账号模型适用
	Remark      string `json:"remark,omitempty"`      // 备注：用于用户自己需要的一些备注信息
	VaultId     string `json:"vaultId,omitempty"`     // 金库id
}

type WalletTransactionSendDataWAASParam struct {
	WalletId    string `json:"walletId,omitempty"`    // 钱包id
	RequestId   string `json:"requestId,omitempty"`   // 请求方交易的requestId
	ChainSymbol string `json:"chainSymbol,omitempty"` // 链标识
	AssetId     string `json:"assetId,omitempty"`     // 资产id
	From        string `json:"from,omitempty"`        // from 地址
	To          string `json:"to,omitempty"`          // to地址
	ToTag       string `json:"toTag,omitempty"`       // 交易的memo
	Amount      string `json:"amount,omitempty"`      // 金额
	Fee         string `json:"fee,omitempty"`         // 手续费 对于 UTXO 类的非EVM兼容链的交易,自设置fee, 如参数为 UTXO 资产转账提供，表示每字节的手续费
	GasPrice    string `json:"gasPrice,omitempty"`    // gasprice
	GasLimit    string `json:"gasLimit,omitempty"`    // gaslimit
	Remark      string `json:"remark,omitempty"`      // 备注：用于用户自己需要的一些备注信息
	InputData   string `json:"inputData,omitempty"`   // 以太坊交易data
	VaultId     string `json:"vaultId,omitempty"`     // 金库id
}

type WalletTransactionSpeedupWAASParam struct {
	RequestId   string `json:"requestId,omitempty"`   // 请求方交易的requestId
	GasLimit    string `json:"gasLimit,omitempty"`    // 交易gasLimit，燃料上限，ETH 账号模型适用
	GasPrice    string `json:"gasPrice,omitempty"`    // 交易gasPrice，燃料价格，ETH 账号模型适用，单位为 wei
	Fee         string `json:"fee,omitempty"`         // 手续费 对于 UTXO 类的非EVM兼容链的交易,自设置fee, 如参数为 UTXO 资产转账提供，表示每字节的手续费
	ChainSymbol string `json:"chainSymbol,omitempty"` // 链标识
	AssetId     string `json:"assetId,omitempty"`     // 资产id
}

type WalletTransactionCancelWAASParam struct {
	RequestId   string `json:"requestId,omitempty"`   // 请求方的requestId
	ChainSymbol string `json:"chainSymbol,omitempty"` // 链标识
	AssetId     string `json:"assetId,omitempty"`     // 资产id
	GasLimit    string `json:"gasLimit,omitempty"`    // 交易gasLimit，燃料上限，ETH 账号模型适用
	GasPrice    string `json:"gasPrice,omitempty"`    // 交易gasPrice，燃料价格，ETH 账号模型适用，单位为 wei
}

type WalletTransactionQueryWAASParam struct {
	PageIndex       int    `json:"pageIndex,omitempty"`       // 当前页码，首页为0,默认0
	PageSize        int    `json:"pageSize,omitempty"`        // 每页数据条数（不得小于1,不得大于50）
	Address         string `json:"address,omitempty"`         // 链地址
	SinoIds         string `json:"sinoIds,omitempty"`         // sinoId生成的交易id
	RequestIds      string `json:"requestIds,omitempty"`      // 请求方交易的requestId
	TxHash          string `json:"txHash,omitempty"`          // 交易hash
	ChainSymbol     string `json:"chainSymbol,omitempty"`     // 链标识
	AssetId         string `json:"assetId,omitempty"`         // 资产id
	StartUpdateTime string `json:"startUpdateTime,omitempty"` // 根据更新时间查询,开始时间 传了开始时间,开始结束也得带上 格式 "2022-02-02 00:00:00"
	EndUpdateTime   string `json:"endUpdateTime,omitempty"`   // 根据更新时间查询,结束时间  传了结束时间,开始时间也得带上 格式 "2022-02-02 00:00:00"
}

type WalletTransactionQueryWAASRequestIdParam struct {
	RequestIds string `json:"requestIds,omitempty"` // 通过requestIds查询获取已确认交易记录列表，requestId以逗号分割，不能为空且不能大于50个
}

type WalletTransactionQueryBySinoIdParam struct {
	SinoId string `json:"sinoId,omitempty"` // 通过sinoId查询获取已确认交易记录列表
}

type WalletTransactionQueryWAASTxHashdParam struct {
	TxHash      string `json:"txHash,omitempty"`      // 通过交易txHash，查询获取已确认交易记录列表
	ChainSymbol string `json:"chainSymbol,omitempty"` // 链标识
}

type SignMessageParam struct {
	RequestId     string `json:"requestId,omitempty"`     // 请求方交易的requestId
	ChainSymbol   string `json:"chainSymbol,omitempty"`   // 链
	AssetId       string `json:"assetId,omitempty"`       // 资产id
	HdPath        string `json:"hdPath,omitempty"`        // bip32、bip44的推导路径
	SignAlgorithm string `json:"signAlgorithm,omitempty"` // 支持签名算法: （personal_sign、signTypedData、eth_signTypedData_v3、eth_signTypedData_v4）
	Message       string `json:"message,omitempty"`       // 待签名的字符串信息
}

type WaasSignRawDataParam struct {
	RawData   string `json:"rawData,omitempty"`   // 签名数据
	RequestId string `json:"requestId,omitempty"` // 唯一id 用户自己生成的请求唯一id, 用于重试
	WalletId  string `json:"walletId,omitempty"`  // 钱包id
	HdPath    string `json:"hdPath,omitempty"`    // 地址对应的path eth 示例 m/1/1/60/0
}

type WaasAddressPathParam struct {
	Index         int    `json:"index,omitempty"`         // 用于区分同一个钱包的同一个cointype 下的不同地址
	AlgorithmType int    `json:"algorithmType,omitempty"` // 算法类型(0: "ECDSA:secp256k1", 1: "EdDSA:ed25519")
	CoinType      int    `json:"coinType,omitempty"`      // 参考slip-44 https://github.com/satoshilabs/slips/blob/master/slip-0044.md 中定义的coin type常量，使用none-hardened 的版本。根据业界常规做法，对于所有 eth-like 公链，可公钥以太坊的 coin type 60
	WalletId      string `json:"walletId,omitempty"`      // 钱包id
}

type WaasUpdateWalletParam struct {
	WalletId        string `json:"walletId,omitempty"`        // 钱包id 当指定了以后就是根据钱包开关,否则就是金库级别
	AdvancedEnabled int    `json:"advancedEnabled,omitempty"` // 高级功能开关 (关：0，开：1) 开了以后支持的功能: 根据指定的路径创建地址 ,原始数据签名
}

type WaasMpcNodeExecRecordParam struct {
	BusinessExecType   int    `json:"businessExecType,omitempty"`   // 业务执行类型（KeyGen 类型请求：1，signTx 类型请求：2，signMessage 类型请求：3，signRawData 类型请求：4）
	BusinessExecStatus int    `json:"businessExecStatus,omitempty"` // 业务执行状态 (进行中：0，成功：1，失败：2）
	SinoId             string `json:"sinoId,omitempty"`             // sinoId,不传按照分页查询
	PageIndex          int    `json:"pageIndex,omitempty"`          // 当前页码，首页为0,默认0
	PageSize           int    `json:"pageSize,omitempty"`           // 每页数据条数（不得小于1,不得大于50）
}
