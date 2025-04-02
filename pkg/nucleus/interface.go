package nucleus

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type IClient interface {
	Get(ctx context.Context, endpoint string, queryParams map[string]string) (interface{}, error)
	Post(ctx context.Context, endpoint string, jsonData []byte) (interface{}, error)
	GetAddressBook() map[int64]NetworkData
}

type ICalldataQueue interface {
	AddCall(targetAddress common.Address, calldata []byte, value *big.Int)
	GetCalldata(ctx context.Context) (*Calldata, error)
	Execute(ctx context.Context) (string, error)
}
