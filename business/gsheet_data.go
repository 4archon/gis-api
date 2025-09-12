package business

import (
	"time"
)


type GSheetBase struct {
	Points			[]AllDataPoint		`json:"points"`
}

type GSheetDoneWorks struct {
	Works			[]GSheetWork		`json:"works"`
}

type GSheetWork struct {
	Login				*string			`json:"login"`
	PointID				int				`json:"pointID"`
	Long				*string			`json:"long"`
	Lat					*string			`json:"lat"`
	Address				*string			`json:"address"`
	Owner				*string			`json:"owner"`
	ServiceID			int				`json:"serviceID"`
	UserID				[]string		`json:"userID"`
	ExecutionDate		*time.Time		`json:"executionDate"`
	SentBy				*int			`json:"sentBy"`
	WithoutTask			*bool			`json:"withoutTask"`
	Work				*string			`json:"work"`
	Arc					*int			`json:"arc"`
}

type GSheetDoneVisits struct {
	Visits			[]GSheetVisit		`json:"visits"`
}

type GSheetVisit struct {
	Login				*string			`json:"login"`
	PointID				int				`json:"pointID"`
	Long				*string			`json:"long"`
	Lat					*string			`json:"lat"`
	Address				*string			`json:"address"`
	Owner				*string			`json:"owner"`
	ServiceID			int				`json:"serviceID"`
	UserID				[]string		`json:"userID"`
	ExecutionDate		*time.Time		`json:"executionDate"`
	SentBy				*int			`json:"sentBy"`
	WithoutTask			*bool			`json:"withoutTask"`
}