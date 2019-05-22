package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmCityResource struct {
	BmCityStorage *BmDataStorage.BmCityStorage
}

func (c BmCityResource) NewResource(args []BmDataStorage.BmStorage) BmCityResource {
	var cs *BmDataStorage.BmCityStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmCityStorage" {
			cs = arg.(*BmDataStorage.BmCityStorage)
		}
	}
	return BmCityResource{BmCityStorage: cs}
}

// FindAll Citys
func (c BmCityResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	result := c.BmCityStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c BmCityResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmCityStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmCityResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.City)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmCityStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmCityResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmCityStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmCityResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.City)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmCityStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
