package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"fmt"
)

type BmMarketdimensionResource struct {
	BmMarketdimensionStorage *BmDataStorage.BmMarketdimensionStorage
}

func (c BmMarketdimensionResource) NewMarketdimensionResource(args []BmDataStorage.BmStorage) BmMarketdimensionResource {
	var cs *BmDataStorage.BmMarketdimensionStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmMarketdimensionStorage" {
			cs = arg.(*BmDataStorage.BmMarketdimensionStorage)
		}	
	}
	return BmMarketdimensionResource{BmMarketdimensionStorage: cs}
}

// FindAll Marketdimensions
func (c BmMarketdimensionResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var results []string
	var list BmModel.List
	_, ok := r.QueryParams["infomation"]
	if ok{
		result := c.BmMarketdimensionStorage.GetAll(r,-1,-1)
		for _,info := range result{
			tmpstr := info.Market+":" + fmt.Sprintf("%d",info.Ym)
			results = append(results,tmpstr)
		}
		list.Results=results
		return &Response{Res: list}, nil
	}
	result := c.BmMarketdimensionStorage.GetAll(r,-1,-1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c BmMarketdimensionResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmMarketdimensionStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmMarketdimensionResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Marketdimension)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmMarketdimensionStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmMarketdimensionResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmMarketdimensionStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmMarketdimensionResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Marketdimension)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmMarketdimensionStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
