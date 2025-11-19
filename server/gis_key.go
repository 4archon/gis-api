package server

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type keyForForm struct {
	Key		string
}


func (s Server) getChangeGisKey(response http.ResponseWriter, req *http.Request) {
	id, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	if id != 1 {
		http.Redirect(response, req, "main", http.StatusFound)
		return
	}

	var key keyForForm
	key.Key = *s.GisApi
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	tmpl, err := template.ParseFiles("server/static/change_gis_key/change_gis_key.html")
	if err != nil {
		log.Println(err)
		return
	}
	tmpl.Execute(response, key)
}

func (s Server) postChangeGisKey(response http.ResponseWriter, req *http.Request) {
	id, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	if id != 1 {
		return
	}

	err = req.ParseForm()
	if err != nil {
		return
	}
	key := req.FormValue("key")
	*s.GisApi = key
	err = writeNewKeyToFile(key)
	if err != nil {
		return
	}

	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	http.Redirect(response, req, "main", http.StatusFound)
}

func writeNewKeyToFile(key string) error {
	err := os.WriteFile("config/2gis.key", []byte(key), 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}