package Resp

import (


	"github.com/jmoiron/sqlx"
	"fmt"
	"testing"
)
var dbx *sqlx.DB


func TestInsertSo(t *testing.T) {

	// Setup

	MockSo := Saleorder{}
	lcDoc := "W01-SCV5905-0016"
	err := MockSo.GetByDocno(lcDoc,dbtest)
	fmt.Println(MockSo)
	// Tear down

	//clear old data before
	lccommand := `delete bcnp.dbo.bcsaleorder where docno = ?`
	_,err = dbtest.Exec(lccommand,lcDoc)

	if err != nil{
		fmt.Println("error : ",err.Error())
	}

	lccommand = `delete bcnp.dbo.bcsaleordersub where docno =?`
	_,err = dbtest.Exec(lccommand,lcDoc)
	if err != nil{
		fmt.Println("error : ",err.Error())
	}


	newDoc,err := MockSo.Insert(dbtest)

	if err != nil {
		t.Error(err)
		return
	}
	//fmt.Println()
	t.Logf("Success insert new record : ",newDoc)
}
