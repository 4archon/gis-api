package business

type UserInfo struct {
	ID			int				`json:"id"`
	Login		*string			`json:"login"`
	Role		*string			`json:"role"`
	Active		*bool			`json:"active"`
	Name		*string			`json:"name"`
	Surname		*string			`json:"surname"`
	Patronymic	*string			`json:"patronymic"`
	TgID		*int			`json:"tgID"`
	Subgroup	*string			`json:"subgroup"`
	Trust		*bool			`json:"trust"`
}

type UsersInfo struct {
	Info		[]UserInfo		`json:"users"`
}

type User struct {
	ID			int				`json:"id"`
	Login		*string			`json:"login"`
	Password	*string			`json:"password"`
	Role		*string			`json:"role"`
	Active		*bool			`json:"active"`
	Name		*string			`json:"name"`
	Surname		*string			`json:"surname"`
	Patronymic	*string			`json:"patronymic"`
	TgID		*int			`json:"tgID"`
	Subgroup	*string			`json:"subgroup"`
	Trust		*bool			`json:"trust"`
}