package server

import (
	"fmt"
	"html/template"
	"map/authentication"
	"map/database"
	"net/http"
	"map/point"
	"github.com/gorilla/mux"
)


type Server struct {
	Host string
	Port string
	GisApi string
	DB database.DB
	Auth authentication.Auth
}

type dataMain struct {
	GisApiKey	string
	Points		[]point.Point
}

func (s Server) rootPage(response http.ResponseWriter, req *http.Request) {
	http.Redirect(response, req, "main", http.StatusFound)
}

func (s Server) mainPage(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}
	var data dataMain
	data.GisApiKey = s.GisApi
	data.Points = s.DB.GetPoints()
	tmpl, _ := template.ParseFiles("server/templates/main/main.html")
	tmpl.Execute(response, data)
}

func (s Server) blockFileServer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, req *http.Request) {
		_, _, err := s.checkUser(response, req)
		if err != nil {
			return
		}
		next.ServeHTTP(response, req)
	})
}

func (s Server) Run() {
	router := mux.NewRouter()

	fsBootstrap := http.FileServer(http.Dir("server/static/bootstrap"))
	router.PathPrefix("/bootstrap/").Handler(http.StripPrefix("/bootstrap", fsBootstrap))

	fsMedia := http.FileServer(http.Dir("server/static/media"))
	router.PathPrefix("/media/").Handler(http.StripPrefix("/media", fsMedia))

	fs := http.FileServer(http.Dir("server/static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", s.blockFileServer(fs)))

	router.HandleFunc("/", s.rootPage)
	router.HandleFunc("/main", s.mainPage)

	router.HandleFunc("/auth", s.authentication)
	router.HandleFunc("/logout", s.logout)

	router.HandleFunc("/employees", s.getEmployees)
	router.HandleFunc("/new_employee", s.newEmployee)
	router.HandleFunc("/edit_employee/{id}", s.editEmployee)
	router.HandleFunc("/profile", s.profile)

	router.HandleFunc("/analytics", s.analytics)

	router.HandleFunc("/distribute_tasks", s.distribute)
	router.HandleFunc("/assign_tasks", s.assignTasks)

	router.HandleFunc("/tasks", s.tasks)
	router.HandleFunc("/inspection/{reportID}/{id}", s.inspection)
	router.HandleFunc("/service/{reportID}/{id}", s.service)
	router.HandleFunc("/change_point/{reportID}/{id}", s.changePoint)
	router.HandleFunc("/deactivate/{reportID}/{id}", s.deactivate)

	router.HandleFunc("/points", s.getPoints)
	router.HandleFunc("/account/login", s.getAccountLogin)
	router.HandleFunc("/account/role", s.getAccountRole)

	fmt.Println("Server is running")
	http.ListenAndServe(fmt.Sprintf("%s:%s", s.Host, s.Port), router)
}