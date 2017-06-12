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
	app.POST("/saleorder",ctrl.PostNewSaleorder)
	//	app.POST("/saleorder", ctrl.PostSaleOrder)
	app.GET("/quotation",ctrl.GetQuotation)
	//	app.get("/quotation", ctrl.GetQuotation)
	app.GET("/customer",ctrl.GetCustomer)
	app.GET("/customers",ctrl.GetCustomerList)
	app.GET("/employee",ctrl.GetEmployee)
	app.GET("/employees",ctrl.GetEmployeeList)
	app.GET("/warehouses",ctrl.GetWarehouseList)
	app.GET("/warehouse",ctrl.GetWarehouse)
	app.GET("/shelfcodes",ctrl.GetShelfcodeList)
	app.GET("/shelfcode",ctrl.GetShelfcode)
	app.GET("/items",ctrl.GetItemList)
	app.GET("/item",ctrl.GetItem)
	app.Run(":8000")

}

func getVersion(c *gin.Context){
	c.JSON(http.StatusOK,"OK I'm Running on Docker !!!!")
}
