package dto

import (
	"hka-server-login/model"
)

type User struct {
	Username    string `json:"username"`
	LastName    string `json:"lastName"`
	Course      string `json:"course"`
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
	Ldap        string `json:"ldap"`
	FirstName   string `json:"firstName"`
	Email       string `json:"email"`
	Faculty     string `json:"faculty"`
	Role        string `json:"role"`
	RoomNumber  string `json:"roomNumber"`
}

func (user *User) FromModel(userModel *model.User) {
	user.Username = userModel.Username
	user.LastName = userModel.LastName
	user.Course = userModel.Course
	user.Description = userModel.Description
	user.DisplayName = userModel.DisplayName
	user.Ldap = userModel.Ldap
	user.FirstName = userModel.FirstName
	user.Email = userModel.Email
	user.Faculty = userModel.Faculty
	user.Role = userModel.Role
	user.RoomNumber = userModel.RoomNumber
}
