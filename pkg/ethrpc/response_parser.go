package ethrpc

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	ErrResponseParserIsNotRegistered = fmt.Errorf("requestParser is not registered")
	ErrReturnDataIsNotMatched        = fmt.Errorf("number of return data is not match with number of call")
	ErrUnpackMulticallFailed         = fmt.Errorf("unpack multicall failed")
)

var (
	responseParserRegistry = map[RequestMethod]ResponseParser{}
)

func RegisterResponseParser(method RequestMethod, parser ResponseParser) {
	responseParserRegistry[method] = parser
}

func GetResponseParser(method RequestMethod) (ResponseParser, error) {
	parser, ok := responseParserRegistry[method]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrResponseParserIsNotRegistered, method)
	}

	return parser, nil
}

func init() {
	RegisterResponseParser(RequestMethodCall, ResponseParserCall)
	RegisterResponseParser(RequestMethodAggregate, ResponseParserAggregate)
	RegisterResponseParser(RequestMethodTryBlockAndAggregate, ResponseParserTryBlockAndAggregate)
}

type ResponseParser func(executor RequestExecutor, req *Response) error

// ResponseParserCall parses raw call msg for RequestMethodCall
func ResponseParserCall(_ RequestExecutor, resp *Response) error {
	if len(resp.Request.Calls) != 1 {
		return ErrWrongCallParams
	}

	call := resp.Request.Calls[0]

	if err := call.ABI.UnpackIntoInterface(call.Output[0], call.Method, resp.RawResponse); err != nil {
		return err
	}

	return nil
}

// ResponseParserAggregate parse raw call msg for RequestMethodAggregate
func ResponseParserAggregate(executor RequestExecutor, resp *Response) error {
	var (
		multicallABI = executor.GetMulticallABI()

		result AggregateResult
	)

	if err := multicallABI.UnpackIntoInterface(&result, string(RequestMethodAggregate), resp.RawResponse); err != nil {
		return err
	}

	if len(result.ReturnData) != len(resp.Request.Calls) {
		return ErrReturnDataIsNotMatched
	}

	for i, c := range resp.Request.Calls {
		// result will always be true if it can reach this far
		resp.Result = append(resp.Result, true)

		if err := c.ABI.UnpackIntoInterface(c.Output[0], c.Method, result.ReturnData[i]); err != nil {
			return fmt.Errorf("%w: %w", ErrUnpackMulticallFailed, err)
		}
	}
	resp.BlockNumber = result.BlockNumber

	return nil
}

// ResponseParserTryBlockAndAggregate parse raw call msg for RequestMethodTryBlockAndAggregate
func ResponseParserTryBlockAndAggregate(executor RequestExecutor, resp *Response) error {
	var (
		multicallABI = executor.GetMulticallABI()

		result TryBlockAndAggregateResult
	)

	if err := multicallABI.UnpackIntoInterface(&result, string(RequestMethodTryBlockAndAggregate), resp.RawResponse); err != nil {
		return err
	}

	for i, c := range resp.Request.Calls {
		resp.Result = append(resp.Result, result.ReturnData[i].Success)

		if !result.ReturnData[i].Success {
			continue
		}

		if unpackSucceed := tryUnpack(c.UnpackABI, c.Output, c.Method, result.ReturnData[i].ReturnData); !unpackSucceed {
			return ErrUnpackMulticallFailed
		}
	}

	resp.BlockNumber = result.BlockNumber

	return nil
}

// tryUnpack receives a list of ABIs, try to unpack and returns true if it can unpack successfully
func tryUnpack(unpackABIs []*abi.ABI, output []any, method string, data []byte) bool {
	for i, unpackABI := range unpackABIs {
		if err := unpackABI.UnpackIntoInterface(output[i], method, data); err != nil {
			continue
		}

		return true
	}

	return false
}
