package business

import(
	"time"
)

type MainPoint struct {
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
	Coordinates	[]*string		`json:"coordinates"`
	Appointed	bool			`json:"appointed"`
}

type Main struct {
	Points		[]AnalyticsPoint	`json:"points"`
	GisKey		string				`json:"gisKey"`
}