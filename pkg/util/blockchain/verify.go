package blockchain

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySignature(fromAddress, signatureHex, message string) error {
	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		return err
	}

	signature[crypto.RecoveryIDOffset] -= 27

	messageHash := accounts.TextHash([]byte(message))

	pubKey, err := crypto.SigToPub(messageHash, signature)
	if err != nil {
		return err
	}

	if common.HexToAddress(fromAddress) != crypto.PubkeyToAddress(*pubKey) {
		return fmt.Errorf("failed to verify signature")
	}

	return nil
}
