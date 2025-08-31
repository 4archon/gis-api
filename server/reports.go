package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func (s Server) getReports(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	if role != "admin" {
		http.Redirect(response, req, "main", http.StatusFound)
		return
	}

	http.Redirect(response, req, "/reports/1", http.StatusFound)
}

func (s Server) getReportsPage(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	if role != "admin" {
		http.Redirect(response, req, "main", http.StatusFound)
		return
	}

	http.ServeFile(response, req, "server/static/reports/reports.html")
}

func (s Server) postReportsPage(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	if role != "admin" {
		http.Redirect(response, req, "main", http.StatusFound)
		return
	}

	vars := mux.Vars(req)
	pageStr := vars["page"]
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	
	services, err := s.DB.GetAllServices(10, page - 1)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	if services.CurrentPage > services.LastPage {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := json.Marshal(services)
	if err != nil {
		log.Println(err)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(result)
}