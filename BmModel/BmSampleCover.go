package BmModel

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"github.com/manyminds/api2go/jsonapi"
)

type SampleCover struct {
	ID                 string        `json:"-"`
	Id_                bson.ObjectId `json:"-" bson:"_id"`
	InfoId             string        `json:"info-id" bson:"INFO_ID"`
	CityId             string        `json:"city-id" bson:"CITY_ID"`
	PrivateHospitalNum int32         `json:"private-hospital-num" bson:"PRIVATE_HOSPITAL_NUM"`
	UniverseNum        int32         `json:"universe-num" bson:"UNIVERSE_NUM"`
	SampleNum          int32         `json:"sample-num" bson:"SAMPLE_NUM"`
	CoverageRatio      float64       `json:"coverage-ratio" bson:"COVERAGE_RATIO"`

	City *City `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a SampleCover) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *SampleCover) SetID(id string) error {
	a.ID = id
	return nil
}

func (a *SampleCover) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "info-id":
			rst["INFO_ID"] = v[0]
		}
	}
	return rst
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (a SampleCover) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "cities",
			Name: "city",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (a SampleCover) GetReferencedIDs() []jsonapi.ReferenceID {
	var result []jsonapi.ReferenceID
	result = append(result, jsonapi.ReferenceID{
		ID:   a.CityId,
		Type: "cities",
		Name: "city",
	})
	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (a SampleCover) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	var result []jsonapi.MarshalIdentifier
	if a.City != nil {
		result = append(result, a.City)
	}
	return result
}

func (a *SampleCover) SetToOneReferenceID(name, ID string) error {
	if name == "city" {
		a.CityId = ID
		return nil
	}
	return errors.New("There is no to-one relationship with the name " + name)
}
