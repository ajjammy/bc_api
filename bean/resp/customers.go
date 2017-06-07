package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type Customer struct {
	Code string `json:"ar_code" db:"code"`
	Name1 string `json:"ar_name" db:"name1"`
	BillAddress string `json:"address" db:"billaddress"`
	Telephone string `json:"ar_telephone" db:"telephone"`
	DebtLimit1 float32 `json:"limit_amount" db:"debtlimit1"`
	DebtLimitBal float32 `json:"balance_amount" db:"debtlimitbal"`
	DebtAmount float32 `json:"debt_amount" db:"debtamount"`
	SumOfMark1 float32 `json:"point"  db:"sumofmark1"`
	ImageProfile string `json:"img_profile" db:"impageprofile"`
}

func(c *Customer)GetByKeyWord(keyword string,db *sqlx.DB)(cc []Customer,err error){
	lcCommand := "select top 10 isnull(code,'') as code,isnull(name1,'') as name1" +
		" ,isnull(billaddress,'') as billaddress,isnull(telephone,'') as telephone" +
		" ,isnull(debtlimit1,0) as debtlimit1,isnull(debtlimitbal,0) as debtlimitbal" +
		" ,isnull(debtamount,0) as debtamount,isnull(sumofmark1,0) as sumofmark1,'' as impageprofile" +
		" from bcnp.dbo.bcar where isnull(code,'') like '%"+keyword+"%'" +
		" or isnull(name1,'') like '%"+keyword+"%' or isnull(telephone,'') like '%"+keyword+"%'" +
		" or isnull(billaddress,'') like '%"+keyword+"%' order by code"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	//cc = []Customer{}
	err = db.Select(&cc,lcCommand)

	if err !=nil{
		return nil,err
	}

	fmt.Println(cc)

	return cc,nil
}

func(c *Customer)GetCustomerCode(keyword string,db *sqlx.DB)(cc []Customer,err error){
	lcCommand := "select top 10 isnull(code,'') as code,isnull(name1,'') as name1" +
		" ,isnull(billaddress,'') as billaddress,isnull(telephone,'') as telephone" +
		" ,isnull(debtlimit1,0) as debtlimit1,isnull(debtlimitbal,0) as debtlimitbal" +
		" ,isnull(debtamount,0) as debtamount,isnull(sumofmark1,0) as sumofmark1,'' as impageprofile" +
		" from bcnp.dbo.bcar where code = '"+keyword+"' order by code"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	//cc = []Customer{}
	err = db.Select(&cc,lcCommand)

	if err !=nil{
		return nil,err
	}

	fmt.Println(cc)

	return cc,nil
}
