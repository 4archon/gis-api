package server

import (
	"fmt"
	"map/authentication"
	"map/database"
	"net/http"
	"github.com/gorilla/mux"
)

type Server struct {
	Host string
	Port string
	GisApi string
	DB database.DB
	Auth authentication.Auth
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

	router.HandleFunc("/", s.rootPage).Methods("GET")
	router.HandleFunc("/main", s.getMain).Methods("GET")
	router.HandleFunc("/main", s.postMain).Methods("POST")

	router.HandleFunc("/auth", s.authentication)
	router.HandleFunc("/logout", s.logout)

	router.HandleFunc("/account/login", s.getAccountLogin)
	router.HandleFunc("/account/role", s.getAccountRole)

	router.HandleFunc("/employees", s.getUsers).Methods("GET")
	router.HandleFunc("/employees", s.postUsers).Methods("POST")
	router.HandleFunc("/create_new_user", s.postCreateNewUser).Methods("POST")
	router.HandleFunc("/change_user", s.postChangeUser).Methods("POST")
	router.HandleFunc("/change_user_profile", s.postChangeUserProfile).Methods("POST")

	router.HandleFunc("/profile", s.getProfile).Methods("GET")
	router.HandleFunc("/profile", s.postProfile).Methods("POST")

	router.HandleFunc("/history", s.history).Methods("POST")
	router.HandleFunc("/recent_media", s.postPointRecentMedia).Methods("POST")
	router.HandleFunc("/current_tasks", s.postPointCurrentTasks).Methods("POST")

	router.HandleFunc("/distribute_tasks", s.getDistributeTasks).Methods("GET")
	router.HandleFunc("/distribute_tasks", s.postDistributeTasks).Methods("POST")
	router.HandleFunc("/new_task", s.postApplyTaskToPoints).Methods("POST")
	router.HandleFunc("/appoint", s.postAppointUsersToPoints).Methods("POST")

	router.HandleFunc("/report/decline", s.postReportDecline).Methods("POST")
	router.HandleFunc("/report/service", s.postReportService).Methods("POST")
	router.HandleFunc("/report/inspection", s.postReportInspection).Methods("POST")
	router.HandleFunc("/report/media", s.postReportMedia).Methods("POST")
	
	router.HandleFunc("/analytics", s.getAnalytics).Methods("GET")
	router.HandleFunc("/analytics", s.postAnalytics).Methods("POST")


	fmt.Println("Server is running")
	http.ListenAndServe(fmt.Sprintf("%s:%s", s.Host, s.Port), router)
}