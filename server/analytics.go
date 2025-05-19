package server

import (
	"encoding/json"
	"log"
	"map/business"
	"net/http"
)


func (s Server) getAnalytics(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if role != "admin" {
		http.Redirect(response, req, "/main", http.StatusFound)
	}

	http.ServeFile(response, req, "server/static/analytics/analytics.html")
}

func (s Server) postAnalytics(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if role != "admin" {
		http.Redirect(response, req, "/main", http.StatusFound)
	}

	var analytics business.Analytics
	analytics.GisKey = s.GisApi
	analytics.Points, err = s.DB.GetPointsForAnalytics()
	if err != nil {
		return
	}

	result, err := json.Marshal(analytics)
	if err != nil {
		log.Println(err)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(result)
}