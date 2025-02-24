package database

import (
	"map/point"
)

type DB interface {
	GetPoints() []point.Point
	GetPointsDesc([]int) []point.PointDesc
	GetAuth(string, string) (int, string, error)
	CheckActiveAuth(int, string) bool
	GetUserLogin(int) string
	GetUsersInfo() []point.User
	GetUserInfo(int) (point.User, error)
}
