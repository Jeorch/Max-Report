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

// BmSalesRecordStorage stores all SalesRecords
type BmSalesRecordStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmSalesRecordStorage) NewSalesRecordStorage(args []BmDaemons.BmDaemon) *BmSalesRecordStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmSalesRecordStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmSalesRecordStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.SalesRecord {
	in := BmModel.SalesRecord{}
	var out []BmModel.SalesRecord
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.SalesRecord
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.SalesRecord)
	}
}

// GetOne model
func (s BmSalesRecordStorage) GetOne(id string) (BmModel.SalesRecord, error) {
	in := BmModel.SalesRecord{ID: id}
	model := BmModel.SalesRecord{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("SalesRecord for id %s not found", id)
	return BmModel.SalesRecord{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmSalesRecordStorage) Insert(c BmModel.SalesRecord) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmSalesRecordStorage) Delete(id string) error {
	in := BmModel.SalesRecord{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("SalesRecord with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmSalesRecordStorage) Update(c BmModel.SalesRecord) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("SalesRecord with id does not exist")
	}

	return nil
}

func (s *BmSalesRecordStorage) Count(req api2go.Request, c BmModel.SalesRecord) int {
	r, _ := s.db.Count(req, &c)
	return r
}
