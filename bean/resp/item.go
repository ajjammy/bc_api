package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	//"github.com/mrtomyum/paybox_web/model"
	//"strconv"
)

type Item struct {
	Roworder int64 `json:"id"`
	Code string `json:"item_code"`
	Name string `json:"item_name" db:"name1"`
	UnitCode string `json:"unit_code" db:"defstkunitcode"`
	StockQty float32 `json:"stock_qty"`
	Stocks []*Stock `json:"stock_list"`
	Units []*Unit `json:"units"`
	SoQty float32 `json:"so_qty" db:"remainoutqty"`
	RoQty float32 `json:"ro_qty" db:"reserveqty"`
	PoQty	float32 `json:"po_qty" db:"remaininqty"`
	MyGrade string `json:"my_grade" db:"mygrade"`
	Comm float32 `json:"comm"`
	ImageProfile string `json:"img_profile" db:"picfilename1"`

	//“sale_qty”:”12”,
	//“reserve_qty”:”3”,
	//“po_qty”:”30”,
	//“my_grade”:”B”,
	//“comm”:”1” //ณ เวลาที่ค้นหา,
	//“img_profile”:”59110”

}

type Stock struct {
	Qty float32 `json:"qty" db:"qty"`
	UnitCode string `json:"unit_code" db:"unitcode"`
	WhCode string `json:"wh_code" db:"whcode"`
	ShelfCode string `json:"shelf_code"`
}

type Unit struct {
	UnitId int `json:"unit_id" db:"id"`
	UnitCode string `json:"unit_code" db:"unitcode"`
	UnitName string `json:"unit_name" db:"unitname"`
	PackingRate1 float32 `json:"packing_rate" db:"packing_rate"`
	Price float32 `json:"price" db:"price"`
	//“unit_id”:”0”,
	//”unit_name”:”ชิ้น”,
	//“packing_rate”:”1”,
	//“price”:”100”

}

func(i *Item)GetByCode(itemcode string,db *sqlx.DB)(err error){

	lcCommand := "select top 10  roworder, code,name1,defstkunitcode," +
		"isnull(stockqty,0) as stockqty," +
		"isnull(remainoutqty,0) as remainoutqty," +
		"isnull(reserveqty,0) as reserveqty," +
		"isnull(remaininqty,0) as remaininqty," +
		"isnull(mygrade,'-') as mygrade," +
		"isnull(picfilename1,'') as picfilename1 " +
		"from bcitem where code = '"+itemcode+"'"
	//fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	err = db.Get(i,lcCommand)
	fmt.Println(itemcode)
	sqlsub := `select qty,unitcode,whcode,shelfcode from dbo.bcstklocation where qty > 0 and shelfcode = '-'  and whcode not like '%TRN%'  and whcode not like '%ISP%' and  itemcode=?`
	fmt.Println(sqlsub)
	err = db.Select(&i.Stocks,sqlsub,i.Code)
	if len(i.Stocks) == 0 {
		stk := Stock{0,"","",""}
		//stks := []Stock{}
		i.Stocks = append(i.Stocks,&stk)
		fmt.Println("Null stock is ",stk)

	}


	if err !=nil{
	return err
	}
	fmt.Println(i)

	unitsub := `select  a.roworder as id,a.unitcode,b.name as unitname ,
		isnull(c.rate1,1) as packing_rate,
		saleprice1 as price
		from dbo.bcpricelist a
			left join dbo.bcitemunit b on a.unitcode=b.code
		 	left join dbo.bcstkpacking c on a.unitcode=c.unitcode and a.itemcode=c.itemcode
		 	where a.itemcode=?  and saletype = 0 and transporttype=0`
	//fmt.Println(unitsub)
	err = db.Select(&i.Units,unitsub,i.Code)

	if len(i.Units) ==0 {
		u := Unit{0,"","",0,0}
		i.Units = append(i.Units,&u)
	}

	return nil
}

func(i *Item)GetByKeyword(keyword string,db *sqlx.DB)(items []Item,err error){

	lcCommand := "select top 10 roworder , code,name1,defstkunitcode," +
		"isnull(stockqty,0) as stockqty," +
		"isnull(remainoutqty,0) as remainoutqty," +
		"isnull(reserveqty,0) as reserveqty," +
		"isnull(remaininqty,0) as remaininqty," +
		"isnull(mygrade,'-') as mygrade," +
		"isnull(picfilename1,'') as picfilename1 " +
		"from dbo.bcitem where code like '%"+keyword+"%' or name1 like '%"+keyword+"%' or code in " +
		"(select itemcode  from dbo.bcbarcodemaster where barcode like '%"+keyword+"%') "
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	err = db.Select(&items,lcCommand)
	if err !=nil{
		return nil,err
	}
	fmt.Println("start loop for range items")
	for idx, item := range items {
		//todo: add child node
			fmt.Println("item stock loop")
			sqlsub := "select qty,unitcode,whcode,shelfcode " +
				"from dbo.bcstklocation " +
				"where qty>0 and shelfcode ='-' and whcode not like '%ISP%'  and whcode not like '%TRN%' and  itemcode='"+item.Code+"'"
			fmt.Println(item.Code)
			err = db.Select(&items[idx].Stocks,sqlsub)

			if len(items[idx].Stocks) == 0 {
				stk := Stock{0,"","",""}
				//stks := []Stock{}
				items[idx].Stocks = append(items[idx].Stocks,&stk)
				fmt.Println("Null stock is ",stk)

			}
			fmt.Println(sqlsub)

		// bind : unit & Price
		unitsub := "select  a.roworder as id,a.unitcode,b.name as unitname ," +
			"isnull(c.rate1,1) as packing_rate," +
			"saleprice1 as price "+
			"from .dbo.bcpricelist a left join dbo.bcitemunit b on a.unitcode=b.code "+
		 	"left join dbo.bcstkpacking c on a.unitcode=c.unitcode and a.itemcode=c.itemcode"+
		 	" where a.itemcode='"+item.Code+"'  and saletype = 0 and transporttype=0"
		fmt.Println(unitsub)

			err = db.Select(&items[idx].Units,unitsub)
		if len(items[idx].Units) ==0 {
			u := Unit{0,"","",0,0}
			items[idx].Units = append(items[idx].Units,&u)
		}

	}

	fmt.Println(i)
	return items,nil
}



