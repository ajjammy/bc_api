package Resp
import (
	"testing"
	"fmt"
)

func TestQuotation_GetByDocno(t *testing.T) {
	e := "W02-QCV6003-0066"
	qt := Quotation{}
	qt.DocNo = e
	_ = qt.GetByDocno(e,dbtest)
	fmt.Println("docno = ",qt.DocNo)

	if qt.DocNo != e{
		t.Fatalf("expect docno %s but got %s",e,qt.DocNo)
	}
}