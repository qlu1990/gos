package gos

import (
	"context"
	"net/http"
	"strings"
)
//method id
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
//Route route struct
type Route struct {
	Routers []*node
	Uses    []Middleware
}

// IRoute interface for route
type IRoute interface {
	Get(url string, f HandlerFunc)
	Post(url string, f HandlerFunc)
	Head(url string, f HandlerFunc)
	Delete(url string, f HandlerFunc)
}

//NewRoute get new *Route
func NewRoute() *Route {
	r := &Route{
		Routers: make([]*node, 9),
		Uses:    make([]Middleware, 0),
	}
	for i := 0; i < END; i++ {
		n := new(node)
		n.nodeType = ROOT
		r.Routers[i] = n
	}
	return r
}

//HandlerFunc func type for handle request
type HandlerFunc func(*Context)

//Middleware middleware struct
type Middleware struct {
	Name        string
	HandlerFunc HandlerFunc
}


func (ru *Route) get(url string, f HandlerFunc) {
	paths := GetPaths(url)
	n := getMatchOne(ru.Routers[GET], paths)
	if n != nil && n.handlerFunc != nil {
		Fatal("Get url :", url, "duplicate")
	}
	ru.Routers[GET].AddRoute(url, f)
}

func (ru *Route) post(url string, f HandlerFunc) {
	paths := GetPaths(url)
	n := getMatchOne(ru.Routers[POST], paths)
	if n != nil && n.handlerFunc != nil {
		Fatal("Post url :", url, "duplicate")
	}
	ru.Routers[POST].AddRoute(url, f)
}

func (ru *Route) head(url string, f HandlerFunc) {
	paths := GetPaths(url)
	n := getMatchOne(ru.Routers[HEAD], paths)
	if n != nil && n.handlerFunc != nil {
		Fatal("Head url :", url, "duplicate")
	}
	ru.Routers[HEAD].AddRoute(url, f)
}

func (ru *Route) delete(url string, f HandlerFunc) {
	paths := GetPaths(url)
	n := getMatchOne(ru.Routers[DELETE], paths)
	if n != nil && n.handlerFunc != nil {
		Fatal("Delete url :", url, "duplicate")
	}
	ru.Routers[DELETE].AddRoute(url, f)
}

//ServeHTTP interface func
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

func (ru *Route) use(m Middleware) {
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
		Params:         make(map[string]string),
	}
	for _, v := range ru.Uses {
		v.HandlerFunc(c)
		if !c.Next() {
			return
		}

	}
	paths := GetPaths(r.RequestURI)
	n := getMatchOne(ru.Routers[method], paths)
	routePaths := GetPaths(n.fullPath)
	for i, v := range paths {
		if strings.HasPrefix(routePaths[i], ":") {
			c.Params[string(routePaths[i][1:])] = v
		}
	}
	if n != nil {
		n.handlerFunc(c)
	} else {
		StatusNotFound(c)
	}
}

// StatusNotFound 404 notFound
func StatusNotFound(c *Context) {
	http.Error(c.ResponseWriter, "404 NotFound", http.StatusNotFound)
}
