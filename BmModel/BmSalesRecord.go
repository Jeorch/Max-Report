package BmModel

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"github.com/manyminds/api2go/jsonapi"
	"strconv"
)

type SalesRecord struct {
	ID     string        `json:"-"`
	Id_    bson.ObjectId `json:"-" bson:"_id"`
	InfoId string        `json:"info-id" bson:"INFO_ID"`

	// 1 => Month ; 2 => Quarter ; 3 => Year ; 4 => YTD ; 5 => MAT ;
	DateType int32  `json:"date-type" bson:"DATE_TYPE"`
	Date     string `json:"date" bson:"DATE"`

	// 0 => Total ; 1 => City ; 2 => Province ; 3 => Region ; 4 => Hospital ;
	AddressType int32  `json:"address-type" bson:"ADDRESS_TYPE"`
	AddressId   string `json:"address-id" bson:"ADDRESS_ID"`

	// 0 => Total ; 1 => Product ; 2 => Mole ; 3 => Corp_Name ; 4 => Market ; 5 => OAD
	GoodsType int32  `json:"goods-type" bson:"GOODS_TYPE"`
	GoodsId   string `json:"goods-id" bson:"GOODS_ID"`

	// 1 => Salse ; 2 => Growth ; 3 => Mkt_Share ; 4 => Mkt_Share_Growth ; 5 => Mkt_EI ;
	// 							6 => Mole_Share ; 7 => Mole_Share_Growth ; 8 => Mole_EI ;
	ValueType int32   `json:"value-type" bson:"VALUE_TYPE"`
	Value     float64 `json:"value" bson:"VALUE"`

	City    *City    `json:"-"`
	Product *Product `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a SalesRecord) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *SalesRecord) SetID(id string) error {
	a.ID = id
	return nil
}

func (a *SalesRecord) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	ymr := make(map[string]interface{})

	for k, v := range parameters {
		switch k {
		case "info-id":
			rst["INFO_ID"] = v[0]
		case "date-type":
			rst["DATE_TYPE"], _ = strconv.Atoi(v[0])
		case "date":
			rst["DATE"] = v[0]
		case "lte[date]":
			ymr["$lte"] = v[0]
			rst["DATE"] = ymr
		case "gte[date]":
			ymr["$gte"] = v[0]
			rst["DATE"] = ymr
		case "address-type":
			rst["ADDRESS_TYPE"], _ = strconv.Atoi(v[0])
		case "address-id":
			rst["ADDRESS_ID"] = v[0]
		case "goods-type":
			rst["GOODS_TYPE"], _ = strconv.Atoi(v[0])
		case "goods-id":
			r := make(map[string]interface{})
			var values []string
			for i := 0; i < len(v); i++ {
				values = append(values, v[i])
			}
			r["$in"] = values
			rst["GOODS_ID"] = r
		case "value-type":
			r := make(map[string]interface{})
			var values []int
			for i := 0; i < len(v); i++ {
				value, _ := strconv.Atoi(v[i])
				values = append(values, value)
			}
			r["$in"] = values
			rst["VALUE_TYPE"] = r
		}
	}
	return rst
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (a SalesRecord) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "cities",
			Name: "city",
		},
		{
			Type: "products",
			Name: "product",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (a SalesRecord) GetReferencedIDs() []jsonapi.ReferenceID {
	var result []jsonapi.ReferenceID

	if a.AddressType == 1 {
		result = append(result, jsonapi.ReferenceID{
			ID:   a.AddressId,
			Type: "cities",
			Name: "city",
		})
	}

	if a.GoodsType == 1 {
		result = append(result, jsonapi.ReferenceID{
			ID:   a.GoodsId,
			Type: "products",
			Name: "product",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (a SalesRecord) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	var result []jsonapi.MarshalIdentifier

	if a.AddressType == 1 && a.City != nil {
		result = append(result, a.City)
	}

	if a.GoodsType == 1 && a.Product != nil {
		result = append(result, a.Product)
	}

	return result
}

func (a *SalesRecord) SetToOneReferenceID(name, ID string) error {
	if name == "city" {
		a.AddressId = ID
		return nil
	} else if name == "product" {
		a.GoodsId = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}
