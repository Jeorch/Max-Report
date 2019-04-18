package BmModel

import (
	//"errors"
	//"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
	//"strings"
	//"strconv"
)

type List struct {
	ID						  string        `json:"-"`
	Id_						  bson.ObjectId `json:"-" bson:"_id"`
	Results				  	  []string		`json:"results" bson:"results"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a List) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *List) SetID(id string) error {
	a.ID = id
	return nil
}
func (a *List) GetConditionsBsonM(parameters map[string][]string) bson.M {
	return bson.M{}
}
