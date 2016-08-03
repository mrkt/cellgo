//|------------------------------------------------------------------
//|        __
//|     __/  \
//|  __/  \__/_
//| /  \__/    \
//|/\__/CellGo /_
//|\/_/NetFW__/  \
//|  /\__ _/  \__/
//|  \/_/  \__/_/
//|    /\__/_/
//|    \/_/
//| ------------------------------------------------------------------
//| Cellgo Framework tool/rsa file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-08-03

package tool

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var Rsa = &rsaTool{PrivateKey: privateKey, PublicKey: publicKey}

type rsaTool struct {
	PrivateKey []byte
	PublicKey  []byte
}

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCjx9aymJ3UhAOW20IVhJYvBi6/JsKo9vhWpOGlSbMmF7mSSS6W
E1yY4797NCiwq5HB76aIiaYcSot/JaEk0jzv9hPfV0yhqIQNNQrViYssPENEeiJv
d2bBSZCx5kZt3Jvqxmz+lLNPYf7B+Y8HWKeinKMAIiCjlbgjsWL3ldGJ/QIDAQAB
AoGAWwpk8BYh9dYEYME0tN1k1nLrXVpgbqgKV6+DtuuG6C/b+dMwUEAnAt3mvMe7
rqlQdquOuOs7KRLPBDiYoO007ZLN0qT4EhSAoyDGRwSxR4duxBeWFGpZugbl5lpe
ORhnyFBUCoy2Nn5v/qjbMWpDzUnej74YWsjLBHXUP/FrvAECQQDWMzRRmXXPNnvY
Xp9jMrDbZ1KNB4Uas4dPuYCvozHB/NowzFMW4eiMb2J9pNK5/NYvcsezon7e02sB
hlD4goR9AkEAw73UyUZ0GSWZj1CQhsjUs+sVu/SzoaZWAo4H0kyxHbif5Vr/Dait
gkxz18LWYAhOJoYt+HawPOrIuygYoC2TgQJATCZqDDR1eJRTFQoWugp0a5vg8jhh
Lqvyh/pX8JkhAGknHMAXLgRkS0DyD97/95UWhEnXC1rSpd8dpK4erSqhdQJBAIN3
j2A0KqvtHgsssDVm072XqjxYKQHsRx5mKIitq9PrebFBAcc5wHegQ7npibRsP5kO
S/fyN4JiqrlRP+DtW4ECQDCjPVuFjIo1c3jk+BSaKx5oW/g+iqXZtZ3pGMgpeM7x
3TDnZMf41GKGjHXgUC5pQlyiV8jWrqKDPiPG2Px+Xkg=
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCjx9aymJ3UhAOW20IVhJYvBi6/
JsKo9vhWpOGlSbMmF7mSSS6WE1yY4797NCiwq5HB76aIiaYcSot/JaEk0jzv9hPf
V0yhqIQNNQrViYssPENEeiJvd2bBSZCx5kZt3Jvqxmz+lLNPYf7B+Y8HWKeinKMA
IiCjlbgjsWL3ldGJ/QIDAQAB
-----END PUBLIC KEY-----
`)

// Rsa encrypt
func (r *rsaTool) RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(r.PublicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// Rsa Decrypt
func (r *rsaTool) RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(r.PrivateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
