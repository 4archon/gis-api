package server

import (
	"html/template"
	"log"
	"map/point"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)


func (s Server) tasks(response http.ResponseWriter, req *http.Request) {
	id, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if role == "admin" {
		data, err := s.DB.GetTasksInfo()
		if err != nil {
			return
		}
		tmpl, _ := template.ParseFiles("server/templates/tasks/tasks_admin.html")
		tmpl.Execute(response, data)
	} else {
		data, err := s.DB.GetUserTasksInfo(id)
		if err != nil {
			return
		}
		tmpl, _ := template.ParseFiles("server/templates/tasks/tasks_worker.html")
		tmpl.Execute(response, data)
	}
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

type serviceSaved struct {
	ServiceInfo		[]point.ServiceReport
	ReportID		int
	ServiceID		int
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
			serviceID, err := strconv.Atoi(serviceIDStr)
			if err != nil {
				log.Println(err)
				return
			}
			serviceInfo, err := s.DB.GetService(serviceID)
			if err != nil {
				return
			}
			var serviceSaved serviceSaved
			serviceSaved.ReportID = reportID
			serviceSaved.ServiceInfo = serviceInfo
			serviceSaved.ServiceID = serviceSaved.ServiceInfo[0].ID
			tmpl, _ := template.ParseFiles("server/templates/tasks/service.html")
			tmpl.Execute(response, serviceSaved)
		}
	} else if req.Method == "POST" {
		if serviceIDStr == "new" {
			err = s.createNewServiceReport(reportID, req)
			if err != nil {
				return
			}
			http.Redirect(response, req, "/tasks", http.StatusFound)
		} else {
			_, err := strconv.Atoi(serviceIDStr)
			if err != nil {
				log.Println(err)
				return
			}

			err = s.DB.DeleteService(reportID)
			if err != nil {
				return
			}
			http.Redirect(response, req, "/tasks", http.StatusFound)
		}
	} else {
		return
	}
}

type changePoint struct {
	ReportID			int
	ChangePointID		int
	ChangePointInfo		point.ChangeReport
}

func (s Server) newChangePoint(reportID int, req *http.Request) error {
	var newChange point.ChangeReport
	newChange.Long = req.FormValue("long")
	newChange.Lat = req.FormValue("lat")
	newChange.PointAddress = req.FormValue("address")
	newChange.District = req.FormValue("district")
	numArc, err := strconv.Atoi(req.FormValue("number-arc"))
	if err != nil {
		log.Println(err)
		return err
	}
	newChange.NumberArc = numArc
	newChange.ArcType = req.FormValue("arc-type")
	newChange.Carpet = req.FormValue("carpet")
	newChange.Comment = req.FormValue("comment")
	newChange.ChangeDate = time.Now()
	err = s.DB.NewChangePoint(reportID, newChange)
	if err != nil {
		return err
	}
	return nil
}


func (s Server) changePoint(response http.ResponseWriter, req *http.Request) {
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
	changeIDStr := vars["id"]

	if req.Method == "GET" {
		if changeIDStr == "new" {
			change, err := s.DB.GetPointFromReport(reportID)
			if err != nil {
				return
			}
			var changePoint changePoint
			changePoint.ChangePointInfo = change
			changePoint.ReportID = reportID
			tmpl, _ := template.ParseFiles("server/templates/tasks/change_point_new.html")
			tmpl.Execute(response, changePoint)
		} else {
			changeID, err := strconv.Atoi(changeIDStr)
			if err != nil {
				log.Println(err)
				return
			}
			change, err := s.DB.GetPointInfo(changeID)
			if err != nil {
				return
			}
			var changePoint changePoint
			changePoint.ChangePointInfo = change
			changePoint.ReportID = reportID
			changePoint.ChangePointID = changeID
			tmpl, _ := template.ParseFiles("server/templates/tasks/change_point.html")
			tmpl.Execute(response, changePoint)
		}
	} else if req.Method == "POST" {
		if changeIDStr == "new" {
			err = s.newChangePoint(reportID, req)
			if err != nil {
				return
			}
			http.Redirect(response, req, "/tasks", http.StatusFound)
		} else {
			_, err := strconv.Atoi(changeIDStr)
			if err != nil {
				log.Println(err)
				return
			}

			err = s.DB.DeleteChangePoint(reportID)
			if err != nil {
				return
			}
			http.Redirect(response, req, "/tasks", http.StatusFound)
		}
	} else {
		return
	}
}

