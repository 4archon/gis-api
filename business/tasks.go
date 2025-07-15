package business

import (
	"time"
)

type Task struct {
	ID			int				`json:"id"`
	Type		string			`json:"type"`
	Comment		*string			`json:"comment"`
	Deadline	*time.Time		`json:"deadline"`
	Customer	*string			`json:"customer"`
	EntryDate	*time.Time		`json:"entryDate"`
	Done		*bool			`json:"done"`
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
