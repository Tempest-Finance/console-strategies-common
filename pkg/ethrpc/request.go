package ethrpc

import (
	"context"
	"github.com/Tempest-Finance/console-strategies-common/pkg/ethrpc/adapter/types"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Request struct {
	executor RequestExecutor

	Method         RequestMethod
	RequireSuccess bool
	Calls          []*Call
	ctx            context.Context
	RawCallMsg     *types.CallMsg

	BlockNumber *big.Int
	BlockHash   common.Hash
}

type RequestMethod string

var (
	RequestMethodCall                 RequestMethod = "call"
	RequestMethodAggregate            RequestMethod = "aggregate"
	RequestMethodTryBlockAndAggregate RequestMethod = "tryBlockAndAggregate"
)

// Context method returns the Context if it's already set in request
// otherwise it creates new one using `context.Background()`.
func (r *Request) Context() context.Context {
	if r.ctx == nil {
		return context.Background()
	}

	return r.ctx
}

// SetContext method sets the context.Context for current Request. It allows
// to interrupt the request execution if ctx.Done() channel is closed.
// See https://blog.golang.org/context article and the "context" package
// documentation.
func (r *Request) SetContext(ctx context.Context) *Request {
	r.ctx = ctx
	return r
}

// AddCall adds a call to the request
// it will autofill the UnpackABI in case it's not set
func (r *Request) AddCall(c *Call, output []interface{}) *Request {
	c.autofillUnpackABI()
	c.SetOutput(output)
	r.Calls = append(r.Calls, c)

	return r
}

func (r *Request) SetRequireSuccess(requireSuccess bool) *Request {
	r.RequireSuccess = requireSuccess

	return r
}

func (r *Request) SetBlockNumber(blockNumber *big.Int) *Request {
	r.BlockNumber = blockNumber

	return r
}

func (r *Request) SetBlockHash(blockHash common.Hash) *Request {
	r.BlockHash = blockHash

	return r
}

func (r *Request) Execute(method RequestMethod) (*Response, error) {
	r.Method = method

	return r.executor.Execute(r)
}

func (r *Request) Call() (*Response, error) {
	return r.Execute(RequestMethodCall)
}

func (r *Request) Aggregate() (*Response, error) {
	return r.Execute(RequestMethodAggregate)
}

func (r *Request) TryBlockAndAggregate() (*Response, error) {
	return r.Execute(RequestMethodTryBlockAndAggregate)
}
