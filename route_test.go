package gos

import (
	"fmt"
	"net/http"
	"testing"
)

var (
	route = NewRoute()
	mtest = Middleware{
		Name: "test",
		HandlerFunc: func(c *Context) {
			fmt.Fprintln(c.ResponseWriter, c.Request.Method, "  helloworld!")
		},
	}
	r = new(http.Request)
	w = response{}
)

func TestRouteUse(t *testing.T) {
	route.Use(mtest)
}

func TestGET(t *testing.T) {
	route.GET("/hello", funcation)
}
func TestPOST(t *testing.T) {
	route.POST("/hello", funcation)
}
func TestHEAD(t *testing.T) {
	route.HEAD("/hello", funcation)
}
func TestDELETE(t *testing.T) {
	route.DELETE("/hello", funcation)
}

func TestGetLongPathFunc(t *testing.T) {
	hello := getLongPathFunc(route.Routers[GET], "/hello")
	if hello == nil {
		t.Error("getlongpathFunc run error")
	}
}

func TestCall(t *testing.T) {
	r.RequestURI = "/hello"
	route.call(w, r, GET)
}

func TestServeHTTP(t *testing.T) {
	r.Method = http.MethodGet
	route.ServeHTTP(w, r)
	r.Method = http.MethodHead
	route.ServeHTTP(w, r)
	r.Method = http.MethodDelete
	route.ServeHTTP(w, r)
	r.Method = http.MethodPost
	route.ServeHTTP(w, r)
}
