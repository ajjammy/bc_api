package ctrl

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	so "github.com/satit13/bc_api/bean/resp"
	api "github.com/satit13/bc_api/bean/resp"
	"log"
	//"strconv"
)

func GetSaleorderList(c *gin.Context){
	so := so.Saleorder{}
	//result := so.GetByDocno("test")
	//fmt.Println("result object : ",result)
	fmt.Println("call so.GetByDocno :",dbx)
	ss,err := so.GetByKeyWord("W01-SCV5905-02",dbx)
	if err != nil{
		log.Println(err.Error())
	}
	fmt.Println(ss)
	c.JSON(http.StatusOK,ss)
}


func GetSaleorder(c *gin.Context){
	log.Println("call GET Saleorder By Docno")
	c.Header("Server", "BC_API")
	c.Header("Host", "nopadol.net:8001")
	c.Header("Access-Control-Allow-Origin", "*")
	//keyword := c.Param("keyword")
	//token := c.Param("token")
	//param1 := c.URL.Query().Get("param1")
	token := c.Request.URL.Query().Get("token")
	keyword := c.Request.URL.Query().Get("keyword")


	fmt.Println("token = ",token)
	so := so.Saleorder{}

	//result := so.GetByDocno("test")
	//fmt.Println("result object : ",result)
	fmt.Println("call so.GetByDocno :",keyword)


	ss,err := so.GetByKeyWord(keyword,dbx)
	if err != nil{
		log.Println(err.Error())
	}
	fmt.Println(ss)

	rs := api.Response{}

	if err != nil {
		rs.Status = api.ERROR
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = api.SUCCESS
		rs.Data = ss
		//rs.Link.Self = config.API_HOST + "/v1/users/"
		c.JSON(http.StatusOK, rs)
	}

	//c.JSON(http.StatusOK,rs)

}

func PostSaleorder(c *gin.Context){

}