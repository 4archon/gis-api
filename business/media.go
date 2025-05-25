package business

type Media struct {
	ID			int		`json:"id"`
	MediaType	string	`json:"type"`
	MediaName	string	`json:"name"`
}

type PointMedias struct {
	ID			int			`json:"id"`
	Medias		[]Media		`json:"medias"`
}