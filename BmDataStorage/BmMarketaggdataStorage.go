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

// BmMarketaggdataStorage stores all Marketaggdataes
type BmMarketaggdataStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmMarketaggdataStorage) NewStorage(args []BmDaemons.BmDaemon) *BmMarketaggdataStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmMarketaggdataStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmMarketaggdataStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Marketaggregation {
	in := BmModel.Marketaggregation{}
	var out []BmModel.Marketaggregation
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Marketaggregation
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Marketaggregation)
	}
}

// GetOne model
func (s BmMarketaggdataStorage) GetOne(id string) (BmModel.Marketaggregation, error) {
	in := BmModel.Marketaggregation{ID: id}
	model := BmModel.Marketaggregation{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Marketaggregation for id %s not found", id)
	return BmModel.Marketaggregation{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmMarketaggdataStorage) Insert(c BmModel.Marketaggregation) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmMarketaggdataStorage) Delete(id string) error {
	in := BmModel.Marketaggregation{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Marketaggregation with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmMarketaggdataStorage) Update(c BmModel.Marketaggregation) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Marketaggregation with id does not exist")
	}

	return nil
}

func (s *BmMarketaggdataStorage) Count(req api2go.Request, c BmModel.Marketaggregation) int {
	r, _ := s.db.Count(req, &c)
	return r
}
