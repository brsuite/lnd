package mock

import (
	"github.com/brsuite/brond/bronec"
	"github.com/brsuite/brond/chaincfg/chainhash"
	"github.com/brsuite/lnd/keychain"
)

// SecretKeyRing is a mock implementation of the SecretKeyRing interface.
type SecretKeyRing struct {
	RootKey *bronec.PrivateKey
}

// DeriveNextKey currently returns dummy values.
func (s *SecretKeyRing) DeriveNextKey(
	_ keychain.KeyFamily) (keychain.KeyDescriptor, error) {

	return keychain.KeyDescriptor{
		PubKey: s.RootKey.PubKey(),
	}, nil
}

// DeriveKey currently returns dummy values.
func (s *SecretKeyRing) DeriveKey(
	_ keychain.KeyLocator) (keychain.KeyDescriptor, error) {

	return keychain.KeyDescriptor{
		PubKey: s.RootKey.PubKey(),
	}, nil
}

// DerivePrivKey currently returns dummy values.
func (s *SecretKeyRing) DerivePrivKey(
	_ keychain.KeyDescriptor) (*bronec.PrivateKey, error) {

	return s.RootKey, nil
}

// ECDH currently returns dummy values.
func (s *SecretKeyRing) ECDH(_ keychain.KeyDescriptor,
	_ *bronec.PublicKey) ([32]byte, error) {

	return [32]byte{}, nil
}

// SignMessage signs the passed message and ignores the KeyDescriptor.
func (s *SecretKeyRing) SignMessage(_ keychain.KeyLocator,
	msg []byte, doubleHash bool) (*bronec.Signature, error) {

	var digest []byte
	if doubleHash {
		digest = chainhash.DoubleHashB(msg)
	} else {
		digest = chainhash.HashB(msg)
	}
	return s.RootKey.Sign(digest)
}

// SignMessageCompact signs the passed message.
func (s *SecretKeyRing) SignMessageCompact(_ keychain.KeyLocator,
	msg []byte, doubleHash bool) ([]byte, error) {

	var digest []byte
	if doubleHash {
		digest = chainhash.DoubleHashB(msg)
	} else {
		digest = chainhash.HashB(msg)
	}
	return bronec.SignCompact(bronec.S256(), s.RootKey, digest, true)
}
