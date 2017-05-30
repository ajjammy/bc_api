package ctrl

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	so "github.com/satit13/bc_api/bean/resp"
	"log"
)

func GetSaleorder(c *gin.Context){
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