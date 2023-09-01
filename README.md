# 关于Sinohope WaaS API

[toc]

## 概述

欢迎使用 Sinohope Waas（Wallet as a Service） API。

[Sinohope 平台](https://www.sinohope.com/)是一站式的数字资产托管平台，同时支持自托管（基于MPC-TSS 技术）/全托管（基于 HSM 技术）产品，也支持以 API 的方式来使用相应的产品，满足广大客户的各类需求。

简要介绍 Sinohope 产品优势：

安全

- 综合 Tier 1 MPC-TSS 技术服务提供商 基于全球最新 MPC-CMP 协议创新研发 多重灾备方案，多场景分片恢复机制
- 完备的安全体系 软件与硬件结合双结合（MPC-TSS+TEE） 密码学和完备安全机制双保障

经验

- 十年数字资管经验 10年+ 全球 Top3 交易所资管经验 团队成员来自全球顶级高校与企业
- 百亿资金托管经验 150亿+ USD 数字资产托管经验 数十亿笔交易处理经验

资质：

- 多监管区域持牌 隶属香港上市公司，代码：1611 持有香港4/9号牌/TCSP 牌照 获得 SOC 2 Type1 和 Type2 审计报告
- 交易安全性 平台集成 AML/KYT 实时监控能力 交易目标黑白名单/交易复核/规则设置

功能：

- 易用且强大的产品功能 10秒创建钱包，功能易上手，Web2 体验 多角色/多层级/多节点审批流/自定义交易规则引擎 多种托管模式：协管/全托管/合规托管
- 多应用场景解决方案 企业财务多层审批解决方案 钱包基础设施/支付解决方案/国库解决方案

扩展：

- 多种访问方式 APP + 网页版，无缝使用 API + SDK，应对客户多种开发需求
- 面向 Web3 未来 DApp/DeFi/NFT 支持 链上资管/链上投融资 支持

## 术语定义

### sinoId

对于涉及到MPC私钥分片的业务，包括开发者请求的交易及签名类业务（如发起交易、发起签名等），Sinohope均会给该业务产生一个唯一标识，即 sinoId。

### chainSymbol的规则

1、chainSymbol是链配置的唯一标识；

2、区分测试网和正式网；

3、使用大写字母命名。

> 例如：
>
> - 主网使用常见简称，如`ETH`, `BTC` 分别是 `Ethereum 主网`/ `Bitcoin主网`
> - 对于有特定名字的测试网使用其测试网名字，如用`GOERLI` 表示`ETH Goerli 测试网`
> - 对于无特定名字的测试网，使用`_TEST`加以区分，如 `BTC_TEST` 表示 `Bitcoin testnet3测试网`
> - 注意：系统支持的链，以`/v1/waas/common/get_supported_chains`接口返回值为准

### assetId的规则

1、assetId是Asset配置表唯一标识；

2、命名方式前币(coin)后链(chain)，中间以下划线分隔；

3、使用大写字母命名。

> 例如：USDT_GOERLI 是 ETH Goerli 测试网上的的一个 USDT 币种

## 开始使用

本章节是一个快速使用教程，将帮助您快速了解并开始使用Sinohope WaaS Golang SDK。SDK 主要提供以下功能：

- 统一封装接口请求签名处理；
- 封装API接口，方便开发者以方法调用的方式完成对WaaS API的请求。

### 代码架构

```
.
├── go.mod
├── go.sum
├── core/
│   ├── signer/
│   │   ├── signer_test.go
│   │   ├── ecdsa.go
│   │   ├── signer.go
│   ├── http/
│   │   ├── http_test.go
│   │   ├── http.go
│   │   ├── utilities.go
│   ├── sdk/
│   │   ├── common_api_test.go
│   │   ├── common_api.go
│   │   ├── mpc_api.go
│   │   ├── mpc_node_api_test.go
│   │   ├── advance_api.go
│   │   ├── mpc_node_api.go
│   │   ├── mpc_api_test.go
├── app/
│   ├── main.go
├── features/
│   ├── http.go
│   ├── sinohope_waas_api.go
│   ├── signer.go
├── common/
│   ├── types.go
├── log/
│   ├── log.go
│   ├── error.go
│   ├── json_formatter.go
│   ├── types.go
│   ├── console_hook.go
│   ├── file_hook.go
```

架构说明：

+ app
+ core
  + signer
    + signer.go: 根据sinohope网关的规则生成POST请求签名；
  + http
    + http.go: 调用signer生成请求签名，然后构建POST请求；
  + sdk
    + common_api.go: 通用API的实现；
    + mpc_api.go: 帐户&地址API的实现；
    + advance_api.go: 高级功能API的实现；
    + mpc_node_api.go: 查询MPC Node节点记录与状态的API的实现；
+ features：接口抽象
  + http: 负责向sinohope发起请求和解析返回结果，依赖signer；
  + sinohope_waas_api: 按模块划分定义了当前支持的所有接口；
  + signer: http模块发起请求前会调用signer，根据sinohope网关签名规则生成请求签名；
+ common：全局数据结构定义
+ log：日志模块

### 使用示例

项目的main方法中包含了MPC Node节点API的使用示例：

```go
package main

import (
	"flag"
	"fmt"
	"github.com/sinohope/sinohope-golang-sdk/core/http"
	"github.com/sinohope/sinohope-golang-sdk/core/sdk"
	"github.com/sinohope/sinohope-golang-sdk/core/signer"
	"github.com/sinohope/sinohope-golang-sdk/log"
	"github.com/sirupsen/logrus"

	"github.com/sinohope/sinohope-golang-sdk/common"
)

var (
	private = flag.String("private", common.FakePrivateKey, "Your private key.")
	url     = flag.String("url", common.BaseUrl, "URL to mpc-proxy.")
)

func main() {
	flag.Parse()

  // 启动参数检查
	if err := check(); err != nil {
		panic(err)
	}

  // 日志模块初始化
	a := log.App{
		Name: "sinohope-golang-sdk",
	}
	l := log.Log{
		Stdout: struct {
			Enable bool `toml:"enable"`
			Level  int  `toml:"level"`
		}{
			Enable: true,
			Level:  5,
		},
		File: struct {
			Enable bool   `toml:"enable"`
			Level  int    `toml:"level"`
			Path   string `toml:"path"`
			MaxAge int    `toml:"max-age"`
		}{
			Enable: true,
			Level:  5,
			Path:   "./logs/sinohope-golang-sdk.log",
			MaxAge: 7,
		},
	}
	log.SetLogDetailsByConfig(a, l)

  // 调用MPC Node API的示例
	checkMPCNodeStatus()
}

func check() error {
	if *private == "" {
		return fmt.Errorf("private key can not be empty")
	}
	if *url == "" {
		return fmt.Errorf("mpc-proxy url can not be empty")
	}
	return nil
}
```

调用MPC Node API的示例：

```go
func checkMPCNodeStatus() {
	//public, private, err := generateECDSAPrivateKey()
	//if err != nil {
	//	logrus.Errorf("create new ecdsa failed, %v", err)
	//	return
	//}
	public := common.FakePublicKey
	private := common.FakePrivateKey
	logrus.
		WithField("public", public).
		WithField("private", private).
		Infof("after generate ECDSA keypair")
  
  // 构建signer对象需要提供ECDSA的私钥
	s, err := signer.NewSigner(private)
	if err != nil {
		logrus.Errorf("create new signer failed, %v", err)
		return
	}
	p, err := http.NewHTTP(common.BaseUrl, s)
	if err != nil {
		logrus.Errorf("create http failed, %v", err)
		return
	}
  // 初始化MPC Node API SDK
	m, err := sdk.NewMPCNodeAPI(p)
	if err != nil {
		logrus.Errorf("create mpc node sdk failed, %v", err)
		return
	}
	request := &common.WaasMpcNodeExecRecordParam{
		BusinessExecType:   1,
		BusinessExecStatus: 10,
		SinoId:             "fake sino id",
		PageIndex:          0,
		PageSize:           40,
	}
  // 查询MPC Node协议执行记录
	var result *common.PageData
	if result, err = m.ListMPCRequests(request); err != nil {
		logrus.Errorf("list mpc requests failed, %v", err)
	} else {
		// TODO: do something with result
		fmt.Printf("-----------> [%s]", result.List)
	}
  // 查询MPC Node状态
	var status *common.WaaSMpcNodeStatusDTOData
	if status, err = m.Status(); err != nil {
		logrus.Errorf("get mpc node status failed, %v", err)
	} else {
		logrus.Infof("get mpc node status success, %v", status)
	}
}
```

生成ECDSA密钥对：

```
func generateECDSAPrivateKey() (string, string, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", "", err
	}
	pkcs8Bytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return "", "", err
	}
	pubKeyBytes := elliptic.Marshal(privateKey.PublicKey.Curve, privateKey.PublicKey.X, privateKey.PublicKey.Y)
	return hex.EncodeToString(pubKeyBytes), hex.EncodeToString(pkcs8Bytes), nil
}
```

说明：

1. 新生成的ECDSA密钥对请先在 Sinohope Web Console 完成 API Key 的绑定，否则会返回如下错误：code: 1004 msg: APIKey，