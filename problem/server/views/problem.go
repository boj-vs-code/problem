package views

import (
	"github.com/hwangseonu/gin-restful"
	"github.com/moreal/boj-vs-code-api-server/problem"
)

type Problems struct {
	*gin_restful.Resource
}

func InitProblemsResource() Problems {
	return Problems{gin_restful.InitResource()}
}

func (r Problems) Get(id int) (*problem.ProblemModel, int) {
	return problem.FindProblemById(id), 200
}