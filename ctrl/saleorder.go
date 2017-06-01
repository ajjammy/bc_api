package ctrl

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	so "github.com/satit13/bc_api/bean/resp"
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
	keyword := c.Param("keyword")



	so := so.Saleorder{}

	//result := so.GetByDocno("test")
	//fmt.Println("result object : ",result)
	fmt.Println("call so.GetByDocno :",keyword)


	ss,err := so.GetByKeyWord(keyword,dbx)
	if err != nil{
		log.Println(err.Error())
	}
	fmt.Println(ss)
	c.JSON(http.StatusOK,ss)

}

func PostSaleorder(c *gin.Context){

}