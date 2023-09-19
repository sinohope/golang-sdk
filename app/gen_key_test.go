package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	public, private, err := generateECDSAPrivateKey()
	if err != nil {
		t.Errorf("create new ecdsa failed, %v", err)
		return
	}
	fmt.Println("privateKeyString:", private)
	fmt.Println("publicKeyString:", public)
}

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
	// fmt.Println("private base64:", base64.StdEncoding.EncodeToString(pkcs8Bytes))
	// fmt.Println("public base64:", base64.StdEncoding.EncodeToString(pubKeyBytes))
	return hex.EncodeToString(pubKeyBytes), hex.EncodeToString(pkcs8Bytes), nil
}
