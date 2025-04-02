package crypto

import (
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Signer struct {
	address    common.Address
	privateKey *ecdsa.PrivateKey
}

func NewSigner(privateKey string) (*Signer, error) {
	ecdsaKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	return &Signer{
		address:    crypto.PubkeyToAddress(ecdsaKey.PublicKey),
		privateKey: ecdsaKey,
	}, nil
}

func (s *Signer) Sign(ctx context.Context, hash []byte) ([]byte, error) {
	signature, err := crypto.Sign(hash, s.privateKey)
	if err != nil {
		return nil, err
	}

	if signature[crypto.RecoveryIDOffset] == 0 || signature[crypto.RecoveryIDOffset] == 1 {
		signature[crypto.RecoveryIDOffset] += 27 // Transform yellow paper V from 27/28 to 0/1
	}

	return signature, nil
}

func (s *Signer) PublicKey() string {
	return s.address.String()
}
