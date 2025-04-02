package http

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func NewClient(retryCount, retryWaitTimeSeconds, retryMaxWaitTimeSeconds int) *resty.Client {
	var client = resty.New()
	client.
		SetRetryCount(retryCount).
		SetRetryWaitTime(time.Duration(retryWaitTimeSeconds) * time.Second).
		SetRetryMaxWaitTime(time.Duration(retryMaxWaitTimeSeconds) * time.Second).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				if r == nil {
					return false
				}
				return r.StatusCode() >= 500
			},
		)
	return client
}

func R[R, E any](c *resty.Client) *Request[R, E] {
	restyReq := c.R()
	r := &Request[R, E]{
		LogReqRes: true,
		Request:   restyReq,
	}
	return r
}
