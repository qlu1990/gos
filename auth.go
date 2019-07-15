package gos

import (
	"net/http"
	"strings"
)

// Auth struct for auth verify
type Auth struct {
	Sessions []string
	white    []string
}

// NewAuth get a new *Auth
func NewAuth() *Auth {
	return &Auth{
		Sessions: make([]string, 0),
	}
}

// Verification hanlderfunc for Verifiy
func (a *Auth) Verification(c *Context) {
	for _, v := range a.white {
		if strings.Contains(c.Request.RequestURI, v) {
			return
		}
	}
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

// MVerify get vertify  Middleware type
func (a *Auth) MVerify() Middleware {
	return Middleware{
		Name:        "auth",
		HandlerFunc: a.Verification,
	}
}

// AddVerifySession  add Session
func (a *Auth) AddVerifySession(Session string) {
	a.Sessions = append(a.Sessions, Session)
}