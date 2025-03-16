package server

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	// "map/point"
)


func saveFile(req *http.Request, nameDir string, formName string, name string, id int) error {
	photo, _, err := req.FormFile(formName)
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
	err := saveFile(req, "inspection", "photo_left", "photo_left", logDataID)
	if err != nil {
		return err
	}
	err = saveFile(req, "inspection", "photo_right", "photo_right", logDataID)
	if err != nil {
		return err
	}
	err = saveFile(req, "inspection", "photo_front", "photo_front", logDataID)
	if err != nil {
		return err
	}
	err = saveFile(req, "inspection", "video", "video", logDataID)
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


func saveFileService(req *http.Request, logDataID int, i int, extra bool) error {
	index := strconv.Itoa(i)
	err := saveFile(req, "service", "photo_before" + index, "photo_before", logDataID)
	if err != nil {
		return err
	}
	err = saveFile(req, "service", "photo_left" + index, "photo_left", logDataID)
	if err != nil {
		return err
	}
	err = saveFile(req, "service", "photo_right" + index, "photo_right", logDataID)
	if err != nil {
		return err
	}
	err = saveFile(req, "service", "photo_front" + index, "photo_front", logDataID)
	if err != nil {
		return err
	}
	err = saveFile(req, "service", "video" + index, "video", logDataID)
	if err != nil {
		return err
	}
	if extra {
		err = saveFile(req, "service", "photo_extra" + index, "photo_extra", logDataID)
		if err != nil {
			return err
	}
	}
	return nil
}

func (s Server) createNewServiceReport(reportID int, req *http.Request) error {
	serviceCounter := req.FormValue("serviceCounter")
	counter, err := strconv.Atoi(serviceCounter)
	if err != nil {
		log.Println(err)
		return err
	}

	serviceLogID, err := s.DB.CreateService(reportID)
	if err != nil {
		return err
	}

	for i := 1; i <= counter; i++ {
		serviceType := req.FormValue("inputServiceType" + strconv.Itoa(i))
		subtype := req.FormValue("inputSubtype" + strconv.Itoa(i))
		comment := req.FormValue("inputComment" + strconv.Itoa(i))
		serviceLogDataID, err := s.DB.CreateServiceLogData(serviceLogID, serviceType, subtype, comment)
		if err != nil {
			return err
		}
		_, _, err = req.FormFile("photo_extra" + strconv.Itoa(i))
		extra := err == nil
		err = saveFileService(req, serviceLogDataID, i, extra)
		if err != nil {
			return err
		}
		err = s.DB.ApproveServiceReport(reportID, serviceLogID, serviceLogDataID, extra)
		if err != nil {
			return err
		}
	}
	return nil
}