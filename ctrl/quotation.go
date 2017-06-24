package ctrl

import (
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	qt "github.com/satit13/bc_api/bean/resp"
	api "github.com/satit13/bc_api/bean/resp"
	"net/http"
)

func GetQuotation(c *gin.Context){
	log.Println("call GET Quotation By Docno")
	c.Keys=headerKeys
	token := c.Request.URL.Query().Get("token")
	docno := c.Request.URL.Query().Get("docno")

	fmt.Println("token = ",token)
	qto := new(qt.Quotation)
	qto.DocNo = docno

	fmt.Println("call qt.GetByDocno :",docno)

	err :=qto.GetByDocno(docno,dbx)
	fmt.Println("dbx = ",dbx.DB)
	if err != nil{
		log.Println(err.Error())
	}
	//fmt.Println(qt)
	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = qto
		c.JSON(http.StatusOK, rs)
	}
}
