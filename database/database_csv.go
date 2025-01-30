package database

import (
	"os"
	"fmt"
	"encoding/csv"
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

	return records[1:]
}

type CsvDB struct {
	filename string
	data [][]string
}

func (c *CsvDB) Init(filename string) {
	c.filename = filename
	c.data = readCSV(filename)
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

func (c CsvDB) GetPoints() []point.Point {
	var points []point.Point = make([]point.Point, 0)
	for _, i := range c.data {
		points = append(points, getPoint(i[0], i[2]))
	}
	return points
}

func contain(id string, pointsID []string) bool {
	for _, i := range pointsID {
		if id == i {
			return true
		}
	}
	return false
}

func (c CsvDB) GetPointsDesc(pointsID []string) []point.PointDesc {
	var desc []point.PointDesc = make([]point.PointDesc, 0)
	for _, i := range c.data {
		if contain(i[0], pointsID) {
			var pointDesc point.PointDesc
			pointDesc.Init(i[0], i[3], i[4], i[6], i[7])
			desc = append(desc, pointDesc)
		}
	}
	return desc
}

