package brahma

const (
	RegisterExecutorPath                           = "/v1/vendor/automations/executor"
	GetExecutorPath                                = "/v1/vendor/automations/executor/{address}/{chainID}"
	GetSubscriptionsByRegistryIDPath               = "/v1/vendor/automations/executor/{registryID}/subscriptions"
	ExecuteTaskPath                                = "/v1/vendor/automations/tasks/execute/{chainID}"
	GetTaskStatusPath                              = "/v1/vendor/relayer/tasks/status/{taskID}"
	GetSubscriptionsByConsoleAddressAndChainIDPath = "/v1/vendor/automations/subscriptions/console/{address}/{chainId}"
	GetConsoleAccountsPath                         = "/v1/vendor/user/consoles/{eoa}"
)

const (
	TaskStatusCancelled  = "cancelled"
	TaskStatusExecuting  = "executing"
	TaskStatusPending    = "pending"
	TaskStatusSuccessful = "successful"
)

const (
	TaskTimeoutInSecond = 60 * 3
)

const (
	SubscriptionStatusActive   = 2
	SubscriptionStatusInactive = 4
)
