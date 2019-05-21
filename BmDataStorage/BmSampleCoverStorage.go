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

// BmSampleCoverStorage stores all SampleCoveres
type BmSampleCoverStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmSampleCoverStorage) NewSampleCoverStorage(args []BmDaemons.BmDaemon) *BmSampleCoverStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmSampleCoverStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmSampleCoverStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.SampleCover {
	in := BmModel.SampleCover{}
	var out []BmModel.SampleCover
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.SampleCover
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.SampleCover)
	}
}

// GetOne model
func (s BmSampleCoverStorage) GetOne(id string) (BmModel.SampleCover, error) {
	in := BmModel.SampleCover{ID: id}
	model := BmModel.SampleCover{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("SampleCover for id %s not found", id)
	return BmModel.SampleCover{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmSampleCoverStorage) Insert(c BmModel.SampleCover) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmSampleCoverStorage) Delete(id string) error {
	in := BmModel.SampleCover{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("SampleCover with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmSampleCoverStorage) Update(c BmModel.SampleCover) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("SampleCover with id does not exist")
	}

	return nil
}

func (s *BmSampleCoverStorage) Count(req api2go.Request, c BmModel.SampleCover) int {
	r, _ := s.db.Count(req, &c)
	return r
}
