package brahma

import (
	"math/big"
	"time"

	"github.com/Brahma-fi/go-safe/types"
	"github.com/Tempest-Finance/console-strategies-common/pkg/crypto"
	"github.com/ethereum/go-ethereum/common"
)

// Common =============================================================================================================="

type Transaction struct {
	Target common.Address `json:"to"`
	Val    *big.Int       `json:"value"`
	Data   string         `json:"data"`
}

func (t *Transaction) From() common.Address {
	return common.HexToAddress("")
}

func (t *Transaction) To() common.Address {
	return t.Target
}

func (t *Transaction) CallData() string {
	return t.Data
}

func (t *Transaction) Value() *big.Int {
	return t.Val
}

func (t *Transaction) Operation() uint8 {
	// Call Only
	return 0
}

// Executor Signer =============================================================================================================="

type ExecutorSignerConfig struct {
	Address string `json:"address" mapstructure:"address"`
	Name    string `json:"name" mapstructure:"name"`
}

// Executor =============================================================================================================="

type Executor struct {
	Config    ExecutorConfig   `json:"config"`
	Executor  string           `json:"executor"`
	Signature string           `json:"signature"`
	ChainId   int              `json:"chainId"`
	Timestamp int              `json:"timestamp"`
	Metadata  ExecutorMetadata `json:"executorMetadata"`
	Id        string           `json:"id"`
	Status    int              `json:"status"`
}

type ExecutorConfig struct {
	// FeeInBPS fees collected by the automation in BPS; MAX_BPS=4
	FeeInBPS int `json:"feeInBPS"`

	// FeeReceiver address of the automation client's fee receiver
	FeeReceiver string `json:"feeReceiver"`

	// LimitPerExecution boolean denoting if the user should limit the input tokens being pulled by executor per execution or for a specific duration of time; true if the limit is per execution
	LimitPerExecution bool `json:"limitPerExecution"`

	// FeeToken address of the token to collect fees in
	FeeToken string `json:"feeToken"`

	// InputTokens list of addresses of input tokens for the automation
	InputTokens []string `json:"inputTokens"`

	// HopAddresses list of whitelisted addresses through which the token is allowed to be transferred during the automations
	HopAddresses []string `json:"hopAddresses"`
}

type ExecutorMetadata struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type ExecuteParams struct {
	ChainID                  int64
	MultiSendCallOnlyAddress common.Address
	ExecutorAddress          common.Address
	SubAccount               common.Address
	Signer                   crypto.ISigner
	Transactions             []types.Transaction
}

// Subscription =============================================================================================================="

type Subscription struct {
	ChainId           int64             `json:"chainId"`
	CommitHash        string            `json:"commitHash"`
	CreatedAt         time.Time         `json:"createdAt"`
	Duration          int               `json:"duration"`
	FeeAmount         string            `json:"feeAmount"`
	FeeToken          string            `json:"feeToken"`
	Id                string            `json:"id"`
	Metadata          map[string]string `json:"metadata"`
	RegistryId        string            `json:"registryId"`
	Status            int               `json:"status"`
	SubAccountAddress string            `json:"subAccountAddress"`
	TokenInputs       map[string]string `json:"tokenInputs"`
	TokenLimits       map[string]string `json:"tokenLimits"`
}

type Executable struct {
	// CallType type of call to make; CALL(0), DELEGATECALL(1), or STATICCALL(2)
	CallType uint8 `json:"callType"`

	// To address of the target to execute on
	To string `json:"to"`

	// Value amount of ETH to transfer from subscriber to `to`
	Value string `json:"value"`

	// Data calldata to execute on `to`
	Data string `json:"data"`
}

type TaskStatus struct {
	TaskId   string `json:"taskId"`
	Metadata struct {
		Request struct {
			TaskId             string `json:"taskId"`
			To                 string `json:"to"`
			CallData           string `json:"callData"`
			RequestedAt        int    `json:"requestedAt"`
			Timeout            string `json:"timeout"`
			Signer             string `json:"signer"`
			ChainID            string `json:"chainID"`
			UseSafeGasEstimate bool   `json:"useSafeGasEstimate"`
			MaxGasLimit        int    `json:"maxGasLimit"`
			EnableAccessList   bool   `json:"enableAccessList"`
			BackendId          string `json:"backendId"`
			Webhook            string `json:"webhook"`
		} `json:"request"`
		Response struct {
			IsSuccessful    bool        `json:"isSuccessful"`
			Error           string      `json:"error"`
			TransactionHash interface{} `json:"transactionHash"`
		} `json:"response"`
	} `json:"metadata"`
	OutputTransactionHash string    `json:"outputTransactionHash"`
	Status                string    `json:"status"`
	CreatedAt             time.Time `json:"createdAt"`
}

// =====================================================================================================================
// RegisterExecutor ====================================================================================================

type RegisterExecutorRequestBody struct {
	// Config Configuration specifications for client automation
	Config ExecutorConfig `json:"config"`

	// Executor APIBindAddress of executor that runs the automation
	Executor string `json:"executor"`

	// Signature of the client configuration message signed by the executor
	Signature string `json:"signature"`

	// ChainId Chain ID to create client automation on
	ChainId int `json:"chainId"`

	// Timestamp Current timestamp. Must match the timestamp that is signed by executor
	Timestamp int64 `json:"timestamp"`

	// ExecutorMetadata Arbitrary metadata JSON of data that the executor wishes to accept from users at time of subscription
	ExecutorMetadata any `json:"executorMetadata"`
}

type RegisterExecutorResult struct {
	Data *Executor `json:"data"`
}

// =====================================================================================================================
// GetExecutor ========================================================================================================

type GetExecutorResult struct {
	Data *Executor `json:"data"`
}

// =====================================================================================================================
// GetSubscriptionsByRegistryID ====================================================================================================

type GetSubscriptionsResult struct {
	Data []Subscription `json:"data"`
}

// =====================================================================================================================
// ExecuteTask =========================================================================================================

type ExecuteTaskRequest struct {
	ChainID int64                  `json:"-"`
	Task    ExecuteTaskRequestBody `json:"task"`
	Webhook string                 `json:"webhook"`
}

type ExecuteTaskRequestBody struct {
	SubAccount        string     `json:"subaccount"`
	Executor          string     `json:"executor"`
	ExecutorSignature string     `json:"executorSignature"`
	Executable        Executable `json:"executable"`
}

type ExecuteTaskResult struct {
	Data struct {
		Data struct {
			TaskID                string `json:"taskId"`
			OutputTransactionHash string `json:"outputTransactionHash"`
		} `json:"data"`
	}
}

// =====================================================================================================================
// GetTaskStatus =========================================================================================================

type GetTaskStatusResult struct {
	Data *TaskStatus `json:"data"`
}

// =====================================================================================================================
// GetConsoleAccount ===================================================================================================

type GetConsoleAccountResult struct {
	Data []ConsoleInfo `json:"data"`
}

type ConsoleInfo struct {
	ConsoleAccount string    `json:"consoleAddress"`
	EOA            string    `json:"eoa"`
	ChainId        int64     `json:"chainId"`
	CreatedAt      time.Time `json:"createdAt"`
}

type TaskInfo struct {
	TaskId string `json:"taskId"`
	TxHash string `json:"TxHash"`
}
