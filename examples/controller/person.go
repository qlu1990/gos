package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/qlu1990/gos/gos"
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

// ListPersons list all person
func ListPersons(c *gos.Context) {
	persons, err := model.List()
	if err != nil {
		gos.Error(err)
		http.Error(c.ResponseWriter, "404 NotFound ", http.StatusNotFound)
		return
	}
	personsJSON, err := json.Marshal(persons)
	if err != nil {
		gos.Error(err)
		http.Error(c.ResponseWriter, "404 NotFound ", http.StatusNotFound)
		return
	}
	gos.Response(c.ResponseWriter, string(personsJSON), http.StatusOK)
}

// GetPerson get person by name
func GetPerson(c *gos.Context) {
	name := c.Param("name")
	persons, err := model.GetPersonByName(name)
	if err != nil {
		gos.Error(err)
		http.Error(c.ResponseWriter, "404 NotFound ", http.StatusNotFound)
		return
	}
	personsJSON, err := json.Marshal(persons)
	if err != nil {
		gos.Error(err)
		http.Error(c.ResponseWriter, "404 NotFound ", http.StatusNotFound)
		return
	}
	gos.Response(c.ResponseWriter, string(personsJSON), http.StatusOK)
}
