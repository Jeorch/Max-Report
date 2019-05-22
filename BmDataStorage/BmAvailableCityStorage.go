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

// BmAvailableCityStorage stores all AvailableCitys
type BmAvailableCityStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmAvailableCityStorage) NewStorage(args []BmDaemons.BmDaemon) *BmAvailableCityStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmAvailableCityStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmAvailableCityStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.AvailableCity {
	in := BmModel.AvailableCity{}
	var out []BmModel.AvailableCity
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.AvailableCity
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.AvailableCity)
	}
}

// GetOne model
func (s BmAvailableCityStorage) GetOne(id string) (BmModel.AvailableCity, error) {
	in := BmModel.AvailableCity{ID: id}
	model := BmModel.AvailableCity{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("AvailableCity for id %s not found", id)
	return BmModel.AvailableCity{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmAvailableCityStorage) Insert(c BmModel.AvailableCity) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmAvailableCityStorage) Delete(id string) error {
	in := BmModel.AvailableCity{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("AvailableCity with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmAvailableCityStorage) Update(c BmModel.AvailableCity) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("AvailableCity with id does not exist")
	}

	return nil
}

func (s *BmAvailableCityStorage) Count(req api2go.Request, c BmModel.AvailableCity) int {
	r, _ := s.db.Count(req, &c)
	return r
}
