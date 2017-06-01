package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	//"github.com/mrtomyum/paybox_web/model"
	//"strconv"
)

type Item struct {
	Code string
	Name string `json:"item_name" db:"name1"`
	Stock []Stock
}

type Stock struct {
	qty float32
	unitcode string
	whcode string `json:"wh_code"`
}

func(i *Item)GetByCode(itemcode string,db *sqlx.DB)(err error){

	lcCommand := "select code,name1 from bcnp.dbo.bcitem where code = '"+itemcode+"'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	err = db.Get(i,lcCommand)
	fmt.Println(itemcode)
	sqlsub := "select qty,unitcode,whcode from bcnp.dbo.bcstkwarehouse where itemcode='"+i.Code+"'"
	fmt.Println(i.Code)
	err = db.Get(&i.Stock,sqlsub)

	if err !=nil{
	return err
	}
	fmt.Println(i)
	return nil
}

func(i *Item)GetByKeyword(keyword string,db *sqlx.DB)(items []Item,err error){

	lcCommand := "select code,name1 from bcnp.dbo.bcitem where code like '%"+keyword+"%' or name1 like '%"+keyword+"%'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	err = db.Select(&items,lcCommand)
	if err !=nil{
		return nil,err
	}

	for _, item := range items {
		//todo: add child node
			fmt.Println("item stock loop")
			sqlsub := `select qty,unitcode,whcode from bcnp.dbo.bcstkwarehouse where itemcode=?`
			fmt.Println(item.Code)
			err = db.Select(&item.Stock,sqlsub,item.Code)
			fmt.Println(sqlsub)

	}

	fmt.Println(i)
	return items,nil
}



