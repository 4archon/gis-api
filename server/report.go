package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"map/business"
	"net/http"
	"os"
	"strconv"
)

func (s Server) postReportDecline(response http.ResponseWriter, req *http.Request) {
	id, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err);
		return
	}

	var report business.DeclineReport
	err = json.Unmarshal(body, &report)
	if err != nil {
		log.Println(err);
		return
	}

	serviceID, err := s.DB.NewDeclineReport(id, report)
	if err != nil {
		return
	}

	var data business.ServiceID
	data.ID = serviceID
	resutl, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	response.Header().Set("Content-Type", "applicaton/json")
	response.WriteHeader(http.StatusOK)
	response.Write(resutl)
}

func (s Server) postReportService(response http.ResponseWriter, req *http.Request) {
	id, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err);
		return
	}

	var report business.ServiceReport
	err = json.Unmarshal(body, &report)
	if err != nil {
		log.Println(err);
		return
	}


	serviceID, err := s.DB.NewServiceReport(id, report)
	if err != nil {
		return
	}

	var data business.ServiceID
	data.ID = serviceID
	resutl, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	response.Header().Set("Content-Type", "applicaton/json")
	response.WriteHeader(http.StatusOK)
	response.Write(resutl)
}

func (s Server) postReportInspection(response http.ResponseWriter, req *http.Request) {
	id, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err);
		return
	}

	var report business.InspectionReport
	err = json.Unmarshal(body, &report)
	if err != nil {
		log.Println(err);
		return
	}

	serviceID, err := s.DB.NewInspectionReport(id, report)
	if err != nil {
		return
	}

	var data business.ServiceID
	data.ID = serviceID
	resutl, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	response.Header().Set("Content-Type", "applicaton/json")
	response.WriteHeader(http.StatusOK)
	response.Write(resutl)
}

func (s Server) postReportMedia(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	err = s.saveMedia(req)
	if err != nil {
		return
	}

	response.WriteHeader(http.StatusOK)
}


func (s Server) saveMedia(req *http.Request) (error) {
	var media business.Media
	defer req.Body.Close()
	count, err := strconv.Atoi(req.FormValue("count"))
	if err != nil {
		log.Println(err)
		return err
	}
	media.ServiceID, err = strconv.Atoi(req.FormValue("id"))
	if err != nil {
		log.Println(err)
		return err
	}

	for i := 0; i < count; i++ {
		media.MediaName = req.FormValue(fmt.Sprintf("name%d", i))
		media.MediaType = req.FormValue(fmt.Sprintf("type%d", i))
		media.ID, err = s.DB.NewMedia(media)
		if err != nil {
			return err
		}
		s.saveFile(i, strconv.Itoa(media.ID), media.MediaType, req)
	}
	return nil
}

func (s Server) saveFile(index int, name string, mediType string, req *http.Request) (error) {
	file, _, err := req.FormFile(fmt.Sprintf("file%d", index))
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	dst, err := os.Create(fmt.Sprintf("server/static/media/%s.%s", name, mediType))
	if err != nil {
		log.Println(err)
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}