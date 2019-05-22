package BmModel

import (
	//"errors"
	//"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
)

type Productaggregation struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	ProductName        string  `json:"product-name" bson:"PRODUCT_NAME"`
	CorpName           string  `json:"corp-name" bson:"CORP_NAME"`
	MinProduct         string  `json:"min-product" bson:"MIN_PRODUCT"`
	Ym                 int32   `json:"ym" bson:"YM"`
	YmType             string  `json:"ym-type" bson:"YM_TYPE"`
	Address            string  `json:"address" bson:"ADDRESS"`
	AddressType        string  `json:"address-type" bson:"ADDRESS_TYPE"`
	Market             string  `json:"market" bson:"MARKET"`
	Sales              float64 `json:"sales" bson:"SALES"`
	CompanyID          string  `json:"company-id" bson:"COMPANY_ID"`
	SalesSom           float64 `json:"sales-som" bson:"SALES_SOM"`
	SalesRank          int32   `json:"sales-rank" bson:"SALES_RANK"`
	SalesYearGrowth    float64 `json:"sales-year-growth" bson:"SALES_YEAR_GROWTH"`
	SalesSomYearGrowth float64 `json:"sales-som-year-growth" bson:"SALES_SOM_YEAR_GROWTH"`
	SalesEI            float64 `json:"sales-ei" bson:"SALES_EI"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a Productaggregation) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *Productaggregation) SetID(id string) error {
	a.ID = id
	return nil
}
func (a *Productaggregation) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{}, 0)
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
		case "address":
			k = strings.ToUpper(k)
			rst[k] = v[0]
		case "address_type":
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
		case "ym_type":
			k = strings.ToUpper(k)
			rst[k] = v[0]
		case "lte[sales_rank]":
			val, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err.Error())
			}
			rankr["$lte"] = val
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
		case "min-product-names":
			r:=make(map[string]interface{})
			r["$in"]=v
			rst["MIN_PRODUCT"] = r
		}

	}
	return rst
}
