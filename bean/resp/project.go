package Resp

import (
	"github.com/jmoiron/sqlx"
//	"fmt"
)

type Project struct {
	Id int64
	Code string
	Name string
}

func(p *Project)GetAll(db *sqlx.DB) (pjs []Project, err error) {
	lcCommand := "select roworder as  id,code,name from bcproject"

	err = db.Select(&pjs,lcCommand)
	if err != nil{
		return nil,err
	}
	return pjs,nil
}



func(p *Project)GetByCode(db *sqlx.DB,code string) (pj Project,err error) {
	lcCommand := "select roworder as  id,code,name from bcproject where code = '"+code+"'"
	//pj := Project{}
	//fmt.Println(lcCommand)
	err = db.Get(&pj,lcCommand)
	if err != nil{
		return  pj,err
	}
	return pj,nil
}