package gos

import (
	"net/http"
	"strings"
)

type Auth struct {
	Sessions []string
}

func NewAuth() *Auth {
	return &Auth{
		Sessions: make([]string, 0),
	}
}
func (a *Auth) Verification(c *Context) {
	authkey := c.Request.Header.Get("Authorization")
	if strings.Compare(authkey, "") != 0 {
		for _, v := range a.Sessions {
			if strings.Compare(authkey, v) == 0 {
				return
			}
		}
	}
	Error("auth error", authkey)
	http.Error(c.ResponseWriter, "403 Forbidden", http.StatusForbidden)
	c.Fail()
}
func (a *Auth) MVerify() Middleware {
	return Middleware{
		Name:        "auth",
		HandlerFunc: a.Verification,
	}
}
func (a *Auth) AddVerifySession(Session string) {
	a.Sessions = append(a.Sessions, Session)
}
