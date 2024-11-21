package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"map/server"
	"map/config"
	"map/point"
	"regexp"
	"strconv"
)

func readCSV(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return records
}

func getPoint(ID string, csvString string) point.Point {
	var p point.Point
	
	re := regexp.MustCompile(`\d+\.\d+`)
	res := re.FindAllString(csvString, -1)
	id, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println("not int conv")
	}
	p.Init(id, res[1], res[0])
	return p
}

func main() {
	var conf config.Config
	conf.Init()
	fmt.Println(conf.GisApi)

	fileName := "baza.csv"
	res := readCSV(fileName)
	res = res[1:]

	var points []point.Point = make([]point.Point, 0)
	for _, i := range res {
		points = append(points, getPoint(i[0], i[2]))
	}

	var serv server.Server;
	serv.Host = "127.0.0.1"
	serv.Port = "56001"
	serv.Conf = conf
	serv.Points = points

	serv.Run()

}