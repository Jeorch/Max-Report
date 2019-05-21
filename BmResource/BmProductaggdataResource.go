package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"strconv"
)

type BmProductaggdataResource struct {
	BmProductaggdataStorage *BmDataStorage.BmProductaggdataStorage
}

func (c BmProductaggdataResource) NewResource(args []BmDataStorage.BmStorage) BmProductaggdataResource {
	var cs *BmDataStorage.BmProductaggdataStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmProductaggdataStorage" {
			cs = arg.(*BmDataStorage.BmProductaggdataStorage)
		}	
	}
	return BmProductaggdataResource{BmProductaggdataStorage: cs}
}

func (c BmProductaggdataResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	skipArgs, sok := r.QueryParams["skip"]
	takeArgs, tok := r.QueryParams["take"]
	if sok && tok {
		skip, err:= strconv.Atoi(skipArgs[0])
		bmerror.PanicError(err)
		take, err:= strconv.Atoi(takeArgs[0])
		bmerror.PanicError(err)
		result := c.BmProductaggdataStorage.GetAll(r, skip, take)
		return &Response{Res: result}, nil
	}
	result := c.BmProductaggdataStorage.GetAll(r,-1,-1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c BmProductaggdataResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmProductaggdataStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmProductaggdataResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Productaggregation)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmProductaggdataStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmProductaggdataResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmProductaggdataStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmProductaggdataResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Productaggregation)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmProductaggdataStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
