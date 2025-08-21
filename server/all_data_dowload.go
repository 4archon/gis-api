package server

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	vars := mux.Vars(req)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	maxID := len(allData) / pieceSize

	if id < 0 || id > maxID {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var result []byte

	if id == maxID {
		result = allData[id * pieceSize:]
	} else {
		result = allData[id * pieceSize:(id + 1) * pieceSize]
	}

	response.Header().Set("Content-Type", "applicaton/octet-stream")
	response.WriteHeader(http.StatusOK)
	response.Write(result)
}

var allData = []byte{}
var pieceSize = 1 * 1024 * 1024

type dataInfo struct {
	Urls	[]int	`json:"urls"`
}

func (s Server) allDataInfo(response http.ResponseWriter, req *http.Request) {
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

	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	gz.Write(dataJson)
	gz.Close()

	allData = buf.Bytes()

	var info dataInfo
	for i := 0; i <= len(allData) / pieceSize; i++ {
		info.Urls = append(info.Urls, i)
	}

	result, err := json.Marshal(info)
	if err != nil {
		log.Println(err)
		return
	}

	response.Header().Set("Content-Type", "applicaton/json")
	response.WriteHeader(http.StatusOK)
	response.Write(result)
}