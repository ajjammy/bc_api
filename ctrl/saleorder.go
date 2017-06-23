package ctrl

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	so  "github.com/satit13/bc_api/bean/resp"
	api "github.com/satit13/bc_api/bean/resp"
	"log"
	//"strconv"
	//_ "npdl.co/it/BC_API/bean/resp"
	//"npdl.co/it/BC_API/bean/resp"
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


func PutSaleorder(c *gin.Context){
	log.Println("call PostNewSaleOrder()")
	c.Header("Server", "BC_API ")
	c.Header("Host", "nopadol.net:8000")
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	//todo: delete old data
	fmt.Println("call PutSaleorder()")
	so := so.Saleorder{}
	rs := api.Response{}

	if err := c.BindJSON(&so); err != nil{
		fmt.Println(so)
		log.Println("Error decode.Decode(&so) >>", err)
		rs.Status = "fail"
		rs.Message = err.Error()
		c.JSON(http.StatusOK,rs)
		return
	}

	existdocno := so.CheckExists(dbx,so.Docno)
	fmt.Println(existdocno)
	if !existdocno{
		rs.Status = "fail"
		rs.Message = "ไม่มีเอกสารเลชที่นี้อยู่ ไม่สามารถ Update ได้"
		return
	}
	// DELETE
	err :=  so.Delete(dbx,so.Docno)

	if err != nil {  // cannot delete
		//  มีรายการแล้ว
		rs.Status="fail"
		rs.Message = err.Error()
		c.JSON(http.StatusConflict,rs)
		return
	}

	// INSERT SO UPDATE
	_,err = so.Insert(dbx)
	if err != nil{
		rs.Status="fail"
		rs.Message = err.Error()
		c.JSON(http.StatusConflict,rs)
		return
	}

	rs.Status = "success"
	rs.Data = so
	c.JSON(http.StatusOK,rs)
	return
}


func PostNewSaleorder(c *gin.Context){
		log.Println("call PostNewSaleOrder()")
		c.Header("Server", "BC_API ")
		c.Header("Host", "nopadol.net:8000")
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")

		fmt.Println("Ctrl.PostNewSaleorder")
		so := so.Saleorder{}
		rs := api.Response{}
		if err := c.BindJSON(&so); err != nil{
			fmt.Println(so)
			log.Println("Error decode.Decode(&so) >>", err)
			rs.Status = "fail"
			rs.Message = err.Error()
			c.JSON(http.StatusOK,rs)

		} else {
		if so.CheckExists(dbx,so.Docno) == true {
			//  มีรายการแล้ว
			rs.Status="fail"
			rs.Message="SaleOrder : "+so.Docno+" Aready exists"
			c.JSON(http.StatusConflict,rs)
			return
		}

		newSoNumber,err := so.Insert(dbx)
		fmt.Println("<---------------Start insert code")
		if err != nil {
			fmt.Println("Error Insert DB:", err)
			rs.Status = "fail"
			rs.Message = "Error Insert Saleorder :"+err.Error()
			c.JSON(http.StatusBadRequest,rs)
			return
		} else {
			rs.Status = "success"
			rs.Data = newSoNumber
		}
	}

	//-------------------

	rs.Status="success"
	// todo : mock data saleorder & Test
	// todo : mock data saleordersub & test
	// todo : Post Saleorder
	// todo : Post Saleordersub
	c.JSON(http.StatusOK,rs)
}

func VoidSaleorder(c *gin.Context){
	log.Println("call PostNewSaleOrder()")
	c.Header("Server", "BC_API ")
	c.Header("Host", "nopadol.net:8000")
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")


	token := c.Request.URL.Query().Get("token")
	fmt.Println(token)
	soNumber := c.Request.URL.Query().Get("soNumber")

	rs := api.Response{}
	so := so.Saleorder{}
	err := so.Void(dbx,soNumber)
	if err != nil{
		rs.Status="fail"
		rs.Message=err.Error()
		c.JSON(http.StatusBadRequest,rs)
		return
	}
	rs.Status = "success"
	c.JSON(http.StatusOK,rs)
}