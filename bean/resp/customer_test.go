package Resp
import (
	"testing"
	"fmt"
)

func TestCustomer_GetCustomerCode(t *testing.T) {
	e := "41054"
	cus := Customer{}
	cus.Code = e
	cuss,_ := cus.GetCustomerCode(e,dbtest)

	fmt.Println("name 1 = ",cuss[0].Name1)
	if cuss[0].Name1 !="นาย สาธิต  โฉมวัฒนา" {
		t.Fatalf("exect (สาธิต โฉมวัฒนา) But Got value -> " ,cus.Name1)
	}



}