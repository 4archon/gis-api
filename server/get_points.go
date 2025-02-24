package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

func parseBody(body string) []int {
	res := []int{}
	last := 0
	for index, i := range body {
		if i == ',' {
			id, err := strconv.Atoi(body[last:index])
			if err == nil {
				res = append(res, id)
			}
			last = index + 1
		}
	}
	l := len(body)
	id, err := strconv.Atoi(body[last:l])
	if err == nil {
		res = append(res, id)
	}
	return res
}

func (s Server) getPoints(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err.Error())
	}

	pointsID := parseBody(string(body))
	desc := s.DB.GetPointsDesc(pointsID)
	descJson, err := json.Marshal(desc)
	if err != nil {
		log.Println(err.Error())
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(descJson)
}