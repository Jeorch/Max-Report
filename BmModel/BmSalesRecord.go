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

	// 1 => Month ; 2 => Quarter ; 3 => Year ; 4 => YTD ; 5 => MAT
	DateType int32  `json:"date-type" bson:"DATE_TYPE"`
	Date     string `json:"date" bson:"DATE"`

	// 0 => Total ; 1 => City ; 2 => Province ; 3 => Region ;
	AddressType int32  `json:"address-type" bson:"ADDRESS_TYPE"`
	AddressId   string `json:"address-id" bson:"ADDRESS_ID"`

	// 0 => Total ; 1 => Product ;
	GoodsType int32  `json:"goods-type" bson:"GOODS_TYPE"`
	GoodsId   string `json:"goods-id" bson:"GOODS_ID"`

	// 1 => Salse ; 2 => growth ; 3 => Share ;
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
	for k, v := range parameters {
		switch k {
		case "info-id":
			rst["INFO_ID"] = v[0]
		case "date-type":
			rst["DATE_TYPE"], _ = strconv.Atoi(v[0])
		case "date":
			rst["DATE"] = v[0]
		case "address-type":
			rst["ADDRESS_TYPE"], _ = strconv.Atoi(v[0])
		case "address-id":
			rst["ADDRESS_ID"] = v[0]
		case "goods-type":
			rst["GOODS_TYPE"], _ = strconv.Atoi(v[0])
		case "goods-id":
			rst["GOODS_ID"] = v[0]
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
			Type: "goods",
			Name: "goods",
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
			Type: "goods",
			Name: "goods",
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
	} else if name == "goods" {
		a.GoodsId = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}
