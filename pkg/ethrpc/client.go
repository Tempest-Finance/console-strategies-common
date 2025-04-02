package ethrpc

import (
	"context"
	"github.com/Tempest-Finance/console-strategies-common/pkg/ethrpc/adapter"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var zeroHash = common.Hash{}

type Client struct {
	adapter adapter.EthClientAdapter

	options ClientOptions

	requestMiddlewares  []RequestMiddleware
	responseMiddlewares []ResponseMiddleware
}

type ClientOptions struct {
	MultiCallContractAddress common.Address
	MultiCallABI             *abi.ABI
}

func NewClient(options ...func(*Client)) *Client {
	client := &Client{}

	for _, o := range options {
		o(client)
	}

	return client
}

func WithEthClientAdapter(ethClientAdapter adapter.EthClientAdapter) func(*Client) {
	return func(client *Client) {
		client.adapter = ethClientAdapter
	}
}

func WithMulticall(multiCallContractAddress common.Address, multiCallABI *abi.ABI) func(*Client) {
	return func(adapter *Client) {
		adapter.options.MultiCallContractAddress = multiCallContractAddress
		adapter.options.MultiCallABI = multiCallABI
	}
}

func WithRequestMiddlewares(middlewares ...RequestMiddleware) func(*Client) {
	return func(adapter *Client) {
		adapter.requestMiddlewares = middlewares
	}
}

func WithResponseMiddlewares(middlewares ...ResponseMiddleware) func(*Client) {
	return func(adapter *Client) {
		adapter.responseMiddlewares = middlewares
	}
}

func (a *Client) NewRequest() *Request {
	return &Request{
		executor: a,
	}
}

func (a *Client) Execute(req *Request) (*Response, error) {
	for _, f := range a.requestMiddlewares {
		if err := f(a, req); err != nil {
			return nil, err
		}
	}

	rawResponse, err := a.callContract(req)
	if err != nil {
		return nil, err
	}

	resp := &Response{
		Request:     req,
		RawResponse: rawResponse,
	}

	for _, f := range a.responseMiddlewares {
		if executeErr := f(a, resp); executeErr != nil {
			return nil, executeErr
		}
	}

	return resp, err
}

func (a *Client) GetMulticallContractAddress() common.Address {
	return a.options.MultiCallContractAddress
}

func (a *Client) GetMulticallABI() *abi.ABI {
	return a.options.MultiCallABI
}

func (a *Client) BlockNumber(ctx context.Context) (uint64, error) {
	return a.adapter.BlockNumber(ctx)
}

func (a *Client) callContract(req *Request) ([]byte, error) {
	if req.BlockHash != zeroHash {
		return a.adapter.CallContractAtHash(req.Context(), req.RawCallMsg, req.BlockHash)
	}

	return a.adapter.CallContract(req.Context(), req.RawCallMsg, req.BlockNumber)
}
