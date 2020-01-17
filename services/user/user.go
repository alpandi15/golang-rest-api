package userservice

import (
	"log"
	"project1/config"
	Model "project1/models/user"
)

func Find() []Model.Student {
	var users Model.Student
	var arrUser []Model.Student
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
