package price

import (
	"context"
	"fmt"
	"time"

	"github.com/Tempest-Finance/console-strategies-common/pkg/goerrors"
	"github.com/Tempest-Finance/console-strategies-common/pkg/http"
	"github.com/Tempest-Finance/console-strategies-common/pkg/logger"
	"github.com/go-resty/resty/v2"
)

type IClient interface {
	GetRealtimeTokenPriceUsd(ctx context.Context, chainID int64, tokenAddresses []string, timestamp int64) (map[string]Token, *goerrors.Error)
}

type Client struct {
	httpClient *resty.Client
	config     *Config
}

var client *Client

func NewClient(httpClient *resty.Client, getPriceUrl, apiKey string) *Client {
	if client == nil {
		client = &Client{
			httpClient: httpClient,
			config: &Config{
				GetPriceUrl: getPriceUrl,
				ApiKey:      apiKey,
			},
		}
	}
	return client
}

func (c *Client) GetRealtimeTokenPriceUsd(
	ctx context.Context,
	chainID int64,
	tokenAddresses []string,
	timestamp int64,
) (map[string]Token, *goerrors.Error) {
	if len(tokenAddresses) == 0 {
		return map[string]Token{}, nil
	}

	reqBody := make([]Request, 0, len(tokenAddresses))
	for _, addr := range tokenAddresses {
		reqBody = append(reqBody, Request{
			ChainID:      fmt.Sprint(chainID),
			TokenAddress: addr,
			Timestamp:    timestamp,
		})
	}

	result := make(map[string]Token, len(tokenAddresses))

	// retry up to 3 times
	for i := 0; i < 3; i++ {
		_, resp, errRes, err := http.
			R[GetTokenPriceRes, string](c.httpClient).
			SetHeader("Content-Type", "application/json").
			SetHeader("x-api-key", c.config.ApiKey).
			SetBody(reqBody).
			Post(ctx, c.config.GetPriceUrl)

		if err != nil {
			if i == 2 {
				logger.Error(ctx, err)
				return nil, goerrors.NewErrUnknown(err)
			}
			time.Sleep(time.Second)
			continue
		}

		if errRes != nil {
			if i == 2 {
				logger.Error(ctx, errRes)
				goErr := goerrors.NewErrUnknown(fmt.Errorf("PriceService: %v", errRes))
				logger.Error(ctx, goErr)
				return nil, goErr
			}
			time.Sleep(time.Second)
			continue
		}

		for _, p := range resp.Data {
			key := fmt.Sprintf("%d:%s", chainID, p.Address)
			result[key] = p
		}

		break
	}

	return result, nil
}
