package http

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/Tempest-Finance/console-strategies-common/pkg/logger"
	"github.com/go-resty/resty/v2"
)

type Request[R, E any] struct {
	LogReqRes bool

	*resty.Request
}

func (r *Request[R, E]) SetLogReqRes(val bool) *Request[R, E] {
	r.LogReqRes = val
	return r
}

// Execute execute the request
// @return HTTP Status Code, Successful response body, Error response body, error
func (r *Request[R, E]) Execute(ctx context.Context, method string, url string) (int, *R, *E, error) {
	r.SetContext(ctx)
	reqBodyJson, _ := json.Marshal(r.Body)

	res, err := r.Request.Execute(method, url)

	if err != nil {
		if r.LogReqRes {
			logger.Warnf(ctx, "[HTTP Client] Send %s request to %s error, request params: %s, request body: %s, err: %s", method, url, r.QueryParam.Encode(), reqBodyJson, err)
		} else {
			logger.Warnf(ctx, "[HTTP Client] Send %s request to %s error, err: %s", method, url, err)
		}
		return 0, nil, nil, err
	}

	if res.StatusCode() >= http.StatusBadRequest {
		if r.LogReqRes {
			logger.Debugf(ctx, "[HTTP Client] Send %s request to %s error, request params: %s, request body: %s, response status: %s, response body: %s", method, url, r.QueryParam.Encode(), reqBodyJson, res.Status(), res.Body())
		} else {
			logger.Debugf(ctx, "[HTTP Client] Send %s request to %s error, status: %s. ", method, url, res.Status())
		}

		var errRes E
		if err := json.Unmarshal(res.Body(), &errRes); err != nil {
			var ok bool
			if errRes, ok = any(string(res.Body())).(E); !ok {
				if r.LogReqRes {
					logger.Debugf(ctx, "[HTTP Client] Can't unmarshal response body, res: %s", res.Body())
				} else {
					logger.Debugf(ctx, "[HTTP Client] Can't unmarshal response body")
				}
				return 0, nil, nil, err
			}
		}
		return res.StatusCode(), nil, &errRes, nil
	}

	var resData R
	data := res.Body()
	if err := json.Unmarshal(data, &resData); err != nil {
		var ok bool
		if resData, ok = any(string(res.Body())).(R); !ok {
			if r.LogReqRes {
				logger.Debugf(ctx, "[HTTP Client] Can't unmarshal response body, res: %s", res.Body())
			} else {
				logger.Debugf(ctx, "[HTTP Client] Can't unmarshal response body")
			}
			return 0, nil, nil, err
		}
	}

	if r.LogReqRes {
		logger.Debugf(ctx, "[HTTP Client] Send %s request to %s successfully, request params: %s, request body: %s, response status: %s, response body: %s", method, url, r.QueryParam.Encode(), reqBodyJson, res.Status(), res.Body())
	} else {
		logger.Debugf(ctx, "[HTTP Client] Send %s request to %s successfully", method, url)
	}

	return res.StatusCode(), &resData, nil, nil
}

func (r *Request[R, E]) Get(ctx context.Context, url string) (int, *R, *E, error) {
	return r.Execute(ctx, http.MethodGet, url)
}

func (r *Request[R, E]) Post(ctx context.Context, url string) (int, *R, *E, error) {
	return r.Execute(ctx, http.MethodPost, url)
}

func (r *Request[R, E]) Put(ctx context.Context, url string) (int, *R, *E, error) {
	return r.Execute(ctx, http.MethodPut, url)
}

func (r *Request[R, E]) Patch(ctx context.Context, url string) (int, *R, *E, error) {
	return r.Execute(ctx, http.MethodPatch, url)
}

func (r *Request[R, E]) Delete(ctx context.Context, url string) (int, *R, *E, error) {
	return r.Execute(ctx, http.MethodDelete, url)
}

// SetHeader override resty.Request.SetHeader
func (r *Request[R, E]) SetHeader(header, value string) *Request[R, E] {
	r.Request.SetHeader(header, value)
	return r
}

// SetHeaders override resty.Request.SetHeaders
func (r *Request[R, E]) SetHeaders(headers map[string]string) *Request[R, E] {
	r.Request.SetHeaders(headers)
	return r
}

// SetHeaderMultiValues override resty.Request.SetHeaderMultiValues
func (r *Request[R, E]) SetHeaderMultiValues(headers map[string][]string) *Request[R, E] {
	r.Request.SetHeaderMultiValues(headers)
	return r
}

// SetHeaderVerbatim override resty.Request.SetHeaderVerbatim
func (r *Request[R, E]) SetHeaderVerbatim(header, value string) *Request[R, E] {
	r.Request.SetHeaderVerbatim(header, value)
	return r
}

// SetQueryParam override resty.Request.SetQueryParam
func (r *Request[R, E]) SetQueryParam(param, value string) *Request[R, E] {
	r.Request.SetQueryParam(param, value)
	return r
}

// SetQueryParams override resty.Request.SetQueryParams
func (r *Request[R, E]) SetQueryParams(params map[string]string) *Request[R, E] {
	r.Request.SetQueryParams(params)
	return r
}

// SetQueryParamsFromValues override resty.Request.SetQueryParamsFromValues
func (r *Request[R, E]) SetQueryParamsFromValues(params url.Values) *Request[R, E] {
	r.Request.SetQueryParamsFromValues(params)
	return r
}

// SetQueryString override resty.Request.SetQueryString
func (r *Request[R, E]) SetQueryString(query string) *Request[R, E] {
	r.Request.SetQueryString(query)
	return r
}

