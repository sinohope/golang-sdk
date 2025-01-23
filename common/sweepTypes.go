package common

type CollectionStrategyReq struct {
	VaultId                 string `json:"vaultId"`
	ToAddress               string `json:"toAddress"`
	AssetId                 string `json:"assetId"`
	MinimumCollectionAmount int64  `json:"minimumCollectionAmount"`
	MinimumReserveAmount    int64  `json:"minimumReserveAmount,omitempty"`
	FeeLimit                int    `json:"feeLimit,omitempty"`
	Enabled                 int    `json:"enabled,omitempty"`
	AutoFueling             int    `json:"autoFueling"` //必填，0-关闭；1-开启
}

type CollectionStrategyListsReq struct {
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
	Offset    int    `json:"offset"`
	VaultId   string `json:"vaultId"`
}

type PageCollectStrategyRes struct {
	BaseWaasParam
	List []*CollectStrategyRes `json:"list,omitempty"`
}

type CollectStrategyRes struct {
	VaultId                 string `json:"vaultId"`
	ToAddress               string `json:"toAddress"`
	AssetId                 string `json:"assetId"`
	ChainSymbol             string `json:"chainSymbol"`
	MinimumCollectionAmount int64  `json:"minimumCollectionAmount"`
	MinimumReserveAmount    int64  `json:"minimumReserveAmount"`
	FeeLimit                int64  `json:"feeLimit"`
	Enabled                 int    `json:"enabled"`
	Type                    string `json:"type"`
	AutoFueling             int    `json:"autoFueling"`
}

type SetGasStationReq struct {
	ChainSymbol string `json:"chainSymbol"`
	Address     string `json:"address"`
}

type GasStationListsReq struct {
	ChainSymbol string `json:"chainSymbol"`
}

type GasStationRes struct {
	ChainSymbol string `json:"chainSymbol"`
	Address     string `json:"address"`
}
