package server

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s Server) getViewReport(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	vars := mux.Vars(req)
	reportType := vars["type"]
	IDStr := vars["id"]
	id, err := strconv.Atoi(IDStr)
	if err != nil {
		log.Println(err)
		return
	}

	if reportType == "service" {
		serviceInfo, err := s.DB.GetService(id)
		if err != nil {
			return
		}
		tmpl, _ := template.ParseFiles("server/templates/view/service.html")
		tmpl.Execute(response, serviceInfo)
	} else if reportType == "inspection" {
		inspectionInfo, err := s.DB.GetInspection(id)
		if err != nil {
			return
		}
		tmpl, _ := template.ParseFiles("server/templates/view/inspection.html")
		tmpl.Execute(response, inspectionInfo)
	} else if reportType == "change_point" {
		change, err := s.DB.GetPointInfo(id)
		if err != nil {
			return
		}
		tmpl, _ := template.ParseFiles("server/templates/view/change_point.html")
		tmpl.Execute(response, change)
	} else if reportType == "deactivate" {
		active, err := s.DB.GetActiveStatus(id)
		if err != nil {
			return
		}
		tmpl, _ := template.ParseFiles("server/templates/view/deactivate.html")
		tmpl.Execute(response, active)
	} else {
		http.Redirect(response, req, "/tasks", http.StatusFound)
	}
}