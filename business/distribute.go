package business

import(
	"time"
)

type DistibutePoint struct {
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
	Coordinates	[]*string		`json:"coordinates"`
	Tasks		[]Task			`json:"tasks"`
	Works		[]Work			`json:"works"`
	Appoint		[]AppointUser	`json:"appoint"`
	Marks		[]Mark			`json:"marks"`
}

type Distibute struct {
	Points		[]DistibutePoint	`json:"points"`
	GisKey		string				`json:"gisKey"`
}


type ApplyTask struct {
	Task		*string			`json:"task"`
	Customer	*string			`json:"customer"`
	Deadline	*time.Time		`json:"deadline"`
	Comment		*string			`json:"comment"`
	Points		[]int			`json:"points"`
}

type Appoint struct {
	Users		[]int			`json:"users"`
	Points		[]int			`json:"points"`
}

type AppointUser struct {
	ID			int				`json:"id"`
	Subgroup	*string			`json:"subgroup"`
}

type ChangePoint struct {
	ID			int				`json:"id"`
	Long		*float64		`json:"long"`
	Lat			*float64		`json:"lat"`
	Address		*string			`json:"address"`
	District	*string			`json:"district"`
	Active		*bool			`json:"active"`
	Status		*string			`json:"status"`
	ExternalID	*string			`json:"externalID"`
	Carpet		*string			`json:"carpet"`
	NumberArc	*int			`json:"numberArc"`
	ArcType		*string			`json:"arcType"`
	Owner		*string			`json:"owner"`
	Operator	*string			`json:"operator"`
	Comment		*string			`json:"comment"`
	Marks		[]Mark			`json:"marks"`
}