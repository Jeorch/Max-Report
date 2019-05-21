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

// BmSalesDataStorage stores all SalesDataes
type BmSalesDataStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmSalesDataStorage) NewSalesDataStorage(args []BmDaemons.BmDaemon) *BmSalesDataStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmSalesDataStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmSalesDataStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.SalesData {
	in := BmModel.SalesData{}
	var out []BmModel.SalesData
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.SalesData
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.SalesData)
	}
}

// GetOne model
func (s BmSalesDataStorage) GetOne(id string) (BmModel.SalesData, error) {
	in := BmModel.SalesData{ID: id}
	model := BmModel.SalesData{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("SalesData for id %s not found", id)
	return BmModel.SalesData{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmSalesDataStorage) Insert(c BmModel.SalesData) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmSalesDataStorage) Delete(id string) error {
	in := BmModel.SalesData{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("SalesData with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmSalesDataStorage) Update(c BmModel.SalesData) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("SalesData with id does not exist")
	}

	return nil
}

func (s *BmSalesDataStorage) Count(req api2go.Request, c BmModel.SalesData) int {
	r, _ := s.db.Count(req, &c)
	return r
}
