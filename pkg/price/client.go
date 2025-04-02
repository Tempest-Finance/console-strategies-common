package price

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/Tempest-Finance/console-strategies-common/pkg/goerrors"
	"github.com/Tempest-Finance/console-strategies-common/pkg/http"
	"github.com/Tempest-Finance/console-strategies-common/pkg/logger"
	"github.com/go-resty/resty/v2"
	"github.com/shopspring/decimal"
)

type IClient interface {
	GetRealtimeTokenPriceUsd(ctx context.Context, chainID int64, tokenAddresses []string) (map[string]Token, *goerrors.Error)
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

func (c *Client) GetRealtimeTokenPriceUsd(ctx context.Context, chainID int64, tokenAddresses []string) (map[string]Token, *goerrors.Error) {
	if len(tokenAddresses) == 0 {
		return map[string]Token{}, nil
	}

	result := map[string]Token{}
	params := url.Values{}

	for _, tokenAddress := range tokenAddresses {
		params.Add("ids", fmt.Sprintf("%d_%s", chainID, tokenAddress))
	}

	for i := 0; i < 3; i++ {
		_, res, errRes, err := http.
			R[GetTokenPriceRes, string](c.httpClient).
			SetQueryParamsFromValues(params).
			SetHeader("Content-Type", "application/json").
			SetHeader("x-api-key", c.config.ApiKey).
			Get(ctx, c.config.GetPriceUrl)

		if err != nil {
			if i == 2 {
				logger.Error(ctx, err)
				return nil, goerrors.NewErrUnknown(err)
			} else {
				time.Sleep(time.Second)
				continue
			}
		}
		if errRes != nil {
			if i == 2 {
				logger.Error(ctx, errRes)
				goErr := goerrors.NewErrUnknown(fmt.Errorf("PriceService: %v", errRes))
				logger.Error(ctx, goErr)
				return nil, goErr
			} else {
				time.Sleep(time.Second)
				continue
			}
		}
		for address, price := range res.Data {
			result[address] = Token{
				ChainID:  chainID,
				Address:  strings.Split(address, ":")[1],
				PriceUsd: decimal.NewFromFloat(price),
			}
		}
		break
	}

	return result, nil
}
