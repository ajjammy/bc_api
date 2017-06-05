package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type Shelfcode struct {
	Code string `json:"shelf_code" db:"code"`
	Name string `json:"shelf_name" db:"name"`
	WhCode string `json:"wh_code" db:"whcode"`
}

func(c *Shelfcode)GetByKeyWord(keyword string,db *sqlx.DB)(ss []Shelfcode,err error){
	lcCommand := "select code,name,whcode" +
		" from bcnp.dbo.BCShelf where code like '%"+keyword+"%' or name like '%"+keyword+"%'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	//cc = []Customer{}
	err = db.Select(&ss,lcCommand)
	if err !=nil{
		return nil,err
	}
	fmt.Println(ss)

	return ss,nil
}

func(c *Shelfcode)GetShelfcode(keyword string,db *sqlx.DB)(ss []Shelfcode,err error){
	lcCommand := "select code,name,whcode" +
		" from bcnp.dbo.BCShelf where code = '"+keyword+"'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	//cc = []Customer{}
	err = db.Select(&ss,lcCommand)

	if err !=nil{
		return nil,err
	}
	fmt.Println(ss)
	return ss,nil
}
