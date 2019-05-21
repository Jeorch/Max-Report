package BmDataStorage

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
)

// BmProductaggdataStorage stores all Productaggdataes
type BmProductaggdataStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmProductaggdataStorage) NewStorage(args []BmDaemons.BmDaemon) *BmProductaggdataStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmProductaggdataStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmProductaggdataStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Productaggregation {
	in := BmModel.Productaggregation{}
	var out []BmModel.Productaggregation
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Productaggregation
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Productaggregation)
	}
}

// GetOne model
func (s BmProductaggdataStorage) GetOne(id string) (BmModel.Productaggregation, error) {
	in := BmModel.Productaggregation{ID: id}
	model := BmModel.Productaggregation{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Productaggregation for id %s not found", id)
	return BmModel.Productaggregation{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmProductaggdataStorage) Insert(c BmModel.Productaggregation) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmProductaggdataStorage) Delete(id string) error {
	in := BmModel.Productaggregation{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Productaggregation with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmProductaggdataStorage) Update(c BmModel.Productaggregation) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Productaggregation with id does not exist")
	}

	return nil
}

func (s *BmProductaggdataStorage) Count(req api2go.Request, c BmModel.Productaggregation) int {
	r, _ := s.db.Count(req, &c)
	return r
}
