package server

import (
	"log"
	"net/http"
)

func (s Server) authentication(response http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		_, _, err := s.checkAuth(req)
		if err == nil {
			http.Redirect(response, req, "main", http.StatusTemporaryRedirect)
			return
		}
		http.ServeFile(response, req, "server/templates/auth.html")
		return
	} else if req.Method == "POST" {
		err := req.ParseForm()
		if err != nil {
			log.Println(err.Error())
		}
		// email := req.FormValue("email")
		// password := req.FormValue("password")
		id := 123321
		role := "pidka"
		token, err := s.Auth.GetToken(id, role)
		if err != nil {
			log.Println(err)
		}
		cookie := http.Cookie{Name: "AuthToken", Value: token}
		http.SetCookie(response, &cookie)
		http.Redirect(response, req, "main", http.StatusFound)
		return
	} else {
		log.Println("Wrong method")
		return
	}
}

func (s Server) checkAuth(req *http.Request) (int, string, error) {
	cookie, err := req.Cookie("AuthToken")
	if err != nil {
		return 0, "", err
	}
	jwtString := cookie.Value
	id, role, err := s.Auth.GetPayload(jwtString)
	if err != nil {
		return 0, "", err
	}
	return id, role, nil
}

func (s Server) checkUser(response http.ResponseWriter, req *http.Request) (int, string, error) {
	id, role, err := s.checkAuth(req)
	if err != nil {
		http.Redirect(response, req, "auth", http.StatusTemporaryRedirect)
		return 0, "", err
	}
	return id, role, err
}

func (s Server) logout(response http.ResponseWriter, req *http.Request) {
	cookie := http.Cookie{Name: "AuthToken", Value: "null"}
	http.SetCookie(response, &cookie)
	http.Redirect(response, req, "main", http.StatusTemporaryRedirect)
}