package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"time"
)

type Saleorder struct {
	//Docno string`json:"doc_no"`
	//Arcode string `json:"ar_code"`
	//SumOfItemAmount float32 `json:"sum_of_item_amount"`
	//DiscountAmount float32 `json:"discount_amount"`
	//BeforeTaxAmount float32 `json:"before_tax_amount"`
	//TaxAmount float32 `json:"tax_amount"`
	//TotalAmount float32 `json:"total_amount"`
	Docno string `json:"doc_no"`
	Docdate string `json:"docdate"`
	Taxtype int `json:"taxtype"`
	Billtype int `json:"billtype"`
	Arcode string `json:"arcode"`
	Departcode string `json:"departcode"`
	Creditday int `json:"creditday"`
	Duedate string `json:"duedate"`
	Salecode string `json:"salecode"`
	Taxrate float32 `json:"taxrate"`
	Isconfirm int `json:"isconfirm"`
	Mydescription string `json:"mydescription"`
	Billstatus int `json:"billstatus"`
	Sostatus int `json:"so_status"`
	Holdingstatus int `json:"holding_status"`
	Sumofitemamount float32 `json:"sumofitemamount"`
	Discountword string `json:"discountword"`
	Discountamount float32 `json:"discount_amount" DB:"discountamount" `
	Afterdiscount float32 `json:"afterdiscount"`
	Beforetaxamount float32 `json:"beforetaxamount"`
	Taxamount float32 `json:"taxamount"`
	Totalamount float32 `json:"totalamount"`
	Netamount float32 `json:"netamount"`
	Iscancel int `json:"iscancel"`
	Creatorcode string `json:"creatorcode"`
	Createdatetime string `json:"createdatetime"`
	Lasteditorcode string `json:"lasteditorcode"`
	Lasteditdatet string  `json:"lasteditdatet"`// must to convert to datetime type in sql server
	Confirmcode string `json:"confirmcode"`
	Confirmdatetime string `json:"confirmdatetime"`
	Cancelcode string `json:"cancelcode"`
	Canceldatetime string `json:"canceldatetime"`
	Isconditionsend int `json:"isconditionsend"`
	Deliveryday int `json:"delivery_day"`
	Deliverydate time.Time `json:"delivery_date"`
	//items []*Sosub

	Items  []*Saleordersub `json:"item"`
}

type Saleordersub struct {
	Id int64  `json:"id"`//roworder
	Linenumber int64  `json:"line_number"`
	Itemcode string `json:"item_code"`
	Itemname string `json:"item_name"`
	Qty float32`json:"qty"`
	Unitcode string `json:"unit_code"`
	Price  float32 `json:"price"`
	Amount float32 `json:"amounts"`
	Netamount float32 `json:"net_amount"`
	Packingrate1 float32 `json:"packing_rate_1"`
	Packingrate2 float32 `json:"packing_rate_2"`
}

func(s *Saleorder)GetByDocno(docno string,db *sqlx.DB)(err error){

	lcCommand := "select docno,docdate,billtype,departcode," +
		"duedate,salecode,arcode,sumofitemamount,discountamount,beforetaxamount,taxamount,totalamount," +
		"taxrate,isconfirm,billstatus,creatorcode,createdatetime,lasteditorcode,lasteditdatet,cancelcode,canceldatetime " +
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
	lcCommand := "select top 10 docno,arcode,sumofitemamount,discountamount,beforetaxamount,taxamount,totalamount" +
		" from bcnp.dbo.bcsaleorder where docno like '%"+keyword+"%' "
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
