package Resp

import (
	"testing"
	"fmt"
)



func TestEmployee_GetEmployeeCode(t *testing.T) {
	e := Employee{}

	em,_ := 	e.GetEmployeeCode("41054",dbtest)
	fmt.Println(em)
	if em[0].Name != "สาธิต โฉมวัฒนา"{
			t.Fatalf("Expected สาธิต  but got  : ",em[0].Name )
	}
}

func TestEmployee_GetByKeyWord(t *testing.T) {
	e := Employee{}
	em,_ := e.GetByKeyWord("41054",dbtest)

	if len(em)!=1 {
		t.Fatalf("Exected 1 record but we got , ",len(em))
	}
}