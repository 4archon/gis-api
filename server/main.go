package server

import (
	"encoding/json"
	"log"
	"map/business"
	"net/http"
)

func (s Server) rootPage(response http.ResponseWriter, req *http.Request) {
	http.Redirect(response, req, "main", http.StatusFound)
}

func (s Server) getMain(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	http.ServeFile(response, req, "server/static/main/main.html")
}

func (s Server) postMain(response http.ResponseWriter, req *http.Request) {
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
