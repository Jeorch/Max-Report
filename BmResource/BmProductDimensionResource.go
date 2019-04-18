package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	//"strconv"
	//"gopkg.in/mgo.v2/bson"
)

type BmProductdimensionResource struct {
	BmProductdimensionStorage *BmDataStorage.BmProductdimensionStorage
}

func (c BmProductdimensionResource) NewProductdimensionResource(args []BmDataStorage.BmStorage) BmProductdimensionResource {
	var cs *BmDataStorage.BmProductdimensionStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmProductdimensionStorage" {
			cs = arg.(*BmDataStorage.BmProductdimensionStorage)
		}	
	}
	return BmProductdimensionResource{BmProductdimensionStorage: cs}
}

// FindAll Productdimensions
func (c BmProductdimensionResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var results []*BmModel.Productdimension
	_, rankok := r.QueryParams["lt[sales_rank]"]
	_, gteym := r.QueryParams["gte[ym]"]
	_, lteym := r.QueryParams["lte[ym]"]
	var i int32
	if rankok && gteym && lteym{
		result := c.BmProductdimensionStorage.GetAll(r,-1,-1)
		for i=1;i<=10;i++{
			for _,mark := range result{
				if mark.Sales_Rank == i{
					results = append(results,mark)
				}
			}
		}
		return &Response{Res: results}, nil
	}
	result := c.BmProductdimensionStorage.GetAll(r,-1,-1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c BmProductdimensionResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmProductdimensionStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmProductdimensionResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Productdimension)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmProductdimensionStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmProductdimensionResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmProductdimensionStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmProductdimensionResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Productdimension)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmProductdimensionStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
