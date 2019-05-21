package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmOverallInfoResource struct {
	BmOverallInfoStorage *BmDataStorage.BmOverallInfoStorage
}

func (c BmOverallInfoResource) NewOverallInfoResource(args []BmDataStorage.BmStorage) BmOverallInfoResource {
	var cs *BmDataStorage.BmOverallInfoStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmOverallInfoStorage" {
			cs = arg.(*BmDataStorage.BmOverallInfoStorage)
		}
	}
	return BmOverallInfoResource{BmOverallInfoStorage: cs}
}

// FindAll OverallInfos
func (c BmOverallInfoResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	_, ok := r.QueryParams["market-id"]
	if ok {
		_, ok := r.QueryParams["orderby"]
		if ok {
			r.QueryParams["orderby"] = []string{"-UPDATE_TIME"}
			result := c.BmOverallInfoStorage.GetAll(r, 0, 1)
			if len(result) < 1 {
				return &Response{}, nil
			}
			return &Response{Res: result[0]}, nil
		}
	}
	return &Response{}, nil
}

// FindOne choc
func (c BmOverallInfoResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmOverallInfoStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmOverallInfoResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.OverallInfo)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmOverallInfoStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmOverallInfoResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmOverallInfoStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmOverallInfoResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.OverallInfo)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmOverallInfoStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
