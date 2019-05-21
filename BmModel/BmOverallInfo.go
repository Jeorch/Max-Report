package BmModel

import (
	"gopkg.in/mgo.v2/bson"
)

type OverallInfo struct {
	ID              string        `json:"-"`
	Id_             bson.ObjectId `json:"-" bson:"_id"`
	MarketId        string        `json:"market-id" bson:"MARKET_ID"`
	UpdateTime      int64         `json:"update-time" bson:"UPDATE_TIME"`
	CoverStartTime  int64         `json:"cover-start-time" bson:"COVER_START_TIME"`
	CoverEndTime    int64         `json:"cover-end-time" bson:"COVER_END_TIME"`
	OrganizationNum int32         `json:"organization-num" bson:"ORGANIZATION_NUM"`
	AdminAreasNum   int32         `json:"admin-areas-num" bson:"ADMIN_AREAS_NUM"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a OverallInfo) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *OverallInfo) SetID(id string) error {
	a.ID = id
	return nil
}

func (a *OverallInfo) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "market-id":
			rst["MARKET_ID"] = v[0]
		}
	}
	return rst
}
