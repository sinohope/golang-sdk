package gateway

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"fmt"
)

func loadPrivate(private string) (*ecdsa.PrivateKey, error) {
	pkBytes, err := hex.DecodeString(private)
	if err != nil {
		return nil, err
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(pkBytes)
	if err != nil {
		return nil, err
	}
	ecdsaPrivateKey, ok := privateKey.(*ecdsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("convert PKCS8 to ecdsa private key failed")
	}
	return ecdsaPrivateKey, nil
}

func Sign(private *ecdsa.PrivateKey, message string) (string, error) {
	messageBytes, err := hex.DecodeString(message)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(messageBytes)
	signature, err := ecdsa.SignASN1(rand.Reader, private, hash[:])
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(signature), nil
}
