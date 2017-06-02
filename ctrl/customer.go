package ctrl

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	cus "github.com/satit13/bc_api/bean/resp"
	api "github.com/satit13/bc_api/bean/resp"
)

func GetCustomerList(c *gin.Context){
	log.Println("call GET Customer List")
	c.Header("Server", "BC_API")
	c.Header("Host", "nopadol.net:8001")
	c.Header("Access-Control-Allow-Origin", "*")
	access_token := c.Request.URL.Query().Get("access_token")
	keyword := c.Request.URL.Query().Get("keyword")

	fmt.Println("access_token = ",access_token)
	cust := cus.Customer{}

	fmt.Println("call customer.GetCustomer :",keyword)

	cc,err := cust.GetByKeyWord(keyword,dbx)
	if err != nil{
		log.Println(err.Error())
	}
	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cc
		c.JSON(http.StatusOK, rs)
	}
}

func GetCustomer(c *gin.Context){
	log.Println("call GET Customer List")
	c.Header("Server", "BC_API")
	c.Header("Host", "nopadol.net:8001")
	c.Header("Access-Control-Allow-Origin", "*")
	access_token := c.Request.URL.Query().Get("access_token")
	keyword := c.Request.URL.Query().Get("keyword")

	fmt.Println("access_token = ",access_token)
	cust := cus.Customer{}

	fmt.Println("call customer.GetCustomer :",keyword)

	cc,err := cust.GetCustomerCode(keyword,dbx)
	if err != nil{
		log.Println(err.Error())
	}
	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: "+err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		if cc==nil{
			//fmt.Println("Yes")
			rs.Status = "error"
			rs.Message = "No Content: NotData"
			c.JSON(http.StatusNotFound, rs)
		}else {
			rs.Status = "success"
			rs.Data = cc
			c.JSON(http.StatusOK, rs)
			}
	}
}

