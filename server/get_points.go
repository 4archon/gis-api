package server

import(
	"net/http"
	"log"
	"encoding/json"
	"io"
)

func parseBody(body string) []string {
	res := []string{}
	last := 0
	for index, i := range body {
		if i == ',' {
			res = append(res, body[last:index])
			last = index + 1
		}
	}
	l := len(body)
	res = append(res, body[last:l])
	return res
}

func (s Server) getPoints(response http.ResponseWriter, req *http.Request) {
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