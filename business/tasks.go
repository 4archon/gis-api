package business

type Task struct {
	ID			int		`json:"id"`
	Type		string	`json:"type"`
	Comment		*string	`json:"comment"`
}

type Work struct {
	ID			int				`json:"id"`
	Type		*string			`json:"type"`
	Work		*string			`json:"work"`
	Arc			*int			`json:"arc"`
}

type Tasks struct {
	PointID		int		`json:"id"`
	Tasks		[]Task	`json:"tasks"`
	Works		[]Work	`json:"works"`
}
