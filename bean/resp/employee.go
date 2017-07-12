package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type Employee struct {
	Roworder int64 `json:"id"`
	Code string `json:"sale_code" db:"code"`
	Name string `json:"sale_name" db:"name"`
	Telephone string `json:"sale_telephone" db:"telephone"`
	Profitcenter string `json:"profit_center" db:"profitcenter"`
	ImageProfile string `json:"img_profile" db:"imageprofile"`
}

func(e *Employee)GetByKeyWord(keyword string,db *sqlx.DB)(ee []Employee,err error){
	lcCommand := "select top 5 roworder, isnull(code,'') as code,isnull(name,'') as name" +
		" ,isnull(telephone,'') as telephone,'S01' as profitcenter,'' as imageprofile" +
		" from bcnp.dbo.bcsale where (code like '%"+keyword+"%' or name like '%"+keyword+"%' or telephone like '%"+keyword+"%') " +
		" and name not like '%ลาออก%'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	//cc = []Customer{}
	err = db.Select(&ee,lcCommand)

	if err !=nil{
		return nil,err
	}

	fmt.Println(ee)

	return ee,nil
}

func(e *Employee)GetEmployeeCode(keyword string,db *sqlx.DB)(ee []Employee,err error){
	lcCommand := "select top 5 roworder, isnull(code,'') as code,isnull(name,'') as name" +
		" ,isnull(telephone,'') as telephone,'S01' as profitcenter,'' as imageprofile" +
		" from bcnp.dbo.bcsale where code = '"+keyword+"'"
	fmt.Println(lcCommand)
	err = db.Select(&ee,lcCommand)

	if err !=nil{
		return nil,err
	}
	fmt.Println(ee)
	return ee,nil
}