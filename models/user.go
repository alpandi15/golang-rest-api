package models

type Student struct {
	ID    string `form:"id" json:"id"`
	Name  string `form:"name" json:"name"`
	Age   int    `form:"age" json:"age"`
	Grade int    `form:"grade" json:"grade"`
}
