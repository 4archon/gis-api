package server

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s Server) pointProfile(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	vars := mux.Vars(req)
	pointIDStr := vars["id"]
	pointID, err := strconv.Atoi(pointIDStr)
	if err != nil {
		log.Println(err)
		return
	}
	profile, err := s.DB.GetPointProfile(pointID)
	if err != nil {
		return
	}

	tmpl, _ := template.ParseFiles("server/templates/profile/profile.html")
	tmpl.Execute(response, profile)
}

func (s Server) pointStory(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	vars := mux.Vars(req)
	pointIDStr := vars["id"]
	pointID, err := strconv.Atoi(pointIDStr)
	if err != nil {
		log.Println(err)
		return
	}

	story, err := s.DB.GetPointStory(pointID)
	if err != nil {
		return
	}

	tmpl, _ := template.ParseFiles("server/templates/story/story.html")
	tmpl.Execute(response, story)
}