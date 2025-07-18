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
	GetUserSubgroup(id int) (string, error)
	GetUserTrust(id int) (bool, error)
	
	GetPointsForAnalytics() ([]business.AnalyticsPoint, error)

	GetDataForMain(id int) ([]business.Point, error)

	GetDataForDistribute() ([]business.DistibutePoint, error)
	NewTaskToPoints(data business.ApplyTask) (error)
	AppointPointsToUsers(data business.Appoint) (error)
	
	GetPointHistory(id int) (business.History, error)
	GetPointMedia(id int) (business.PointMedias, error)
	GetPointCurrentTasks(id int) (business.Tasks, error)

	CreateNewUser(user business.User) (int, error)
	ChangeUser(user business.User) error
	ChangeUserProfile(user business.User) error
}
