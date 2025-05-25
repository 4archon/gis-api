package business

import(
	"time"
)


type StoryPoint struct {
	UserIDs		[]string		`json:"userIDs"`
	UserLogins	[]string		`json:"userLogins"`
	Deadline	*time.Time		`json:"deadline"`
	Execution	*time.Time		`json:"execution"`
	Comment		*string			`json:"comment"`
	Status		*string			`json:"status"`
	Sent		*bool			`json:"sent"`
	Works		[]Work			`json:"works"`
	Tasks		[]Task			`json:"tasks"`
}

type History struct {
	ID			int				`json:"id"`
	StoryPoints	[]StoryPoint	`json:"storyPoints"`
}