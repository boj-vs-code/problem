package models

import (
	"fmt"
	"github.com/andrewstuart/goq"
	"github.com/moreal/boj-vs-code-api-server/problem"
	"log"
	"net/http"
)

type ProblemModel struct {
	Id int `json:"id"`
	Title string `json:"title" goquery:"#problem_title" firestore:"title"`
	Description string `json:"description" goquery:"#problem_description" firestore:"description"`
	InputDescription string `json:"inputDescription" goquery:"#problem_input" firestore:"input_description"`
	OutputDescription string `json:"outputDescription" goquery:"#problem_output" firestore:"output_description"`
	Testcases []string `json:"testcases" goquery:".sampledata" firestore:"testcases"`
}

func (p *ProblemModel) Save() {
	problem.connection.Add("problems", p.Id, p)
}

func FindProblemById(id int) *ProblemModel {
	problem := fetchProblemFromDB(id)
	if problem == nil {
		problem = parse(id)
		problem.Save()
	}
	return problem
}

func fetchProblemFromDB(id int) *ProblemModel {
	problem.connection.Initialize()
	return problem.connection.Fetch("problems", id)
}

func parse(id int) *ProblemModel {
	res, err := http.Get(fmt.Sprintf("https://acmicpc.net/problem/%d", id))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var problem ProblemModel

	err = goq.NewDecoder(res.Body).Decode(&problem)
	if err != nil {
		log.Fatal(err)
	}

	problem.Id = id

	return &problem
}