package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)
type Saleorder struct {
	Docno string
	Arcode string
	SumOfItemAmount float32
	DiscountAmount float32
	BeforeTaxAmount float32
	TaxAmount float32
	TotalAmount float32
	Items  []*Saleordersub
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
