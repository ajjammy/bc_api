package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type Customer struct {
	Code string //`json:"ar_code"`
	Name1 string `json:"ar_name"`
	BillAddress string `json:"address"`
	Telephone string `json:"ar_telephone"`
	DebtLimit1 float32 `json:"limit_amount"`
	DebtLimitBal float32 `json:"balance_amount"`
	DebtAmount float32 `json:"debt_amount"`
	SumOfMark1 float32 `json:"point"`
	ImageProfile string `json:"img_profile"`
}

func(c *Customer)GetByKeyWord(keyword string,db *sqlx.DB)(cc []Customer,err error){
	lcCommand := "select top 10 Code,Name1,BillAddress,Telephone,DebtLimit1,DebtLimitBal,DebtAmount,SumOfMark1,'' as ImageProfile" +
		" from bcnp.dbo.bcar where code like '%"+keyword+"%' or name1 like '%"+keyword+"%'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	cc = []Customer{}
	err = db.Select(&cc,lcCommand)
	if err !=nil{
		return nil,err
	}
	fmt.Println(c)
	return cc,nil
}