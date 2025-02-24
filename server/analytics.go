package server

import (
	"net/http"
	"html/template"
)


func (s Server) analytics(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	var data dataMain
	data.GisApiKey = s.GisApi
	data.Points = s.DB.GetPoints()
	tmpl, _ := template.ParseFiles("server/templates/analytics/analytics.html")
	tmpl.Execute(response, data)
}