package BmModel

import (
	//"errors"
	//"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"strconv"
)

type Productdimension struct {
	ID						string        `json:"-"`
	Id_						bson.ObjectId `json:"-" bson:"_id"`

	Company_ID				  string		`json:"company-id" bson:"COMPANY_ID"`
	Market					  string		`json:"market" bson:"MARKET"`
	Ym						  int32	   	 	`json:"ym" bson:"YM"`
	Sales					  float64		`json:"sales" bson:"SALES"`
	Units					  float64		`json:"units" bson:"UNITS"`
	Province_Count 	          float64  	    `json:"province-count" bson:"PROVINCE_COUNT"`
	City_Count 	              int64  	    `json:"city-count" bson:"CITY_COUNT"`
	Product_Count			  int64			`json:"product-count" bson:"PRODUCT_COUNT"`
	Min_Product	  			  string		`json:"min-product" bson:"MIN_PRODUCT"`	
	Sales_Som                 float64		`json:"sales-som" bson:"SALES_SOM"`	
	Sales_Rank                int32			`json:"sales-rank" bson:"SALES_RANK"`	
	SalesRing                 float64		`json:"salesring" bson:"SALESRING"`	
	Sales_Ring_Growth    	  float64		`json:"sales-ring-growth" bson:"SALES_RING_GROWTH"`
	Sales_Ring_Growth_Rank    int32			`json:"sales-ring-growth-rank" bson:"SALES_RING_GROWTH_RANK"`
	Sales_Year_On_Year      float64			`json:"sales-year-on-year" bson:"SALES_YEAR_ON_YEAR"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a Productdimension) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *Productdimension) SetID(id string) error {
	a.ID = id
	return nil
}
func (a *Productdimension) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	ymr := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "company_id":
			k = strings.ToUpper(k)
			rst[k] = v[0]
		case "market":
			k = strings.ToUpper(k)
			rst[k] = v[0]
		case "sales_rank":
			val, err:= strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			k = strings.ToUpper(k)
			rst[k] = val

		case "ym":
			k = strings.ToUpper(k)
			ym,_ := strconv.Atoi(v[0])
			rst[k] = ym
		case "lte[sales_rank]":
			val, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			ymr["$lte"] = val
			rst["SALES_RANK"] = ymr
		case "lte[sales_som]":
			val, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			ymr["$lte"] = val
			rst["SALES_SOM"] = ymr
		case "lte[sales_ring_growth_rank]":
			val, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			ymr["$lte"] = val
			rst["SALES_RING_GROWTH_RANK"] = ymr
		case "gte[sales_rank]":
			val, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			ymr["$gte"] = val
			rst["SALES_RANK"] = ymr
		case "lt[ym]":
			val, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			ymr["$lt"] = val
			rst["YM"] =ymr
		case "lte[ym]":
			val, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			ymr["$lte"] = val
			rst["YM"] = ymr
		case "gt[ym]":
			val, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			ymr["$gt"] = val
			rst["YM"] = ymr
		case "gte[ym]":
			val, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			ymr["$gte"] = val
			rst["YM"] = ymr
		}

	}
	return rst
}
