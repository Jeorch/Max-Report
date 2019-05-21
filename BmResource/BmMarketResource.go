package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmMarketResource struct {
	BmMarketStorage *BmDataStorage.BmMarketStorage
}

func (c BmMarketResource) NewMarketResource(args []BmDataStorage.BmStorage) BmMarketResource {
	var cs *BmDataStorage.BmMarketStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmMarketStorage" {
			cs = arg.(*BmDataStorage.BmMarketStorage)
		}
	}
	return BmMarketResource{BmMarketStorage: cs}
}

// FindAll Markets
func (c BmMarketResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []*BmModel.Market
	_, ok := r.QueryParams["company-id"]
	if ok {
		result = c.BmMarketStorage.GetAll(r, -1, -1)
		return &Response{Res: result}, nil
	}
	return &Response{Res: result}, nil
}

// FindOne choc
func (c BmMarketResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmMarketStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmMarketResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Market)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmMarketStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmMarketResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmMarketStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmMarketResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Market)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmMarketStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
