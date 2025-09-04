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

func (s Server) postReport(response http.ResponseWriter, req *http.Request) {
	id, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	defer req.Body.Close()
	reportJson := req.FormValue("report")
	reportType := req.FormValue("reportType")
	count, err := strconv.Atoi(req.FormValue("count"))
	if err != nil {
		log.Println(err)
		return
	}

	var serviceID int
	switch reportType {
	case "decline":
		var report business.DeclineReport
		err = json.Unmarshal([]byte(reportJson), &report)
		if err != nil {
			log.Println(err);
			return
		}

		serviceID, err = s.DB.NewDeclineReport(id, report)
		if err != nil {
			return
		}
	case "service":
		var report business.ServiceReport
		err = json.Unmarshal([]byte(reportJson), &report)
		if err != nil {
			log.Println(err);
			return
		}

		serviceID, err = s.DB.NewServiceReport(id, report)
		if err != nil {
			return
		}
	case "inspection":
		var report business.InspectionReport
		err = json.Unmarshal([]byte(reportJson), &report)
		if err != nil {
			log.Println(err);
			return
		}

		serviceID, err = s.DB.NewInspectionReport(id, report)
		if err != nil {
			return
		}
	}

	err = s.saveMedia(req, serviceID, count)
	if err != nil {
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (s Server) saveMedia(req *http.Request, serviceID int, count int) (error) {
	var media business.Media
	defer req.Body.Close()
	media.ServiceID = serviceID

	for i := 0; i < count; i++ {
		var err error
		media.MediaName = req.FormValue(fmt.Sprintf("name%d", i))
		media.MediaType = req.FormValue(fmt.Sprintf("type%d", i))
		media.ID, err = s.DB.NewMedia(media)
		if err != nil {
			return err
		}
		err = s.saveFile(i, strconv.Itoa(media.ID), media.MediaType, req)
		if err != nil {
			return err
		}
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