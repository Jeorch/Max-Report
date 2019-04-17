package BmHandler

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"github.com/alfredyang1986/blackmirror/jsonapi/jsonapiobj"
	"reflect"
	//"os"
	//"bytes"
	//"strings"
	//"strconv"
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/julienschmidt/httprouter"
)
type MarketlistHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
}
type PersonInfo struct {
    Name    string
    age     int32
    Sex     bool
    Hobbies []string
}

func (h MarketlistHandler) NewBmMarketlistHandler(args ...interface{}) MarketlistHandler {
	var m *BmMongodb.BmMongodb
	var hm string
	var md string
	var ag []string
	for i, arg := range args {
		if i == 0 {
			sts := arg.([]BmDaemons.BmDaemon)
			for _, dm := range sts {
				tp := reflect.ValueOf(dm).Interface()
				tm := reflect.ValueOf(tp).Elem().Type()
				if tm.Name() == "BmMongodb" {
					m = dm.(*BmMongodb.BmMongodb)
				}
			}
		} else if i == 1 {
			md = arg.(string)
		} else if i == 2 {
			hm = arg.(string)
		} else if i == 3 {
			lst := arg.([]string)
			for _, str := range lst {
				ag = append(ag, str)
			}
		} else {
		}
	}
	return MarketlistHandler{Method: md, HttpMethod: hm, Args: ag, db: m}
}
func (h MarketlistHandler) Marketlist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")
	response := make(map[string]int)
	in := BmModel.Marketdimension{}
	var out []BmModel.Marketdimension
	jso := jsonapiobj.JsResult{}
	cond := bson.M{}
	h.db.FindMultiByCondition(&in,out,cond,"",-1,-1)
	for _,mark := range out{
		response[mark.Market] = 1
	}
	result := map[string]interface{}{
		"status": "",
		"result": response,
		"error":  nil,
	}
	result["status"] = "ok"
	jso.Obj = result
	enc := json.NewEncoder(w)
	enc.Encode(jso.Obj)
	return 0
}