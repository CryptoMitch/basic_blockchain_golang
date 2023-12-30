package GoBlockchain

import "GoBlockchain/ed25519"

const (
	privateKeyLen = 64
	publicKeyLen = 32

)

// Ed25519 is an elliptic curve cryptography (ECC) algorithm that provides digital signatures. 
// It is based on the mathematical properties of a specific elliptic curve called the Ed25519 curve. 
// The algorithm is designed for efficiency, security, and simplicity.

type PrivateKey struct {
	key ed25519.PrivateKey
}

func (p * PrivateKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) Sign(msg []byte) []byte {
	return ed25519.Sign(p.key, msg)
}

func (p *PrivateKey) Public() *PublicKey{
	b := make([]byrw)

}

func PublicKey struct {
	key ed25519.PublicKey