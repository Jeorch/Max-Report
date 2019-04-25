package BmFactory

import (
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmHandler"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/PharbersDeveloper/Max-Report/BmResource"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/PharbersDeveloper/Max-Report/BmMiddleware"
)

type BmTable struct{}

var BLACKMIRROR_MODEL_FACTORY = map[string]interface{}{
	"BmProductdimension":                BmModel.Productdimension{},
	"BmMarketdimension":                 BmModel.Marketdimension{},
	"BmProd_etc":                		 BmModel.Prod_etc{},
}

var BLACKMIRROR_RESOURCE_FACTORY = map[string]interface{}{
	"BmProductdimensionResource":               BmResource.BmProductdimensionResource{},
	"BmMarketdimensionResource":                BmResource.BmMarketdimensionResource{},
}

var BLACKMIRROR_STORAGE_FACTORY = map[string]interface{}{
	"BmProductdimensionStorage":               BmDataStorage.BmProductdimensionStorage{},
	"BmMarketdimensionStorage":                BmDataStorage.BmMarketdimensionStorage{},
	"BmProd_etcStorage":               BmDataStorage.BmProd_etcStorage{},
}

var BLACKMIRROR_MIDDLEWARE_FACTORY = map[string]interface{}{
	"BmCheckTokenMiddleware": BmMiddleware.BmCheckTokenMiddleware{},
}

var BLACKMIRROR_DAEMON_FACTORY = map[string]interface{}{
	 "BmMongodbDaemon": BmMongodb.BmMongodb{},
	 "BmRedisDaemon":   BmRedis.BmRedis{},
}

var BLACKMIRROR_FUNCTION_FACTORY = map[string]interface{}{
	 "BmUserAgentHandler":       	   BmHandler.UserAgentHandler{},
	 "BmGenerateAccessTokenHandler":    BmHandler.BmGenerateAccessTokenHandler{},
	 "BmRefreshAccessTokenHandler":     BmHandler.RefreshAccessTokenHandler{},
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