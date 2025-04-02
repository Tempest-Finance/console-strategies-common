package brahma

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
)

type TypedDataExecutionMessage struct {
	Operation      uint8          `mapstructure:"operation" json:"operation"`
	To             common.Address `mapstructure:"to" json:"to"`
	Account        common.Address `mapstructure:"account" json:"account"`
	Executor       common.Address `mapstructure:"executor" json:"executor"`
	GasToken       common.Address `mapstructure:"gasToken" json:"gasToken"`
	RefundReceiver common.Address `mapstructure:"refundReceiver" json:"refundReceiver"`
	Value          *big.Int       `mapstructure:"value" json:"value"`
	Nonce          *big.Int       `mapstructure:"nonce" json:"nonce"`
	SafeTxGas      *big.Int       `mapstructure:"safeTxGas" json:"safeTxGas"`
	BaseGas        *big.Int       `mapstructure:"baseGas" json:"baseGas"`
	GasPrice       *big.Int       `mapstructure:"gasPrice" json:"gasPrice"`
	Data           []byte         `mapstructure:"data" json:"data"`
}

func (m TypedDataExecutionMessage) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"operation":      fmt.Sprintf("%d", m.Operation),
		"to":             m.To.Hex(),
		"account":        m.Account.Hex(),
		"executor":       m.Executor.Hex(),
		"value":          m.Value,
		"nonce":          m.Nonce,
		"data":           m.Data,
		"gasToken":       m.GasToken.Hex(),
		"refundReceiver": m.RefundReceiver.Hex(),
		"safeTxGas":      m.SafeTxGas,
		"baseGas":        m.BaseGas,
		"gasPrice":       m.GasPrice,
	}
}

func GetExecutableDigest(domain apitypes.TypedDataDomain, message TypedDataExecutionMessage) ([]byte, error) {
	typedData := apitypes.TypedData{
		Types: apitypes.Types{
			"ExecutionParams": {
				{Name: "operation", Type: "uint8"},
				{Name: "to", Type: "address"},
				{Name: "account", Type: "address"},
				{Name: "executor", Type: "address"},
				{Name: "gasToken", Type: "address"},
				{Name: "refundReceiver", Type: "address"},
				{Name: "value", Type: "uint256"},
				{Name: "nonce", Type: "uint256"},
				{Name: "safeTxGas", Type: "uint256"},
				{Name: "baseGas", Type: "uint256"},
				{Name: "gasPrice", Type: "uint256"},
				{Name: "data", Type: "bytes"},
			},
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "ExecutionParams",
		Domain:      domain,
		Message:     message.ToMap(),
	}

	hash, _, err := apitypes.TypedDataAndHash(typedData)
	return hash, err
}
