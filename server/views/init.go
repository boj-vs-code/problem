package views

import "github.com/hwangseonu/gin-restful"

func RegisterViews(api *gin_restful.Api) {
	api.AddResource(HelloWorld{gin_restful.InitResource()}, "/helloworld")
}