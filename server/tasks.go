package server

import (
	"net/http"
	// "html/template"
)


func (s Server) tasks(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	// var data dataMain
	// data.GisApiKey = s.GisApi
	// data.Points = s.DB.GetPoints()
	// tmpl, _ := template.ParseFiles("server/templates/distribute_tasks/distribute.html")
	// tmpl.Execute(response, data)

	http.ServeFile(response, req, "server/templates/tasks/tasks.html")
}