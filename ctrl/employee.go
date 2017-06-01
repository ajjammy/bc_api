package ctrl

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	sale "github.com/satit13/bc_api/bean/resp"
	api "github.com/satit13/bc_api/bean/resp"
)


func GetEmployee(c *gin.Context){
	log.Println("call GET Employee")
	c.Header("Server", "BC_API")
	c.Header("Host", "nopadol.net:8001")
	c.Header("Access-Control-Allow-Origin", "*")
	//keyword := c.Param("keyword")
	//token := c.Param("token")
	//param1 := c.URL.Query().Get("param1")
	access_token := c.Request.URL.Query().Get("access_token")
	keyword := c.Request.URL.Query().Get("keyword")


	fmt.Println("access_token = ",access_token)
	emp := sale.Employee{}

	//result := so.GetByDocno("test")
	//fmt.Println("result object : ",result)
	fmt.Println("call Employee.GetEmployee :",keyword)


	ee,err := emp.GetByKeyWord(keyword,dbx)
	if err != nil{
		log.Println(err.Error())
	}
	//fmt.Println(ee)

	rs := api.Response{}

	if err != nil {
		rs.Status = "false"
		rs.Message = "Search Employee error"
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "true"
		rs.Message = "Search Employee success"
		rs.Data = ee
		//rs.Link.Self = config.API_HOST + "/v1/users/"
		c.JSON(http.StatusOK, rs)
	}

}