package BmModel

import (
	//"errors"
	//"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
)

type Marketaggregation struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	Market             string  `json:"market" bson:"MARKET"`
	Address            string  `json:"address" bson:"ADDRESS"`
	AddressType        string  `json:"address-type" bson:"ADDRESS_TYPE"`
	Ym                 int32   `json:"ym" bson:"YM"`
	YmType             string  `json:"ym-type" bson:"YM_TYPE"`
	ProductCount       int64   `json:"product-count" bson:"PRODUCT_COUNT"`
	Sales              float64 `json:"sales" bson:"SALES"`
	SalesSom           float64 `json:"sales-som" bson:"SALES_SOM"`
	CompanyID          string  `json:"company-id" bson:"COMPANY_ID"`
	SalesYearGrowth    float64 `json:"sales-year-growth" bson:"SALES_YEAR_GROWTH"`
	SalesEI            float64 `json:"sales-ei" bson:"SALES_EI"`
	SalesSomYearGrowth float64 `json:"sales-som-year-growth" bson:"SALES_SOM_YEAR_GROWTH"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a Marketaggregation) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *Marketaggregation) SetID(id string) error {
	a.ID = id
	return nil
}

func (a *Marketaggregation) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	r := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "company_id":
			k = strings.ToUpper(k)
			rst[k] = v[0]
		case "market":
			k = strings.ToUpper(k)
			rst[k] = v[0]
		case "address":
			k = strings.ToUpper(k)
			rst[k] = v[0]
		case "address_type":
			k = strings.ToUpper(k)
			rst[k] = v[0]
		case "ym":
			k = strings.ToUpper(k)
			ym, _ := strconv.Atoi(v[0])
			rst[k] = ym
		case "ym_type":
			k = strings.ToUpper(k)
			rst[k] = v[0]
		case "lt[ym]":
			val, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			r["$lt"] = val
			rst["YM"] = r
		case "lte[ym]":
			val, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			r["$lte"] = val
			rst["YM"] = r
		case "gt[ym]":
			val, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			r["$gt"] = val
			rst["YM"] = r
		case "gte[ym]":
			val, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			r["$gte"] = val
			rst["YM"] = r
		}

	}
	return rst
}
