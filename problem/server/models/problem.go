package models

import (
	"fmt"
	"github.com/andrewstuart/goq"
	"log"
	"net/http"
)

type ProblemModel struct {
	Id int `json:"id"`
	Title string `json:"title" goquery:"#problem_title,html" firestore:"title"`
	Description string `json:"description" goquery:"#problem_description,html" firestore:"description"`
	InputDescription string `json:"inputDescription" goquery:"#problem_input,html" firestore:"input_description"`
	OutputDescription string `json:"outputDescription" goquery:"#problem_output,html" firestore:"output_description"`
	Testcases []string `json:"testcases" goquery:".sampledata,html" firestore:"testcases"`
}

func (p *ProblemModel) Save() {
	connection.Add("problems", p.Id, p)
}

func FindProblemById(id int) *ProblemModel {
	problem := fetchProblemFromDB(id)
	if problem == nil {
		problem := parse(id)
		problem.Save()
	}

	return problem
}

func fetchProblemFromDB(id int) *ProblemModel {
	connection.Initialize()
	return connection.Fetch("problems", id)
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