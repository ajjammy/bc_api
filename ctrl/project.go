package ctrl

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"github.com/satit13/bc_api/bean/resp"
	"log"
)

func GetProjectList(c *gin.Context){
	pjs := Resp.Project{}
	//result := so.GetByDocno("test")
	//fmt.Println("result object : ",result)

	ss,err := pjs.GetAll(dbx)
	if err != nil{
		log.Println(err.Error())
	}
	fmt.Println(ss)

	rs := Resp.Response{}

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


func GetProjectById(c *gin.Context){
	p := Resp.Project{}
	//result := so.GetByDocno("test")
	//fmt.Println("result object : ",result)
	id := c.Param("id")
	pj,err := p.GetById(dbx,id)

	if err != nil{
		log.Println(err.Error())
	}
	fmt.Println(pj)

	rs := Resp.Response{}

	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = pj
		//rs.Link.Self = config.API_HOST + "/v1/users/"
		c.JSON(http.StatusOK, rs)
	}
	//c.JSON(http.StatusOK,rs)
}