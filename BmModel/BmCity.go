package BmModel

import (
	"gopkg.in/mgo.v2/bson"
)

type City struct {
	ID       string        `json:"-"`
	Id_      bson.ObjectId `json:"-" bson:"_id"`
	Title    string        `json:"title" bson:"TITLE"`
	CityTier int32         `json:"city-tier" bson:"CITY_TIER"`
	Reliable string        `json:"reliable" bson:"RELIABLE"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a City) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *City) SetID(id string) error {
	a.ID = id
	return nil
}

func (a *City) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "title":
			rst["TITLE"] = v[0]
		}
	}
	return rst
}
