package ctrl

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	so "github.com/satit13/bc_api/bean/resp"
	api "github.com/satit13/bc_api/bean/resp"
	"log"
	//"strconv"
	//_ "npdl.co/it/BC_API/bean/resp"
	"npdl.co/it/BC_API/bean/resp"
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

	rs := api.Response{}

	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = ss
		//rs.Link.Self = config.API_HOST + "/v1/users/"
		c.JSON(http.StatusOK, rs)
	}
	//c.JSON(http.StatusOK,rs)
}


func GetSaleorder(c *gin.Context){
	log.Println("call GET Saleorder By Docno")
	c.Header("Server", "BC_API")
	c.Header("Host", "nopadol.net:8001")
	c.Header("Access-Control-Allow-Origin", "*")
	token := c.Request.URL.Query().Get("token")
	keyword := c.Request.URL.Query().Get("keyword")


	fmt.Println("token = ",token)
	so := so.Saleorder{}

	//result := so.GetByDocno("test")
	//fmt.Println("result object : ",result)
	fmt.Println("call so.GetByDocno :",keyword)


	err :=so.GetByDocno(keyword,dbx)
	if err != nil{
		log.Println(err.Error())
	}
	fmt.Println(so)

	rs := api.Response{}

	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = so
		//rs.Link.Self = config.API_HOST + "/v1/users/"
		c.JSON(http.StatusOK, rs)
	}

	//c.JSON(http.StatusOK,rs)

}

func PostNewSaleorder(c *gin.Context){
	rs := api.Response{}
	rs.Status="success"
	// todo : mock data saleorder & Test
	// todo : mock data saleordersub & test
	// todo : Post Saleorder
	// todo : Post Saleordersub
	c.JSON(http.StatusOK,rs)
}

