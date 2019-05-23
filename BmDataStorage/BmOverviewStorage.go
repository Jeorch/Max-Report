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

// BmOverviewStorage stores all Overviews
type BmOverviewStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmOverviewStorage) NewStorage(args []BmDaemons.BmDaemon) *BmOverviewStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmOverviewStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmOverviewStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Overview {
	in := BmModel.Overview{}
	var out []BmModel.Overview
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Overview
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Overview)
	}
}

// GetOne model
func (s BmOverviewStorage) GetOne(id string) (BmModel.Overview, error) {
	in := BmModel.Overview{ID: id}
	model := BmModel.Overview{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Overview for id %s not found", id)
	return BmModel.Overview{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmOverviewStorage) Insert(c BmModel.Overview) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmOverviewStorage) Delete(id string) error {
	in := BmModel.Overview{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Overview with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmOverviewStorage) Update(c BmModel.Overview) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Overview with id does not exist")
	}

	return nil
}

func (s *BmOverviewStorage) Count(req api2go.Request, c BmModel.Overview) int {
	r, _ := s.db.Count(req, &c)
	return r
}
