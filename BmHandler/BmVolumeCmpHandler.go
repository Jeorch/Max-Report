package BmHandler

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
	"github.com/alfredyang1986/blackmirror/jsonapi/jsonapiobj"
	"reflect"
	//"strings"
	"gopkg.in/mgo.v2/bson"
	//"github.com/manyminds/api2go"
	"time"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/julienschmidt/httprouter"
)
type VolumeCmpHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
}

func (h VolumeCmpHandler) NewBmVolumeCmpHandler(args ...interface{}) VolumeCmpHandler {
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
	return VolumeCmpHandler{Method: md, HttpMethod: hm, Args: ag, db: m}
}

func (h VolumeCmpHandler) VolumeCmp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")
	in := BmModel.MarketDimension{}
	//var out []BmModel.MarketDimension
	var oneout BmModel.MarketDimension
	jso := jsonapiobj.JsResult{}
	//var sum float64
	response := map[string]interface{}{
		"status": "",
		"sum": nil,
		"same":  nil,
		"ring":  nil,
		"error":  nil,
	}

	t := time.Now()
	tm := t.UTC()
	n := tm.Year()
	y := tm.Month()

	//本年
	ps := fmt.Sprintf("%d-%02d", n,y)
	cond := bson.M{"ym": ps}
	err := h.db.FindOneByCondition(&in,&oneout,cond)
	if err != nil{
		return 0
	}
	Product_Count := oneout.Product_Count
	response["sum"] = fmt.Sprintf("%f", Product_Count)
	
	//同比 
	ln:=n-1
	lps := fmt.Sprintf("%d-%02d", ln,y)
	if len(r.Header["Market"][0])<=0{
		return 0
	}
	cond = bson.M{"ym": lps,"market":r.Header["Market"][0]}
	err = h.db.FindOneByCondition(&in,&oneout,cond)
	if err != nil{
		return 0
	}
	same := Product_Count/oneout.Product_Count
	response["same"] = fmt.Sprintf("%f", same)
	//环比
	ly := y-1
	lps = fmt.Sprintf("%d-%02d", n,ly)
	cond = bson.M{"ym": lps,"market":r.Header["Market"][0]}
	err = h.db.FindOneByCondition(&in,&oneout,cond)
	if err != nil{
		return 0
	}
	ring := Product_Count/oneout.Product_Count
	response["ring"] = fmt.Sprintf("%f", ring)
	response["status"] = "ok"
	jso.Obj = response
	enc := json.NewEncoder(w)
	enc.Encode(jso.Obj)
	return 0
}

func (h VolumeCmpHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h VolumeCmpHandler) GetHandlerMethod() string {
	return h.Method
}

