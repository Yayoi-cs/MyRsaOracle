package RsaCrypt

import (
	"crypto/rsa"
	"math/big"
)

var (
	PrivateKeys *rsa.PrivateKey
	PublicKeys  *rsa.PublicKey
)

func RsaInit() {
	p := big.NewInt(103289)
	q := big.NewInt(103841)
	PrivateKeys = generateRSAKey(p, q)
	PublicKeys = &PrivateKeys.PublicKey
}
