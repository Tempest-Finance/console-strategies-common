package nucleus

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	client  *resty.Client
	chainId int64

	NucleusAPIKey string
	BaseURL       string
	AddressBook   map[int64]NetworkData
}

func NewClient(nucleusAPIKey, baseURL string) (*Client, error) {
	if baseURL == "" {
		baseURL = DefaultBaseUrl
	}

	client := resty.New().SetBaseURL(baseURL).
		SetHeader("Content-Type", "application/json").
		SetHeader("x-api-key", nucleusAPIKey)

	var addressBook map[int64]NetworkData
	resp, err := client.R().SetResult(&addressBook).Get(AddressBookUrl)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New(fmt.Sprintf("Failed to fetch address book: %s", resp.Status()))
	}

	return &Client{
		client:        client,
		NucleusAPIKey: nucleusAPIKey,
		BaseURL:       baseURL,
		AddressBook:   addressBook,
	}, nil
}

func (c *Client) Get(ctx context.Context, endpoint string, queryParams map[string]string) (interface{}, error) {
	return c.request(ctx, "GET", endpoint, nil, queryParams)
}

func (c *Client) Post(ctx context.Context, endpoint string, jsonData []byte) (interface{}, error) {
	return c.request(ctx, "POST", endpoint, jsonData, nil)
}

func (c *Client) GetAddressBook() map[int64]NetworkData {
	return c.AddressBook
}

func (c *Client) request(ctx context.Context, method, endpoint string, jsonData []byte, queryParams map[string]string) (interface{}, error) {
	var result interface{}
	req := c.client.R().
		SetContext(ctx).
		SetResult(&result)

	if jsonData != nil {
		req.SetBody(jsonData)
	}

	if len(queryParams) > 0 {
		req.SetQueryParams(queryParams)
	}

	resp, err := req.Execute(method, endpoint)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, ErrFailedToExecute
	}

	return result, nil
}
