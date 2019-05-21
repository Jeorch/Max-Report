package BmModel

import (
	//"errors"
	//"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
)

type Productdimension struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	ProductId           string  `json:"product-id" bson:"PRODUCT_ID"`
	Ym                  int32   `json:"ym" bson:"YM"`
	Market              string  `json:"market" bson:"MARKET"`
	Sales               float64 `json:"sales" bson:"SALES"`
	CompanyID           string  `json:"company-id" bson:"COMPANY_ID"`
	SalesSom            float64 `json:"sales-som" bson:"SALES_SOM"`
	SalesRank           int32   `json:"sales-rank" bson:"SALES_RANK"`
	SalesRingGrowthRank int32   `json:"sales-ring-growth-rank" bson:"SALES_RING_GROWTH_RANK"`
	SalesRingGrowth     float64 `json:"sales-ring-growth" bson:"SALES_RING_GROWTH"`
	SalesYearGrowth     float64 `json:"sales-year-growth" bson:"SALES_YEAR_GROWTH"`
	ProductName         string  `json:"product-name" bson:"MIN_PRODUCT"`
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
	rst := make(map[string]interface{},0)
	ymr := make(map[string]interface{})
	rankr := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "company_id":
			k = strings.ToUpper(k)
			rst[k] = v[0]
		case "market":
			k = strings.ToUpper(k)
			rst[k] = v[0]
		case "sales_rank":
			val, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			k = strings.ToUpper(k)
			rst[k] = val

		case "ym":
			k = strings.ToUpper(k)
			ym, _ := strconv.Atoi(v[0])
			rst[k] = ym
		case "lte[sales_rank]":
			val, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			rankr["$lte"] = val
			rst["SALES_RANK"] = rankr
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
			rankr["$gte"] = val
			rst["SALES_RANK"] = rankr
		case "lt[ym]":
			val, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			ymr["$lt"] = val
			rst["YM"] = ymr
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
