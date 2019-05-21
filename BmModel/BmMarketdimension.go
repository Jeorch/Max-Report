package BmModel

import (
	//"errors"
	//"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
)

type Marketdimension struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	Market                 string  `json:"market" bson:"MARKET"`
	Ym                     int32   `json:"ym" bson:"YM"`
	ProductCount           int64   `json:"product-count" bson:"PRODUCT_COUNT"`
	Sales                  float64 `json:"sales" bson:"SALES"`
	SalesSom               float64 `json:"sales-som" bson:"SALES_SOM"`
	CompanyID              string  `json:"company-id" bson:"COMPANY_ID"`
	ProductCountRingGrowth float64 `json:"product-count-ring-growth" bson:"PRODUCT_COUNT_RING_GROWTH"`
	ProductCountYearGrowth float64 `json:"product-count-year-growth" bson:"PRODUCT_COUNT_YEAR_GROWTH"`
	ConcentratedSales      float64 `json:"concentrated-sales" bson:"CONCENTRATED_SALES"`
	ConcentratedSom        float64 `json:"concentrated-som" bson:"CONCENTRATED_SOM"`
	ConcentratedRingGrowth float64 `json:"concentrated-ring-growth" bson:"CONCENTRATED_RING_GROWTH"`
	ConcentratedYearGrowth float64 `json:"concentrated-year-growth" bson:"CONCENTRATED_YEAR_GROWTH"`
	SalesSomRingGrowth     float64 `json:"sales-som-ring-growth" bson:"SALES_SOM_RING_GROWTH"`
	SalesSomYearGrowth     float64 `json:"sales-som-year-growth" bson:"SALES_SOM_YEAR_GROWTH"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a Marketdimension) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *Marketdimension) SetID(id string) error {
	a.ID = id
	return nil
}

func (a *Marketdimension) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
		case "ym":
			k = strings.ToUpper(k)
			ym, _ := strconv.Atoi(v[0])
			rst[k] = ym
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
