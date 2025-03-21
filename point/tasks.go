package point

import (
	"time"
)

type TasksRequest struct {
	Points		[]int	`json:"tasks"`
	Employees	[]int	`json:"workers"`
	Deadline	string	`json:"deadline"`
}

type AppliedUser struct {
	ID		int
	Login	string
}

type Task struct {
	TaskID			int
	Users			[]int
	UsersApplied	[]AppliedUser
	ChangeID		int
	ActiveID		int
	ServiceID		int
	InspectionID	int
	PointID			int
	Address			string
	Lat				string
	Long			string
	District		string
	NumberArc		int
	TypeArc			string
	Carpet			string
	Deadline		time.Time
	SentWorker		bool
	Verified		int
}

type InspectionReport struct {
	ID				int
	Checkup			string
	RepairType		string
	PhotoLeft		string
	PhotoRight		string
	PhotoFront		string
	Video			string
	Comment			string
}

type ServiceReport struct {
	ID				int
	ServiceType		string
	Subtype			string
	ActionArc		int
	PhotoBefore		string
	PhotoLeft		string
	PhotoRight		string
	PhotoFront		string
	PhotoExtra		string
	Video			string
	Comment			string
	Index			int
}

type ChangeReport struct {
	Long            string
    Lat             string
    PointAddress	string
    District        string
    NumberArc		int
    ArcType			string
    Carpet          string
    ChangeDate		time.Time
    Comment         string
}

type ActiveReport struct {
	Status			bool
	Comment			string
	ChangeDate		time.Time
}