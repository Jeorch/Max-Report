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

type BmMarketaggdataResource struct {
	BmMarketaggdataStorage *BmDataStorage.BmMarketaggdataStorage
}

func (c BmMarketaggdataResource) NewResource(args []BmDataStorage.BmStorage) BmMarketaggdataResource {
	var cs *BmDataStorage.BmMarketaggdataStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmMarketaggdataStorage" {
			cs = arg.(*BmDataStorage.BmMarketaggdataStorage)
		}	
	}
	return BmMarketaggdataResource{BmMarketaggdataStorage: cs}
}

func (c BmMarketaggdataResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	skipArgs, sok := r.QueryParams["skip"]
	takeArgs, tok := r.QueryParams["take"]
	if sok && tok {
		skip, err:= strconv.Atoi(skipArgs[0])
		bmerror.PanicError(err)
		take, err:= strconv.Atoi(takeArgs[0])
		bmerror.PanicError(err)
		result := c.BmMarketaggdataStorage.GetAll(r, skip, take)
		return &Response{Res: result}, nil
	}
	result := c.BmMarketaggdataStorage.GetAll(r,-1,-1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c BmMarketaggdataResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmMarketaggdataStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmMarketaggdataResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Marketaggregation)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmMarketaggdataStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmMarketaggdataResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmMarketaggdataStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmMarketaggdataResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Marketaggregation)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmMarketaggdataStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
