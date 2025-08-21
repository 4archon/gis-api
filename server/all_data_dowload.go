package server

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
)


func (s Server) allDataDownload(response http.ResponseWriter, req *http.Request) {
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

	data, err := s.DB.GetAllData()
	if err != nil {
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	gz.Write(result)
	gz.Close()

	response.Header().Set("Content-Type", "applicaton/json")
	response.WriteHeader(http.StatusOK)
	response.Write(buf.Bytes())
}