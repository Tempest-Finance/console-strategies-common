package brahma

import "errors"

var (
	ErrRegisterExecutorFailed = errors.New("failed to register executor")
	ErrGetExecutorFailed      = errors.New("failed to get executor")
	ErrGetSubscriptionsFailed = errors.New("failed to get subscriptions")
	ErrExecuteTaskFailed      = errors.New("failed to execute task")
	ErrGetTaskStatusFailed    = errors.New("failed to get task status")
)

var (
	ErrGetConsoleAccountFailed = errors.New("failed to get console account")
)

var (
	ErrSyncNotFoundConsoleAccount = errors.New("sync not found console account")
)
