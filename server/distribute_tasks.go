package server

import (
	"encoding/json"
	"log"
	"map/business"
	"net/http"
	"io"
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
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if role != "admin" {
		http.Redirect(response, req, "/main", http.StatusFound)
	}

	var data business.Distibute
	data.GisKey = s.GisApi
	data.Points, err = s.DB.GetDataForDistribute()
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

func (s Server) postApplyTaskToPoints(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if role != "admin" {
		http.Redirect(response, req, "/main", http.StatusFound)
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err);
		return
	}

	var task business.ApplyTask
	err = json.Unmarshal(body, &task)
	if err != nil {
		log.Println(err);
		return
	}

	err = s.DB.NewTaskToPoints(task)
	if err != nil {
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (s Server) postAppointUsersToPoints(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if role != "admin" {
		http.Redirect(response, req, "/main", http.StatusFound)
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err);
		return
	}

	var appoint business.Appoint
	err = json.Unmarshal(body, &appoint)
	if err != nil {
		log.Println(err);
		return
	}

	err = s.DB.AppointPointsToUsers(appoint)
	if err != nil {
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (s Server) postPointEdit(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if role != "admin" {
		http.Redirect(response, req, "/main", http.StatusFound)
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err);
		return
	}

	var changePoint business.ChangePoint
	err = json.Unmarshal(body, &changePoint)
	if err != nil {
		log.Println(err);
		return
	}

	err = s.DB.ChangePoint(changePoint)
	if err != nil {
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (s Server) postDeletePointTask(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if role != "admin" {
		http.Redirect(response, req, "/main", http.StatusFound)
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err);
		return
	}

	var task business.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		log.Println(err);
		return
	}

	err = s.DB.DeletePointTask(task)
	if err != nil {
		return
	}

	response.WriteHeader(http.StatusOK)
}