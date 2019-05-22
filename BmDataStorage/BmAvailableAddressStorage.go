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

// BmAvailableAddressStorage stores all AvailableAddresss
type BmAvailableAddressStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmAvailableAddressStorage) NewStorage(args []BmDaemons.BmDaemon) *BmAvailableAddressStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmAvailableAddressStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmAvailableAddressStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.AvailableAddress {
	in := BmModel.AvailableAddress{}
	var out []BmModel.AvailableAddress
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.AvailableAddress
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.AvailableAddress)
	}
}

// GetOne model
func (s BmAvailableAddressStorage) GetOne(id string) (BmModel.AvailableAddress, error) {
	in := BmModel.AvailableAddress{ID: id}
	model := BmModel.AvailableAddress{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("AvailableAddress for id %s not found", id)
	return BmModel.AvailableAddress{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmAvailableAddressStorage) Insert(c BmModel.AvailableAddress) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmAvailableAddressStorage) Delete(id string) error {
	in := BmModel.AvailableAddress{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("AvailableAddress with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmAvailableAddressStorage) Update(c BmModel.AvailableAddress) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("AvailableAddress with id does not exist")
	}

	return nil
}

func (s *BmAvailableAddressStorage) Count(req api2go.Request, c BmModel.AvailableAddress) int {
	r, _ := s.db.Count(req, &c)
	return r
}
