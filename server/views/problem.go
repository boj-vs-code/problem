package views

import (
	"github.com/hwangseonu/gin-restful"
	"github.com/moreal/boj-vs-code-api-server/server/models"
)

type Problems struct {
	*gin_restful.Resource
}

func InitProblemsResource() Problems {
	return Problems{gin_restful.InitResource()}
}

func (r Problems) Get(id int) (*models.ProblemModel, int) {
	return models.FindProblemById(id), 200
}