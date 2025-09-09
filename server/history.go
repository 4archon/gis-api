package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)


func (s Server) history(response http.ResponseWriter, req *http.Request) {
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

	data, err := s.DB.GetPointHistory(id)
	if err != nil {
		return
	}


	result, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	response.Header().Set("Content-Type", "applicaton/json")
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

	result, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	response.Header().Set("Content-Type", "applicaton/json")
	response.WriteHeader(http.StatusOK)
	response.Write(result)
}
