package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"map/server"
	"map/config"
	"map/point"
	"regexp"
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

func getCoordinates(csvString string) point.Point {
	var p point.Point
	
	re := regexp.MustCompile(`\d+\.\d+`)
	res := re.FindAllString(csvString, -1)
	p.Init(res[1], res[0])
	return p
}

func main() {
	var conf config.Config
	conf.Init()
	fmt.Println(conf.GisApi)

	fileName := "map_dug.csv"
	res := readCSV(fileName)
	res = res[1:]

	var points []point.Point = make([]point.Point, 0)
	for _, i := range res {
		points = append(points, getCoordinates(i[1]))
	}

	var serv server.Server;
	serv.Host = "127.0.0.1"
	serv.Port = "56001"
	serv.Conf = conf
	serv.Points = points

	serv.Run()

}