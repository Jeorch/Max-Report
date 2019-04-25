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

// BmProd_etcStorage stores all Prod_etces
type BmProd_etcStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmProd_etcStorage) NewProd_etcStorage(args []BmDaemons.BmDaemon) *BmProd_etcStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmProd_etcStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmProd_etcStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Prod_etc {
	in := BmModel.Prod_etc{}
	var out []BmModel.Prod_etc
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Prod_etc
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Prod_etc)
	}
}

// GetOne model
func (s BmProd_etcStorage) GetOne(id string) (BmModel.Prod_etc, error) {
	in := BmModel.Prod_etc{ID: id}
	model := BmModel.Prod_etc{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Prod_etc for id %s not found", id)
	return BmModel.Prod_etc{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmProd_etcStorage) Insert(c BmModel.Prod_etc) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmProd_etcStorage) Delete(id string) error {
	in := BmModel.Prod_etc{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Prod_etc with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmProd_etcStorage) Update(c BmModel.Prod_etc) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Prod_etc with id does not exist")
	}

	return nil
}

func (s *BmProd_etcStorage) Count(req api2go.Request, c BmModel.Prod_etc) int {
	r, _ := s.db.Count(req, &c)
	return r
}
