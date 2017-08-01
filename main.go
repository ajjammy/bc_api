package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//m "github.com/satit13/bc_api/bean"
	//	"github.com/satit13/bc_api/config"
	ctrl "github.com/satit13/bc_api/ctrl"
	//_ "github.com/denisenkom/go-mssqldb"
	"gopkg.in/gin-contrib/cors.v1"

)

func main() {
	fmt.Println("BC API Project")
	// 1 = MsSql server , 0 = MySql

	app := gin.Default()
	app.Use(cors.Default())
	app.GET("/", getVersion)
	app.GET("/branch" , ctrl.GetBranch)

	// SO
	app.GET("/saleorders", ctrl.GetSaleorderList)
	app.GET("/saleorder", ctrl.GetSaleorder)
	app.POST("/saleorder", ctrl.PostNewSaleorder)
	app.PUT("/saleorder", ctrl.PutSaleorder)
	app.DELETE("/saleorder", ctrl.VoidSaleorder)
	// QT
	app.GET("/quotation", ctrl.GetQuotation)
	app.POST("/quotation", ctrl.PostNewQuotation)
	app.PUT("/quotation", ctrl.PutQuotation)
	app.DELETE("/quotation", ctrl.VoidQuotation)
	app.POST("/qt", ctrl.PostQT)


	app.GET("/customer", ctrl.GetCustomer)
	app.GET("/customers", ctrl.GetCustomerList)

	app.GET("/employee", ctrl.GetEmployee)
	app.GET("/employees", ctrl.GetEmployeeList)
	//Warehouse & Shelf
	app.GET("/warehouses", ctrl.GetWarehouseList)
	app.GET("/warehouse", ctrl.GetWarehouse)
	app.GET("/shelfcodes", ctrl.GetShelfcodeList)
	app.GET("/shelfcode", ctrl.GetShelfcode)
	// items
	app.GET("/items", ctrl.GetItemList)
	app.GET("/item", ctrl.GetItem)
	// others master
	app.GET("/allocates", ctrl.GetAllocateList)
	app.GET("/projects", ctrl.GetProjectList)
	app.GET("/project/:id", ctrl.GetProjectById)
	app.GET("/shelfs" , ctrl.GetShelfcodeList)

	app.Run(":8000")

}

func getVersion(c *gin.Context) {
	c.JSON(http.StatusOK, "OK I'm Running on Docker !!!!")
}



