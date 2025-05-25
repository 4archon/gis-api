package server

import (
	"net/http"
)

func (s Server) getAccountLogin(response http.ResponseWriter, req *http.Request) {
	id, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	login := s.DB.GetUserLogin(id)

	response.Header().Set("Content-Type", "text/plain")
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(login))
}

func (s Server) getAccountRole(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	response.Header().Set("Content-Type", "text/plain")
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(role))
}