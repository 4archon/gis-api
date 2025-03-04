package server

import (
	"net/http"
	"html/template"
)


func (s Server) tasks(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	data, err := s.DB.GetTasksInfo()
	if err != nil {
		return
	}

	tmpl, _ := template.ParseFiles("server/templates/tasks/tasks.html")
	tmpl.Execute(response, data)
}