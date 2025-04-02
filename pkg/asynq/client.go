package asynq

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/hibiken/asynq"
)

type IClient interface {
	EnqueueTask(ctx context.Context, taskType string, taskID string, queueID string, payload interface{}, maxRetry, timeoutMinutes int, processAt time.Time, forceEnqueue bool) error
	DeleteQueue(ctx context.Context, queueID string, taskId string) error
	GetTaskInfo(ctx context.Context, queueID string, taskId string) (*asynq.TaskInfo, error)
}

type Client struct {
	client    *asynq.Client
	inspector *asynq.Inspector
}

var client *Client

func NewClient(config Config) (*Client, error) {
	if client == nil {
		redisConnOpt, err := GetAsynqRedisConnectionOption(config)
		if err != nil {
			return nil, err
		}

		client = &Client{
			client:    asynq.NewClient(redisConnOpt),
			inspector: asynq.NewInspector(redisConnOpt),
		}
	}

	return client, nil
}

func (c *Client) EnqueueTask(
	ctx context.Context, taskType string, taskID string, queueID string,
	payload interface{}, maxRetry int, timeoutMinutes int, processAt time.Time, forceEnqueue bool,
) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	var opts []asynq.Option
	if taskID != "" {
		opts = append(opts, asynq.TaskID(taskID))
	}
	if queueID != "" {
		opts = append(opts, asynq.Queue(queueID))
	}
	if maxRetry >= 0 {
		opts = append(opts, asynq.MaxRetry(maxRetry))
	}
	if timeoutMinutes >= 0 {
		opts = append(opts, asynq.Timeout(time.Minute*time.Duration(timeoutMinutes)))
	}
	if !processAt.IsZero() {
		opts = append(opts, asynq.ProcessAt(processAt))
	}

	if forceEnqueue {
		_ = c.DeleteQueue(ctx, queueID, taskID)
	}
	task := asynq.NewTask(taskType, payloadBytes, opts...)
	_, err = c.client.Enqueue(task)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteQueue(_ context.Context, queueID string, taskId string) error {
	return c.inspector.DeleteTask(queueID, taskId)
}

func (c *Client) GetTaskInfo(_ context.Context, queueID string, taskId string) (*asynq.TaskInfo, error) {
	info, err := c.inspector.GetTaskInfo(queueID, taskId)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func GetAsynqRedisConnectionOption(config Config) (asynq.RedisConnOpt, error) {
	redisAddresses := config.InitAddress

	if len(redisAddresses) == 0 {
		return nil, errors.New("redis host is empty")
	}

	if len(redisAddresses) == 1 {
		return asynq.RedisClientOpt{
			Addr:     redisAddresses[0],
			Password: config.Password,
		}, nil
	}

	return asynq.RedisClusterClientOpt{
		Addrs:    redisAddresses,
		Password: config.Password,
	}, nil
}
