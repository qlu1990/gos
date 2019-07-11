package gos

import (
	"context"
	"net/http"
	"strings"
)

const (
	GET = iota
	POST
	PUT
	DELETE
	CONNECTIBNG
	HEAD
	OPTIONS
	PATCH
	TRACE
	END
)

type Route struct {
	Routers []map[string]HandlerFunc
	Uses    []Middleware
}

type IRoute interface {
	GET(url string, f HandlerFunc)
	POST(url string, f HandlerFunc)
	HEAD(url string, f HandlerFunc)
	DELETE(url string, f HandlerFunc)
}
type Handlers map[string]HandlerFunc

func NewRoute() *Route {
	r := &Route{
		Routers: make([]map[string]HandlerFunc, 9),
		Uses:    make([]Middleware, 0),
	}
	for i := 0; i < END; i++ {
		r.Routers[i] = make(Handlers)
	}
	return r
}

type HandlerFunc func(*Context)

type Middleware struct {
	Name        string
	HandlerFunc HandlerFunc
}

func (ru *Route) GET(url string, f HandlerFunc) {
	if _, ok := ru.Routers[GET][url]; ok {
		Fatal("url :", url, "duplicate")
	}
	ru.Routers[GET][url] = f
}

func (ru *Route) POST(url string, f HandlerFunc) {
	if _, ok := ru.Routers[POST][url]; ok {
		Fatal("url :", url, "duplicate")
	}
	ru.Routers[POST][url] = f
}

func (ru *Route) HEAD(url string, f HandlerFunc) {
	if _, ok := ru.Routers[HEAD][url]; ok {
		Fatal("url :", url, "duplicate")
	}
	ru.Routers[HEAD][url] = f
}

func (ru *Route) DELETE(url string, f HandlerFunc) {
	if _, ok := ru.Routers[DELETE][url]; ok {
		Fatal("url :", url, "duplicate")
	}
	ru.Routers[DELETE][url] = f
}

func (ru *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ru.call(w, r, GET)
	case http.MethodPost:
		ru.call(w, r, POST)
	case http.MethodHead:
		ru.call(w, r, HEAD)
	case http.MethodDelete:
		ru.call(w, r, DELETE)
	default:
	}

}

func (ru *Route) Use(m Middleware) {
	exists := false
	for _, v := range ru.Uses {
		if strings.Compare(v.Name, m.Name) == 0 {
			Glog.Error("use func duplicate : ", m.Name)
			break
		}
	}
	if !exists {
		ru.Uses = append(ru.Uses, m)
	}
}

func (ru *Route) call(w http.ResponseWriter, r *http.Request, method int) {
	c := &Context{
		Ctx:            context.Background(),
		ResponseWriter: w,
		Request:        r,
		NextFlag:       true,
	}
	for _, v := range ru.Uses {
		v.HandlerFunc(c)
		if !c.Next() {
			return
		}

	}
	h := getLongPathFunc(ru.Routers[method], r.RequestURI)
	if h != nil {
		h(c)
	}
}

func getLongPathFunc(handlers Handlers, url string) (h HandlerFunc) {
	keyLen := 1000000
	for k, v := range handlers {
		if strings.Contains(k, url) {
			if len(k) < keyLen {
				keyLen = len(k)
				h = v
			}
		}
	}

	return h
}
