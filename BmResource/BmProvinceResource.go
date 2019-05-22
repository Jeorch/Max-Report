package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmProvinceResource struct {
	BmProvinceStorage *BmDataStorage.BmProvinceStorage
}

func (c BmProvinceResource) NewResource(args []BmDataStorage.BmStorage) BmProvinceResource {
	var cs *BmDataStorage.BmProvinceStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmProvinceStorage" {
			cs = arg.(*BmDataStorage.BmProvinceStorage)
		}
	}
	return BmProvinceResource{BmProvinceStorage: cs}
}

// FindAll Provinces
func (c BmProvinceResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	result := c.BmProvinceStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c BmProvinceResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmProvinceStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmProvinceResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Province)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmProvinceStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmProvinceResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmProvinceStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmProvinceResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Province)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmProvinceStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
