package BmFactory

import (
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmHandler"
	"github.com/PharbersDeveloper/Max-Report/BmMiddleware"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/PharbersDeveloper/Max-Report/BmResource"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
)

type BmTable struct{}

var BLACKMIRROR_MODEL_FACTORY = map[string]interface{}{
	"BmProductdimension":   BmModel.Productdimension{},
	"BmMarketdimension":    BmModel.Marketdimension{},
	"BmProductaggregation": BmModel.Productaggregation{},
	"BmMarketaggregation":  BmModel.Marketaggregation{},
	"BmMarket":             BmModel.Market{},
	"BmCity":               BmModel.City{},
	"BmProduct":            BmModel.Product{},
	"BmOverallInfo":        BmModel.OverallInfo{},
	"BmAvailableCity":      BmModel.AvailableCity{},
	"BmAvailableDate":      BmModel.AvailableDate{},
	"BmSampleCover":        BmModel.SampleCover{},
	"BmSalesRecord":        BmModel.SalesRecord{},
}

var BLACKMIRROR_STORAGE_FACTORY = map[string]interface{}{
	"BmProductdimensionStorage": BmDataStorage.BmProductdimensionStorage{},
	"BmMarketdimensionStorage":  BmDataStorage.BmMarketdimensionStorage{},
	"BmProductaggdataStorage":   BmDataStorage.BmProductaggdataStorage{},
	"BmMarketaggdataStorage":    BmDataStorage.BmMarketaggdataStorage{},
	"BmMarketStorage":           BmDataStorage.BmMarketStorage{},
	"BmCityStorage":             BmDataStorage.BmCityStorage{},
	"BmProductStorage":          BmDataStorage.BmProductStorage{},
	"BmOverallInfoStorage":      BmDataStorage.BmOverallInfoStorage{},
	"BmAvailableCityStorage":    BmDataStorage.BmAvailableCityStorage{},
	"BmAvailableDateStorage":    BmDataStorage.BmAvailableDateStorage{},
	"BmSampleCoverStorage":      BmDataStorage.BmSampleCoverStorage{},
	"BmSalesRecordStorage":      BmDataStorage.BmSalesRecordStorage{},
}

var BLACKMIRROR_RESOURCE_FACTORY = map[string]interface{}{
	"BmProductdimensionResource": BmResource.BmProductdimensionResource{},
	"BmMarketdimensionResource":  BmResource.BmMarketdimensionResource{},
	"BmProductaggdataResource":   BmResource.BmProductaggdataResource{},
	"BmMarketaggdataResource":    BmResource.BmMarketaggdataResource{},
	"BmMarketResource":           BmResource.BmMarketResource{},
	"BmOverallInfoResource":      BmResource.BmOverallInfoResource{},
	"BmAvailableCityResource":    BmResource.BmAvailableCityResource{},
	"BmAvailableDateResource":    BmResource.BmAvailableDateResource{},
	"BmSampleCoverResource":      BmResource.BmSampleCoverResource{},
	"BmSalesRecordResource":      BmResource.BmSalesRecordResource{},
}

var BLACKMIRROR_MIDDLEWARE_FACTORY = map[string]interface{}{
	"BmCheckTokenMiddleware": BmMiddleware.BmCheckTokenMiddleware{},
}

var BLACKMIRROR_DAEMON_FACTORY = map[string]interface{}{
	"BmMongodbDaemon": BmMongodb.BmMongodb{},
	"BmRedisDaemon":   BmRedis.BmRedis{},
}

var BLACKMIRROR_FUNCTION_FACTORY = map[string]interface{}{
	"BmUserAgentHandler":           BmHandler.UserAgentHandler{},
	"BmGenerateAccessTokenHandler": BmHandler.BmGenerateAccessTokenHandler{},
	"BmRefreshAccessTokenHandler":  BmHandler.RefreshAccessTokenHandler{},
}

func (t BmTable) GetModelByName(name string) interface{} {
	return BLACKMIRROR_MODEL_FACTORY[name]
}

func (t BmTable) GetResourceByName(name string) interface{} {
	return BLACKMIRROR_RESOURCE_FACTORY[name]
}

func (t BmTable) GetStorageByName(name string) interface{} {
	return BLACKMIRROR_STORAGE_FACTORY[name]
}

func (t BmTable) GetDaemonByName(name string) interface{} {
	return BLACKMIRROR_DAEMON_FACTORY[name]
}

func (t BmTable) GetFunctionByName(name string) interface{} {
	return BLACKMIRROR_FUNCTION_FACTORY[name]
}

func (t BmTable) GetMiddlewareByName(name string) interface{} {
	return BLACKMIRROR_MIDDLEWARE_FACTORY[name]
}
