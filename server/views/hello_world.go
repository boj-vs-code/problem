package views

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