type deactivationReport struct {
	Active			point.ActiveReport
	ReportID		int
	DeactivateID	int
}

func (s Server) deactivate(response http.ResponseWriter, req *http.Request) {
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
	deactivateIDStr := vars["id"]

	if req.Method == "GET" {
		if deactivateIDStr == "new" {
			active, err := s.DB.GetActiveFromReport(reportID)
			if err != nil {
				return
			}
			var activeReport deactivationReport
			activeReport.Active = active
			activeReport.ReportID = reportID
			tmpl, _ := template.ParseFiles("server/templates/tasks/deactivate_new.html")
			tmpl.Execute(response, activeReport)
		} else {
			deactivateID, err := strconv.Atoi(deactivateIDStr)
			if err != nil {
				log.Println(err)
				return
			}
			active, err := s.DB.GetActiveStatus(deactivateID)
			if err != nil {
				return
			}
			var activeReport deactivationReport
			activeReport.Active = active
			activeReport.ReportID = reportID
			tmpl, _ := template.ParseFiles("server/templates/tasks/deactivate.html")
			tmpl.Execute(response, activeReport)
		}
	} else if req.Method == "POST" {
		if deactivateIDStr == "new" {
			status := req.FormValue("status")
			comment := req.FormValue("comment")
			err = s.DB.NewDeactivate(reportID, status, comment)
			if err != nil {
				return
			}
			http.Redirect(response, req, "/tasks", http.StatusFound)
		} else {
			_, err := strconv.Atoi(deactivateIDStr)
			if err != nil {
				log.Println(err)
				return
			}

			err = s.DB.DeleteDeactivation(reportID)
			if err != nil {
				return
			}
			http.Redirect(response, req, "/tasks", http.StatusFound)
		}
	} else {
		return
	}
}


func (s Server) deleteAllReport(response http.ResponseWriter, req *http.Request) {
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

	err = s.DB.DeleteService(reportID)
	if err != nil {
		return
	}
	err = s.DB.DeleteInspection(reportID)
	if err != nil {
		return
	}
	err = s.DB.DeleteChangePoint(reportID)
	if err != nil {
		return
	}
	err = s.DB.DeleteDeactivation(reportID)
	if err != nil {
		return
	}
	http.Redirect(response, req, "/tasks", http.StatusFound)
}


func (s Server) sendReport(response http.ResponseWriter, req *http.Request) {
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
	err = s.DB.SendReport(reportID)
	if err != nil {
		return
	}
	http.Redirect(response, req, "/tasks", http.StatusFound)
}

func (s Server) declineReport(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if role != "admin" {
		http.Redirect(response, req, "/tasks", http.StatusFound)
		return
	}

	vars := mux.Vars(req)
	reportIDStr := vars["reportID"]
	reportID, err := strconv.Atoi(reportIDStr)
	if err != nil {
		log.Println(err)
		return
	}
	err = s.DB.DeclineReport(reportID)
	if err != nil {
		return
	}
	http.Redirect(response, req, "/tasks", http.StatusFound)
}

func (s Server) verifyReport(response http.ResponseWriter, req *http.Request) {
	_, role, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	if role != "admin" {
		http.Redirect(response, req, "/tasks", http.StatusFound)
		return
	}

	vars := mux.Vars(req)
	reportIDStr := vars["reportID"]
	reportID, err := strconv.Atoi(reportIDStr)
	if err != nil {
		log.Println(err)
		return
	}
	err = s.DB.VerifyReport(reportID)
	if err != nil {
		return
	}
	http.Redirect(response, req, "/tasks", http.StatusFound)
}