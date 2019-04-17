package BmHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/alfredyang1986/blackmirror/jsonapi/jsonapiobj"
	"reflect"
	"strconv"
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/julienschmidt/httprouter"
)
type GetMarketHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
}

func (h GetMarketHandler) NewBmGetMarketHandler(args ...interface{}) GetMarketHandler {
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
	return GetMarketHandler{Method: md, HttpMethod: hm, Args: ag, db: m}
}

func (h GetMarketHandler) GetMarket(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")
	in := BmModel.MarketDimension{}
	var out  BmModel.MarketDimension
	jso := jsonapiobj.JsResult{}
	response := map[string]interface{}{
		"status": "",
		"product_count": nil,
		"product_countring":  nil,
		"sales_somring":  nil,
		"sales_som":  nil,
		"product_count_ring_growth":  nil,
		"product_countYear_on_year":  nil,
		"product_count_year_growth":  nil,
		"sales_somyear_on_year":  nil,
		"sales_som_year_growth":  nil,
		"sales_som_ring_growth":  nil,
		"error":  nil,
	}
	//同年同月多个市场
	ym,_:= strconv.Atoi(r.Header["Ym"][0])
	condtmp := bson.M{"YM":ym,
	"MARKET":r.Header["Market"][0],
	"COMPANY_ID":r.Header["Company-Id"][0]}
	fmt.Println(condtmp)
	err := h.db.FindOneByCondition(&in,&out,bson.M{})
	if err != nil{
		return 0
	}
	response["product_count"] = out.City_Count
	response["product_countring"] = out.Product_Countring
	response["sales_somring"] = out.Sales_SomRing
	response["sales_som"] = out.Sales_Som
	response["product_count_ring_growth"] = out.Product_Count_Ring_Growth
	response["product_countYear_on_year"] = out.Product_CountYear_On_Year
	response["product_count_year_growth"] = out.Product_Count_Ring_Growth
	response["sales_somyear_on_year"] = out.Sales_SomYear_On_Year
	response["sales_som_year_growth"] = out.Sales_Som_Year_Growth
	response["sales_som_ring_growth"] = out.Sales_Som_Ring_Growth
	response["status"] = "ok"
	jso.Obj = response
	enc := json.NewEncoder(w)
	enc.Encode(jso.Obj)
	return 0

}

func (h GetMarketHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h GetMarketHandler) GetHandlerMethod() string {
	return h.Method
}


