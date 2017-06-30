package ctrl

import (
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	qt "github.com/satit13/bc_api/bean/resp"
	api "github.com/satit13/bc_api/bean/resp"
	"net/http"
)

func GetQuotation(c *gin.Context){
	log.Println("call GET Quotation By Docno")
	c.Keys=headerKeys
	token := c.Request.URL.Query().Get("token")
	docno := c.Request.URL.Query().Get("docno")

	fmt.Println("token = ",token)
	qto := new(qt.Quotation)
	qto.DocNo = docno

	fmt.Println("call qt.GetByDocno :",docno)

	err :=qto.GetByDocno(docno,dbx)
	fmt.Println("dbx = ",dbx.DB)
	if err != nil{
		log.Println(err.Error())
	}
	//fmt.Println(qt)
	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = qto
		c.JSON(http.StatusOK, rs)
	}
}

func PostNewQuotation(c *gin.Context){
	log.Println("call PostNewSaleOrder()")
	c.Keys=headerKeys

	fmt.Println("Ctrl.PostNewSaleorder")
	qt := qt.Quotation{}
	rs := api.Response{}
	if err := c.BindJSON(&qt); err != nil{
		fmt.Println(qt)
		log.Println("Error decode.Decode(&so) >>", err)
		rs.Status = "fail"
		rs.Message = err.Error()
		c.JSON(http.StatusOK,rs)

	} else {
		if qt.CheckExists(dbx,qt.DocNo) == true {
			//  มีรายการแล้ว
			rs.Status="fail"
			rs.Message="SaleOrder : "+qt.DocNo+" Aready exists"
			c.JSON(http.StatusConflict,rs)
			return
		}

		newQtNumber,err := qt.Insert(dbx)
		fmt.Println("<---------------Start insert code")
		if err != nil {
			fmt.Println("Error Insert DB:", err)
			rs.Status = "fail"
			rs.Message = "Error Insert Quotation :"+err.Error()
			c.JSON(http.StatusBadRequest,rs)
			return
		} else {
			rs.Status = "success"
			rs.Data = newQtNumber
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