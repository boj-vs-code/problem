package models

type Testcase struct {
	Input string `json:"input"`
	Output string `json:"output"`
}

type ProblemModel struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	InputDescription string `json:"inputDescription"`
	OutputDescription string `json:"outputDescription"`
	Testcases []Testcase `json:"testcases"`
}