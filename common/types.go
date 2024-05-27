package common

import "encoding/json"

const (
	FakePublicKey  = "3059301306072a8648ce3d020106082a8648ce3d0301070342000421627f551985a685e00188665a6a2f4fa80d5569fd53b31471785aad62fc9b2e23e8a9ba8a3d03b90d2c8fa5fad26fb2a05a9e7477a1ee5e4228bd85bd660309"
	FakePrivateKey = "308193020100301306072a8648ce3d020106082a8648ce3d030107047930770201010420d59deed8651f9dc130ce12c7ce9ddbda1129a2dc0d57c6e42596188c041a9aa8a00a06082a8648ce3d030107a1440342000421627f551985a685e00188665a6a2f4fa80d5569fd53b31471785aad62fc9b2e23e8a9ba8a3d03b90d2c8fa5fad26fb2a05a9e7477a1ee5e4228bd85bd660309"
	BaseUrl        = "https://api.sinohope.com"
	//BaseUrl         = "https://api-sandbox-qa1.newhuoapps.com/"
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

type WaaSWalletInfoData struct {
	WalletId        string `json:"walletId,omitempty"`        // 钱包id
	WalletName      string `json:"walletName,omitempty"`      // 钱包名称
	AdvancedEnabled int    `json:"advancedEnabled,omitempty"` // 高级功能开关 (关：0，开：1)
}

type WaaSListWalletsResult struct {
	BaseWaasParam
	List []*WaaSWalletInfoData `json:"list,omitempty"` // wallets
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

type WaaSListAddedChainsDTOData struct {
	ChainSymbol      string              `json:"chainSymbol,omitempty"`      // 链名称 简称 链标识
	FirstAddressInfo WaaSAddressInfoData `json:"firstAddressInfo,omitempty"` // 地址
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

type WaaSTransferAddressSwitchDTOData struct {
	RiskControlEnabled bool `json:"riskControlEnabled,omitempty"` // 风控开关 (true 开启, false 关闭)
}

type CreateSettlementTxResData struct {
	SinoId         string      `json:"sinoId,omitempty"`
	RequestId      string      `json:"requestId,omitempty"`
	ReplacedSinoId string      `json:"replacedSinoId,omitempty"`
	ChainSymbol    string      `json:"chainSymbol,omitempty"`
	TxDirection    int         `json:"txDirection,omitempty"`
	Transaction    TxInfo      `json:"transaction,omitempty"`
	Brc20Detail    brc20Detail `json:"brc20Detail,omitempty"`
	State          int         `json:"state,omitempty"`
	Note           string      `json:"note,omitempty"`
	FailReason     string      `json:"failReason,omitempty"`
}
type brc20Detail struct {
	Method        string `json:"method"`
	Ticker        string `json:"ticker"`
	Quantity      string `json:"quantity"`
	InscriptionId string `json:"inscriptionId,omitempty"`
	Step          int    `json:"step"`
	From          string `json:"from,omitempty"`
	To            string `json:"to,omitempty"`
}

type TxInfo struct {
	TxHash            string     `json:"txHash,omitempty"`
	ActionType        string     `json:"actionType,omitempty"`
	From              string     `json:"from,omitempty"`
	FromTag           string     `json:"fromTag,omitempty"`
	To                string     `json:"to,omitempty"`
	ToTag             string     `json:"toTag,omitempty"`
	Amount            string     `json:"amount,omitempty"`
	AssetId           string     `json:"assetId,omitempty"`
	Decimal           int        `json:"decimal,omitempty"`
	ContractAddress   string     `json:"contractAddress,omitempty"`
	FeeAsset          string     `json:"feeAsset,omitempty"`
	FeeAssetDecimal   int        `json:"feeAssetDecimal,omitempty"`
	UsedFee           string     `json:"usedFee,omitempty"`
	Fee               string     `json:"fee,omitempty"`
	GasPrice          string     `json:"gasPrice,omitempty"`
	GasLimit          string     `json:"gasLimit,omitempty"`
	Nonce             string     `json:"nonce,omitempty"`
	InputData         string     `json:"inputData,omitempty"`
	ExpireTime        int        `json:"expireTime,omitempty"`
	ExpireBlockHeight int        `json:"expireBlockHeight,omitempty"`
	ReferBlockHeight  int        `json:"referBlockHeight,omitempty"`
	Vin               []VinItem  `json:"vin,omitempty"`
	Vout              []VoutItem `json:"vout,omitempty"`
}

type VinItem struct {
	Hash      string `json:"hash,omitempty"`
	VoutIndex string `json:"voutIndex,omitempty"`
	Address   string `json:"address,omitempty"`
	Amount    string `json:"amount,omitempty"`
}

type VoutItem struct {
	Address string `json:"address,omitempty"`
	Amount  string `json:"amount,omitempty"`
}

type BaseWaasParam struct {
	PageIndex int `json:"pageIndex,omitempty"`
	PageSize  int `json:"pageSize,omitempty"`
	TotalSize int `json:"totalSize,omitempty"`
}

type TransferHistoryWAASDTO struct {
	BaseWaasParam
	List []CreateSettlementTxResData `json:"list,omitempty"`
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

type WaaSMessageHashResult struct {
	SinoId      string `json:"sinoId,omitempty"`
	MessageHash string `json:"messageHash,omitempty"`
}

// Request params

type WaasChainParam struct {
	ChainSymbol string `json:"chainSymbol"` // 链名称 简称 链标识 具有唯一性
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
	VaultId     string                 `json:"vaultId"`
	WalletInfos []WaaSCreateWalletInfo `json:"walletInfos,omitempty"` // 钱包信息 如果不为空 count值必须 walletInfos的数量一致
	RequestId   string                 `json:"requestId"`
}

type WaaSCreateWalletInfo struct {
	WalletName      string `json:"walletName,omitempty"` // 钱包名称 为空的话用默认值: wallet+随机值 不为空时：部门下钱包名称不能给重复
	AdvancedEnabled int    `json:"advancedEnabled"`      // 高级功能开关 (关：0，开：1) 开了以后支持的功能: 根据指定的路径创建地址 ,原始数据签名
}

type WaaSListAddressesParam struct {
	VaultId     string `json:"vaultId"`             // the vault id
	WalletId    string `json:"walletId"`            // 钱包id
	ChainSymbol string `json:"chainSymbol"`         // 链名称 简称 链标识
	PageIndex   int    `json:"pageIndex,omitempty"` // 当前页码，首页为0,默认0
	PageSize    int    `json:"pageSize,omitempty"`  // 每页数据条数（不得小于1,不得大于50）
}

type WaaSListAddressesResult struct {
	BaseWaasParam
	List []*WaaSAddressInfoData `json:"list,omitempty"`
}

type WaaSGenerateChainAddressParam struct {
	RequestId   string `json:"requestId"`       // 请求ID
	VaultId     string `json:"vaultId"`         // the vault id
	WalletId    string `json:"walletId"`        // 钱包id
	Count       int    `json:"count,omitempty"` // 创建多少个,不传默认为1
	ChainSymbol string `json:"chainSymbol"`     // 链名称 简称 链标识 具有唯一性
	Encoding    int    `json:"encoding"`
}

type WaaSListWalletsParam struct {
	VaultId   string `json:"vaultId"`             // 金库ID
	PageIndex int    `json:"pageIndex,omitempty"` // 当前页码，首页为0,默认0
	PageSize  int    `json:"pageSize,omitempty"`  // 每页数据条数（不得小于1,不得大于50）
}

type WaaSListAddedChainsParam struct {
	VaultId  string `json:"vaultId"`  // 金库ID
	WalletId string `json:"walletId"` // 钱包id
}

type WaaSGetAddressBalanceParam struct {
	AssetId string `json:"assetId"` // 币名称 简称 币标识 具有唯一性
	Address string `json:"address"` // 地址
}

type WaaSAddressCheckParam struct {
	AssetId string `json:"assetId"` // 币种代号 币标识 具有唯一性
	Address string `json:"address"` // 地址
}

type WaaSTransferAddressBookParam struct {
	VaultId     string `json:"vaultId"`             // 金库ID
	ChainSymbol string `json:"chainSymbol"`         // 链名称 简称 链标识 具有唯一性
	PageIndex   int    `json:"pageIndex,omitempty"` // 当前页码，首页为0,默认0
	PageSize    int    `json:"pageSize,omitempty"`  // 每页数据条数（不得小于1,不得大于50）
}

type WaaSAddressBookResult struct {
	BaseWaasParam
	List []*WhiteAddress `json:"list,omitempty"`
}
type WhiteAddress struct {
	Address string `json:"address,omitempty"`
}

type WaaSVaultIdDTO struct {
	VaultId string `json:"vaultId"` // 金库id
}

type WalletTransactionFeeWAASParam struct {
	OperationType string `json:"operationType"`       // 交易类型 必填 TRANSFER ：转账 CONTRACT_CALL ：web3合约调用
	From          string `json:"from"`                // from地址
	To            string `json:"to"`                  // to地址 必填 如果是合约调用，则填写合约地址
	AssetId       string `json:"assetId,omitempty"`   // 资产id，如果交易类型为TRANSFER必填。
	ChainSymbol   string `json:"chainSymbol"`         // 链标识
	Amount        string `json:"amount,omitempty"`    // 转账金额 如果所属链为btc like，则amount必传。
	InputData     string `json:"inputData,omitempty"` // 如果交易类型为CONTRACT_CALL合约调用时，inputData必填
}

type WalletTransactionFeeWAASResponse struct {
	TransactionFee struct {
		SlowFee    string `json:"slowFee,omitempty"`    // 低档交易费
		AverageFee string `json:"averageFee,omitempty"` // 中档交易费
		FastFee    string `json:"fastFee,omitempty"`    // 高档交易费
		FeePerByte string `json:"feePerByte,omitempty"` // 每个字节的费用 btc like链会返回
	} `json:"transactionFee,omitempty"`
	GasFee struct {
		Slow struct {
			GasPrice string `json:"gasPrice"`
			GasLimit string `json:"gasLimit"`
		} `json:"slow"`
		Average struct {
			GasPrice string `json:"gasPrice"`
			GasLimit string `json:"gasLimit"`
		} `json:"average"`
		Fast struct {
			GasPrice string `json:"gasPrice"`
			GasLimit string `json:"gasLimit"`
		} `json:"fast"`
	} `json:"gasFee,omitempty"` // 对于ETH及EVM兼容链，预估费用详情

}

type WalletTransactionSendWAASParam struct {
	VaultId     string `json:"vaultId"`            // 金库id
	WalletId    string `json:"walletId"`           // 钱包id
	RequestId   string `json:"requestId"`          // 请求方交易的requestId
	ChainSymbol string `json:"chainSymbol"`        // 链标识
	AssetId     string `json:"assetId"`            // 资产id
	From        string `json:"from"`               // from 地址
	To          string `json:"to"`                 // to地址
	ToTag       string `json:"toTag,omitempty"`    // memo
	Amount      string `json:"amount"`             // 金额
	Fee         string `json:"fee,omitempty"`      // total transaction fee
	FeeRate     string `json:"feeRate,omitempty"`  // fee per byte
	GasPrice    string `json:"gasPrice,omitempty"` // 交易gasPrice，燃料价格，ETH 账号模型适用，单位为 wei
	GasLimit    string `json:"gasLimit,omitempty"` // 交易gasLimit，燃料上限，ETH 账号模型适用
	Note        string `json:"note,omitempty"`     // 备注：用于用户自己需要的一些备注信息
	UtxoType    string `json:"utxoType,omitempty"` // UTXO类型：btc_only、all(包括铭文)，非必填，空缺情况下是btc_only。请谨慎使用all，这有可能会将有价值的铭文UTXO转走。
}

type WalletTransactionSendDataWAASParam struct {
	VaultId     string `json:"vaultId"`             // 金库id
	WalletId    string `json:"walletId"`            // 钱包id
	RequestId   string `json:"requestId"`           // 请求方交易的requestId
	ChainSymbol string `json:"chainSymbol"`         // 链标识
	From        string `json:"from"`                // from 地址
	To          string `json:"to"`                  // to地址
	ToTag       string `json:"toTag,omitempty"`     // 交易的memo
	Amount      string `json:"amount,omitempty"`    // 金额
	Fee         string `json:"fee,omitempty"`       // 手续费 对于 UTXO 类的非EVM兼容链的交易,自设置fee, 如参数为 UTXO 资产转账提供，表示每字节的手续费
	FeeRate     string `json:"fee_rate,omitempty"`  // 手续费费率 1:快 2:中 3:慢
	GasPrice    string `json:"gasPrice,omitempty"`  // gasprice
	GasLimit    string `json:"gasLimit,omitempty"`  // gaslimit
	Note        string `json:"note,omitempty"`      // 备注：用于用户自己需要的一些备注信息
	InputData   string `json:"inputData,omitempty"` // 以太坊交易data
}

type WalletTransactionSpeedupWAASParam struct {
	RequestId   string `json:"requestId"`             // 当前请求的requestId
	SinoId      string `json:"sinoId"`                // 要加速的交易的sinoId
	ChainSymbol string `json:"chainSymbol,omitempty"` // 链标识
	GasLimit    string `json:"gasLimit,omitempty"`    // 交易gasLimit，燃料上限，ETH 账号模型适用
	GasPrice    string `json:"gasPrice,omitempty"`    // 交易gasPrice，燃料价格，ETH 账号模型适用，单位为 wei
	Fee         string `json:"fee,omitempty"`         // 手续费 对于 UTXO 类的非EVM兼容链的交易,自设置fee, 如参数为 UTXO 资产转账提供，表示每字节的手续费
}

type WalletTransactionCancelWAASParam struct {
	RequestId   string `json:"requestId"`             // 当前请求的requestId
	SinoId      string `json:"sinoId"`                // 要加速的交易的sinoId
	ChainSymbol string `json:"chainSymbol,omitempty"` // 链标识
}

type WalletTransactionQueryWAASParam struct {
	PageIndex       int    `json:"pageIndex,omitempty"`       // 当前页码，首页为 1
	PageSize        int    `json:"pageSize,omitempty"`        // 每页数据条数（不得小于1,不得大于50；默认 10）
	Address         string `json:"address"`                   // 链地址
	SinoIds         string `json:"sinoIds,omitempty"`         // 通过sinoIds查询获取已确认交易记录列表，sinoid是sinohope唯一标识交易id，以逗号分割，sinoid不能大于50个
	RequestIds      string `json:"requestIds,omitempty"`      // 通过requestIds查询获取已确认交易记录列表，requestId以逗号分割，requestId不能大于50个
	TxHash          string `json:"txHash,omitempty"`          // 交易hash
	ChainSymbol     string `json:"chainSymbol"`               // 链标识
	AssetId         string `json:"assetId,omitempty"`         // 资产id
	StartUpdateTime string `json:"startUpdateTime,omitempty"` // 根据更新时间查询,开始时间 传了开始时间,开始结束也得带上 格式 "2022-02-02 00:00:00"
	EndUpdateTime   string `json:"endUpdateTime,omitempty"`   // 根据更新时间查询,结束时间  传了结束时间,开始时间也得带上 格式 "2022-02-02 00:00:00"
}

type WalletTransactionQueryWAASRequestIdParam struct {
	RequestIds string `json:"requestIds"` // 通过requestIds查询获取已确认交易记录列表，requestId以逗号分割，不能为空且不能大于50个
}

type WalletTransactionQueryBySinoIdParam struct {
	SinoIds string `json:"sinoIds"` // 通过sinoId查询获取已确认交易记录列表，sinoid是sinohope唯一标识交易id，以逗号分割，不能为空且不能大于50个
}

type WalletTransactionQueryWAASTxHashdParam struct {
	TxHash      string `json:"txHash,omitempty"`      // 通过交易txHash，查询获取已确认交易记录列表
	ChainSymbol string `json:"chainSymbol,omitempty"` // 链标识
}

type SignMessageParam struct {
	RequestId     string `json:"requestId,omitempty"`     // 请求方交易的requestId
	ChainSymbol   string `json:"chainSymbol,omitempty"`   // 链
	HdPath        string `json:"hdPath,omitempty"`        // bip32、bip44的推导路径
	SignAlgorithm string `json:"signAlgorithm,omitempty"` // 支持签名算法: （personal_sign、signTypedData、eth_signTypedData_v3、eth_signTypedData_v4）
	Message       string `json:"message,omitempty"`       // 待签名的字符串信息
}

type WaaSSignRawDataParam struct {
	VaultId   string `json:"vaultId"`   // 金库id
	RequestId string `json:"requestId"` // 唯一id 用户自己生成的请求唯一id, 用于重试
	WalletId  string `json:"walletId"`  // 钱包id
	HdPath    string `json:"hdPath"`    // 地址对应的path, eth 示例 m/1/1/60/0
	RawData   string `json:"rawData"`   // 签名数据
	/**
	 * AlgorithmType 未提供或值为0 时（当前默认），使用创建地址时所指定的算法及曲线类型（ecdsa-secp256k1/eddsa-ed25519）执行默认的签名，
	 * 值为 1 时（只对 ecdsa-secp256k1 的密钥有效）执行 符合BIP340 标准的 Schnorr 签名（支持 Bitcoin 铭文reveal交易）
	 * 值为 2 时（只对 ecdsa-secp256k1 的密钥有效）执行 Taproot 签名
	 */
	AlgorithmType int `json:"algorithmType"`
	// 非必填 Taproot签名的script树根
	TapScriptRoot string `json:"tapScriptRoot"`
}
type WaaSSignRawDataRes struct {
	SinoId string `json:"sinoId"` // sinohope 对此业务的唯一标识
}

type WaaSAddressPathParam struct {
	VaultId       string `json:"vaultId"`       // 金库id
	WalletId      string `json:"walletId"`      // 钱包id
	Index         int    `json:"index"`         // 用于区分同一个钱包的同一个cointype 下的不同地址
	AlgorithmType int    `json:"algorithmType"` // 算法类型(0: "ECDSA:secp256k1", 1: "EdDSA:ed25519")
	CoinType      int    `json:"coinType"`      // 参考slip-44 https://github.com/satoshilabs/slips/blob/master/slip-0044.md 中定义的coin type常量，使用none-hardened 的版本。根据业界常规做法，对于所有 eth-like 公链，可共用以太坊的 coin type 60
}

type WaaSUpdateWalletParam struct {
	VaultId         string `json:"vaultId"`         // 金库id
	WalletId        string `json:"walletId"`        // 钱包id 当指定了以后就是根据钱包开关,否则就是金库级别
	AdvancedEnabled int    `json:"advancedEnabled"` // 高级功能开关 (关：0，开：1) 开了以后支持的功能: 根据指定的路径创建地址 ,原始数据签名
}

type WaaSUpdateWalletRes struct {
	UpdateWallet bool `json:"updateWallet"` // 修改结果
}

type WaasMpcNodeExecRecordParam struct {
	BusinessExecType   int    `json:"businessExecType,omitempty"`   // 业务执行类型（KeyGen 类型请求：1，signTx 类型请求：2，signMessage 类型请求：3，signRawData 类型请求：4）
	BusinessExecStatus int    `json:"businessExecStatus,omitempty"` // 业务执行状态 (进行中：0，成功：1，失败：2）
	SinoId             string `json:"sinoId,omitempty"`             // sinoId,不传按照分页查询
	PageIndex          int    `json:"pageIndex,omitempty"`          // 当前页码，首页为1
	PageSize           int    `json:"pageSize,omitempty"`           // 每页数据条数（不得小于1,不得大于50）
}

type WaaSMPCNodeRequestRes struct {
	BaseWaasParam
	List []*MPCNodeExecRecord `json:"list,omitempty"`
}
type MPCNodeExecRecord struct {
	RequestID          string `json:"requestId,omitempty"`          //用户发起业务自己生成的唯一请求id
	SinoID             string `json:"sinoId,omitempty"`             // sinohope 根据当前业务生成的唯一id
	RequestTime        string `json:"requestTime,omitempty"`        // 请求时间 格式 "2022-02-02 00:00:00"
	Param              string `json:"param,omitempty"`              // 发给 MPC Node的参数
	BusinessExecType   int    `json:"businessExecType,omitempty"`   // 业务执行类型（KeyGen 类型请求：1，signTx 类型请求：2，signMessage 类型请求：3，signRawData 类型请求：4）
	BusinessExecStatus int    `json:"businessExecStatus,omitempty"` // 业务执行状态 (进行中：0，成功：1，失败：2）
	IsSuccess          bool   `json:"isSuccess,omitempty"`          // 是否成功
	FailedReason       string `json:"failedReason,omitempty"`       // 失败原因
}
