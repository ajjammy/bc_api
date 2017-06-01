package Resp

import "github.com/jmoiron/sqlx"

type Item struct {
	Code string
	Name string `json:"item_name" db:"name1"`
	Stock []Stock
}

type Stock struct {
	qty int32
	unitcode string
	whcode string `json:"wh_code"`
}

func(i *Item)GetByCode(keyword string,db *sqlx.DB)(items []Item,err error){
//	items := []Item{}
return

}
