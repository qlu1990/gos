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
	route.use(mtest)
}

func TestGet(t *testing.T) {
	route.get("/hello", funcation)
}
func TestPost(t *testing.T) {
	route.post("/hello", funcation)
}
func TestHead(t *testing.T) {
	route.head("/hello", funcation)
}
func TestDelete(t *testing.T) {
	route.delete("/hello", funcation)
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
