package BmModel

import (
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	ID            string        `json:"-"`
	Id_           bson.ObjectId `json:"-" bson:"_id"`
	Title         string        `json:"title" bson:"TITLE"`
	CorpName      string        `json:"corp-name" bson:"CORP_NAME"`
	DeliveryWay   string        `json:"delivery-way" bson:"DELIVERY_WAY"`
	DosageName    string        `json:"dosage-name" bson:"DOSAGE_NAME"`
	MoleName      string        `json:"mole-name" bson:"MOLE_NAME"`
	PackageDes    string        `json:"package-des" bson:"PACKAGE_DES"`
	PackageNumber int32         `json:"package-number" bson:"PACKAGE_NUMBER"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a Product) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *Product) SetID(id string) error {
	a.ID = id
	return nil
}

func (a *Product) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "title":
			rst["TITLE"] = v[0]
		}
	}
	return rst
}
