# 说明

本项目是 Sinohope WaaS API 的 golang SDK。关于Sinohope WaaS 服务，详见[docs.sinohope.com](https://docs.sinohope.com/)

## 开始使用

本章节是一个快速使用教程，将帮助您快速了解并开始使用Sinohope WaaS Golang SDK。SDK 主要提供以下功能：

- 统一封装接口请求签名处理；
- 封装API接口，方便开发者以方法调用的方式完成对WaaS API的请求。

### 生成API请求密钥对

在使用之前，需要您自行生成用于代表您的身份访问Sinohope WaaS API的 ECDSA 密钥对，然后将公钥（即 API Key）通过Sinohope Web console进行配置，与您的组织绑定。

可以参考位于`app/gen_key_test.go` 中的代码生成ECDSA密钥对，如下：

```go
func generateECDSAPrivateKey() (string, string, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", "", err
	}
	pkcs8Bytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return "", "", err
	}
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}
	return hex.EncodeToString(pubKeyBytes), hex.EncodeToString(pkcs8Bytes), nil
}
```

说明：

1. 新生成的ECDSA密钥对请先在 Sinohope Web Console 完成 API Key 的绑定，否则会返回如下错误：code: 1004 msg: Invalid API Key


### 使用示例

项目的main方法中包含了MPC Node节点API的使用示例。

基本使用方式如下：

```go
	client, err := sdk.NewApiClient("<baseurl>", "<your api private key, in PKCS #8, ASN.1 DER form, hex string>")
	if err != nil {
		// TODO: handle initialization error
		return
	}
	
	var status *common.WaaSMpcNodeStatusDTOData
	if status, err = client.MPCNode.Status(); err != nil {
		// handle request error
	} else {
		// handle result
	}
```

### 代码架构简介

代码结构及主要内容说明：

+ app： 示例内容
+ core
  + signer
    + signer.go: 根据sinohope网关的规则生成POST请求签名；
  + http
    + http.go: 调用signer生成请求签名，然后构建POST请求；
  + sdk
    + sdk.go: 对所有API调用实现的一个统一封装；
    + common_api.go: 通用API的实现；
    + account_and_address_api.go: 帐户&地址API的实现；
    + transaction_api.go: 交易及签名API的实现；
    + advance_api.go: 高级功能API的实现；
    + mpc_node_api.go: 查询MPC Node节点记录与状态的API的实现；
+ features：接口抽象
  + http: 负责向sinohope发起请求和解析返回结果，依赖signer；
  + sinohope_waas_api: 按模块划分定义了当前支持的所有接口；
  + signer: http模块发起请求前会调用signer，根据sinohope网关签名规则生成请求签名；
+ common：全局数据结构定义
+ log：日志模块



