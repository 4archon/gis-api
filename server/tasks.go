package server

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"map/point"

	"github.com/gorilla/mux"
)


func (s Server) tasks(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	data, err := s.DB.GetTasksInfo()
	if err != nil {
		return
	}

	tmpl, _ := template.ParseFiles("server/templates/tasks/tasks.html")
	tmpl.Execute(response, data)
}

func saveFile(req *http.Request, nameDir string, name string, id int) error {
	photo, _, err := req.FormFile(name)
	if err != nil {
		log.Println(err)
		return err
	}
	defer photo.Close()

	dest := "server/static/media/" + nameDir + "/" + strconv.Itoa(id) + "/"
	err = os.MkdirAll(dest, 0775)
	if err != nil && !os.IsExist(err) {
		log.Println(err)
		return err
	}
	if name != "video" {
		dest += name + ".jpeg"
	} else {
		dest += name + ".mov"
	}
	file, err := os.Create(dest)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, photo)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func saveFileInspection(req *http.Request, logDataID int) error {
	err := saveFile(req, "inspection", "photo_left", logDataID)
	if err != nil {
		return err
	}
	err = saveFile(req, "inspection", "photo_right", logDataID)
	if err != nil {
		return err
	}
	err = saveFile(req, "inspection", "photo_front", logDataID)
	if err != nil {
		return err
	}
	err = saveFile(req, "inspection", "video", logDataID)
	if err != nil {
		return err
	}
	return nil
}

func (s Server) createNewInspectionReport(reportID int, req *http.Request) error {
	checkup := req.FormValue("inputCheckup")
	repairType := req.FormValue("inputRepairType")
	comment := req.FormValue("inputComment")
	logInsID, err := s.DB.CreateInspection(reportID, checkup, repairType, comment)
	if err != nil {
		return err
	}
	err = saveFileInspection(req, logInsID)
	if err != nil {
		return err
	}
	return nil
}

type inspectionSaved struct {
	InspectionInfo	point.InspectionReport
	ReportID		int
}

func (s Server) inspection(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	vars := mux.Vars(req)
	reportIDStr := vars["reportID"]
	reportID, err := strconv.Atoi(reportIDStr)
	if err != nil {
		log.Println(err)
		return
	}
	inspectionIDStr := vars["id"]

	if req.Method == "GET" {
		if inspectionIDStr == "new" {
			tmpl, _ := template.ParseFiles("server/templates/tasks/inspection_new.html")
			tmpl.Execute(response, reportIDStr)
		} else {
			inspectionID, err := strconv.Atoi(inspectionIDStr)
			if err != nil {
				log.Println(err)
				return
			}
			inspectionInfo, err := s.DB.GetInspection(inspectionID)
			if err != nil {
				return
			}
			var insSaved inspectionSaved
			insSaved.ReportID = reportID
			insSaved.InspectionInfo = inspectionInfo
			tmpl, _ := template.ParseFiles("server/templates/tasks/inspection.html")
			tmpl.Execute(response, insSaved)
		}
	} else if req.Method == "POST" {
		if inspectionIDStr == "new" {
			err := s.createNewInspectionReport(reportID, req)
			if err != nil {
				return
			}
			http.Redirect(response, req, "/tasks", http.StatusFound)
		} else {
			_, err := strconv.Atoi(inspectionIDStr)
			if err != nil {
				log.Println(err)
				return
			}
			err = s.DB.DeleteInspection(reportID)
			if err != nil {
				return
			}
			http.Redirect(response, req, "/tasks", http.StatusFound)
		}
	} else {
		return
	}
}


func (s Server) service(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	vars := mux.Vars(req)
	reportIDStr := vars["reportID"]
	reportID, err := strconv.Atoi(reportIDStr)
	if err != nil {
		log.Println(err)
		return
	}
	serviceIDStr := vars["id"]

	if req.Method == "GET" {
		if serviceIDStr == "new" {
			tmpl, _ := template.ParseFiles("server/templates/tasks/service_new.html")
			tmpl.Execute(response, reportIDStr)
		} else {
			inspectionID, err := strconv.Atoi(serviceIDStr)
			if err != nil {
				log.Println(err)
				return
			}
			inspectionInfo, err := s.DB.GetInspection(inspectionID)
			if err != nil {
				return
			}
			var insSaved inspectionSaved
			insSaved.ReportID = reportID
			insSaved.InspectionInfo = inspectionInfo
			tmpl, _ := template.ParseFiles("server/templates/tasks/inspection.html")
			tmpl.Execute(response, insSaved)
		}
	} else if req.Method == "POST" {
		if serviceIDStr == "new" {
			val := req.FormValue("serviceCounter")
			println(val)
			http.Redirect(response, req, "/tasks", http.StatusFound)
		} else {
			_, err := strconv.Atoi(serviceIDStr)
			if err != nil {
				log.Println(err)
				return
			}

			http.Redirect(response, req, "/tasks", http.StatusFound)
		}
	} else {
		return
	}
}