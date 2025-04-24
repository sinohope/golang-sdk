package sdk

import (
	"encoding/json"
	"testing"

	"github.com/sinohope/golang-sdk/common"
)

//	{
//	   "fromAddress":"mqC4NKJw3kDNDdLZpTjxUEAkXgSTWiCoBX,0x8b6b4d95cdc757a9f5edd488983c1a23ed29fa64,0xcfa54817f0d9e2b7b1234e5225e890e2fae30a07,0x5d64ddbc51e2aeb9a2e8ef5f1107a9ddc0a7cbe1,mmt9wanCTbjtVZzRajeJpbKiTMXzJSvPFL,0x8a7278f4c3df49e5a4ad1dd891da1b1dadab8207,n4brx1ZF6eWSx9voVkRVzVYsLvyy9y5fRx",
//	   "assetId": "",
//	   "toAddress": "",
//	   "limits": [
//	       {
//	           "type":2,
//	           "chargeUnit": "usdt",
//	           "limit24TimeAmount":5.5
//	       }
//	   ],
//	   "hitResult": "2",
//	   "vaultId": "624388747146949",
//	   "state": 1
//	}
func TestSetTransferStrategy(t *testing.T) {
	api, err := NewConfigAPI(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatalf("create new mpc api failed, %v", err)
	}

	var limits = []*common.SetTransferStrategyLimit{&common.SetTransferStrategyLimit{
		Type:              1,
		ChargeUnit:        "usdt",
		Limit24TimeAmount: "123",
	},
	}

	param := &common.SetTransferStrategyReq{
		FromAddress: "mqC4NKJw3kDNDdLZpTjxUEAkXgSTWiCoBX,0x8b6b4d95cdc757a9f5edd488983c1a23ed29fa64,0xcfa54817f0d9e2b7b1234e5225e890e2fae30a07,0x5d64ddbc51e2aeb9a2e8ef5f1107a9ddc0a7cbe1,mmt9wanCTbjtVZzRajeJpbKiTMXzJSvPFL,0x8a7278f4c3df49e5a4ad1dd891da1b1dadab8207,n4brx1ZF6eWSx9voVkRVzVYsLvyy9y5fRx",
		AssetId:     "",
		ToAddress:   "",
		Limits:      limits,
		VaultId:     "624388747146949",
		HitResult:   2,
		State:       1,
	}
	res, err := api.SetTransferStrategy(param)
	if err != nil {
		t.Error(err)
	}
	d, _ := json.Marshal(res)
	t.Log(string(d))
}
