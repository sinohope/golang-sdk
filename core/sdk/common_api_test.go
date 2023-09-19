package sdk

import (
	"testing"

	"github.com/sinohope/sinohope-golang-sdk/common"
	"github.com/sinohope/sinohope-golang-sdk/log"
)

func TestCommonAPI(t *testing.T) {
	a := log.App{}
	l := log.Log{}
	log.SetLogDetailsByConfig(a, l)

	c, err := NewCommonAPI(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatalf("create new commonAPI api failed, %v", err)
	}

	// GetSupportedChains
	var supportList []*common.WaasChainData
	if supportList, err = c.GetSupportedChains(); err != nil {
		t.Fatalf("get supported chains failed, %v", err)
	} else {
		t.Logf("supported chains:")
		for _, v := range supportList {
			t.Logf("chainName: %v, chainSymbol: %v", v.ChainName, v.ChainSymbol)
		}
	}
	// GetSupportedCoins
	var supportCoins []*common.WaaSCoinDTOData
	for _, v := range supportList {
		param := &common.WaasChainParam{
			ChainName:   v.ChainName,
			ChainSymbol: v.ChainSymbol,
		}
		if supportCoins, err = c.GetSupportedCoins(param); err != nil {
			t.Fatalf("get supported coins failed, %v", err)
		} else {
			t.Logf("supported coins:")
			for _, v := range supportCoins {
				t.Logf("assetName: %v, assetId: %v, assetDecimal: %v",
					v.AssetName, v.AssetId, v.AssetDecimal)
			}
		}
	}
	// GetVaults
	var vaults []*common.WaaSVaultInfoData
	if vaults, err = c.GetVaults(); err != nil {
		t.Fatalf("get vaults failed, %v", err)
	} else {
		for _, v := range vaults {
			t.Logf("businessType: %v",
				v.BusinessType)
			for _, v2 := range v.VaultInfoOfOpenApiList {
				t.Logf("vaultId: %v, vaultName: %v, authorityType: %v, createTime: %v",
					v2.VaultId, v2.VaultName, v2.AuthorityType, v2.CreateTime)
			}
		}
	}
}
