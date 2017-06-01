package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

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

func(i *Item)GetByCode(itemcode string,db *sqlx.DB)(err error){

	lcCommand := "select code,name1 from bcnp.dbo.bcitem where code = '"+itemcode+"'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	err = db.Get(i,lcCommand)
	if err !=nil{
	return err
	}
	fmt.Println(i)
	return nil
}


