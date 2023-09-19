package sdk

import "github.com/sinohope/sinohope-golang-sdk/features"

type APIClient struct {
	Common            features.CommonAPI
	AccountAndAddress features.AccountAndAddressAPI
	Transaction       features.TransactionAPI
	Advance           features.AdvanceAPI
	MPCNode           features.MPCNodeAPI
}

func NewApiClient(baseUrl, private string) (*APIClient, error) {
	commonAPI, err := NewCommonAPI(baseUrl, private)
	if err != nil {
		return nil, err
	}
	accountAndAddressAPI, err := NewAccountAndAddressAPI(baseUrl, private)
	if err != nil {
		return nil, err
	}
	transactionAPI, err := NewTransactionAPI(baseUrl, private)
	if err != nil {
		return nil, err
	}
	advanceAPI, err := NewAdvanceAPI(baseUrl, private)
	if err != nil {
		return nil, err
	}
	mpcNodeAPI, err := NewMPCNodeAPI(baseUrl, private)
	if err != nil {
		return nil, err
	}
	return &APIClient{
		Common:            commonAPI,
		AccountAndAddress: accountAndAddressAPI,
		Transaction:       transactionAPI,
		Advance:           advanceAPI,
		MPCNode:           mpcNodeAPI,
	}, nil
}
