package RsaCrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
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
	cipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, PublicKeys, []byte(plainText), nil)
	return cipherText, err
}

func RsaDecrypt(cipherText string) ([]byte, error) {
	plainText, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, PrivateKeys, []byte(cipherText), nil)
	return plainText, err
}

func ModInverseA(a, m *big.Int) *big.Int {
	g := new(big.Int).Set(a)
	x := new(big.Int)
	y := new(big.Int).Set(m)
	u := big.NewInt(1)
	v := big.NewInt(0)

	for g.Cmp(big.NewInt(0)) != 0 {
		q := new(big.Int)
		q.Div(y, g)

		t := new(big.Int)
		t.Set(g)
		g.Mod(y, g)
		y.Set(t)

		t.Set(x)
		t.Mul(q, x)
		t2 := new(big.Int)
		t2.Set(u)
		u.Sub(u, t2)
		u.Mul(u, q)
		x.Set(v)
		x.Add(x, u)
		u.Set(t)
		v.Set(t2)
	}

	result := new(big.Int).Add(x, m)
	result.Mod(result, m)

	return result
}
