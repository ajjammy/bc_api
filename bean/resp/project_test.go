package Resp
import (
	"testing"
)

func TestProject_GetAll(t *testing.T) {
	p := Project{}

	a,_ := p.GetAll(dbtest)
	if len(a) ==0 {
		t.Fatalf("Error Empty result")
	}
}

