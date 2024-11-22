package server

import (
	"fmt"
	"html/template"
	"io"
	"map/config"
	"map/database"
	"net/http"
	"encoding/json"
)


type Server struct {
	Host string
	Port string
	Conf config.Config
	DB database.DB
}

func (s Server) test(response http.ResponseWriter, req *http.Request) {
	var mes string = "my fisrt step";
	tmpl, _ := template.ParseFiles("server/templates/test.html")
	tmpl.Execute(response, mes)
}

func (s Server) mainPage(response http.ResponseWriter, req *http.Request) {
	var data dataMain
	data.GisApiKey = s.Conf.GisApi
	data.Points = s.DB.GetPoints()
	tmpl, _ := template.ParseFiles("server/templates/main.html")
	tmpl.Execute(response, data)
}

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
		fmt.Println(err.Error())
	}

	pointsID := parseBody(string(body))
	desc := s.DB.GetPointsDesc(pointsID)
	descJson, err := json.Marshal(desc)
	if err != nil {
		fmt.Println(err.Error())
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(descJson)
}

func (s Server) phonePage(response http.ResponseWriter, req *http.Request) {
	var data dataMain
	data.GisApiKey = s.Conf.GisApi
	data.Points = s.DB.GetPoints()
	tmpl, _ := template.ParseFiles("server/templates/phone.html")
	tmpl.Execute(response, data)
}

func (s Server) Run() {
	fs := http.FileServer(http.Dir("server/static"))
	http.Handle("/", fs)
	http.HandleFunc("/test", s.test)
	http.HandleFunc("/main", s.mainPage)
	http.HandleFunc("/points", s.getPoints)
	http.HandleFunc("/phone", s.phonePage)

	fmt.Println("Server is running")
	http.ListenAndServe(s.Host + ":" + s.Port, nil)
}