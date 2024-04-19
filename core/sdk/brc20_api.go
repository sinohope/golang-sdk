package sdk

import (
	"encoding/json"
	"fmt"
	commonData "github.com/sinohope/sinohope-golang-sdk/common"
	"github.com/sinohope/sinohope-golang-sdk/core/gateway"
	"github.com/sinohope/sinohope-golang-sdk/features"
)

type brc20Api struct {
	gw features.Gateway
}

func NewBrc20API(baseUrl, private string) (features.Brc20Api, error) {
	gw, err := gateway.NewGateway(baseUrl, private)
	if err != nil {
		return nil, fmt.Errorf("create new gateway failed, %v", err)
	}
	return &brc20Api{
		gw: gw,
	}, nil
}

// InscribeDeploy Inscribe deploy inscription
// POST: /v1/waas/mpc/brc20/inscribe_deploy
func (t *brc20Api) InscribeDeploy(param *commonData.InscribeDeployParam) error {
	if response, err := t.gw.Post("/v1/waas/mpc/brc20/inscribe_deploy", param); err != nil {
		return fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		return nil
	}
}

// InscribeMint Inscribe mint inscription
func (t *brc20Api) InscribeMint(param *commonData.InscribeMintParam) error {
	if response, err := t.gw.Post("/v1/waas/mpc/brc20/inscribe_mint", param); err != nil {
		return fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		return nil
	}
}

// InscribeTransfer Inscribe transfer inscription
func (t *brc20Api) InscribeTransfer(param *commonData.InscribeTransferParam) error {
	if response, err := t.gw.Post("/v1/waas/mpc/brc20/inscribe_transfer", param); err != nil {
		return fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		return nil
	}
}

// InscribeTransferById Brc20 inscription transfer
func (t *brc20Api) InscribeTransferById(param *commonData.TransferByIdParam) error {
	if response, err := t.gw.Post("/v1/waas/mpc/brc20/inscribe_transfer_by_id", param); err != nil {
		return fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		return nil
	}
}

// OneStopTransfer One-stop brc20 transfer
func (t *brc20Api) OneStopTransfer(param *commonData.OneStopTransferParam) error {
	if response, err := t.gw.Post("/v1/waas/mpc/brc20/one_stop_transfer", param); err != nil {
		return fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		return nil
	}
}

// QueryInscribeTransfers Query BRC20 transferable inscription information of specified assets at specified address
func (t *brc20Api) QueryInscribeTransfers(param *commonData.WaasBrc20QueryInscribeTransferReq) (*commonData.WaasBrc20QueryInscribeTransfersRes, error) {
	if response, err := t.gw.Post("/v1/waas/mpc/brc20/query_inscribe_transfers", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.WaasBrc20QueryInscribeTransfersRes
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// QueryPageBalanceSummary All BRC20 asset balance information for the specified address (paged query)
func (t *brc20Api) QueryPageBalanceSummary(param *commonData.WaasBrc20PageQueryBalanceSummaryReq) (*commonData.WaasBrc20PageQueryBalanceSummaryRes, error) {
	if response, err := t.gw.Post("/v1/waas/mpc/brc20/page_balance_summary", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.WaasBrc20PageQueryBalanceSummaryRes
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

func (t *brc20Api) AddressTickerInfo(param *commonData.WaasBrc20QueryAddressTickerInfoReq) (*commonData.WaasBrc20QueryAddressTickerInfoRes, error) {
	if response, err := t.gw.Post("/v1/waas/mpc/brc20/address_ticker_info", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.WaasBrc20QueryAddressTickerInfoRes
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}

// AddressBalance Query the details of the BTC balance of the specified address (total balance, non-inscription balance, inscription related balance)
func (t *brc20Api) AddressBalance(param *commonData.WaasBrc20QueryAddressBalanceReq) (*commonData.WaasBrc20QueryAddressBalanceRes, error) {
	if response, err := t.gw.Post("/v1/waas/mpc/brc20/address_balance", param); err != nil {
		return nil, fmt.Errorf("send request failed, %v", err)
	} else if response.Code != commonData.MPCProxyStatusOk {
		return nil, fmt.Errorf("error response, code: %v msg: %v",
			response.Code, response.Message)
	} else {
		var result *commonData.WaasBrc20QueryAddressBalanceRes
		if err := json.Unmarshal(response.Data, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response failed, %v", err)
		}
		return result, nil
	}
}
