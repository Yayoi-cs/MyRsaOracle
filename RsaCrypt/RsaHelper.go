package RsaCrypt

import (
	"crypto/rsa"
	"math/big"
)

func generateRSAKey(p, q *big.Int) *rsa.PrivateKey {
	// N = p * q
	N := new(big.Int).Mul(p, q)

	// φ(N) = (p - 1) * (q - 1)
	pSub1 := new(big.Int).Sub(p, big.NewInt(1))
	qSub1 := new(big.Int).Sub(q, big.NewInt(1))
	phi := new(big.Int).Mul(pSub1, qSub1)

	// Choose E
	E := big.NewInt(65537)
	iE := 65537

	// Calc D
	D := new(big.Int).ModInverse(E, phi)

	// Generate Public Key
	publicKey := rsa.PublicKey{N: N, E: iE}

	// Generate Private Key
	privateKey := rsa.PrivateKey{PublicKey: publicKey, D: D}

	return &privateKey
}

func RsaEncrypt(plainText string) ([]byte, error) {
	c := new(big.Int).SetBytes([]byte(plainText))
	c.Exp(c, big.NewInt(int64(PublicKeys.E)), PublicKeys.N)
	return c.Bytes(), nil
}

func RsaDecrypt(cipherText []byte) ([]byte, error) {
	c := new(big.Int).SetBytes([]byte(cipherText))
	m := new(big.Int).Exp(c, PrivateKeys.D, PrivateKeys.N)
	return m.Bytes(), nil
}
