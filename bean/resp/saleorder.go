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
	Docno string
	Arcode string
	SumOfItemAmount float32
	DiscountAmount float32
	BeforeTaxAmount float32
	TaxAmount float32
	TotalAmount float32

}

func(s *Saleorder)GetByDocno(docno string,db *sqlx.DB)(ss []Saleorder,err error){

	lcCommand := "select docno,arcode,sumofitemamount,discountamount,beforetaxamount,taxamount,totalamount" +
		" from bcnp.dbo.bcsaleorder where docno = '"+docno+"'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	ss = []Saleorder{}
	err = db.Select(&ss,lcCommand)
	if err !=nil{
		return nil,err
	}
	fmt.Println(ss)
	return ss,nil
}

func(s *Saleorder)GetByKeyWord(docno string,db *sqlx.DB)(ss []Saleorder,err error){
	lcCommand := "select docno,arcode,sumofitemamount,discountamount,beforetaxamount,taxamount,totalamount" +
		" from bcnp.dbo.bcsaleorder where docno like '%"+docno+"%'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	//ss = []Saleorder{}
	err = db.Select(&ss,lcCommand)
	if err !=nil{
		return nil,err
	}
	fmt.Println(s)
	return ss,nil
}
