package gateway

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/sirupsen/logrus"
)

type signatureData struct {
	Data      string `json:"data,omitempty"`
	Path      string `json:"path,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Version   string `json:"version,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
}

type Signer struct {
	private *ecdsa.PrivateKey

	privateKeyStr string // x509 PKCS#8 private string
	publicKeyStr  string // x509 ASN.1 der public string
}

func NewSigner(private string) (*Signer, error) {
	sk, err := loadPrivate(private)
	if err != nil {
		return nil, fmt.Errorf("load private key failed, %v", err)
	}
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(sk)
	if err != nil {
		return nil, fmt.Errorf("marshal private key failed, %v", err)
	}
	derBytes, err := x509.MarshalPKIXPublicKey(&sk.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("marshal public key failed, %v", err)
	}
	privateKeyStr := hex.EncodeToString(privateKeyBytes)
	publicKeyStr := hex.EncodeToString(derBytes)
	return &Signer{
		private:       sk,
		privateKeyStr: privateKeyStr,
		publicKeyStr:  publicKeyStr,
	}, nil
}

func (s *Signer) Sign(path, timestamp, payload string) (string, error) {
	logrus.
		WithField("path", path).
		WithField("payload", payload).
		Debugf("prepare to sign")
	messageData := map[string]string{
		"data":      payload,
		"path":      path,
		"timestamp": timestamp,
		"version":   "1.0.0",
	}
	message := generateSignMetaDataAsString(s.PublicKey(), messageData)
	logrus.
		WithField("sign-meta-data", message).
		Debugf("generate sign meta data")
	signature, err := Sign(s.private, hex.EncodeToString([]byte(message)))
	if err != nil {
		return "", fmt.Errorf("sign failed, %v", err)
	}
	logrus.
		WithField("signature", signature).
		Debugf("sign success")
	return signature, err
}

func (s *Signer) PublicKey() string {
	return s.publicKeyStr
}

func (s *Signer) PrivateKey() string {
	return s.privateKeyStr
}

func generateSignMetaDataAsString(appKey string, data map[string]string) string {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signature := ""
	for _, k := range keys {
		signature += k + data[k]
	}
	signature += appKey
	return signature
}
