package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type Customer struct {
	Code string `json:"ar_code"`
	Name1 string `json:"ar_name"`
	BillAddress string `json:"address"`
	Telephone string `json:"ar_telephone"`
	DebtLimit1 float32 `json:"limit_amount"`
	DebtLimitBal float32 `json:"balance_amount"`
	DebtAmount float32 `json:"debt_amount"`
	SumOfMark1 float32 `json:"point"`
	ImageProfile string `json:"img_profile"`
}

func(s *Customer)GetByKeyWord(keyword string,db *sqlx.DB)(ss []Customer,err error){
	lcCommand := "select Code,Name1,BillAddress,Telephone,DebtLimit1,DebtLimitBal,DebtAmount,SumOfMark1,ImageProfile" +
		" from bcnp.dbo.bcar where code like '%"+keyword+"%'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	ss = []Customer{}
	err = db.Select(&ss,lcCommand)
	if err !=nil{
		return nil,err
	}
	fmt.Println(s)
	return ss,nil
}