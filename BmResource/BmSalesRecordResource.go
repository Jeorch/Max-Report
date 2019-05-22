package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"strconv"
)

type BmSalesRecordResource struct {
	BmCityStorage        *BmDataStorage.BmCityStorage
	BmProductStorage     *BmDataStorage.BmProductStorage
	BmSalesRecordStorage *BmDataStorage.BmSalesRecordStorage
}

func (c BmSalesRecordResource) NewSalesRecordResource(args []BmDataStorage.BmStorage) BmSalesRecordResource {
	var cs *BmDataStorage.BmCityStorage
	var ps *BmDataStorage.BmProductStorage
	var sds *BmDataStorage.BmSalesRecordStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmCityStorage" {
			cs = arg.(*BmDataStorage.BmCityStorage)
		} else if tp.Name() == "BmProductStorage" {
			ps = arg.(*BmDataStorage.BmProductStorage)
		} else if tp.Name() == "BmSalesRecordStorage" {
			sds = arg.(*BmDataStorage.BmSalesRecordStorage)
		}
	}
	return BmSalesRecordResource{
		BmCityStorage:        cs,
		BmProductStorage:     ps,
		BmSalesRecordStorage: sds,
	}
}

// FindAll SalesRecords
func (c BmSalesRecordResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []*BmModel.SalesRecord
	_, ok := r.QueryParams["info-id"]
	if ok {
		skipArgs, sok := r.QueryParams["skip"]
		takeArgs, tok := r.QueryParams["take"]
		if sok && tok {
			skip, _ := strconv.Atoi(skipArgs[0])
			take, _ := strconv.Atoi(takeArgs[0])
			result := c.BmSalesRecordStorage.GetAll(r, skip, take)
			for _, sc := range result {
				if sc.AddressType == 1 {
					city, _ := c.BmCityStorage.GetOne(sc.AddressId)
					sc.City = &city
				}

				if sc.GoodsType == 1 {
					product, _ := c.BmProductStorage.GetOne(sc.GoodsId)
					sc.Product = &product
				}
			}
			return &Response{Res: result}, nil
		}
	}
	return &Response{Res: result}, nil
}

// FindOne choc
func (c BmSalesRecordResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmSalesRecordStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmSalesRecordResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.SalesRecord)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmSalesRecordStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmSalesRecordResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmSalesRecordStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmSalesRecordResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.SalesRecord)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmSalesRecordStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
