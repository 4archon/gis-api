package server

import (
	"encoding/json"
	"log"
	"net/http"
	"io"
	"strconv"
)


func (s Server) getProfile(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	http.ServeFile(response, req, "server/static/profile/profile.html")
}


func (s Server) postProfile(response http.ResponseWriter, req *http.Request) {
	id, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	if role == "worker" {
		http.Redirect(response, req, "main", http.StatusFound)
		return
	}
	
	user, err := s.DB.GetUserInfo(id)
	if err != nil {
		return
	}

	result, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(result)
}


func (s Server) postPointCurrentTasks(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return
	}

	id, err := strconv.Atoi(string(body))
	if err != nil {
		log.Println(err)
		return
	}

	data, err := s.DB.GetPointCurrentTasks(id)
	if err != nil {
		return
	}

	resutl, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	response.Header().Set("Content-Type", "applicaton/json")
	response.WriteHeader(http.StatusOK)
	response.Write(resutl)
}