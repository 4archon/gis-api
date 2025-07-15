package business

import(
	"time"
)

type Point struct {
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
	Appoint		[]int			`json:"appoint"`
	Deadline	*time.Time		`json:"deadline"`
}

type Main struct {
	Points		[]Point				`json:"points"`
	GisKey		string				`json:"gisKey"`
	Subgroup	string				`json:"subgroup"`
	Trust		bool				`json:"trust"`
}