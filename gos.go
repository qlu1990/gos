package gos

import (
	"net/http"
	"time"
)

type Gos struct {
	Server *http.Server
	Route  *Route
}

func NewGos() *Gos {
	handler := NewRoute()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &Gos{
		Server: server,
		Route:  handler,
	}
}

func (g *Gos) Use(m Middleware) {
	g.Route.Use(m)
}

func (g *Gos) AddGet(url string, f HandlerFunc) {
	g.Route.GET(url, f)
}

func (g *Gos) AddPost(url string, f HandlerFunc) {
	g.Route.POST(url, f)
}

func (g *Gos) AddHead(url string, f HandlerFunc) {
	g.Route.HEAD(url, f)
}

func (g *Gos) AddDetele(url string, f HandlerFunc) {
	g.Route.DELETE(url, f)
}

func (web *Gos) Run(args ...string) {
	if len(args) > 0 {
		web.Server.Addr = args[0]
	}
	Glog.Info("server is running  ", web.Server.Addr)
	web.Server.ListenAndServe()

}
