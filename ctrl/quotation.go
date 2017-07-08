package ctrl

import (
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	//qt "github.com/satit13/bc_api/bean/resp"
	m "github.com/satit13/bc_api/bean/resp"
	"net/http"
)

func GetQuotation(c *gin.Context){
	log.Println("call GET Quotation By Docno")
	c.Keys=headerKeys
	token := c.Request.URL.Query().Get("token")
	docno := c.Request.URL.Query().Get("docno")

	fmt.Println("token = ",token)
	qto := new(m.Quotation)
	qto.DocNo = docno

	fmt.Println("call qt.GetByDocno :",docno)

	err :=qto.GetByDocno(docno,dbx)
	fmt.Println("dbx = ",dbx.DB)
	if err != nil{
		log.Println(err.Error())
	}
	//fmt.Println(qt)
	rs := m.Response{}
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
	log.Println("call PostNewQuotation()")
	c.Keys=headerKeys

	qt := m.Quotation{}
	rs := m.Response{}
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
			rs.Message="Quotation : "+qt.DocNo+" Aready exists"
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


func VoidQuotation(c *gin.Context){

	log.Println("call Void_Quotation()")
	c.Keys=headerKeys
	token := c.Request.URL.Query().Get("token")
	docno := c.Request.URL.Query().Get("docno")
	fmt.Println("token = ",token)
	qto := m.Quotation{}
	qto.DocNo = docno
	rs := m.Response{}

	msg,result,err := qto.Void(dbx,qto.DocNo,"User")

	if err != nil{
		rs.Message=msg
		rs.Status="fail"
		c.JSON(http.StatusConflict,rs)
		return
	}

	// void ไม่ได้เนื่องจาก ติด Refer เอกสาร
	if result==false {
		rs.Message=msg
		rs.Status="fail"
		c.JSON(http.StatusConflict,rs)
		return
	}


	rs.Status ="success"
	rs.Message = msg
	c.JSON(http.StatusOK,rs)


}

func PutQuotation(c *gin.Context){
	log.Println("call edit: Quotation()")
	c.Keys=headerKeys

	//todo: delete old data
	qt := m.Quotation{}
	rs := m.Response{}
	if err := c.BindJSON(&qt); err != nil{
		fmt.Println(qt)
		log.Println("Error decode.Decode(&qt) >>", err)
		rs.Status = "fail"
		rs.Message = err.Error()
		c.JSON(http.StatusOK,rs)
		return
	}


	existdocno := qt.CheckExists(dbx,qt.DocNo)
	fmt.Println(existdocno)
	if !existdocno{
		rs.Status = "fail"
		rs.Message = "ไม่มีเอกสารเลชที่นี้อยู่ ไม่สามารถ Update ได้"
		return
	}

	msg,err := qt.Update(dbx)
	if err != nil{
		rs.Status="fail"
		rs.Message = msg+err.Error()
		c.JSON(http.StatusConflict,rs)
		return
	}

	rs.Status = "success"
	rs.Data = qt
	c.JSON(http.StatusOK,rs)
	return



}