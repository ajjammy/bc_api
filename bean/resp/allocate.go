package Resp

import (
	"github.com/jmoiron/sqlx"

)

type Allocate struct {
	Id int64
	Code string
	Name string
}

func(a *Allocate)GetAll(db *sqlx.DB) (as []Allocate, err error) {
	lcCommand := "select roworder as id,code,name from dbo.bcallocate"

	err = db.Select(&as,lcCommand)
	if err != nil{
		return nil,err
	}
	return as,nil
}