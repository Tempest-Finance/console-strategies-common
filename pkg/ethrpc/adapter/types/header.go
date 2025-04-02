package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Header struct {
	Hash       common.Hash `json:"hash"`
	ParentHash common.Hash `json:"parentHash"`
	Number     *big.Int    `json:"number"`
	Time       uint64      `json:"timestamp"`
}
