package Resp

import (


	"github.com/jmoiron/sqlx"
	"fmt"
	"testing"
//	"net/http"
)
var dbx *sqlx.DB


func TestInsertSo(t *testing.T) {

	// Setup
	e := true

	MockSo := Saleorder{}
	lcDoc := "W01-SCV5905-0022"
	err := MockSo.GetByDocno(lcDoc,dbtest)
	//fmt.Println(MockSo)
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


	_,err = MockSo.Insert(dbtest)

	if err != nil {
		t.Error(err)
		t.Fatalf("expect %s but Got Error! %s",e,err.Error())
	}

}
