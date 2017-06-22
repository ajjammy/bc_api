package Resp

import "github.com/jmoiron/sqlx"

type Project struct {
	Code string
	Name string
}

func(p *Project)GetAll(db *sqlx.DB) (pjs []Project, err error) {
	lccommand := "select code,name from bcnp.dbo.bcproject"

	err = db.Select(&pjs,lccommand)
	if err != nil{
		return nil,err
	}
	return pjs,nil
}