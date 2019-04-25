package BmModel

import (
	//"errors"
	//"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"strconv"
)

type Marketdimension struct {
	ID						  string        `json:"-"`
	Id_						  bson.ObjectId `json:"-" bson:"_id"`

	Market					  string		`json:"market" bson:"MARKET"`
	Ym						  int32	   	 	`json:"ym" bson:"YM"`
	Product_Count			  int64			`json:"product-count" bson:"PRODUCT_COUNT"`
	Sales					  float64		`json:"sales" bson:"SALES"`
	Sales_Som                 float64		`json:"sales-som" bson:"SALES_SOM"`	
	Company_ID				  string		`json:"company-id" bson:"COMPANY_ID"`
	Product_Count_Ring_Growth		  float64			`json:"product-count-ring-growth" bson:"PRODUCT_COUNT_RING_GROWTH"`	
	Product_Count_Year_Growth		  float64			`json:"product-count-year-growth" bson:"PRODUCT_COUNT_YEAR_GROWTH"`	
	Concentrated_Sales             float64		`json:"concentrated-sales" bson:"CONCENTRATED_SALES"`	
	Concentrated_Som 			float64		`json:"concentrated-som" bson:"CONCENTRATED_SOM"`	
	Concentrated_Ring_Growth float64			`json:"concentrated-ring-growth" bson:"CONCENTRATED_RING_GROWTH"`	
	Concentrated_Year_Growth float64		`json:"concentrated-year-growth" bson:"CONCENTRATED_YEAR_GROWTH"`

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
			ym,_ := strconv.Atoi(v[0])
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
