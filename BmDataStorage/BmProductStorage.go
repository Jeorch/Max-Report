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

// BmProductStorage stores all Products
type BmProductStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmProductStorage) NewProductStorage(args []BmDaemons.BmDaemon) *BmProductStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmProductStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmProductStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Product {
	in := BmModel.Product{}
	var out []BmModel.Product
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Product
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Product)
	}
}

// GetOne model
func (s BmProductStorage) GetOne(id string) (BmModel.Product, error) {
	in := BmModel.Product{ID: id}
	model := BmModel.Product{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Product for id %s not found", id)
	return BmModel.Product{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmProductStorage) Insert(c BmModel.Product) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmProductStorage) Delete(id string) error {
	in := BmModel.Product{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Product with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmProductStorage) Update(c BmModel.Product) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Product with id does not exist")
	}

	return nil
}

func (s *BmProductStorage) Count(req api2go.Request, c BmModel.Product) int {
	r, _ := s.db.Count(req, &c)
	return r
}
