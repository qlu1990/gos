package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/qlu1990/gos"
	"github.com/qlu1990/gos/examples/model"
)

func AddPerson(c *gos.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		gos.Error(err)
		http.Error(c.ResponseWriter, "error get data", http.StatusNotFound)
		return
	}
	var person = &model.Person{}
	err = json.Unmarshal(data, person)
	if err != nil {
		gos.Error(err)
		http.Error(c.ResponseWriter, "error json ", http.StatusNotFound)
		return
	}
	err = person.Add()
	if err != nil {
		gos.Error(err)
		http.Error(c.ResponseWriter, "fail add person ", http.StatusNotFound)
		return
	}
	gos.Response(c.ResponseWriter, "success", http.StatusOK)

}
