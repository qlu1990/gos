package gos

import (
	"context"
	"net/http"
)

// Context context for handle func
type Context struct {
	Ctx            context.Context
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	NextFlag       bool
	Params         map[string]string
}

// Next get flag for runnext
func (c *Context) Next() bool {
	return c.NextFlag
}

// Fail Set fail on flag do'nt run next handle
func (c *Context) Fail() {
	c.NextFlag = false
}

// Param get param in url
func (c *Context) Param(name string) string {
	result, _ := c.Params[name]
	return result
}
