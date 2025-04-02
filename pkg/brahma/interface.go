package brahma

import (
	"context"

	"github.com/Tempest-Finance/console-strategies-common/pkg/crypto"
)

// IConsole is an interface that defines methods for interacting with the Brahma Console.
// It acts as a service object that encapsulates encoding, signing, and transformation logic
// to build transactions and requests. It uses an IClient implementation to connect to the Brahma server.
type IConsole interface {
	// Execute builds and executes a task using the provided parameters.
	// It encodes, signs, and transforms the transaction data, then sends the request
	// to the Brahma server via the IClient.
	Execute(ctx context.Context, params *ExecuteParams) (*TaskInfo, error)
}

// IClient is the interface for interacting with the Brahma server
// https://docs.brahma.fi/brahma-builder-kit/api-reference
type IClient interface {
	// RegisterExecutor registers an executor (an automation that users can subscribe to) with the Brahma server
	// POST /v1/vendor/automations/executor
	RegisterExecutor(ctx context.Context, reqBody *RegisterExecutorRequestBody) (*RegisterExecutorResult, error)

	// GetExecutor retrieves an executor by its address and chain ID
	// GET /v1/vendor/automations/executor/:address/:chainID
	GetExecutor(ctx context.Context, address string, chainID int) (*Executor, error)

	// GetConsoleAccounts GetConsoleAccountsPath retrieves console accounts
	// GET /v1/vendor/user/consoles/:eoa
	GetConsoleAccounts(ctx context.Context, eoa string) ([]ConsoleInfo, error)

	// GetSubscriptionsByConsoleAccountAndChainID retrieves the subscriptions for a given registry ID
	// GET /v1/vendor/automations/subscriptions/console/:address/:chainId
	GetSubscriptionsByConsoleAccountAndChainID(ctx context.Context, consoleAccount string, chainId int64) ([]Subscription, error)

	// GetSubscriptionsByRegistryID retrieves the subscriptions for a given registry ID
	// GET /v1/vendor/automations/executor/:registryID/subscriptions
	GetSubscriptionsByRegistryID(ctx context.Context, registryID string) ([]Subscription, error)

	// ExecuteTask pass an executable for a subscriber's account and execute it using Console Relayer, if it complies with the policy
	// POST /v1/vendor/automations/tasks/execute/:chainID
	ExecuteTask(ctx context.Context, chainID int64, reqBody *ExecuteTaskRequestBody) (string, error)

	// GetTaskStatus retrieves the status of a task by its ID
	// GET /v1/vendor/relayer/tasks/status/:taskId
	GetTaskStatus(ctx context.Context, taskID string) (*TaskStatus, error)
}

type IExecutorSigner interface {
	GetExecutorSigner(address string) (crypto.ISigner, bool)
}
