package database

import (
	"map/point"
)

type DB interface {
	GetPoints() []point.Point
	GetPointsDesc([]string) []point.PointDesc
}
