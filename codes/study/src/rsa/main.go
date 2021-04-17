package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

func RsaEncrypt(origData, publicKeyData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKeyData)
	if block == nil {
		return nil, errors.New("public key error")
	}

	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := key.(*rsa.PublicKey)

	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func RsaDecrypt(cipherText, privateKeyData []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		return nil, errors.New("private key error")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, key, cipherText)
}

func main() {
	data, _ := RsaEncrypt([]byte("hello world"))
	fmt.Println(base64.StdEncoding.EncodeToString(data))
	origData, _ := RsaDecrypt(data)
	fmt.Println(string(origData))
}
