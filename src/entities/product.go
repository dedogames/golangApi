package entities

import "strconv"

type ProductBody struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *ProductBody) ToString() string {
	return " Id: " + strconv.Itoa(p.Id) +
		" Name: " + p.Name +
		" Price: " + strconv.FormatFloat(p.Price, 'f', 2, 64)
}
