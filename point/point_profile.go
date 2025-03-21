package point

import (
	"time"
)

type PointProfile struct {
	ID					int
	Status				bool
	StatusLastChange	time.Time
	Long				string
	Lat					string
	Address				string
	District			string
	NumberArc			int
	ArcType				string
	Carpet				string
	PointLastChange		time.Time
	Service				[]ServiceReport
	ServiceLast			time.Time
	Inspection			InspectionReport
	InspectionLast		time.Time
}


type StoryPoint struct {
	TaskID			int
	Users			[]int
	UsersApplied	[]AppliedUser
	ChangeID		int
	ActiveID		int
	ServiceID		int
	InspectionID	int
	PointID			int
	Deadline		time.Time
	Appointment		time.Time
	Submission		time.Time
	SentWorker		bool
	Verified		int
}