package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-restful"
	"github.com/moreal/boj-vs-code-api-server/problem"
)

func CreateServer() *gin.Engine {
	router := gin.Default()
	api := gin_restful.NewApi(router,"/")
	problem.RegisterViews(api)
	return router
}