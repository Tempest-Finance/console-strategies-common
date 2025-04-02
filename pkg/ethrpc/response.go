package ethrpc

import (
	"math/big"
)

type Response struct {
	Request     *Request
	RawResponse []byte
	BlockNumber *big.Int

	// Result is an array that contains response result for all calls in the request
	Result []bool
}
