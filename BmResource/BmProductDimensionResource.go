package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"strconv"
	"gopkg.in/mgo.v2/bson"
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
	var result []BmModel.Productdimension
	hospitalConfigsID, pciok := r.QueryParams["lt[sales_rank]"]
	if pciok {
		modelRootID := hospitalConfigsID[0]
		ranks,_ := strconv.Atoi(modelRootID)
		yml,_ := strconv.Atoi(r.QueryParams["gte[ym]"][0])
		ymn,_ := strconv.Atoi(r.QueryParams["lte[ym]"][0])
		s := make(map[string]interface{})
		s["$lte"] = ymn
		s["$gte"] = yml
		for i:=1;i<=ranks;i++{
			cond := bson.M{
							"COMPANY_ID":r.QueryParams["company_id"][0],"MARKET":r.QueryParams["market"][0],
							"YM":s,"SALES_RANK":i,
						}
			modelRoot:= c.BmProductdimensionStorage.GetAllByCond(cond,-1,-1)
			for _, model := range modelRoot {
				result = append(result,*model)
			}
			
		}
		// model, err := c.BmProductdimensionStorage.GetOne(modelRoot.HospitalID)
		// if err != nil {
		// 	return &Response{}, err
		// }
		return &Response{Res: result}, nil
	}
	results := c.BmProductdimensionStorage.GetAll(r,-1,-1)
	return &Response{Res: results}, nil
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
