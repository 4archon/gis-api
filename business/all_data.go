package business

import (
	"time"
)


type AllData struct {
	Users		[]User				`json:"users"`
	Points		[]AllDataPoint		`json:"points"`
	PointsLog	[]AllDataLogPoint	`json:"points_log"`
	Marks		[]AllDataMark		`json:"marks"`
	Service		[]AllDataService	`json:"service"`
	Works		[]AllDataWork		`json:"works"`
	Media		[]Media				`json:"media"`
	Tasks		[]AllDataTask		`json:"tasks"`
}

type AllDataPoint struct {
	ID			int				`json:"id"`
	Active		*bool			`json:"active"`
	Long		*string			`json:"long"`
	Lat			*string			`json:"lat"`
	Address		*string			`json:"address"`
	District	*string			`json:"district"`
	NumberArc	*string			`json:"numberArc"`
	ArcType		*string			`json:"arcType"`
	Carpet		*string			`json:"carpet"`
	ChangeDate	*time.Time		`json:"changeDate"`
	Comment		*string			`json:"comment"`
	Status		*string			`json:"status"`
	Owner		*string			`json:"owner"`
	Operator	*string			`json:"operator"`
	ExternalID	*string			`json:"externalID"`
}

type AllDataLogPoint struct {
	ID			int				`json:"id"`
	PointID		int				`json:"pointID"`
	Active		*bool			`json:"active"`
	Long		*string			`json:"long"`
	Lat			*string			`json:"lat"`
	Address		*string			`json:"address"`
	District	*string			`json:"district"`
	NumberArc	*string			`json:"numberArc"`
	ArcType		*string			`json:"arcType"`
	Carpet		*string			`json:"carpet"`
	ChangeDate	*time.Time		`json:"changeDate"`
	Comment		*string			`json:"comment"`
	Status		*string			`json:"status"`
	Owner		*string			`json:"owner"`
	Operator	*string			`json:"operator"`
	ExternalID	*string			`json:"externalID"`
}

type AllDataMark struct {
	ID			int				`json:"id"`
	PointID		int				`json:"pointID"`
	Number		*string			`json:"number"`
	Type		*string			`json:"type"`
	Active		*bool			`json:"active"`
}

type AllDataService struct {
	ID					int				`json:"id"`
	PointID				int				`json:"pointID"`
	UserID				[]string		`json:"userID"`
	AppointmentDate		*time.Time		`json:"appointmentDate"`
	ExecutionDate		*time.Time		`json:"executionDate"`
	Comment				*string			`json:"comment"`
	Status				*string			`json:"status"`
	Sent				*bool			`json:"sent"`
	SentBy				*int			`json:"sentBy"`
	WithoutTask			*bool			`json:"withoutTask"`
}

type AllDataWork struct {
	ID			int				`json:"id"`
	ServiceID	int				`json:"serviceID"`
	Type		*string			`json:"type"`
	Work		*string			`json:"work"`
	Arc			*int			`json:"arc"`
}

type AllDataTask struct {
	ID			int				`json:"id"`
	PointID		int				`json:"pointID"`
	Type		string			`json:"type"`
	Comment		*string			`json:"comment"`
	ServiceID	*int			`json:"serviceID"`
	Customer	*string			`json:"customer"`
	EntryDate	*time.Time		`json:"entryDate"`
	Deadline	*time.Time		`json:"deadline"`
	Done		*bool			`json:"done"`
}