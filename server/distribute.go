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
	Points		[]point.Point
	Employees	[]point.User
}


func (s Server) distribute(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	var data dataDistribute
	data.GisApiKey = s.GisApi
	data.Points = s.DB.GetPoints()
	data.Employees = s.DB.GetUsersInfo()
	tmpl, _ := template.ParseFiles("server/templates/distribute_tasks/distribute.html")
	tmpl.Execute(response, data)
}

func (s Server) assignTasks(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
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