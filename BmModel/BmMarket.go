package BmModel

import (
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type Market struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	Market    string `json:"market" bson:"MARKET"`
	CompanyID string `json:"company-id" bson:"COMPANY_ID"`
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

func (u *Market) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "COMPANY_ID":
			k = strings.ToUpper(k)
			rst[k] = v[0]
		}

	}
	return rst
}
