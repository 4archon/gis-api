package server

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"map/point"

	"github.com/gorilla/mux"
	// "log"
)

func (s Server) getEmployees(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	if role == "worker" {
		http.Redirect(response, req, "main", http.StatusFound)
		return
	}
	users := s.DB.GetUsersInfo()
	tmpl, _ := template.ParseFiles("server/templates/employees/employees.html")
	tmpl.Execute(response, users)
}

func (s Server) newEmployee(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	if role == "worker" {
		http.Redirect(response, req, "main", http.StatusFound)
		return
	}

	if req.Method == "GET" {
		http.ServeFile(response, req, "server/templates/employees/new_employee.html")
	} else if req.Method == "POST"{
		var user point.NewUser
		user.Init(req)
		err = s.DB.NewUser(user)
		if err != nil {
			return
		}

		http.Redirect(response, req, "employees", http.StatusFound)
	} else {
		log.Println("Wrong Method")
		return
	}
}

func (s Server) editEmployee(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	if role == "worker" {
		http.Redirect(response, req, "main", http.StatusFound)
		return
	}

	if req.Method == "GET" {
		vars := mux.Vars(req)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return
		}
		if id < 10 {
			http.Redirect(response, req, "/", http.StatusFound)
		}
		user, err := s.DB.GetUserInfo(id)
		if err != nil {
			return
		}

		tmpl, _ := template.ParseFiles("server/templates/employees/edit_employee.html")
		tmpl.Execute(response, user)
	} else if req.Method == "POST" {
		vars := mux.Vars(req)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return
		}
		var user point.NewUser
		user.Init(req)
		err = s.DB.ChangeUserAllInfo(id, user)
		if err != nil {
			return
		}
		if user.Password != "" {
			err = s.DB.ChangeUserPassword(id, user.Password)
			if err != nil {
				return
			}
		}
		http.Redirect(response, req, "/edit_employee/" + strconv.Itoa(id), http.StatusFound)
	} else {
		log.Println("Wrong Method")
		return
	}
}

func (s Server) profile(response http.ResponseWriter, req *http.Request) {
	id, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if req.Method == "GET" {
		user, err := s.DB.GetUserInfo(id)
		if err != nil {
			return
		}

		tmpl, _ := template.ParseFiles("server/templates/employees/profile.html")
		tmpl.Execute(response, user)
	} else if req.Method == "POST" {
		name := req.FormValue("inputName")
		surname := req.FormValue("inputSurname")
		patronymic := req.FormValue("inputPatronymic")
		tgID := req.FormValue("inputTgID")
		err = s.DB.ChangeUserInfo(id, name, surname, patronymic, tgID)
		if err != nil {
			return
		}

		password := req.FormValue("inputPassword")
		if password != "" {
			err = s.DB.ChangeUserPassword(id, password)
			if err != nil {
				return
			}
		}

		http.Redirect(response, req, "profile", http.StatusFound)
	} else {
		log.Println("Wrong Method")
		return
	}

}