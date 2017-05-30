package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//m "github.com/satit13/bc_api/bean"
//	"github.com/satit13/bc_api/config"
	mrq "github.com/satit13/bc_api/bean/resp"
	_ "github.com/denisenkom/go-mssqldb"
	//"github.com/jmoiron/sqlx"
	//"github.com/revel/modules/db/app"
)


func main(){
	fmt.Println("BC API Project")
	// 1 = MsSql server , 0 = MySql
//	dsn := config.LoadDSN("config.json",1)

	// 1 = MsSql server
	//db := m.NewDB(dsn,1)
	//err := db.Ping()
	//if err != nil {
	//	fmt.Println("Error connect database : ",err.Error())
	//}
	//fmt.Println("dsn : " ,dsn)
	app := gin.Default()
	app.GET("/",getVersion)
	app.GET("/saleorder", getSaleorder)
	app.Run(":8000")

}

func getVersion(c *gin.Context){
	c.JSON(http.StatusOK,"OK I'm Running on Docker !!!!")
}

func getSaleorder(c *gin.Context){
	so := mrq.Saleorder{}
	//result := so.GetByDocno("test")
	//fmt.Println("result object : ",result)
	so.GetByDocno("s01-shv5902-0021")
	fmt.Println(so)

	c.JSON(http.StatusOK,so)

}