package views

import (
	"github.com/hwangseonu/gin-restful"
	"github.com/moreal/boj-vs-code-api-server/problem"
)

func RegisterViews(api *gin_restful.Api) {
	api.AddResource(problem.InitProblemsResource(), "/problem")
}