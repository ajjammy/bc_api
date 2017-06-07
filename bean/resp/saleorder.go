package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)
type Saleorder struct {
	Docno string`json:"doc_no"`
	Arcode string `json:"ar_code"`
	SumOfItemAmount float32 `json:"sum_of_item_amount"`
	DiscountAmount float32 `json:"discount_amount"`
	BeforeTaxAmount float32 `json:"before_tax_amount"`
	TaxAmount float32 `json:"tax_amount"`
	TotalAmount float32 `json:"total_amount"`
	Items  []*Saleordersub `json:"items"`
}

type Saleordersub struct {
	Id int64  `json:"id"`//roworder
	Linenumber int64  `json:"linenumber"`
	Itemcode string `json:"itemcode"`
	Itemname string `json:"itemname"`
	Qty float32`json:"qty"`
	Unitcode string `json:"unitcode"`
	Price  float32 `json:"price"`
	Amount float32 `json:"amount"`
	Netamount float32 `json:"netamount"`
	Packingrate1 float32 `json:"packingrate_1"`
	Packingrate2 float32 `json:"packingrate_2"`

}

func(s *Saleorder)GetByDocno(docno string,db *sqlx.DB)(err error){

	lcCommand := "select docno,arcode,sumofitemamount,discountamount,beforetaxamount,taxamount,totalamount" +
		" from bcnp.dbo.bcsaleorder where docno = '"+docno+"'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	//ss = []Saleorder{}
	err = db.Get(s,lcCommand)
	if err !=nil{
		return err
	}


	fmt.Println(s)
	// todo: add Node sub details

	sosub := `select  a.roworder as id,a.linenumber,a.itemcode,a.itemname,a.qty
	 		,a.unitcode,a.price,a.amount,a.netamount,a.packingrate1,a.packingrate2
		from bcnp.dbo.bcsaleordersub a
			where a.docno=? `
	fmt.Println(sosub)
	err = db.Select(&s.Items,sosub,docno)
	return err
}

func(s *Saleorder)GetByKeyWord(keyword string,db *sqlx.DB)(ss []Saleorder,err error){
	lcCommand := "select docno,arcode,sumofitemamount,discountamount,beforetaxamount,taxamount,totalamount" +
		" from bcnp.dbo.bcsaleorder where docno like '%"+keyword+"%'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	//ss = []Saleorder{}
	err = db.Select(&ss,lcCommand)
	if err !=nil{
		return nil,err
	}
	// todo : add child node in for loop
	for i,so := range ss{
		sosub := `select  a.roworder as id,a.linenumber,a.itemcode,a.itemname,a.qty
	 		,a.unitcode,a.price,a.amount,a.netamount,a.packingrate1,a.packingrate2
		from bcnp.dbo.bcsaleordersub a
			where a.docno=? `
		fmt.Println(sosub)
		err = db.Select(&ss[i].Items,sosub,so.Docno)

	}

	//fmt.Println(ss)
	return ss,nil
}
