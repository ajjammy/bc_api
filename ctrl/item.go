package ctrl

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/satit13/bc_api/bean/resp"
	api "github.com/satit13/bc_api/bean/resp"
)

func GetItem(c *gin.Context){
	log.Println("Call GET Item")
	c.Header("Server", "BC_API")
	c.Header("Host", "nopadol.net:8001")
	c.Header("Access-Control-Allow-Origin", "*")
	//keyword := c.Param("keyword")
	//token := c.Param("token")
	//param1 := c.URL.Query().Get("param1")
	token := c.Request.URL.Query().Get("token")
	keyword := c.Request.URL.Query().Get("keyword")


	fmt.Println("token = ",token)
	item := Resp.Item{}
	//result := so.GetByDocno("test")
	//fmt.Println("result object : ",result)
	fmt.Println("call Item.GetByCode :",keyword)


	items,err := item.GetByKeyword(keyword,dbx)
	if err != nil{
		log.Println(err.Error())
	}
	fmt.Println(items)

	rs := api.Response{}

	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = items
		//rs.Link.Self = config.API_HOST + "/v1/users/"
		c.JSON(http.StatusOK, rs)
	}

	//c.JSON(http.StatusOK,rs)

}


//c.JSON(http.StatusOK,rs)