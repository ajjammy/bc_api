package Resp
import (
	"testing"
)

func TestShelfcode_GetShelfcode(t *testing.T) {
	sh := Shelfcode{}
	_,err := sh.GetShelfcode("-",dbtest)
	if err != nil{
		t.Fatalf("Expect Pass but Error with : ",err.Error())
	}
}


func TestShelfcode_GetByKeyWord(t *testing.T) {
	sh := Shelfcode{}
	_,err := sh.GetByKeyWord("A",dbtest)
	if err != nil{
		t.Fatalf("Expect Pass but Error with : ",err.Error())
	}
}