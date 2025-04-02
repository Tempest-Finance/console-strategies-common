package brahma

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
}

func NewClient(url string, apiKey string) *Client {
	return &Client{
		client: resty.New().
			SetBaseURL(url).
			SetHeader("Content-Type", "application/json").
			SetHeader("x-api-key", apiKey),
	}
}

func (c *Client) RegisterExecutor(ctx context.Context, reqBody *RegisterExecutorRequestBody) (*RegisterExecutorResult, error) {
	var result RegisterExecutorResult
	req := c.client.R().SetContext(ctx).SetBody(reqBody)
	resp, err := req.SetResult(&result).Post(RegisterExecutorPath)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, ErrRegisterExecutorFailed
	}

	return &result, nil
}

func (c *Client) GetExecutor(ctx context.Context, address string, chainID int) (*Executor, error) {
	var result GetExecutorResult
	req := c.client.R().SetContext(ctx)
	resp, err := req.SetResult(&result).
		SetPathParam("address", address).
		SetPathParam("chainID", strconv.Itoa(chainID)).
		Get(GetExecutorPath)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, ErrGetExecutorFailed
	}

	return result.Data, nil
}

func (c *Client) GetSubscriptionsByRegistryID(ctx context.Context, registryID string) ([]Subscription, error) {
	var result GetSubscriptionsResult
	req := c.client.R().SetContext(ctx)
	resp, err := req.SetResult(&result).
		SetPathParam("registryID", registryID).
		Get(GetSubscriptionsByRegistryIDPath)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, ErrGetSubscriptionsFailed
	}

	subscriptions := removeDuplicateSubscriptions(result.Data)

	return subscriptions, nil
}

func (c *Client) ExecuteTask(ctx context.Context, chainID int64, reqBody *ExecuteTaskRequestBody) (string, error) {
	var result ExecuteTaskResult
	req := c.client.R().SetContext(ctx)

	resp, err := req.SetResult(&result).
		SetPathParam("chainID", strconv.FormatInt(chainID, 10)).
		SetBody(&ExecuteTaskRequest{
			ChainID: chainID,
			Task: ExecuteTaskRequestBody{
				SubAccount:        reqBody.SubAccount,
				Executor:          reqBody.Executor,
				ExecutorSignature: reqBody.ExecutorSignature,
				Executable:        reqBody.Executable,
			},
			Webhook: "",
		}).
		Post(ExecuteTaskPath)
	if err != nil {
		return "", err
	}

	if !resp.IsSuccess() {
		return "", fmt.Errorf("failed to execute task: %s", resp.String())
	}

	return result.Data.Data.TaskID, nil
}

func (c *Client) GetTaskStatus(ctx context.Context, taskID string) (*TaskStatus, error) {
	var result GetTaskStatusResult
	req := c.client.R().SetContext(ctx)
	resp, err := req.SetResult(&result).
		SetPathParam("taskID", taskID).
		Get(GetTaskStatusPath)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, ErrGetTaskStatusFailed
	}

	return result.Data, nil
}

func (c *Client) GetConsoleAccounts(ctx context.Context, eoa string) ([]ConsoleInfo, error) {
	var result GetConsoleAccountResult
	req := c.client.R().SetContext(ctx)
	resp, err := req.SetResult(&result).
		SetPathParam("eoa", eoa).
		Get(GetConsoleAccountsPath)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, ErrGetConsoleAccountFailed
	}

	return result.Data, nil
}

func (c *Client) GetSubscriptionsByConsoleAccountAndChainID(ctx context.Context, consoleAccount string, chainID int64) ([]Subscription, error) {
	var result GetSubscriptionsResult
	req := c.client.R().SetContext(ctx)
	resp, err := req.SetResult(&result).
		SetPathParam("address", consoleAccount).
		SetPathParam("chainId", strconv.Itoa(int(chainID))).
		Get(GetSubscriptionsByConsoleAddressAndChainIDPath)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, ErrGetSubscriptionsFailed
	}

	return result.Data, nil
}

func removeDuplicateSubscriptions(subscriptions []Subscription) []Subscription {
	subscriptionMap := make(map[string]Subscription)
	for _, subscription := range subscriptions {
		if _, ok := subscriptionMap[subscription.SubAccountAddress]; !ok {
			subscriptionMap[subscription.SubAccountAddress] = subscription
		} else {
			if subscription.CreatedAt.After(subscriptionMap[subscription.SubAccountAddress].CreatedAt) {
				subscriptionMap[subscription.SubAccountAddress] = subscription
			}
		}
	}

	var uniqueSubscriptions []Subscription
	for _, subscription := range subscriptionMap {
		uniqueSubscriptions = append(uniqueSubscriptions, subscription)
	}

	return uniqueSubscriptions
}
