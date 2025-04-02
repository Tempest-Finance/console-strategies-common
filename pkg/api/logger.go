package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap/zapcore"
	"strings"

	"github.com/Tempest-Finance/console-strategies-common/pkg/logger"
)

func buildLogFields(c *gin.Context) (zapcore.Field, zapcore.Field) {
	traceIDField := zapcore.Field{
		Key:    CtxTraceIDKey,
		Type:   zapcore.StringType,
		String: uuid.New().String(),
	}

	builder := strings.Builder{}
	builder.WriteString(c.Request.Method)
	builder.WriteString(" ")
	builder.WriteString(c.Request.URL.Path)
	raw := c.Request.URL.RawQuery
	if raw != "" {
		builder.WriteString("?")
		builder.WriteString(raw)
	}
	apiField := zapcore.Field{
		Key:    CtxAPIRequestKey,
		Type:   zapcore.StringType,
		String: builder.String(),
	}
	return traceIDField, apiField
}

// Logger add a logger to gin context with metadata like traceID, etc.
func Logger(c *gin.Context) {
	traceIDField, apiField := buildLogFields(c)
	l := logger.L().With(traceIDField).With(apiField)

	c.Set(CtxLoggerKey, l)
	c.Set(CtxTraceIDKey, traceIDField.String)

	// Process request
	c.Next()
}
