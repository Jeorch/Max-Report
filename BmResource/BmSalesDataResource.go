package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmSalesDataResource struct {
	BmCityStorage      *BmDataStorage.BmCityStorage
	BmProductStorage   *BmDataStorage.BmProductStorage
	BmSalesDataStorage *BmDataStorage.BmSalesDataStorage
}

func (c BmSalesDataResource) NewSalesDataResource(args []BmDataStorage.BmStorage) BmSalesDataResource {
	var cs *BmDataStorage.BmCityStorage
	var ps *BmDataStorage.BmProductStorage
	var sds *BmDataStorage.BmSalesDataStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmCityStorage" {
			cs = arg.(*BmDataStorage.BmCityStorage)
		} else if tp.Name() == "BmProductStorage" {
			ps = arg.(*BmDataStorage.BmProductStorage)
		} else if tp.Name() == "BmSalesDataStorage" {
			sds = arg.(*BmDataStorage.BmSalesDataStorage)
		}
	}
	return BmSalesDataResource{
		BmCityStorage:      cs,
		BmProductStorage:   ps,
		BmSalesDataStorage: sds,
	}
}

// FindAll SalesDatas
func (c BmSalesDataResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []*BmModel.SalesData
	_, ok := r.QueryParams["info-id"]
	if ok {
		result = c.BmSalesDataStorage.GetAll(r, -1, -1)
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
	return &Response{Res: result}, nil
}

// FindOne choc
func (c BmSalesDataResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmSalesDataStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmSalesDataResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.SalesData)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmSalesDataStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmSalesDataResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmSalesDataStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmSalesDataResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.SalesData)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmSalesDataStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
