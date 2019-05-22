package BmModel

import (
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type AvailableDate struct {
	ID       string        `json:"-"`
	Id_      bson.ObjectId `json:"-" bson:"_id"`
	InfoId   string        `json:"info-id" bson:"INFO_ID"`
	DateType int32         `json:"date-type" bson:"DATE_TYPE"`
	Date     string        `json:"date" bson:"DATE"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a AvailableDate) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *AvailableDate) SetID(id string) error {
	a.ID = id
	return nil
}

func (a *AvailableDate) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "info-id":
			rst["INFO_ID"] = v[0]
		case "date-type":
			rst["DATE_TYPE"], _ = strconv.Atoi(v[0])
		}
	}
	return rst
}
