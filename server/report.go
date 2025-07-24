package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"map/business"
	"net/http"
	"os"
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

func (s Server) postReportMedia(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}


	fmt.Println(req.FormValue("count"))
	file, header, err := req.FormFile("file0")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	fmt.Println(header.Filename)
	dst, err := os.Create("server/static/media/1.jpeg")
	if err != nil {
		log.Println(err)
		return
	}
	defer dst.Close()
	io.Copy(dst, file)

	response.WriteHeader(http.StatusOK)
}