package BmModel

import (
	"gopkg.in/mgo.v2/bson"
	)

type Market struct {
	ID        string        `json:"-"`
	Id_       bson.ObjectId `json:"-" bson:"_id"`
	CompanyId string        `json:"company-id" bson:"COMPANY_ID"`
	Market    string        `json:"market" bson:"MARKET"`
	Desc      string        `json:"desc" bson:"DESC"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a Market) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *Market) SetID(id string) error {
	a.ID = id
	return nil
}

func (a *Market) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "company-id":
			rst["COMPANY_ID"] = v[0]
		case "market":
			rst["MARKET"] = v[0]
		}
	}
	return rst
}
