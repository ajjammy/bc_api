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
	lcCommand := "select code,name1,billaddress,telephone,debtlimit1,debtlimitbal,debtamount,sumofmark1,'' as impageprofile" +
		" from bcnp.dbo.bcar where code like '%"+keyword+"%' or name1 like '%"+keyword+"%'"
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