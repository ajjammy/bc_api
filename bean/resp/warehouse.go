package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type Warehouse struct {
	Code string `json:"warehouse_code" db:"code"`
	Name string `json:"warehouse_name" db:"name"`
}

func(c *Warehouse)GetByKeyWord(keyword string,db *sqlx.DB)(ww []Warehouse,err error){
	lcCommand := "select code,name" +
		" from dbo.BCWarehouse where code like '%"+keyword+"%' or name like '%"+keyword+"%'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	//cc = []Customer{}
	err = db.Select(&ww,lcCommand)
	if err !=nil{
		return nil,err
	}
	fmt.Println(ww)

	return ww,nil
}

func(c *Warehouse)GetWarehouseCode(keyword string,db *sqlx.DB)(ww []Warehouse,err error){
	lcCommand := "select code,name" +
		" from dbo.BCWarehouse where code = '"+keyword+"'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	//cc = []Customer{}
	err = db.Select(&ww,lcCommand)

	if err !=nil{
		return nil,err
	}
	fmt.Println(ww)
	return ww,nil
}