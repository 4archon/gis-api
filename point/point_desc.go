package point

import (
	"fmt"
	"strconv"
)

type PointDesc struct {
	ID 		int		`json:"id"`
	Address string	`json:"address"`
	Date 	string	`json:"date"`
	Img 	string	`json:"img"`
	Amount 	int		`json:"amount"`
}

func (p *PointDesc) Init(id string, address string, date string, amount string, img string) {
	idNum, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err.Error())
	}
	p.ID = idNum
	p.Address = address
	p.Date = date
	amountNum, err := strconv.Atoi(amount)
	if err != nil {
		fmt.Println(err.Error())
	}
	p.Amount = amountNum
	p.Img = img
}