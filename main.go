package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PharbersDeveloper/Max-Report/BmFactory"
	"github.com/PharbersDeveloper/Max-Report/BmMaxDefine"
	"github.com/alfredyang1986/BmServiceDef/BmApiResolver"
	"github.com/alfredyang1986/BmServiceDef/BmConfig"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
)

func main() {
	version := "v2"
	fmt.Println("pod archi begins")

	fac := BmFactory.BmTable{}
	var pod = BmMaxDefine.Pod{ Name: "swl test", Factory:fac }
	pod.RegisterSerFromYAML("resource/def.yaml")

	var bmRouter BmConfig.BmRouterConfig
	//本地测试使用
	os.Setenv("MAXVIEW_HOME", ".")

	bmRouter.GenerateConfig("MAXVIEW_HOME")
	addr := bmRouter.Host + ":" + bmRouter.Port
	fmt.Println("Listening on ", addr)
	api := api2go.NewAPIWithResolver(version, &BmApiResolver.RequestURL{Addr: addr})
	pod.RegisterAllResource(api)
	pod.RegisterAllFunctions(version, api)
	pod.RegisterAllMiddleware(api)
	handler := api.Handler().(*httprouter.Router)
	//pod.RegisterPanicHandler(handler)
	http.ListenAndServe(":"+bmRouter.Port, handler)

	fmt.Println("pod archi ends")
}
