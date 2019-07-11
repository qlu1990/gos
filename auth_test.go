package gos

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

var auth *Auth

type response struct {
}

func (r response) Header() http.Header {
	return http.Header{}
}
func (r response) Write(data []byte) (int, error) {
	l := len(data)
	fmt.Println(string(data[:]))
	return l, nil
}
func (r response) WriteHeader(statusCode int) {
	fmt.Println(statusCode)
}
func TestNewAuth(t *testing.T) {
	auth = NewAuth()
}

func TestAddVerifySession(t *testing.T) {
	auth.AddVerifySession("abc")
}

func TestVerification(t *testing.T) {
	r := &http.Request{
		Header: http.Header{},
	}
	r.Header.Set("Authorization", "aaa")
	c := &Context{
		Ctx:            context.Background(),
		Request:        r,
		ResponseWriter: response{},
	}
	auth.Verification(c)
	if c.NextFlag {
		t.Error("auth verify fail")
	}
	c.NextFlag = true
	c.Request.Header.Set("Authorization", "abc")
	auth.Verification(c)
	if !c.NextFlag {
		t.Error("auth verify fail")
	}
}

func TestMVerify(t *testing.T) {
	_, ok := interface{}(auth.MVerify()).(Middleware)
	if !ok {
		t.Error("eroor MVerify")
	}
}
