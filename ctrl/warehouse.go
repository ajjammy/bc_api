package ctrl

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	wh "github.com/satit13/bc_api/bean/resp"
	api "github.com/satit13/bc_api/bean/resp"
)


func GetWarehouseList(c *gin.Context){
	log.Println("call GET Warehouse List")
	c.Header("Server", "BC_API")
	c.Header("Host", "nopadol.net:8001")
	c.Header("Access-Control-Allow-Origin", "*")
	//keyword := c.Param("keyword")
	//token := c.Param("token")
	//param1 := c.URL.Query().Get("param1")
	access_token := c.Request.URL.Query().Get("access_token")
	keyword := c.Request.URL.Query().Get("keyword")

	fmt.Println("access_token = ",access_token)
	whcode := wh.Warehouse{}
	//result := so.GetByDocno("test")
	//fmt.Println("result object : ",result)
	fmt.Println("call Warehouse.GetWarehouseList :",keyword)

	ww,err := whcode.GetByKeyWord(keyword,dbx)
	if err != nil{
		log.Println(err.Error())
	}
	//fmt.Println(ee)
	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = ww
		//rs.Link.Self = config.API_HOST + "/v1/users/"
		c.JSON(http.StatusOK, rs)
	}
}

func GetWarehouse(c *gin.Context){
	log.Println("call GET Warehouse")
	c.Header("Server", "BC_API")
	c.Header("Host", "nopadol.net:8001")
	c.Header("Access-Control-Allow-Origin", "*")
	access_token := c.Request.URL.Query().Get("access_token")
	keyword := c.Request.URL.Query().Get("keyword")

	fmt.Println("access_token = ",access_token)
	whcode := wh.Warehouse{}

	fmt.Println("call Warehouse.GetWarehouse :",keyword)

	ww,err := whcode.GetWarehouseCode(keyword,dbx)
	if err != nil{
		log.Println(err.Error())
	}
	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		if ww==nil{
			//fmt.Println("Yes")
			rs.Status = "error"
			rs.Message = "No Content: NotData"
			c.JSON(http.StatusNotFound, rs)
		}else {
			rs.Status = "success"
			rs.Data = ww
			c.JSON(http.StatusOK, rs)
		}
	}
}
