package context

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	spanCtx context.Context // tracing span context.
}

func (c *Context) SpanCtx() context.Context {
	return c.spanCtx
}

func New(c *gin.Context) *Context {
	if c == nil {
		return NewDefault()
	}
	return &Context{
		Context: c,
		spanCtx: c.Request.Context(),
	}
}

func NewDefault() *Context {
	return &Context{
		Context: &gin.Context{},
		spanCtx: context.Background(),
	}
}

func (c *Context) New(spanCtx context.Context) *Context {
	return &Context{
		Context: c.Context,
		spanCtx: spanCtx,
	}
}

func (c *Context) Abstract() context.Context {
	return c
}

func Parse(ctx context.Context) (*Context, bool) {
	if c, ok := ctx.(*Context); ok {
		return c, true
	}
	return nil, false
}

func ChildCtx(ctx context.Context, spanCtx context.Context) (context.Context, bool) {
	c, ok := Parse(ctx)
	if !ok {
		return spanCtx, false
	}
	return c.New(spanCtx), true
}
