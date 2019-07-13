package main

import (
	"fmt"

	"github.com/qlu1990/gos"
	"github.com/qlu1990/gos/examples/conf"
	"github.com/qlu1990/gos/examples/controller"
	"github.com/qlu1990/gos/examples/model"
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
	model.SetUp(conf.Cfg.Mongodb.Host)
	r := gos.NewGos()
	r.AddGet("/hello", hello)
	r.AddGet("/bye", goodbye)
	r.AddGet("/hello/person", person)
	r.AddPost("/person/add", controller.AddPerson)
	r.AddGet("/person/list", controller.ListPersons)
	auth := gos.NewAuth()
	auth.AddVerifySession("aaaa")
	r.Use(gos.Mlog)
	r.Use(auth.MVerify())
	r.Run(conf.Cfg.Server.Address)
}
