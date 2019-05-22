package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmAvailableCityResource struct {
	BmCityStorage          *BmDataStorage.BmCityStorage
	BmAvailableCityStorage *BmDataStorage.BmAvailableCityStorage
}

func (c BmAvailableCityResource) NewResource(args []BmDataStorage.BmStorage) BmAvailableCityResource {
	var cs *BmDataStorage.BmCityStorage
	var acs *BmDataStorage.BmAvailableCityStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmCityStorage" {
			cs = arg.(*BmDataStorage.BmCityStorage)
		} else if tp.Name() == "BmAvailableCityStorage" {
			acs = arg.(*BmDataStorage.BmAvailableCityStorage)
		}
	}
	return BmAvailableCityResource{
		BmCityStorage:          cs,
		BmAvailableCityStorage: acs,
	}
}

// FindAll AvailableCitys
func (c BmAvailableCityResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []*BmModel.AvailableCity
	_, ok := r.QueryParams["info-id"]
	if ok {
		result = c.BmAvailableCityStorage.GetAll(r, -1, -1)
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
func (c BmAvailableCityResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmAvailableCityStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmAvailableCityResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.AvailableCity)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmAvailableCityStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmAvailableCityResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmAvailableCityStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmAvailableCityResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.AvailableCity)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmAvailableCityStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
