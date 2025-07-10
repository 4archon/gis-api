package server

import (
	"encoding/json"
	"io"
	"log"
	"map/business"
	"net/http"
	"strconv"
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

func (s Server) postCreateNewUser(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	if role != "admin" {
		http.Redirect(response, req, "main", http.StatusFound)
		return
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err);
		return
	}
	var user business.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err);
		return
	}

	id, err := s.DB.CreateNewUser(user)
	if err != nil {
		return
	}

	response.Header().Set("Content-Type", "text/plain")
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(strconv.Itoa(id)))
}

func (s Server) postChangeUser(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	if role != "admin" {
		http.Redirect(response, req, "main", http.StatusFound)
		return
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err);
		return
	}
	var user business.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err);
		return
	}

	err = s.DB.ChangeUser(user)
	if err != nil {
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (s Server) postChangeUserProfile(response http.ResponseWriter, req *http.Request) {
	id, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err);
		return
	}
	var user business.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err);
		return
	}

	user.ID = id
	err = s.DB.ChangeUserProfile(user)
	if err != nil {
		return
	}

	response.WriteHeader(http.StatusOK)
}