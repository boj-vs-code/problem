package models

import (
	"fmt"
	"github.com/andrewstuart/goq"
	"log"
	"net/http"
)

var exampleProblems = []ProblemModel {{
	1,
	"Title",
	"Description",
	"InputDescription",
	"OutputDescription",
	[]string{"input-case","output-case"},
}}

type Testcase struct {
	Input string `json:"input" goquery:"#sample-input-1"`
	Output string `json:"output" goquery:"#sample-output-1"`
}

type Problem struct {
	Id int `json:"id"`
	Title string `json:"title" goquery:"#problem_title"`
	Description string `json:"description" goquery:"#problem_description"`
	InputDescription string `json:"inputDescription" goquery:"#problem_input"`
	OutputDescription string `json:"outputDescription" goquery:"#problem_output"`
	Testcases []string `json:"testcases" goquery:".sampledata"`
}

type ProblemModel struct {
	Problem
	Testcases []Testcase `json:"testcases"`
}

func (p *ProblemModel) Save() {
	// TODO: Save to CloudSQL
}


func FindProblemById(id int) *ProblemModel {
	// TODO: Fetch from CloudSQL or, Parsing and Save
	parse(1000)
	return &exampleProblems[0]
}

func parse(id int) {
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

	var p ProblemModel

	//return problem
}