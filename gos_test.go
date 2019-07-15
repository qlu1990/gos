package gos



import (
	"fmt"
	"testing"
)

var (
	gos = NewGos()
	m   = Middleware{
		Name: "test",
		HandlerFunc: func(c *Context) {
			fmt.Println("test middlerware ")
		},
	}
	funcation = func(c *Context) {
		fmt.Fprintf(c.ResponseWriter, "hello world!")
	}
)

func TestUse(t *testing.T) {
	gos.Use(m)
}

func TestAddGet(t *testing.T) {
	gos.AddGet("/hello", funcation)
	if getMatchOne(gos.Route.Routers[GET],GetPaths("/hello")) == nil {
		t.Error("GET method map len error")
	}
}

func TestAddHead(t *testing.T) {
	gos.AddHead("/hello", funcation)
	if getMatchOne(gos.Route.Routers[HEAD],GetPaths("/hello")) == nil {
		t.Error("HEAD method map len error")
	}
}

func TestAddDelete(t *testing.T) {
	gos.AddDelete("/hello", funcation)
	if getMatchOne(gos.Route.Routers[DELETE],GetPaths("/hello")) == nil {
		t.Error("DELETE method map len error")
	}
}

func TestAddPost(t *testing.T) {
	gos.AddPost("/hello", funcation)
	if getMatchOne(gos.Route.Routers[POST],GetPaths("/hello")) == nil {
		t.Error("POST method map len error")
	}
}

func TestResponse(t *testing.T) {
	w = response{}
	Response(w, "test Response ", 200)
}
