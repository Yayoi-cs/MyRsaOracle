package RsaCrypt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

var (
	PrivateKeys *rsa.PrivateKey
	PublicKeys  *rsa.PublicKey
)

func RsaInit() {
	privateKeyPEM, err := ioutil.ReadFile("private_key.pem")
	if err != nil {
		fmt.Println("Error reading private key file:", err)
		return
	}

	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		fmt.Println("Failed to parse PEM block containing the private key")
		return
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		return
	}
	var ok bool
	PrivateKeys, ok = privateKey.(*rsa.PrivateKey)
	if !ok {
		fmt.Println("Private key is not an RSA key")
		return
	}

	publicKeyPEM, err := ioutil.ReadFile("public_key.pem")
	if err != nil {
		fmt.Println("Error reading public key file:", err)
		return
	}

	block, _ = pem.Decode(publicKeyPEM)
	if block == nil {
		fmt.Println("Failed to parse PEM block containing the public key")
		return
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing public key:", err)
		return
	}

	PublicKeys, ok = publicKey.(*rsa.PublicKey)
	if !ok {
		fmt.Println("Public key is not an RSA key")
		return
	}
}
