package common

type StakingCreateParam struct {
	RequestId     string `json:"requestId"`
	ChainSymbol   string `json:"chainSymbol"`
	From          string `json:"from"`
	FeeRate       string `json:"feeRate,omitempty"`
	Note          string `json:"note,omitempty"`
	StakingAmount string `json:"stakingAmount"`
	StakingTime   string `json:"stakingTime"` //blocks
}

type StakingRes struct {
	SinoId string `json:"sinoId"`
}

type UnbondCreateParam struct {
	RequestId    string `json:"requestId"`
	OriSinoId    string `json:"oriSinoId"`
	OriRequestId string `json:"oriRequestId"`
	FeeRate      string `json:"feeRate,omitempty"`
	Note         string `json:"note,omitempty"`
}

type SpendingTimeLockPathTxParam struct {
	RequestId    string `json:"requestId"`
	OriSinoId    string `json:"oriSinoId"`
	OriRequestId string `json:"oriRequestId"`
	FeeRate      string `json:"feeRate,omitempty"`
	Note         string `json:"note,omitempty"`
}
