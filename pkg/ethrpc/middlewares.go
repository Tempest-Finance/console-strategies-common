package ethrpc

type (
	// RequestMiddleware type is for request middleware, called before a request is sent
	RequestMiddleware func(RequestExecutor, *Request) error

	// ResponseMiddleware type is for response middleware, called after a response has been received
	ResponseMiddleware func(RequestExecutor, *Response) error
)

func ParseRequestMiddleware(executor RequestExecutor, req *Request) error {
	parser, err := GetRequestParser(req.Method)
	if err != nil {
		return err
	}

	return parser(executor, req)
}

func ParseResponseMiddleware(executor RequestExecutor, resp *Response) error {
	parser, err := GetResponseParser(resp.Request.Method)
	if err != nil {
		return err
	}

	return parser(executor, resp)
}
