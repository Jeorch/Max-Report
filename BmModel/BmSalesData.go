package BmModel

import (
	"gopkg.in/mgo.v2/bson"
)

type SalesData struct {
	ID        string        `json:"-"`
	Id_       bson.ObjectId `json:"-" bson:"_id"`
	InfoId    string        `json:"info-id" bson:"INFO_ID"`
	Date      string        `json:"DATE" bson:"DATE"`
	DateType  string        `json:"date-type" bson:"DATE_TYPE"`
	AddressId string        `json:"address-id" bson:"ADDRESS_ID"`
	// 0 => City ; 1 => Province ; 2 => Region ; 3 => Total
	AddressType string `json:"address-type" bson:"ADDRESS_TYPE"`
	GoodsId     string `json:"goods-id" bson:"GOODS_ID"`
	// 0 => Product ; 1 => Region ; 2 => Province ; 3 => City
	GoodsType string `json:"goods-type" bson:"GOODS_TYPE"`
	Value     string `json:"value" bson:"VALUE"`
	// 0 => Salse ; 1 => growth ; 2 => Share ; 3 => City
	ValueType string `json:"value-type" bson:"VALUE_TYPE"`

	City    *City    `json:"-"`
	Product *Product `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a SalesData) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *SalesData) SetID(id string) error {
	a.ID = id
	return nil
}

func (a *SalesData) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
