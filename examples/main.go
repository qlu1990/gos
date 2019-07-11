package main

import (
	"fmt"

	"github.com/qlu1990/gos"
	"github.com/qlu1990/gos/examples/conf"
)

func hello(c *gos.Context) {
	fmt.Fprintln(c.ResponseWriter, "hello world")
}
func goodbye(c *gos.Context) {
	fmt.Fprintln(c.ResponseWriter, "Goodbye")
}
func person(c *gos.Context) {
	fmt.Fprintln(c.ResponseWriter, "ni hao")
}
func main() {

	conf.LoadConf()
	r := gos.NewGos()
	r.AddGet("/hello", hello)
	r.AddGet("/bye", goodbye)
	r.AddGet("/hello/person", person)
	auth := gos.NewAuth()
	auth.AddVerifySession("aaaa")
	r.Use(gos.Mlog)
	r.Use(auth.MVerify())
	r.Run(conf.Cfg.Server.Address)
}
