package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmAvailableAddressResource struct {
	BmCityStorage          *BmDataStorage.BmCityStorage
	BmAvailableAddressStorage *BmDataStorage.BmAvailableAddressStorage
}

func (c BmAvailableAddressResource) NewResource(args []BmDataStorage.BmStorage) BmAvailableAddressResource {
	var cs *BmDataStorage.BmCityStorage
	var acs *BmDataStorage.BmAvailableAddressStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmCityStorage" {
			cs = arg.(*BmDataStorage.BmCityStorage)
		} else if tp.Name() == "BmAvailableAddressStorage" {
			acs = arg.(*BmDataStorage.BmAvailableAddressStorage)
		}
	}
	return BmAvailableAddressResource{
		BmCityStorage:          cs,
		BmAvailableAddressStorage: acs,
	}
}

// FindAll AvailableAddresss
func (c BmAvailableAddressResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []*BmModel.AvailableAddress
	_, ok := r.QueryParams["info-id"]
	if ok {
		result = c.BmAvailableAddressStorage.GetAll(r, -1, -1)
		for _, sc := range result {
			if sc.AddressType == 1 {
				city, _ := c.BmCityStorage.GetOne(sc.AddressId)
				sc.City = &city
			}
		}
		return &Response{Res: result}, nil
	}
	return &Response{Res: result}, nil
}

// FindOne choc
func (c BmAvailableAddressResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmAvailableAddressStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmAvailableAddressResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.AvailableAddress)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmAvailableAddressStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmAvailableAddressResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmAvailableAddressStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmAvailableAddressResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.AvailableAddress)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmAvailableAddressStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
