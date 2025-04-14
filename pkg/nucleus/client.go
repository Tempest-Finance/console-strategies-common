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

	nucleusAPIKey string
	baseURL       string
	addressBook   map[int64]NetworkData
}

func NewClient(nucleusAPIKey, baseURL string) (*Client, error) {
	if baseURL == "" {
		baseURL = DefaultBaseUrl
	}

	client := resty.New().SetBaseURL(baseURL).
		SetHeader("Content-Type", "application/json").
		SetHeader("x-api-key", nucleusAPIKey)

	resp, err := client.R().Get(AddressBookUrl)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New(fmt.Sprintf("Failed to fetch address book: %s", resp.Status()))
	}

	addressBook, err := parseAddressBook(resp.Body())
	if err != nil {
		return nil, err
	}

	return &Client{
		client:        client,
		nucleusAPIKey: nucleusAPIKey,
		baseURL:       baseURL,
		addressBook:   addressBook,
	}, nil
}

func (c *Client) Get(ctx context.Context, endpoint string, queryParams map[string]string) (interface{}, error) {
	return c.request(ctx, "GET", endpoint, nil, queryParams)
}

func (c *Client) Post(ctx context.Context, endpoint string, jsonData []byte) (interface{}, error) {
	return c.request(ctx, "POST", endpoint, jsonData, nil)
}

func (c *Client) GetAddressBook() map[int64]NetworkData {
	return c.addressBook
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
