package common

import "encoding/json"

type Response struct {
	Code      int             `json:"code,omitempty"`
	Message   string          `json:"message,omitempty"`
	Success   bool            `json:"success,omitempty"`
	RequestID string          `json:"requestId,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
}

type ListMPCRequestsResponse struct {
	PageIndex int              `json:"pageIndex,omitempty"`
	PageSize  int              `json:"pageSize,omitempty"`
	TotalSize int              `json:"totalSize,omitempty"`
	List      []MPCRequestList `json:"list,omitempty"`
}

type MPCRequestList struct {
	RequestID          string `json:"requestId,omitempty"`
	SinoID             string `json:"sinoId,omitempty"`
	RequestTime        int64  `json:"requestTime,omitempty"`
	Param              string `json:"param,omitempty"`
	BusinessExecType   int    `json:"businessExecType,omitempty"`
	BusinessExecStatus int    `json:"businessExecStatus,omitempty"`
	IsSuccess          bool   `json:"isSuccess,omitempty"`
	FailedReason       string `json:"failedReason,omitempty"`
}

type MPCNodeStatusResponse struct {
	NodeId       string `json:"nodeId,omitempty"`
	OnlineStatus int    `json:"onlineStatus,omitempty"`
}
