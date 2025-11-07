package business

import(
	"time"
)

type PointAppoints struct {
	Appoints			[]PointAppoint		`json:"appoints"`
}

type PointAppoint struct {
	ID					int					`json:"id"`
	Users				[]PointAppointUser	`json:"users"`
	AppointmentDate		time.Time			`json:"date"`
}

type PointAppointUser struct {
	ID					int					`json:"id"`
	Login				*string				`json:"login"`
	Subgroup			*string				`json:"subgroup"`
	Trust				*bool				`json:"trust"`
}