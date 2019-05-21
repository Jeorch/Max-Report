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

// BmMarketStorage stores all Markets
type BmMarketStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmMarketStorage) NewMarketStorage(args []BmDaemons.BmDaemon) *BmMarketStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmMarketStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmMarketStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Market {
	in := BmModel.Market{}
	var out []BmModel.Market
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Market
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Market)
	}
}

// GetOne model
func (s BmMarketStorage) GetOne(id string) (BmModel.Market, error) {
	in := BmModel.Market{ID: id}
	model := BmModel.Market{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Market for id %s not found", id)
	return BmModel.Market{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmMarketStorage) Insert(c BmModel.Market) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmMarketStorage) Delete(id string) error {
	in := BmModel.Market{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Market with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmMarketStorage) Update(c BmModel.Market) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Market with id does not exist")
	}

	return nil
}

func (s *BmMarketStorage) Count(req api2go.Request, c BmModel.Market) int {
	r, _ := s.db.Count(req, &c)
	return r
}
