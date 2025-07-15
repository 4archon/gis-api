package business

import(
	"time"
)


type StoryPoint struct {
	ID			int				`json:"id"`
	UserIDs		[]string		`json:"userIDs"`
	UserLogins	[]string		`json:"userLogins"`
	Execution	*time.Time		`json:"execution"`
	Comment		*string			`json:"comment"`
	Status		*string			`json:"status"`
	Sent		*bool			`json:"sent"`
	Works		[]Work			`json:"works"`
	Tasks		[]Task			`json:"tasks"`
	Medias		[]Media			`json:"medias"`
}

type History struct {
	ID			int				`json:"id"`
	StoryPoints	[]StoryPoint	`json:"storyPoints"`
}