package database

import (
	"map/business"
)

type DB interface {
	GetUserLogin(int) string
	ChangeUserPassword(id int, password string) error

	GetAuth(string, string) (int, string, error)
	CheckActiveAuth(int, string) bool

	GetUsersInfo() (business.UsersInfo, error)
	GetUserInfo(id int) (business.UserInfo, error)
	
	GetPointsForAnalytics() ([]business.AnalyticsPoint, error)

	GetDataForMain(id int) ([]business.MainPoint, error)
	
	GetPointHistory(id int) (business.History, error)
	GetPointMedia(id int) (business.PointMedias, error)
	GetPointCurrentTasks(id int) (business.Tasks, error)
}
