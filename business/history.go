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

type AllServices struct {
	CurrentPage	int					`json:"currentPage"`
	LastPage	int					`json:"lastPage"`
	Services	[]ServiceStoryPoint	`json:"services"`
}

type ServiceStoryPoint struct {
	ID				int			`json:"id"`
	Point			PurePoint	`json:"point"`
	Users			[]UserInfo	`json:"users"`
	Appoint			*time.Time	`json:"appoint"`
	Execution		*time.Time	`json:"execution"`
	Comment			*string		`json:"comment"`
	Status			*string		`json:"status"`
	Sent			*bool		`json:"sent"`
	SentBy			*UserInfo	`json:"sentBy"`
	WithoutTasks	*bool		`json:"withoutTask"`
	Works			[]Work		`json:"works"`
	Tasks			[]Task		`json:"tasks"`
	Medias			[]Media		`json:"medias"`
}

type PurePoint struct {
	ID			int				`json:"id"`
	Active		*bool			`json:"active"`
	Long		*string			`json:"long"`
	Lat			*string			`json:"lat"`
	Address		*string			`json:"address"`
	District	*string			`json:"district"`
	NumberArc	*int			`json:"numberArc"`
	ArcType		*string			`json:"arcType"`
	Carpet		*string			`json:"carpet"`
	ChangeDate	*time.Time		`json:"changeDate"`
	Comment		*string			`json:"comment"`
	Status		*string			`json:"status"`
	Owner		*string			`json:"owner"`
	Operator	*string			`json:"operator"`
	ExternalID	*string			`json:"externalID"`
}