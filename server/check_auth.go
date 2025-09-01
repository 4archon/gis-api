package server

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"time"
)

type badLoginPass struct {
	Login string
	Password string
}

func (s Server) authentication(response http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		_, _, err := s.checkAuth(req)
		if err == nil {
			http.Redirect(response, req, "main", http.StatusTemporaryRedirect)
			return
		}
		http.ServeFile(response, req, "server/templates/auth/auth.html")
		return
	} else if req.Method == "POST" {
		err := req.ParseForm()
		if err != nil {
			log.Println(err.Error())
		}
		email := req.FormValue("email")
		password := req.FormValue("password")
		id, role, err := s.DB.GetAuth(email, password)
		if err != nil {
			var data badLoginPass
			data.Login = email
			data.Password = password
			tmpl, _ := template.ParseFiles("server/templates/auth/bad_auth.html")
			tmpl.Execute(response, data)
			return
		}
		token, err := s.Auth.GetToken(id, role)
		if err != nil {
			log.Println(err)
			http.ServeFile(response, req, "server/templates/auth/auth.html")
			return
		}
		cookie := http.Cookie{Name: "AuthToken", Value: token, Expires: time.Now().Add(30 * 24 * time.Hour)}
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

	if !s.DB.CheckActiveAuth(id, role) {
		return 0, "", errors.New("inactive account")
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
	http.Redirect(response, req, "auth", http.StatusTemporaryRedirect)
}