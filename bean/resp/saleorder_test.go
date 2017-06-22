package Resp

import (


	"github.com/jmoiron/sqlx"
	"fmt"
	"testing"
//	"net/http"
)
var dbx *sqlx.DB

func TestSaleorder_CheckExists(t *testing.T) {
	e := true
	lcDoc := "W01-SCV5905-0022"
<<<<<<< HEAD
	fmt.Println("Begin testSaleorder_CheckExists")
=======
	fmt.Println("<---------------Begin testSaleorder_CheckExists")
>>>>>>> dev
	MockSo := Saleorder{}
	a := MockSo.CheckExists(dbtest,lcDoc)
	fmt.Println(a)
	if e != a{
		t.Fatalf("Expected %s But got %s",e,a)
	}
}

func TestInsertSo(t *testing.T) {

	// Setup
	e := true
	fmt.Println("<---------------Begin testSaleorder_Insert")
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
