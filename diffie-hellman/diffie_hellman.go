package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

// PrivateKey returns a pointer to a `i big.Int, 1 < i < p` .
func PrivateKey(p *big.Int) *big.Int {
	if p == nil || p.Cmp(big.NewInt(3)) <= 0 {
		return nil
	}

	two := big.NewInt(2)
	upper := new(big.Int).Sub(p, two) // p - 2

	n, err := rand.Int(rand.Reader, upper) // 0 <= n < p-2
	if err != nil {
		return nil
	}

	return n.Add(n, two) // shift to 2 <= x < p
}

// PublicKey returns a pointer to a a `i big.Int, i=g**a mod |p|`.
func PublicKey(private, p *big.Int, g int64) *big.Int {
	base := big.NewInt(g)
	return new(big.Int).Exp(base, private, p)
}

// NewPair returns the pointers to the private and the public keys as a tuple.
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
	return private, PublicKey(private, p, g)
}

// PublicKey returns a pointer to a a `i big.Int, i=public2**private1 mod |p|`.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
