package server

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "log"
)

func (s Server) getEmployees(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	users := s.DB.GetUsersInfo()
	tmpl, _ := template.ParseFiles("server/templates/employees/employees.html")
	tmpl.Execute(response, users)
}

func (s Server) newEmployee(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if req.Method == "GET" {
		http.ServeFile(response, req, "server/templates/employees/new_employee.html")
	}
}

func (s Server) editEmployee(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if req.Method == "GET" {
		vars := mux.Vars(req)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return
		}
		user, err := s.DB.GetUserInfo(id)
		if err != nil {
			return
		}

		tmpl, _ := template.ParseFiles("server/templates/employees/edit_employee.html")
		tmpl.Execute(response, user)
	}
}

func (s Server) profile(response http.ResponseWriter, req *http.Request) {
	id, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	user, err := s.DB.GetUserInfo(id)
	if err != nil {
		return
	}

	tmpl, _ := template.ParseFiles("server/templates/employees/profile.html")
		tmpl.Execute(response, user)

}