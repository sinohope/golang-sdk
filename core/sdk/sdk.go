package sdk

import "github.com/sinohope/sinohope-golang-sdk/features"

type sdkImpl struct {
	features.CommonAPI
	features.AccountAndAddressAPI
	features.TransactionAPI
	features.AdvanceAPI
	features.MPCNodeAPI
}

func NewSDK(baseUrl, private string) (features.SinohopeWaaSAPI, error) {
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
	return &sdkImpl{
		CommonAPI:            commonAPI,
		AccountAndAddressAPI: accountAndAddressAPI,
		TransactionAPI:       transactionAPI,
		AdvanceAPI:           advanceAPI,
		MPCNodeAPI:           mpcNodeAPI,
	}, nil
}
