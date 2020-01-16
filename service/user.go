package service

import (
	"log"
	"project1/config"
	"project1/models"
)

type Meta struct {
	TotalData int `json:"totalData"`
}

type Response struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Meta    Meta             `json:"meta"`
	Data    []models.Student `json:"data"`
}

type ResponseOne struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Meta    Meta             `json:"meta"`
	Data    []models.Student `json:"data"`
}

func GetAll() []models.Student {
	var users models.Student
	var arrUser []models.Student
	db, err := config.Connect()
	defer db.Close()

	rows, err := db.Query("select * from tb_student")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.ID, &users.Name, &users.Age, &users.Grade); err != nil {
			log.Fatal(err.Error())

		} else {
			arrUser = append(arrUser, users)
		}
	}

	return arrUser
}
