package main

import (
	"fmt"
	"map/server"
	"map/config"
	"map/database"
)


func main() {
	var conf config.Config
	conf.Init()
	fmt.Println(conf.GisApi)

	fileName := "baza.csv"
	var db database.DB = &database.CsvDB{}
	db.Init(fileName)

	var serv server.Server;
	serv.Host = "127.0.0.1"
	serv.Port = "56001"
	serv.Conf = conf
	serv.DB = db

	serv.Run()

}