// SetFormData override resty.Request.SetFormData
func (r *Request[R, E]) SetFormData(data map[string]string) *Request[R, E] {
	r.Request.SetFormData(data)
	return r
}

// SetFormDataFromValues override resty.Request.SetFormDataFromValues
func (r *Request[R, E]) SetFormDataFromValues(data url.Values) *Request[R, E] {
	r.Request.SetFormDataFromValues(data)
	return r
}

// SetBody override resty.Request.SetBody
func (r *Request[R, E]) SetBody(body interface{}) *Request[R, E] {
	r.Request.SetBody(body)
	return r
}

// SetResult override resty.Request.SetResult
func (r *Request[R, E]) SetResult(res interface{}) *Request[R, E] {
	r.Request.SetResult(res)
	return r
}

// SetError override resty.Request.SetError
func (r *Request[R, E]) SetError(err interface{}) *Request[R, E] {
	r.Request.SetError(err)
	return r
}

// SetFile override resty.Request.SetFile
func (r *Request[R, E]) SetFile(param, filePath string) *Request[R, E] {
	r.Request.SetFile(param, filePath)
	return r
}

// SetFiles override resty.Request.SetFiles
func (r *Request[R, E]) SetFiles(files map[string]string) *Request[R, E] {
	r.Request.SetFiles(files)
	return r
}

// SetFileReader override resty.Request.SetFileReader
func (r *Request[R, E]) SetFileReader(param, fileName string, reader io.Reader) *Request[R, E] {
	r.Request.SetFileReader(param, fileName, reader)
	return r
}

// SetMultipartFormData override resty.Request.SetMultipartFormData
func (r *Request[R, E]) SetMultipartFormData(data map[string]string) *Request[R, E] {
	r.Request.SetMultipartFormData(data)
	return r
}

// SetMultipartField override resty.Request.SetMultipartField
func (r *Request[R, E]) SetMultipartField(param, fileName, contentType string, reader io.Reader) *Request[R, E] {
	r.Request.SetMultipartField(param, fileName, contentType, reader)
	return r
}

// SetMultipartFields override resty.Request.SetMultipartFields
func (r *Request[R, E]) SetMultipartFields(fields ...*resty.MultipartField) *Request[R, E] {
	r.Request.SetMultipartFields(fields...)
	return r
}

// SetContentLength override request.Request.SetContentLength
func (r *Request[R, E]) SetContentLength(l bool) *Request[R, E] {
	r.Request.SetContentLength(l)
	return r
}

// SetBasicAuth override resty.Request.SetBasicAuth
func (r *Request[R, E]) SetBasicAuth(username, password string) *Request[R, E] {
	r.Request.SetBasicAuth(username, password)
	return r
}

// SetAuthToken override resty.Request.SetAuthToken
func (r *Request[R, E]) SetAuthToken(token string) *Request[R, E] {
	r.Request.SetAuthToken(token)
	return r
}

// SetAuthScheme override resty.Request.SetAuthScheme
func (r *Request[R, E]) SetAuthScheme(scheme string) *Request[R, E] {
	r.Request.SetAuthScheme(scheme)
	return r
}

// SetOutput override resty.Request.SetOutput
func (r *Request[R, E]) SetOutput(file string) *Request[R, E] {
	r.Request.SetOutput(file)
	return r
}

// SetSRV override resty.Request.SetSRV
func (r *Request[R, E]) SetSRV(srv *resty.SRVRecord) *Request[R, E] {
	r.Request.SetSRV(srv)
	return r
}

// SetDoNotParseResponse override resty.Request.SetDoNotParseResponse
func (r *Request[R, E]) SetDoNotParseResponse(parse bool) *Request[R, E] {
	r.Request.SetDoNotParseResponse(parse)
	return r
}

// SetPathParam override resty.Request.SetPathParam
func (r *Request[R, E]) SetPathParam(param, value string) *Request[R, E] {
	r.Request.SetPathParam(param, value)
	return r
}

// SetPathParams override resty.Request.SetPathParams
func (r *Request[R, E]) SetPathParams(params map[string]string) *Request[R, E] {
	r.Request.SetPathParams(params)
	return r
}

// ExpectContentType override resty.Request.ExpectContentType
func (r *Request[R, E]) ExpectContentType(contentType string) *Request[R, E] {
	r.Request.ExpectContentType(contentType)
	return r
}

// ForceContentType override resty.Request.ForceContentType
func (r *Request[R, E]) ForceContentType(contentType string) *Request[R, E] {
	r.Request.ForceContentType(contentType)
	return r
}

// SetJSONEscapeHTML override resty.Request.SetJSONEscapeHTML
func (r *Request[R, E]) SetJSONEscapeHTML(b bool) *Request[R, E] {
	r.Request.SetJSONEscapeHTML(b)
	return r
}

// SetCookie override resty.Request.SetCookie
func (r *Request[R, E]) SetCookie(hc *http.Cookie) *Request[R, E] {
	r.Request.SetCookie(hc)
	return r
}

// SetCookies override resty.Request.SetCookies
func (r *Request[R, E]) SetCookies(rs []*http.Cookie) *Request[R, E] {
	r.Request.SetCookies(rs)
	return r
}

// AddRetryCondition override resty.Request.AddRetryCondition
func (r *Request[R, E]) AddRetryCondition(condition resty.RetryConditionFunc) *Request[R, E] {
	r.Request.AddRetryCondition(condition)
	return r
}
