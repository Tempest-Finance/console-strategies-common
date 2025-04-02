package response

import (
	"github.com/Tempest-Finance/console-strategies-common/pkg/goerrors"
	"github.com/Tempest-Finance/console-strategies-common/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginationResponse struct {
	TotalItems int64 `json:"totalItems"`
}

func RespondSuccess(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code:    goerrors.ClientErrCodeOK,
		Message: goerrors.ClientErrMsgOK,
		Data:    data,
	})
}

func RespondError(c *gin.Context, apiErr *goerrors.RestAPIError) {
	traceID, ok := c.Value(CtxTraceIDKey).(string)
	if ok && apiErr != nil {
		apiErr.Details = append(apiErr.Details, struct {
			TraceID string `json:"traceId"`
		}{
			TraceID: traceID,
		})
	}
	logger.Error(c, apiErr)
	c.AbortWithStatusJSON(apiErr.HttpStatus, apiErr)
}
