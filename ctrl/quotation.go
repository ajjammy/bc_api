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
	c.Header("Server", "BC_API")
	c.Header("Host", "nopadol.net:8001")
	c.Header("Access-Control-Allow-Origin", "*")
	token := c.Request.URL.Query().Get("token")
	docno := c.Request.URL.Query().Get("docno")

	fmt.Println("token = ",token)
	qto := qt.Quotation{}

	fmt.Println("call qt.GetByDocno :",docno)

	err :=qto.GetByDocno(docno,dbx)
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
