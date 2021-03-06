package usercontroller

import (
	"encoding/json"
	"net/http"
	userModel "project1/models/user"
	userService "project1/services/user"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	var Response userModel.Response
	var Meta userModel.Meta

	user := userService.Find()
	Meta.TotalData = len(user)

	Response.Status = 1
	Response.Message = "Success"
	Response.Meta = Meta
	Response.Data = user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response)
}

/*
func FindAll(c echo.Context) error {
	var Response userModel.Response
	var Meta userModel.Meta

	user := userService.Find()
	Meta.TotalData = len(user)

	Response.Status = 1
	Response.Message = "Success"
	Response.Meta = Meta
	Response.Data = user

	return c.JSON(http.StatusOK, Response)
}

func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	return c.String(http.StatusOK, name)
}
*/
