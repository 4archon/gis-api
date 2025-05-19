package server

import (
	"encoding/json"
	"log"
	"net/http"
)


func (s Server) getUsers(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	if role == "worker" {
		http.Redirect(response, req, "main", http.StatusFound)
		return
	}
	
	http.ServeFile(response, req, "server/static/users/employees.html")
}

func (s Server) postUsers(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	if role == "worker" {
		http.Redirect(response, req, "main", http.StatusFound)
		return
	}
	
	users, err := s.DB.GetUsersInfo()
	if err != nil {
		return
	}

	result, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(result)
}

