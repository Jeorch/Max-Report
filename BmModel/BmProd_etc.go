package BmModel

import (
	//"errors"
	//"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
	//"strings"
	//"strconv"
)

type Prod_etc struct {
	ID						  string        `json:"-"`
	Id_						  bson.ObjectId `json:"-" bson:"_id"`
	Product_ID					string		`json:"-" bson:"PRODUCT_ID"`
	Product_Name				  	  string		`json:"-" bson:"PH_PRODUCT_NAME"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a Prod_etc) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *Prod_etc) SetID(id string) error {
	a.ID = id
	return nil
}
func (a *Prod_etc) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "PRODUCT_ID":
			rst[k] = v[0]
		}
	}
	return rst
}
