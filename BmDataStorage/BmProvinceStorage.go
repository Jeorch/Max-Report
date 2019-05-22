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

// BmProvinceStorage stores all Provinces
type BmProvinceStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmProvinceStorage) NewStorage(args []BmDaemons.BmDaemon) *BmProvinceStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmProvinceStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmProvinceStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Province {
	in := BmModel.Province{}
	var out []BmModel.Province
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Province
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Province)
	}
}

// GetOne model
func (s BmProvinceStorage) GetOne(id string) (BmModel.Province, error) {
	in := BmModel.Province{ID: id}
	model := BmModel.Province{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Province for id %s not found", id)
	return BmModel.Province{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmProvinceStorage) Insert(c BmModel.Province) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmProvinceStorage) Delete(id string) error {
	in := BmModel.Province{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Province with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmProvinceStorage) Update(c BmModel.Province) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Province with id does not exist")
	}

	return nil
}

func (s *BmProvinceStorage) Count(req api2go.Request, c BmModel.Province) int {
	r, _ := s.db.Count(req, &c)
	return r
}
