package server

import (
	"fmt"
	"html/template"
	"log"
	"map/authentication"
	"map/database"
	"net/http"
)


type Server struct {
	Host string
	Port string
	GisApi string
	DB database.DB
	Auth authentication.Auth
}

func (s Server) rootPage(response http.ResponseWriter, req *http.Request) {
	http.Redirect(response, req, "main", http.StatusTemporaryRedirect)
}

func (s Server) mainPage(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	var data dataMain
	data.GisApiKey = s.GisApi
	data.Points = s.DB.GetPoints()
	tmpl, _ := template.ParseFiles("server/templates/main.html")
	tmpl.Execute(response, data)
}

func (s Server) blockFileServer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, req *http.Request){
		cookie, err := req.Cookie("AuthToken")
		if err != nil {
			log.Println(err.Error())
		}
		jwtString := cookie.Value
		id, role, err := s.Auth.GetPayload(jwtString)
		if err != nil {
			log.Println(err.Error())
			http.Redirect(response, req, "auth", http.StatusTemporaryRedirect)
		}
		fmt.Println(id)
		fmt.Println(role)
		next.ServeHTTP(response, req)
	})
}

func (s Server) Run() {
	fs := http.FileServer(http.Dir("server/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", s.rootPage)
	http.HandleFunc("/main", s.mainPage)

	http.HandleFunc("/auth", s.authentication)
	http.HandleFunc("/logout", s.logout)
	http.HandleFunc("/points", s.getPoints)

	fmt.Println("Server is running")
	http.ListenAndServe(fmt.Sprintf("%s:%s", s.Host, s.Port), nil)
}