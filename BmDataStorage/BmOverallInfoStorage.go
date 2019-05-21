package BmDataStorage

import (
	"errors"
	"fmt"
	"net/http"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
)

// BmOverallInfoStorage stores all OverallInfoes
type BmOverallInfoStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmOverallInfoStorage) NewOverallInfoStorage(args []BmDaemons.BmDaemon) *BmOverallInfoStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmOverallInfoStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmOverallInfoStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.OverallInfo {
	in := BmModel.OverallInfo{}
	var out []BmModel.OverallInfo
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.OverallInfo
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.OverallInfo)
	}
}

// GetOne model
func (s BmOverallInfoStorage) GetOne(id string) (BmModel.OverallInfo, error) {
	in := BmModel.OverallInfo{ID: id}
	model := BmModel.OverallInfo{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("OverallInfo for id %s not found", id)
	return BmModel.OverallInfo{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmOverallInfoStorage) Insert(c BmModel.OverallInfo) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmOverallInfoStorage) Delete(id string) error {
	in := BmModel.OverallInfo{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("OverallInfo with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmOverallInfoStorage) Update(c BmModel.OverallInfo) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("OverallInfo with id does not exist")
	}

	return nil
}

func (s *BmOverallInfoStorage) Count(req api2go.Request, c BmModel.OverallInfo) int {
	r, _ := s.db.Count(req, &c)
	return r
}
