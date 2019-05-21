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

// BmCityStorage stores all Citys
type BmCityStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmCityStorage) NewStorage(args []BmDaemons.BmDaemon) *BmCityStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmCityStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmCityStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.City {
	in := BmModel.City{}
	var out []BmModel.City
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.City
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.City)
	}
}

// GetOne model
func (s BmCityStorage) GetOne(id string) (BmModel.City, error) {
	in := BmModel.City{ID: id}
	model := BmModel.City{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("City for id %s not found", id)
	return BmModel.City{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmCityStorage) Insert(c BmModel.City) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmCityStorage) Delete(id string) error {
	in := BmModel.City{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("City with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmCityStorage) Update(c BmModel.City) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("City with id does not exist")
	}

	return nil
}

func (s *BmCityStorage) Count(req api2go.Request, c BmModel.City) int {
	r, _ := s.db.Count(req, &c)
	return r
}
