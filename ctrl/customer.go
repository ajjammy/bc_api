package ctrl

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	c "github.com/satit13/bc_api/bean/resp"
)

func GetCustomer(c *gin.Context){
	log.Println("call GET Customer")
	c.Header("Server", "BC_API")
	c.Header("Host", "nopadol.net:8001")
	c.Header("Access-Control-Allow-Origin", "*")
	//keyword := c.Param("keyword")
	//token := c.Param("token")
	//param1 := c.URL.Query().Get("param1")
	token := c.Request.URL.Query().Get("token")
	keyword := c.Request.URL.Query().Get("keyword")


	fmt.Println("token = ",token)
	cust := c.Customer{}

	//result := so.GetByDocno("test")
	//fmt.Println("result object : ",result)
	fmt.Println("call so.GetCustomer :",keyword)


	cc,err := cust.GetByKeyWord(keyword,dbx)
	if err != nil{
		log.Println(err.Error())
	}
	fmt.Println(cc)
	c.JSON(http.StatusOK,cc)

}
