package Resp

import "github.com/jmoiron/sqlx"

type Allocate struct {
	Code string
	Name string
}

func(a *Allocate)GetAll(db *sqlx.DB) (as []Allocate, err error) {
	lccommand := "select code,name from bcnp.dbo.bcallocate"

	err = db.Select(&as,lccommand)
	if err != nil{
		return nil,err
	}
	return as,nil
}