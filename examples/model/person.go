package model

import (
	"context"

	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	id   bson.ObjectId "_id,omitempty"
	Name string        `bson:"name",json:"name"`
	Sex  string        `bson:"sex",json:"sex"`
	Age  int           `bson:"age",json:"age"`
}

func (p *Person) Add() error {
	ctx, _ := context.WithTimeout(Mongo.Ctx, Mongo.Timeout)
	_, err := Mongo.Client.Database("test").Collection("person").InsertOne(ctx, p)
	return err
}
