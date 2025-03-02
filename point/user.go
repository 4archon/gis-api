package point

import (
	"net/http"
)

type User struct {
	ID 			int
	Login		string
	Role		string
	Active		string
	Name		string
	Surname		string
	Patronymic	string
	TgID		int64
}

type NewUser struct {
	Login		string
	Password	string
	Role		string
	Active		string
	Name		string
	Surname		string
	Patronymic	string
	TgID		string
}

func (user *NewUser) Init(req *http.Request) {
	user.Login = req.FormValue("inputEmail")
	user.Password = req.FormValue("inputPassword")
	user.Role = req.FormValue("inputRole")
	user.Active = req.FormValue("inputActive")
	user.Name = req.FormValue("inputName")
	user.Surname = req.FormValue("inputSurname")
	user.Patronymic = req.FormValue("inputPatronymic")
	user.TgID = req.FormValue("inputTgID")
}