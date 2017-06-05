package ctrl


import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	sh "github.com/satit13/bc_api/bean/resp"
	api "github.com/satit13/bc_api/bean/resp"
)

func GetShelfcodeList(c *gin.Context){
	log.Println("call GET Shelfcode List")
	c.Header("Server", "BC_API")
	c.Header("Host", "nopadol.net:8001")
	c.Header("Access-Control-Allow-Origin", "*")
	//keyword := c.Param("keyword")
	//token := c.Param("token")
	//param1 := c.URL.Query().Get("param1")
	access_token := c.Request.URL.Query().Get("access_token")
	keyword := c.Request.URL.Query().Get("keyword")

	fmt.Println("access_token = ",access_token)
	shelf := sh.Shelfcode{}
	//result := so.GetByDocno("test")
	//fmt.Println("result object : ",result)
	fmt.Println("call Shelfcode.GetShelfcode :",keyword)

	ss,err := shelf.GetByKeyWord(keyword,dbx)
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
		rs.Data = ss
		//rs.Link.Self = config.API_HOST + "/v1/users/"
		c.JSON(http.StatusOK, rs)
	}
}

func GetShelfcode(c *gin.Context){
	log.Println("call GET Shelfcode")
	c.Header("Server", "BC_API")
	c.Header("Host", "nopadol.net:8001")
	c.Header("Access-Control-Allow-Origin", "*")
	access_token := c.Request.URL.Query().Get("access_token")
	keyword := c.Request.URL.Query().Get("keyword")

	fmt.Println("access_token = ",access_token)
	shelf := sh.Shelfcode{}

	fmt.Println("call Shelfcode.GetShelfcode :",keyword)

	ss,err := shelf.GetShelfcode(keyword,dbx)
	if err != nil{
		log.Println(err.Error())
	}
	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		if ss==nil{
			rs.Status = "error"
			rs.Message = "No Content: NotData"
			c.JSON(http.StatusNotFound, rs)
		}else {
			rs.Status = "success"
			rs.Data = ss
			c.JSON(http.StatusOK, rs)
		}
	}
}

