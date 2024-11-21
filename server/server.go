package server

import (
	"fmt"
	"net/http"
	"html/template"
	"map/config"
	"map/point"
)


type Server struct {
	Host string
	Port string
	Conf config.Config
	Points []point.Point
}

func (s Server) test(response http.ResponseWriter, req *http.Request) {
	var mes string = "my fisrt step";
	tmpl, _ := template.ParseFiles("server/templates/test.html")
	tmpl.Execute(response, mes)
}

func (s Server) mainPage(response http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles("server/templates/main.html")
	tmpl.Execute(response, s)
}

func (s Server) Run() {
	fs := http.FileServer(http.Dir("server/static"))
	http.Handle("/", fs)
	http.HandleFunc("/test", s.test)
	http.HandleFunc("/main", s.mainPage)

	fmt.Println("Server is running")
	http.ListenAndServe(s.Host + ":" + s.Port, nil)
}