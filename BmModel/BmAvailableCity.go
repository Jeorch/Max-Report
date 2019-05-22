package BmModel

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"github.com/manyminds/api2go/jsonapi"
)

type AvailableCity struct {
	ID          string        `json:"-"`
	Id_         bson.ObjectId `json:"-" bson:"_id"`
	InfoId      string        `json:"info-id" bson:"INFO_ID"`
	AddressType int32         `json:"address-type" bson:"ADDRESS_TYPE"`
	AddressId   string        `json:"address-id" bson:"ADDRESS_ID"`

	City *City `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a AvailableCity) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *AvailableCity) SetID(id string) error {
	a.ID = id
	return nil
}

func (a *AvailableCity) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "title":
			rst["TITLE"] = v[0]
		}
	}
	return rst
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (a AvailableCity) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "cities",
			Name: "city",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (a AvailableCity) GetReferencedIDs() []jsonapi.ReferenceID {
	var result []jsonapi.ReferenceID

	if a.AddressType == 1 {
		result = append(result, jsonapi.ReferenceID{
			ID:   a.AddressId,
			Type: "cities",
			Name: "city",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (a AvailableCity) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	var result []jsonapi.MarshalIdentifier

	if a.AddressType == 1 && a.City != nil {
		result = append(result, a.City)
	}

	return result
}

func (a *AvailableCity) SetToOneReferenceID(name, ID string) error {
	if name == "city" {
		a.AddressId = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}