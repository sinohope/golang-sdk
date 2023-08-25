package common

type ListMPCRequestsParam struct {
	BusinessExecType   int    `json:"businessExecType"`
	BusinessExecStatus int    `json:"businessExecStatus"`
	SinoID             string `json:"sinoId"`
	PageIndex          int    `json:"pageIndex"`
	PageSize           int    `json:"pageSize"`
}
