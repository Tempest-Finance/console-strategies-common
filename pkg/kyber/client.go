package kyber

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Tempest-Finance/console-strategies-common/pkg/goerrors"
	"github.com/Tempest-Finance/console-strategies-common/pkg/http"
	"github.com/Tempest-Finance/console-strategies-common/pkg/logger"
	"github.com/Tempest-Finance/console-strategies-common/pkg/util/blockchain"
	"github.com/go-resty/resty/v2"
)

type IClient interface {
	GetRealtimeTokenPriceUsd(ctx context.Context, chainID int64, tokenAddresses []string) (map[string]Token, *goerrors.Error)
}

type Client struct {
	httpClient *resty.Client
	config     *Config
}

var client *Client

func NewClient(httpClient *resty.Client, getPriceUrl string, data string) (*Client, error) {
	if client == nil {
		var mappings []PriceMapping
		if err := json.Unmarshal([]byte(data), &mappings); err != nil {
			return nil, err
		}
		client = &Client{
			httpClient: httpClient,
			config: &Config{
				GetPriceUrl:  getPriceUrl,
				PriceMapping: mappings,
			},
		}
	}
	return client, nil
}

func (c *Client) GetRealtimeTokenPriceUsd(ctx context.Context, chainID int64, tokenAddresses []string) (map[string]Token, *goerrors.Error) {
	if len(tokenAddresses) == 0 {
		return map[string]Token{}, nil
	}

	result := map[string]Token{}
	priceRequest := map[int64][]string{}

	for _, tokenAddress := range tokenAddresses {
		mapping := c.getPriceMapping(chainID, tokenAddress)
		if len(mapping) > 0 {
			for _, m := range mapping {
				if _, ok := priceRequest[m.ChainID]; !ok {
					priceRequest[m.ChainID] = []string{}
				}
				priceRequest[m.ChainID] = append(priceRequest[m.ChainID], m.Address)
			}
		}
		priceRequest[chainID] = append(priceRequest[chainID], tokenAddress)
	}

	for i := 0; i < 3; i++ {
		_, res, errRes, err := http.
			R[GetTokenPriceRes, string](c.httpClient).
			SetBody(priceRequest).
			SetLogReqRes(false).
			Post(ctx, c.config.GetPriceUrl)

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
				goErr := goerrors.NewErrUnknown(fmt.Errorf("KyberClient: %v", errRes))
				logger.Error(ctx, goErr)
				return nil, goErr
			} else {
				time.Sleep(time.Second)
				continue
			}
		}
		tokens := res.ToValueObjects()
		for _, token := range tokens {
			reverseMapping := c.reversePriceMapping(token.ChainID, token.Address)
			for _, m := range reverseMapping {
				if _, ok := result[blockchain.ConcatChainIDAddress(m.ChainID, m.Address)]; !ok {
					result[blockchain.ConcatChainIDAddress(m.ChainID, m.Address)] = Token{
						ChainID:  m.ChainID,
						Address:  blockchain.NormalizeAddress(m.Address),
						PriceUsd: token.PriceUsd,
					}
				}
			}
			result[blockchain.ConcatChainIDAddress(token.ChainID, token.Address)] = Token{
				ChainID:  token.ChainID,
				Address:  blockchain.NormalizeAddress(token.Address),
				PriceUsd: token.PriceUsd,
			}
		}
		break
	}

	return result, nil
}

func (c *Client) getPriceMapping(chainID int64, tokenAddress string) []TokenConversion {
	result := make([]TokenConversion, 0)

	for _, priceMapping := range c.config.PriceMapping {
		priceMappingSplit := strings.Split(priceMapping.From, ":")
		if priceMappingSplit[0] == strconv.FormatInt(chainID, 10) && priceMappingSplit[1] == tokenAddress {
			priceMappingSplitTo := strings.Split(priceMapping.To, ":")
			toChainID, _ := strconv.ParseInt(priceMappingSplitTo[0], 10, 64)
			result = append(result, TokenConversion{
				ChainID: toChainID,
				Address: priceMappingSplitTo[1],
			})
		}
	}

	return result
}

func (c *Client) reversePriceMapping(chainID int64, tokenAddress string) []TokenConversion {
	result := make([]TokenConversion, 0)

	for _, priceMapping := range c.config.PriceMapping {
		priceMappingSplit := strings.Split(priceMapping.To, ":")
		if priceMappingSplit[0] == strconv.FormatInt(chainID, 10) && priceMappingSplit[1] == tokenAddress {
			priceMappingSplitTo := strings.Split(priceMapping.From, ":")
			toChainID, _ := strconv.ParseInt(priceMappingSplitTo[0], 10, 64)
			result = append(result, TokenConversion{
				ChainID: toChainID,
				Address: priceMappingSplitTo[1],
			})
		}
	}

	return result
}
