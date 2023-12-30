package GoBlockchain

import "GoBlockchain/ed25519"

const (
	privateKeyLen = 64
	publicKeyLen = 32
	seedLen = 32
)

// Ed25519 is an elliptic curve cryptography (ECC) algorithm that provides digital signatures. 
// It is based on the mathematical properties of a specific elliptic curve called the Ed25519 curve, designed for efficiency, security, and simplicity.

type PrivateKey struct {
	key ed25519.PrivateKey
}

func GeneratePrivateKey() *PrivateKey{
	// Create a byte slice 'seed' with a length of seedLen.
	seed := make([]byte, seedLen)
	// Use the 'rand.Reader' as a source of cryptographic randomness to fill the 'seed' slice.
	_, err := io.ReadFull(rand.Reader, seed)
	// Check for errors during the reading process.
	if err != nil {
		// If an error occurs, trigger a panic with the error message.
    	// Panicking in this context indicates a critical failure, and the program will terminate.
		panic(err)
	}
}

// Bytes returns the raw bytes of the private key.
func (p * PrivateKey) Bytes() []byte {
	return p.key
}

// Sign generates a digital signature for the given message using the private key.
func (p *PrivateKey) Sign(msg []byte) []byte {
	return ed25519.Sign(p.key, msg)
}

/*
In the Ed25519 elliptic curve algorithm, a private key is a 256-bit (32-byte) random number, 
and the corresponding public key is derived from part of this private key. 
The public key is essentially the second half (from byte 32 onwards) of the private key.
*/

func (p *PrivateKey) Public() *PublicKey{
	// Create a byte slice 'b' with a length of publicKeyLen
	b := make([]byte, publicKeyLen)
	// Copy the second half of the private key 'p.key' to 'b'
	copy(b, p.key[32:])
	// Create and return a new PublicKey object with the bytes 'b'
	return &PublicKey{
		key: b,
	}
}

// PublicKey represents the Ed25519 public key.
func PublicKey struct {
	key ed25519.PublicKey
}