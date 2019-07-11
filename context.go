package gos

import (
	"context"
	"net/http"
)

type Context struct {
	Ctx            context.Context
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	NextFlag       bool
}

func (c *Context) Next() bool {
	return c.NextFlag
}

func (c *Context) Fail() {
	c.NextFlag = false
}
