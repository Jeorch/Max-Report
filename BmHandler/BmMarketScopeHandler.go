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
type MarketScopeHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
}

func (h MarketScopeHandler) NewBmMarketScopeHandler(args ...interface{}) MarketScopeHandler {
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
	return MarketScopeHandler{Method: md, HttpMethod: hm, Args: ag, db: m}
}

func (h MarketScopeHandler) MarketScope(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")
	in := BmModel.MarketDimension{}
	var out []BmModel.MarketDimension
	var oneout BmModel.MarketDimension
	jso := jsonapiobj.JsResult{}
	var sum float64
	response := map[string]interface{}{
		"status": "",
		"result": nil,
		"error":  nil,
	}

	t := time.Now()
	tm := t.UTC()
	n := tm.Year()
	y := tm.Month()

	//同年同月多个市场
	ps := fmt.Sprintf("%d-%02d", n,y)
	condtmp := bson.M{"ym": ps}
	err := h.db.FindMultiByCondition(&in,&out,condtmp,"")
	if err != nil{
		return 0
	}
	for _,mark:=range out{
		sum+=mark.Sales
	}
	cond := bson.M{"ym": ps,"market":r.Header["Market"][0]}
	err = h.db.FindOneByCondition(&in,&oneout,cond)
	if err != nil{
		return 0
	}
	sale := oneout.Sales
	this := sale/sum
	sum=0

	//同比 
	ln:=n-1
	lps := fmt.Sprintf("%d-%02d", ln,y)
	condtmp = bson.M{"ym": lps}
	err = h.db.FindMultiByCondition(&in,&out,condtmp,"")
	if err != nil{
		return 0
	}
	for _,mark:=range out{
		sum+=mark.Sales
	}
	if len(r.Header["Market"][0])<=0{
		return 0
	}
	cond = bson.M{"ym": lps,"market":r.Header["Market"][0]}
	err = h.db.FindOneByCondition(&in,&oneout,cond)
	if err != nil{
		return 0
	}
	sale = oneout.Sales
	last := sale/sum
	same := this/last
	sum=0

	//环比
	ly := y-1
	lps = fmt.Sprintf("%d-%02d", n,ly)
	condtmp = bson.M{"ym": lps}
	err = h.db.FindMultiByCondition(&in,&out,condtmp,"")
	if err != nil{
		return 0
	}
	for _,mark:=range out{
		sum+=mark.Sales
	}
	cond = bson.M{"ym": lps,"market":r.Header["Market"][0]}
	err = h.db.FindOneByCondition(&in,&oneout,cond)
	sale = oneout.Sales
	if err != nil{
		return 0
	}
	last = sale/sum
	ring := this/last

	result := this + same +ring
	results:=fmt.Sprintf("%f", result)
	
	if err != nil && err.Error() == "not found" {
		return 0
	}

	response["status"] = "ok"
	response["result"] = results
	jso.Obj = response
	enc := json.NewEncoder(w)
	enc.Encode(jso.Obj)
	return 0
}

func (h MarketScopeHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h MarketScopeHandler) GetHandlerMethod() string {
	return h.Method
}


