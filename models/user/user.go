package usermodel

type Student struct {
	ID    string `form:"id" json:"id"`
	Name  string `form:"name" json:"name"`
	Age   int    `form:"age" json:"age"`
	Grade int    `form:"grade" json:"grade"`
}

type Meta struct {
	TotalData int `json:"totalData"`
}

type Response struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Meta    Meta      `json:"meta"`
	Data    []Student `json:"data"`
}

type ResponseOne struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Meta    Meta    `json:"meta"`
	Data    Student `json:"data"`
}
