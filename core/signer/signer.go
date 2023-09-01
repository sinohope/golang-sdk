package signer

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/sinohope/sinohope-golang-sdk/features"
	"github.com/sirupsen/logrus"
	"sort"
)

type signatureData struct {
	Data      string `json:"data,omitempty"`
	Path      string `json:"path,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Version   string `json:"version,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
}

type signer struct {
	private *ecdsa.PrivateKey
}

func NewSigner(private string) (features.Signer, error) {
	pk, err := loadPrivate(private)
	if err != nil {
		return nil, fmt.Errorf("load private key failed, %v", err)
	}
	return &signer{
		private: pk,
	}, nil
}

func (s *signer) Sign(path, timestamp, payload string) (string, error) {
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

func (s *signer) PublicKey() string {
	derBytes, err := x509.MarshalPKIXPublicKey(&s.private.PublicKey)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(derBytes)
}

func (s *signer) PrivateKey() string {
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(s.private)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(privateKeyBytes)
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
