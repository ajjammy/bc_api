package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type Project struct {
	Id int64
	Code string
	Name string
}

func(p *Project)GetAll(db *sqlx.DB) (pjs []Project, err error) {
	lcCommand := "select roworder as  id,code,name from bcnp.dbo.bcproject"

	err = db.Select(&pjs,lcCommand)
	if err != nil{
		return nil,err
	}
	return pjs,nil
}



func(p *Project)GetById(db *sqlx.DB,id string) (pj Project,err error) {
	lcCommand := "select roworder as  id,code,name from bcnp.dbo.bcproject where roworder = "+id
	//pj := Project{}
	fmt.Println(lcCommand)
	err = db.Get(&pj,lcCommand)
	if err != nil{
		return  pj,err
	}
	return pj,nil
}