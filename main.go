package main

import (
	"map/server"
	"map/config"
	"map/database"
	"map/authentication"
	"log"
)


func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile);

	var conf config.Config
	conf.Init()

	var pdb database.PostgresDB
	pdb.Init(conf)
	defer pdb.Close()
	var db database.DB = &pdb

	var jwt authentication.JwtToken
	jwt.Init(conf.JwtSecretKey)
	var auth authentication.Auth = &jwt

	var serv server.Server;
	serv.Host = "127.0.0.1"
	serv.Port = "56001"
	serv.GisApi = conf.GisApi
	serv.DB = db
	serv.Auth = auth
	serv.AllDataSecretKey = conf.AllDataSecretKey

	serv.Run()

}