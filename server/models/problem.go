package models

import (
	"encoding/json"
	"fmt"
	"github.com/andrewstuart/goq"
	"log"
	"net/http"
)

type ProblemModel struct {
	Id int `json:"id"`
	Title string `json:"title" goquery:"#problem_title"`
	Description string `json:"description" goquery:"#problem_description"`
	InputDescription string `json:"inputDescription" goquery:"#problem_input"`
	OutputDescription string `json:"outputDescription" goquery:"#problem_output"`
	Testcases []string `json:"testcases" goquery:".sampledata"`
}

func (p *ProblemModel) Save() {
	b, err := json.Marshal(p)
	if err != nil {
		log.Panic("Error")
	}
	connection.Exec("INSERT INTO problems (id, data) VALUES (?)", b)
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
	conn := CreateConnectionFromEnvironmentVariables()
	pre, err := conn.Prepare("SELECT * FROM problems WHERE id=?")
	if err != nil {
		log.Panic("Prepare statement create error")
	}

	rows, err := pre.Query(id)
	if err != nil {
		log.Panic("Prepare statement create error")
	}

	var problem ProblemModel

	if !rows.Next() {
		return nil
	} else {
		rows.Scan(&problem)
		return &problem
	}
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
	return &problem
}