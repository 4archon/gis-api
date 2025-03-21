package server

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"map/point"
	"net/http"
)


type dataDistribute struct {
	GisApiKey	string
	Points		[]point.FilterPoint
	Employees	[]point.User
}


func (s Server) distribute(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if role != "admin" {
		http.Redirect(response, req, "/main", http.StatusFound)
	}

	var data dataDistribute
	data.GisApiKey = s.GisApi
	data.Points, err = s.DB.GetFiltredPoints()
	if err != nil {
		return
	}
	data.Employees = s.DB.GetWorkersInfo()
	tmpl, _ := template.ParseFiles("server/templates/distribute_tasks/distribute.html")
	tmpl.Execute(response, data)
}

func (s Server) assignTasks(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if role != "admin" {
		http.Redirect(response, req, "/main", http.StatusFound)
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var tasks point.TasksRequest
	err = json.Unmarshal(body, &tasks)
	if err != nil {
		log.Println(err)
		return
	}

	err = s.DB.AssignTasks(tasks)
	if err != nil {
		log.Println(err)
		return
	}
	
	response.WriteHeader(http.StatusOK)
}