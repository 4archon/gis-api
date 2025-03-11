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
	ChangeUserInfo(id int, name string, surname string, patronymic string, tgID string) error
	ChangeUserPassword(id int, password string) error
	NewUser(user point.NewUser) error
	ChangeUserAllInfo(id int, user point.NewUser) error
	
	AssignTasks(data point.TasksRequest) error
	GetTasksInfo() ([]point.Task, error)

	GetPointIDFromReport(id int) (int, error)
	CreateInspection(reportID int, checkup string, repairType string, comment string) (int, error)
	GetInspection(inspectionID int) (point.InspectionReport, error)
	DeleteInspection(reportID int) error
}
