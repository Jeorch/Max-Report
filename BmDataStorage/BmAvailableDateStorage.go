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

// BmAvailableDateStorage stores all AvailableDates
type BmAvailableDateStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmAvailableDateStorage) NewStorage(args []BmDaemons.BmDaemon) *BmAvailableDateStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmAvailableDateStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmAvailableDateStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.AvailableDate {
	in := BmModel.AvailableDate{}
	var out []BmModel.AvailableDate
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.AvailableDate
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.AvailableDate)
	}
}

// GetOne model
func (s BmAvailableDateStorage) GetOne(id string) (BmModel.AvailableDate, error) {
	in := BmModel.AvailableDate{ID: id}
	model := BmModel.AvailableDate{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("AvailableDate for id %s not found", id)
	return BmModel.AvailableDate{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmAvailableDateStorage) Insert(c BmModel.AvailableDate) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmAvailableDateStorage) Delete(id string) error {
	in := BmModel.AvailableDate{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("AvailableDate with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmAvailableDateStorage) Update(c BmModel.AvailableDate) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("AvailableDate with id does not exist")
	}

	return nil
}

func (s *BmAvailableDateStorage) Count(req api2go.Request, c BmModel.AvailableDate) int {
	r, _ := s.db.Count(req, &c)
	return r
}
