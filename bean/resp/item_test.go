package Resp
import (
	"testing"
	"fmt"
)

func TestItem_GetByCode(t *testing.T) {
	e := "2120250"
	i := Item{}
	_ = i.GetByCode("2120250",dbtest)

	if i.Code != e {
		t.Fatalf("expect code = %s but got =%s",e,i.Code)
	}
}

func TestItem_GetByKeyword(t *testing.T) {
	fmt.Println("start test get item by keyword")
	e := "2120250"

	i := Item{}
	a,_:= i.GetByKeyword(e,dbtest)
	if a[0].Code !=e {
		t.Fatalf("exect to record set but got empty")
	}
}

