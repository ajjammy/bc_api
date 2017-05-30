package ctrl

import (
	"github.com/jmoiron/sqlx"
	db "github.com/mrtomyum/bc_api/bean"
	config "github.com/mrtomyum/bc_api/config"
	"log"
	"fmt"
)

var dbx *sqlx.DB
//var rs *Response

func init() {
	// Read configuration file from "config.json"
	//dsn := GetConfig("./model/config.json") // เปิดใช้งานจริงเมื่อ Docker Container run --link ตรงเข้า mariadb เท่านั้น

	dsn := config.LoadDSN("config.json",1)
	fmt.Println("ctrl.init.go -> start connect in init()")
	//1 = MsSql server
	dbx = db.NewDB(dsn,1)
	err := dbx.Ping()

	if err != nil {
		fmt.Println("Error connect database : ",err.Error())
	}

	fmt.Println("dsn : " ,dsn)

	log.Println("Connected db: ", dbx)
}

