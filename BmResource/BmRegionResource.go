package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmRegionResource struct {
	BmRegionStorage *BmDataStorage.BmRegionStorage
}

func (c BmRegionResource) NewResource(args []BmDataStorage.BmStorage) BmRegionResource {
	var cs *BmDataStorage.BmRegionStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmRegionStorage" {
			cs = arg.(*BmDataStorage.BmRegionStorage)
		}
	}
	return BmRegionResource{BmRegionStorage: cs}
}

// FindAll Regions
func (c BmRegionResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	result := c.BmRegionStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c BmRegionResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmRegionStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmRegionResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Region)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmRegionStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmRegionResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmRegionStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmRegionResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Region)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmRegionStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
