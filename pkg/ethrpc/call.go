package ethrpc

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Call struct {
	ABI       abi.ABI
	UnpackABI []*abi.ABI
	Target    string
	Method    string
	Params    []any
	Output    []any
}

func (c *Call) SetOutput(output []any) *Call {
	c.Output = output

	return c
}

// autofillUnpackABI fills the call's UnpackABI in case it's not set
func (c *Call) autofillUnpackABI() {
	if c.UnpackABI == nil {
		c.UnpackABI = []*abi.ABI{&c.ABI}

		return
	}

	if len(c.UnpackABI) == 0 {
		c.UnpackABI = append(c.UnpackABI, &c.ABI)

		return
	}
}
