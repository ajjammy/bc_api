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
