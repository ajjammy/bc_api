package Resp

import (
	_ "github.com/jmoiron/sqlx"
)
type Saleorder struct {
	Docno string
	Docdate string
	Arcode string
	SumOfItemAmount float32
	DiscountAmount float32
	BetoreTaxAmount float32
	TaxAmount float32
	TotalAmount float32

}

func(s *Saleorder)GetByDocno(keyword string){
	s.Docno = keyword
	//return s
}