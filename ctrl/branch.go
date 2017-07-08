package ctrl

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/satit13/bc_api/bean/resp"
)
//[{"id":1,"code":"S01","name":"สำนักงานใหญ่"},{"id":2,"code":"S02","name":"สาขาสันกำแพง"},{"id":3,"code":"S01","name":"สำนักงานใหญ่"}]
type Branch struct {
	Code string
	Name string
}

func GetBranch(c *gin.Context){

	log.Println("call Void_Quotation()")
	c.Keys=headerKeys
	Branchs := []Branch {
		{"S01", "สำนักงานใหญ่"},
		{"S02", "สาขาสันกำแพง"},
	}

	rs := Resp.Response{}
	rs.Status = "success"
	rs.Data = Branchs

	c.JSON(http.StatusOK, rs)
}
