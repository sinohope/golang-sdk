package features

import "github.com/sinohope/sinohope-golang-sdk/common"

type MPCNode interface {
	// ListMPCRequests ...
	// POST: /v1/waas/mpc/mpcnode/list_mpc_requests
	ListMPCRequests(param *common.ListMPCRequestsParam) (*common.Response, error)

	// Status ...
	// POST: /v1/waas/mpc/mpcnode/status
	Status() (*common.Response, error)
}
