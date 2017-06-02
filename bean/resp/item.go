package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	//"github.com/mrtomyum/paybox_web/model"
	//"strconv"
)

type Item struct {
	Code string `json:"code"`
	Name string `json:"item_name" db:"name1"`
	Stocks []*Stock `json:"stocks"`
}

type Stock struct {
	Qty float32 `json:"qty" db:"qty"`
	UnitCode string `json:"unit_code" db:"unitcode"`
	WhCode string `json:"wh_code" db:"whcode"`
}

func(i *Item)GetByCode(itemcode string,db *sqlx.DB)(err error){

	lcCommand := "select code,name1 from bcnp.dbo.bcitem where code = '"+itemcode+"'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	err = db.Get(i,lcCommand)
	fmt.Println(itemcode)
	sqlsub := `select qty,unitcode,whcode from bcnp.dbo.bcstkwarehouse where itemcode=?`
	fmt.Println(sqlsub)
	err = db.Select(&i.Stocks,sqlsub,i.Code)

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

	for idx, item := range items {
		//todo: add child node
			fmt.Println("item stock loop")
			sqlsub := "select qty,unitcode,whcode from bcnp.dbo.bcstkwarehouse where itemcode='"+item.Code+"'"
			fmt.Println(item.Code)
			err = db.Select(&items[idx].Stocks,sqlsub)
			fmt.Println(sqlsub)

	}

	fmt.Println(i)
	return items,nil
}



