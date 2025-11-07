package server

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"map/business"
	"mime/multipart"
	"net/http"
	"strconv"
)


func (s Server) getDistributeTasks(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if role != "admin" {
		http.Redirect(response, req, "/main", http.StatusFound)
	}

	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
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

func (s Server) postNewPoints(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if role != "admin" {
		return
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err);
		return
	}

	var newPoints business.NewPoints
	err = json.Unmarshal(body, &newPoints)
	if err != nil {
		log.Println(err);
		return
	}

	err = s.DB.NewPoints(newPoints)
	if err != nil {
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (s Server) postNewPointsByFile(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if role != "admin" {
		return
	}

	defer req.Body.Close()
	file, _, err := req.FormFile("file")
	if err != nil {
		log.Println(err)
		return
	}

	defer file.Close()
	newPoints, err := newPointsFromCsv(file)
	if err != nil {
		log.Println(err)
		return
	}
	
	err = s.DB.NewPoints(newPoints)
	if err != nil {
		return
	}

	response.WriteHeader(http.StatusOK)
}

func newPointsFromCsv(file multipart.File) (business.NewPoints, error) {
	var result business.NewPoints
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return result, err
	}
	records = records[1:]

	for _, record := range(records) {
		var newPoint business.NewPoint
		lat, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			return result, err
		}
		newPoint.Lat = &lat
		long, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return result, err
		}
		newPoint.Long = &long
		newPoint.Address = &record[2]
		newPoint.District = &record[3]
		newPoint.ExternalID = &record[4]
		newPoint.Carpet = &record[5]
		numArc, err := strconv.Atoi(record[6])
		if err != nil {
			return result, err
		}
		newPoint.NumberArc = &numArc
		newPoint.ArcType = &record[7]
		newPoint.Owner = &record[8]
		newPoint.Operator = &record[9]
		newPoint.Customer = &record[10]
		newPoint.Comment = &record[11]

		result.NewPoints = append(result.NewPoints, newPoint)
	}

	return result, nil
}


func (s Server) postDeletePointAppoint(response http.ResponseWriter, req *http.Request) {
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

	appoint, err := strconv.Atoi(string(body))
	if err != nil {
		log.Println(err);
		return
	}

	err = s.DB.DeletePointAppoint(appoint)
	if err != nil {
		return
	}

	response.WriteHeader(http.StatusOK)
}