package BmModel

import (
	//"errors"
	//"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type MarketDimension struct {
	ID						  string        `json:"-"`
	Id_						  bson.ObjectId `json:"-" bson:"_id"`

	Company_ID				  string		`json:"company-id" bson:"COMPANY_ID"`
	Market					  string		`json:"market" bson:"MARKET"`
	Ym						  int32	   	 	`json:"ym" bson:"YM"`
	Sales					  float64		`json:"sales" bson:"SALES"`
	Units					  float64		`json:"units" bson:"UNITS"`
	Product_Count			  int64			`json:"product-count" bson:"PRODUCT_COUNT"`
	Product_Countring		  int64			`json:"product-countring" bson:"PRODUCT_COUNTRING"`		
	Sales_SomRing             float64		`json:"sales-somring" bson:"SALES_SOMRING"`	
	Sales_Som                 float64		`json:"sales-som" bson:"SALES_SOM"`	
	Product_Count_Ring_Growth float64		`json:"product-count-ring-growth" bson:"PRODUCT_COUNT_RING_GROWTH"`	
	Product_CountYear_On_Year int64			`json:"product-countYear-on-year" bson:"PRODUCT_COUNTYEAR_ON_YEAR"`	
	Product_Count_Year_Growth float64		`json:"product-count-year-growth" bson:"PRODUCT_COUNT_YEAR_GROWTH"`
	Sales_SomYear_On_Year     float64		`json:"sales-somyear-on-year" bson:"SALES_SOMYEAR_ON_YEAR"`
	Sales_Som_Year_Growth     float64		`json:"sales-som-year-growth" bson:"SALES_SOM_YEAR_GROWTH"`
	Sales_Som_Ring_Growth     float64		`json:"sales-som-ring-growth" bson:"SALES_SOM_RING_GROWTH"`
	Province_Count 	          int64  	    `json:"province-count" bson:"PROVINCE_COUNT"`
	City_Count 	              int64  	    `json:"city-count" bson:"CITY_COUNT"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a MarketDimension) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *MarketDimension) SetID(id string) error {
	a.ID = id
	return nil
}
func (a *MarketDimension) GetConditionsBsonM(parameters map[string][]string) bson.M {
	return bson.M{}
	/*
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "ids":
			r := make(map[string]interface{})
			var ids []bson.ObjectId
			for i := 0; i < len(v); i++ {
				ids = append(ids, bson.ObjectIdHex(v[i]))
			}
			r["$in"] = ids
			rst["_id"] = r
		case "scenario-id":
			rst[k] = v[0]
		}
	}
	return rst
	*/
}
