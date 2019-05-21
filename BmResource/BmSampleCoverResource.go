package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmSampleCoverResource struct {
	BmSampleCoverStorage *BmDataStorage.BmSampleCoverStorage
	BmCityStorage        *BmDataStorage.BmCityStorage
}

func (c BmSampleCoverResource) NewSampleCoverResource(args []BmDataStorage.BmStorage) BmSampleCoverResource {
	var scs *BmDataStorage.BmSampleCoverStorage
	var cs *BmDataStorage.BmCityStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmSampleCoverStorage" {
			scs = arg.(*BmDataStorage.BmSampleCoverStorage)
		} else if tp.Name() == "BmCityStorage" {
			cs = arg.(*BmDataStorage.BmCityStorage)
		}
	}

	return BmSampleCoverResource{
		BmSampleCoverStorage: scs,
		BmCityStorage:        cs,
	}
}

// FindAll SampleCovers
func (c BmSampleCoverResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []*BmModel.SampleCover
	_, ok := r.QueryParams["info-id"]
	if ok {
		r.QueryParams["orderby"] = []string{"-COVERAGE_RATIO"}
		result = c.BmSampleCoverStorage.GetAll(r, -1, -1)
		for _, sc := range result {
			city, _ := c.BmCityStorage.GetOne(sc.CityId)
			sc.City = &city
		}
		return &Response{Res: result}, nil
	}
	return &Response{Res: result}, nil
}

// FindOne choc
func (c BmSampleCoverResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmSampleCoverStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmSampleCoverResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.SampleCover)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmSampleCoverStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmSampleCoverResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmSampleCoverStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmSampleCoverResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.SampleCover)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmSampleCoverStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
