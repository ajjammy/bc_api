package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//m "github.com/satit13/bc_api/bean"
//	"github.com/satit13/bc_api/config"
	 "github.com/satit13/bc_api/ctrl"
	_ "github.com/denisenkom/go-mssqldb"
	//"github.com/jmoiron/sqlx"
	//"github.com/revel/modules/db/app"
)

func main(){
	fmt.Println("BC API Project")
	// 1 = MsSql server , 0 = MySql

	app := gin.Default()
	app.GET("/",getVersion)
	app.GET("/saleorders", ctrl.GetSaleorderList)
	//app.GET("/saleorder/:keyword/:token",ctrl.GetSaleorder)
	app.GET("/saleorder",ctrl.GetSaleorder)
	app.POST("/saleorder", ctrl.PostSaleorder)
	app.GET("/customers",ctrl.GetCustomer)
	app.GET("/item",ctrl.GetItem)
	app.Run(":8000")

}

func getVersion(c *gin.Context){
	c.JSON(http.StatusOK,"OK I'm Running on Docker !!!!")
}
