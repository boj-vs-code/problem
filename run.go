package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-restful"
	"net/http"
)

type HelloWorld struct {
	*gin_restful.Resource
}

func (r HelloWorld) Get(Id int) (gin.H, int) {
	return gin.H{"id": Id}, http.StatusOK
}

func main() {
	router := gin.Default()
	api := gin_restful.NewApi(router,"/")
	api.AddResource(
		HelloWorld{gin_restful.InitResource()}, "/")
	router.Run(":7000")
}