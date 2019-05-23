package BmModel

import (
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type Overview struct {
	ID             string        `json:"-"`
	Id_            bson.ObjectId `json:"-" bson:"_id"`
	CompanyId      string        `json:"company-id" bson:"COMPANY_ID"`
	Market         string        `json:"market" bson:"MARKET"`
	UpdateTime     int64         `json:"update-time" bson:"UPDATE_TIME"`
	CoverStartTime int64         `json:"cover-start-time" bson:"COVER_START_TIME"`
	CoverEndTime   int64         `json:"cover-end-time" bson:"COVER_END_TIME"`
	HospCount      int32         `json:"hosp-count" bson:"HOSP_COUNT"`
	RegionCount    int32         `json:"region-count" bson:"REGION_COUNT"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a Overview) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *Overview) SetID(id string) error {
	a.ID = id
	return nil
}

func (a *Overview) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "company_id":
			k = strings.ToUpper(k)
			rst[k] = v[0]
		case "market":
			rst["MARKET"] = v[0]
		}
	}
	return rst
}
