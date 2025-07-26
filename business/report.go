package business


type DeclineReport struct {
	Appoint			[]int			`json:"appoint"`
	PointID			int				`json:"pointID"`
	Reason			string			`json:"reason"`
	Yourself		*bool			`json:"yourself"`
	Comment			*string			`json:"comment"`
	Duplicate		*duplicate		`json:"duplicate"`
	Tasks			[]Task			`json:"tasks"`
}

type duplicate struct {
	Duplicate	int		`json:"duplicate"`
	Original	int		`json:"original"`
}

type ServiceID struct {
	ID			int		`json:"id"`
}

type ServiceReport struct {
	Appoint			[]int			`json:"appoint"`
	PointID			int				`json:"pointID"`
	Tasks			[]Task			`json:"tasks"`
	NewLocation		[]float64		`json:"location"`
	Done			[]serviceWorks	`json:"done"`
	Required		[]serviceWorks	`json:"required"`
	Comment			*string			`json:"comment"`
}

type serviceWorks struct {
	WorkType		string			`json:"type"`
	Count			int				`json:"count"`
	Number			*string			`json:"number"`
	MarksID			[]int			`json:"selectedMarks"`
}


type InspectionReport struct {
	Appoint			[]int					`json:"appoint"`
	PointID			int						`json:"pointID"`
	Tasks			[]Task					`json:"tasks"`
	Required		[]inspectionWorks		`json:"required"`
	PaintCount		*int					`json:"paint"`
	Comment			*string					`json:"comment"`
}


type inspectionWorks struct {
	WorkType		string			`json:"type"`
	Count			int				`json:"count"`
}