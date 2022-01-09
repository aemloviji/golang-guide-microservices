package domain

import (
	"fmt"
	"net/http"

	"github.com/aemloviji/golang-guide-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Fade", LastName: "Leon", Email: "myemail.gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}