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