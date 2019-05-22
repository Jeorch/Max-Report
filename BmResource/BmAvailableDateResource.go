package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmAvailableDateResource struct {
	BmAvailableDateStorage *BmDataStorage.BmAvailableDateStorage
}

func (c BmAvailableDateResource) NewResource(args []BmDataStorage.BmStorage) BmAvailableDateResource {
	var cs *BmDataStorage.BmAvailableDateStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmAvailableDateStorage" {
			cs = arg.(*BmDataStorage.BmAvailableDateStorage)
		}
	}
	return BmAvailableDateResource{BmAvailableDateStorage: cs}
}

// FindAll AvailableDates
func (c BmAvailableDateResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []*BmModel.AvailableDate
	_, ok := r.QueryParams["info-id"]
	if ok {
		result = c.BmAvailableDateStorage.GetAll(r, -1, -1)
		return &Response{Res: result}, nil
	}
	return &Response{Res: result}, nil
}

// FindOne choc
func (c BmAvailableDateResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmAvailableDateStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmAvailableDateResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.AvailableDate)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmAvailableDateStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmAvailableDateResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmAvailableDateStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmAvailableDateResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.AvailableDate)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmAvailableDateStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
