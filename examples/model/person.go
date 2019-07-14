package model

import (
	"context"

	"gopkg.in/mgo.v2/bson"
)

// Person model
type Person struct {
	id   bson.ObjectId `json:"id",bosn:"_id,omitempty"`
	Name string        `bson:"name",json:"name"`
	Sex  string        `bson:"sex",json:"sex"`
	Age  int           `bson:"age",json:"age"`
}

// Add add
func (p *Person) Add() error {
	ctx, _ := context.WithTimeout(Mongo.Ctx, Mongo.Timeout)
	_, err := Mongo.Client.Database("test").Collection("person").InsertOne(ctx, p)
	return err
}

// List get persons
func List() ([]*Person, error) {
	ctx, _ := context.WithTimeout(Mongo.Ctx, Mongo.Timeout)
	filter := make(map[string]interface{})
	// findOptions := new(options.FindOptions)
	result := make([]*Person, 0)
	cur, err := Mongo.Client.Database("test").Collection("person").Find(ctx, filter)
	if err != nil {
		return result, err
	}
	for cur.Next(ctx) {
		person := new(Person)
		err = cur.Decode(person)
		if err != nil {
			return result, err
		}
		result = append(result, person)
	}
	return result, err
}
func GetPersonByName(name string) ([]*Person, error) {
	ctx, _ := context.WithTimeout(Mongo.Ctx, Mongo.Timeout)
	filter := make(map[string]interface{})
	filter["name"] = name
	// findOptions := new(options.FindOptions)
	result := make([]*Person, 0)
	cur, err := Mongo.Client.Database("test").Collection("person").Find(ctx, filter)
	if err != nil {
		return result, err
	}
	for cur.Next(ctx) {
		person := new(Person)
		err = cur.Decode(person)
		if err != nil {
			return result, err
		}
		result = append(result, person)
	}
	return result, err
}
