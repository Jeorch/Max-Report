package BmDataStorage

import (
	"errors"
	"fmt"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
)

// BmProductdimensionStorage stores all Productdimensiones
type BmProductdimensionStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmProductdimensionStorage) NewProductdimensionStorage(args []BmDaemons.BmDaemon) *BmProductdimensionStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmProductdimensionStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmProductdimensionStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Productdimension {
	in := BmModel.Productdimension{}
	var out []BmModel.Productdimension
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Productdimension
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Productdimension)
	}
}
func (s BmProductdimensionStorage) GetAllByCond(r bson.M, skip int, take int) []*BmModel.Productdimension {
	in := BmModel.Productdimension{}
	var out []BmModel.Productdimension
	err := s.db.FindMultiByCondition( &in, &out,r ,"",skip, take)
	if err == nil {
		var tmp []*BmModel.Productdimension
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Productdimension)
	}
}

// GetOne model
func (s BmProductdimensionStorage) GetOne(id string) (BmModel.Productdimension, error) {
	in := BmModel.Productdimension{ID: id}
	model := BmModel.Productdimension{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Productdimension for id %s not found", id)
	return BmModel.Productdimension{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmProductdimensionStorage) Insert(c BmModel.Productdimension) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmProductdimensionStorage) Delete(id string) error {
	in := BmModel.Productdimension{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Productdimension with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmProductdimensionStorage) Update(c BmModel.Productdimension) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Productdimension with id does not exist")
	}

	return nil
}

func (s *BmProductdimensionStorage) Count(req api2go.Request, c BmModel.Productdimension) int {
	r, _ := s.db.Count(req, &c)
	return r
}
