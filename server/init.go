package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-restful"
	"github.com/moreal/boj-vs-code-api-server/server/views"
)

func CreateServer() *gin.Engine {
	router := gin.Default()
	api := gin_restful.NewApi(router,"/")
	views.RegisterViews(api)
	return router
}