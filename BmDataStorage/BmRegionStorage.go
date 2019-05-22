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

// BmRegionStorage stores all Regions
type BmRegionStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmRegionStorage) NewStorage(args []BmDaemons.BmDaemon) *BmRegionStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmRegionStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmRegionStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Region {
	in := BmModel.Region{}
	var out []BmModel.Region
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Region
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Region)
	}
}

// GetOne model
func (s BmRegionStorage) GetOne(id string) (BmModel.Region, error) {
	in := BmModel.Region{ID: id}
	model := BmModel.Region{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Region for id %s not found", id)
	return BmModel.Region{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmRegionStorage) Insert(c BmModel.Region) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmRegionStorage) Delete(id string) error {
	in := BmModel.Region{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Region with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmRegionStorage) Update(c BmModel.Region) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Region with id does not exist")
	}

	return nil
}

func (s *BmRegionStorage) Count(req api2go.Request, c BmModel.Region) int {
	r, _ := s.db.Count(req, &c)
	return r
}
