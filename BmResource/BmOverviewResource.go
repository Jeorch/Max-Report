package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmOverviewResource struct {
	BmOverviewStorage *BmDataStorage.BmOverviewStorage
}

func (c BmOverviewResource) NewResource(args []BmDataStorage.BmStorage) BmOverviewResource {
	var cs *BmDataStorage.BmOverviewStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmOverviewStorage" {
			cs = arg.(*BmDataStorage.BmOverviewStorage)
		}
	}
	return BmOverviewResource{BmOverviewStorage: cs}
}

// FindAll Overviews
func (c BmOverviewResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	_, cok := r.QueryParams["company_id"]
	_, mok := r.QueryParams["market"]
	if cok && mok {
		_, ok := r.QueryParams["orderby"]
		if ok {
			result := c.BmOverviewStorage.GetAll(r, 0, 1)
			if len(result) < 1 {
				return &Response{}, nil
			}
			return &Response{Res: result[0]}, nil
		} else {
			r.QueryParams["orderby"] = []string{"-UPDATE_TIME"}
			result := c.BmOverviewStorage.GetAll(r, 0, 1)
			if len(result) < 1 {
				return &Response{}, nil
			}
			return &Response{Res: result[0]}, nil
		}
	}
	return &Response{}, nil
}

// FindOne choc
func (c BmOverviewResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmOverviewStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmOverviewResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Overview)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmOverviewStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmOverviewResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmOverviewStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmOverviewResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Overview)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmOverviewStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
