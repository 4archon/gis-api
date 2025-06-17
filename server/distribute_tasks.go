package server

import (
	"encoding/json"
	"log"
	"map/business"
	"net/http"
)


func (s Server) getDistributeTasks(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if role != "admin" {
		http.Redirect(response, req, "/main", http.StatusFound)
	}

	http.ServeFile(response, req, "server/static/distibute_tasks/distribute.html")
}

func (s Server) postDistributeTasks(response http.ResponseWriter, req *http.Request) {
	id, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	var data business.Main
	data.GisKey = s.GisApi
	data.Points, err = s.DB.GetDataForMain(id)
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