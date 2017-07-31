package Resp
import (
	"testing"
)

func TestWarehouse_GetByKeyWord(t *testing.T) {
	w :=Warehouse{}
	e := "S1-A"
	a,_ := w.GetByKeyWord(e,dbtest)

	if e != a[0].Code {
		t.Fatalf("error expect value %s but got %s",e,a)
	}
}

func TestWarehouse_GetWarehouseCode(t *testing.T) {
	w :=Warehouse{}
	e := "S1-A"
	rs,_ := w.GetWarehouseCode(e,dbtest)

	a := rs[0].Code
	if e != a {
		t.Fatalf("error expect value %s but got %s",e,a)
	}
}