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
	if len(gos.Route.Routers[GET]) != 0 {
		t.Error("GET method map len error")
	}
	gos.AddGet("/hello", funcation)
	if len(gos.Route.Routers[GET]) == 0 {
		t.Error("GET method map len error")
	}
}

func TestAddHead(t *testing.T) {
	if len(gos.Route.Routers[HEAD]) != 0 {
		t.Error("HEAD method map len error")
	}
	gos.AddHead("/hello", funcation)
	if len(gos.Route.Routers[HEAD]) == 0 {
		t.Error("HEAD method map len error")
	}
}

func TestAddDelete(t *testing.T) {
	if len(gos.Route.Routers[DELETE]) != 0 {
		t.Error("DELETE method map len error")
	}
	gos.AddDetele("/hello", funcation)
	if len(gos.Route.Routers[DELETE]) == 0 {
		t.Error("DELETE method map len error")
	}
}

func TestAddPost(t *testing.T) {
	if len(gos.Route.Routers[POST]) != 0 {
		t.Error("POST method map len error")
	}
	gos.AddPost("/hello", funcation)
	if len(gos.Route.Routers[POST]) == 0 {
		t.Error("POST method map len error")
	}
}
