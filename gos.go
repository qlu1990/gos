package gos

import (
	"fmt"
	"net/http"
	"time"
)

//Gos gos struct
type Gos struct {
	Server *http.Server
	Route  *Route
}

//NewGos get new *Gos
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
//Use use middleware
func (g *Gos) Use(m Middleware) {
	g.Route.use(m)
}
//AddGet add get route func
func (g *Gos) AddGet(url string, f HandlerFunc) {
	g.Route.get(url, f)
}
//AddPost add Post route func
func (g *Gos) AddPost(url string, f HandlerFunc) {
	g.Route.post(url, f)
}
//AddHead add Head route func
func (g *Gos) AddHead(url string, f HandlerFunc) {
	g.Route.head(url, f)
}
//AddDelete add delete route func
func (g *Gos) AddDelete(url string, f HandlerFunc) {
	g.Route.delete(url, f)
}
//Run run server listen
func (g*Gos) Run(args ...string) {
	if len(args) > 0 {
		g.Server.Addr = args[0]
	}
	Glog.Info("server is running  ", g.Server.Addr)
	g.Server.ListenAndServe()

}
//Response respnse with status
func Response(w http.ResponseWriter, data string, status int) {
	w.WriteHeader(status)
	fmt.Fprintln(w, data)
}
