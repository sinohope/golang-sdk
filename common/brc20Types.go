package common

type InscribeDeployParam struct {
	RequestId    string `json:"requestId"`
	Ticker       string `json:"ticker"`
	TotalSupply  string `json:"totalSupply"`
	LimitPerMint string `json:"limitPerMint"`
	Decimals     string `json:"decimals"`
	ChainSymbol  string `json:"chainSymbol"`
	From         string `json:"from"`
}

type InscribeRes struct {
	SinoId string `json:"sinoId"`
}
type InscribeMintParam struct {
	RequestId   string `json:"requestId"`
	Ticker      string `json:"ticker"`
	Amount      string `json:"amount"`
	ChainSymbol string `json:"chainSymbol"`
	From        string `json:"from"`
}

type InscribeTransferParam struct {
	RequestId   string `json:"requestId"`
	Ticker      string `json:"ticker"`
	Amount      string `json:"amount"`
	ChainSymbol string `json:"chainSymbol"`
	From        string `json:"from"`
}

type TransferByIdParam struct {
	RequestId     string `json:"requestId"`
	ChainSymbol   string `json:"chainSymbol"`
	From          string `json:"from"`
	To            string `json:"to"`
	InscriptionId string `json:"inscriptionId"`
	Ticker        string `json:"ticker"`
}

type OneStopTransferParam struct {
	RequestId   string `json:"requestId"`
	ChainSymbol string `json:"chainSymbol"`
	Amount      string `json:"amount"`
	From        string `json:"from"`
	To          string `json:"to"`
	Ticker      string `json:"ticker"`
}

type WaasBrc20QueryInscribeTransferReq struct {
	Ticker      string `json:"ticker"`
	ChainSymbol string `json:"chainSymbol"`
	Address     string `json:"address"`
}

type WaasBrc20QueryInscribeTransfersRes struct {
	Height int `json:"height"`
	Total  int `json:"total"`
	Start  int `json:"start"`
	Detail []struct {
		Data struct {
			Op      string `json:"op"`
			Tick    string `json:"tick"`
			Lim     string `json:"lim"`
			Amt     string `json:"amt"`
			Decimal string `json:"decimal"`
		} `json:"data"`
		InscriptionId string `json:"inscriptionId"`
		Satoshi       string `json:"satoshi"`
		Confirmations int    `json:"confirmations"`
	} `json:"detail"`
}

type WaasBrc20QueryAddressBalanceReq struct {
	ChainSymbol string `json:"chainSymbol"`
	Address     string `json:"address"`
}

type WaasBrc20QueryAddressBalanceRes struct {
	Address                   string      `json:"address"`
	Satoshi                   interface{} `json:"satoshi"`
	PendingSatoshi            interface{} `json:"pendingSatoshi"`
	UtxoCount                 interface{} `json:"utxoCount"`
	BtcSatoshi                interface{} `json:"btcSatoshi"`
	BtcPendingSatoshi         interface{} `json:"btcPendingSatoshi"`
	BtcUtxoCount              interface{} `json:"btcUtxoCount"`
	InscriptionSatoshi        interface{} `json:"inscriptionSatoshi"`
	InscriptionPendingSatoshi interface{} `json:"inscriptionPendingSatoshi"`
	InscriptionUtxoCount      interface{} `json:"inscriptionUtxoCount"`
}

type WaasBrc20PageQueryBalanceSummaryReq struct {
	Address     string `json:"address"`
	ChainSymbol string `json:"chainSymbol"`
	Start       int    `json:"start"`
	Limit       int    `json:"limit"`
}

type WaasBrc20PageQueryBalanceSummaryRes struct {
	Height int `json:"height"`
	Total  int `json:"total"`
	Start  int `json:"start"`
	Detail []struct {
		OverallBalance         string `json:"overallBalance"`
		Ticker                 string `json:"ticker"`
		TransferableBalance    string `json:"transferableBalance"`
		AvailableBalance       string `json:"availableBalance"`
		AvailableBalanceSafe   string `json:"availableBalanceSafe"`
		AvailableBalanceUnSafe string `json:"availableBalanceUnSafe"`
		Decimal                int    `json:"decimal"`
	} `json:"detail"`
}

type WaasBrc20QueryAddressTickerInfoReq struct {
	Ticker      string `json:"ticker"`
	ChainSymbol string `json:"chainSymbol"`
	Address     string `json:"address"`
}

type WaasBrc20QueryAddressTickerInfoRes struct {
	Ticker                   string `json:"ticker"`
	OverallBalance           string `json:"overallBalance"`
	TransferableBalance      string `json:"transferableBalance"`
	AvailableBalance         string `json:"availableBalance"`
	AvailableBalanceSafe     string `json:"availableBalanceSafe"`
	AvailableBalanceUnSafe   string `json:"availableBalanceUnSafe"`
	TransferableCount        string `json:"transferableCount"`
	TransferableInscriptions []struct {
		Data struct {
			Op      string `json:"op"`
			Tick    string `json:"tick"`
			Lim     string `json:"lim"`
			Amt     string `json:"amt"`
			Decimal string `json:"decimal"`
		} `json:"data"`
		InscriptionNumber int    `json:"inscriptionNumber"`
		InscriptionId     string `json:"inscriptionId"`
		Satoshi           string `json:"satoshi"`
		Confirmations     int    `json:"confirmations"`
	} `json:"transferableInscriptions"`
	HistoryCount        int `json:"historyCount"`
	HistoryInscriptions []struct {
		Data struct {
			Op      string `json:"op"`
			Tick    string `json:"tick"`
			Lim     string `json:"lim"`
			Amt     string `json:"amt"`
			Decimal string `json:"decimal"`
		} `json:"data"`
		InscriptionNumber int    `json:"inscriptionNumber"`
		InscriptionId     string `json:"inscriptionId"`
		Satoshi           int    `json:"satoshi"`
		Confirmations     int    `json:"confirmations"`
	} `json:"historyInscriptions"`
}

type RuneTransferParam struct {
	RequestId   string `json:"requestId"`
	ChainSymbol string `json:"chainSymbol"`
	From        string `json:"from"`
	To          string `json:"to"`
	RuneId      string `json:"runeId"`
	Amount      string `json:"amount"`
	FeeRate     int    `json:"feeRate"`
	Note        string `json:"note"`
}

type RunePageBalanceSummaryParam struct {
	ChainSymbol string `json:"chainSymbol"`
	Start       int    `json:"start"`
	Limit       int    `json:"limit"`
	Address     string `json:"address"`
}

type RunePageBalanceSummaryRes struct {
	Height int `json:"height"`
	Total  int `json:"total"`
	Start  int `json:"start"`
	Detail []struct {
		Ticker                 string `json:"ticker"`
		OverallBalance         string `json:"overallBalance"`
		TransferableBalance    string `json:"transferableBalance"`
		AvailableBalance       string `json:"availableBalance"`
		AvailableBalanceSafe   string `json:"availableBalanceSafe"`
		AvailableBalanceUnSafe string `json:"availableBalanceUnSafe"`
		Decimal                int    `json:"decimal"`
	} `json:"detail"`
}

type RuneQueryBalanceParam struct {
	ChainSymbol string `json:"chainSymbol"`
	Address     string `json:"address"`
	RuneId      string `json:"runeId"`
}

type RuneQueryBalanceRes struct {
	Rune         string `json:"rune"`
	Runeid       string `json:"runeid"`
	SpacedRune   string `json:"spacedRune"`
	Amount       string `json:"amount"`
	Symbol       string `json:"symbol"`
	Divisibility int    `json:"divisibility"`
}
