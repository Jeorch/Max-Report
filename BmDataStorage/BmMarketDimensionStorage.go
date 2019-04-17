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

// BmMarketdimensionStorage stores all Marketdimensiones
type BmMarketdimensionStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmMarketdimensionStorage) NewMarketdimensionStorage(args []BmDaemons.BmDaemon) *BmMarketdimensionStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmMarketdimensionStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmMarketdimensionStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Marketdimension {
	in := BmModel.Marketdimension{}
	var out []BmModel.Marketdimension
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Marketdimension
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Marketdimension)
	}
}

// GetOne model
func (s BmMarketdimensionStorage) GetOne(id string) (BmModel.Marketdimension, error) {
	in := BmModel.Marketdimension{ID: id}
	model := BmModel.Marketdimension{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Marketdimension for id %s not found", id)
	return BmModel.Marketdimension{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmMarketdimensionStorage) Insert(c BmModel.Marketdimension) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmMarketdimensionStorage) Delete(id string) error {
	in := BmModel.Marketdimension{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Marketdimension with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmMarketdimensionStorage) Update(c BmModel.Marketdimension) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Marketdimension with id does not exist")
	}

	return nil
}

func (s *BmMarketdimensionStorage) Count(req api2go.Request, c BmModel.Marketdimension) int {
	r, _ := s.db.Count(req, &c)
	return r
}
