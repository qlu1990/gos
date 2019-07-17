package gos

import (
	"context"
	"encoding/json"
	"io/ioutil"
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

// URIParam get param in uri
func (c *Context) URIParam(name string) string {
	result, _ := c.Params[name]
	return result
}

// Param get param in url
func (c *Context) Param(name string) []string {
	c.Request.ParseForm()
	return c.Request.Form[name]
}

//PostParam get postParam
func (c *Context) PostParam(name string) []string {
	c.Request.ParseForm()
	return c.Request.PostForm[name]
}

//PostBodyBind Unmarshal body body is point for struct or map
func (c *Context) PostBodyBind(body interface{}) error {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, body)
	return err
}
