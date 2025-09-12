package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func (s Server) postGSheetBase(response http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return
	}

	userSecret := string(body)
	if s.AllDataSecretKey != userSecret {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := s.DB.GetGSheetBase()
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

func (s Server) postGSheetDoneWorks(response http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return
	}

	type oprionsJSON struct {
		Start	time.Time	`json:"start"`
		End		time.Time	`json:"end"`
		Secret	string		`json:"secret"`
	}

	var options oprionsJSON
	err = json.Unmarshal(body, &options)
	if err != nil {
		log.Println(err);
		return
	}

	if s.AllDataSecretKey != options.Secret {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := s.DB.GetGSheetDoneWorks(options.Start, options.End)
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