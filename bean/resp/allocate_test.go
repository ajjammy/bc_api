package Resp
import (
"testing"
)

func TestGetAllocateList(t *testing.T) {
	a := Allocate{}
	_,err := 	a.GetAll(dbtest)
	if err != nil {
		t.Fatalf("Expected Not Error  but error : ",err.Error() )
	}
}
