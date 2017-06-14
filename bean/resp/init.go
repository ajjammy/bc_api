package Resp



import (
"github.com/jmoiron/sqlx"
db "github.com/satit13/bc_api/bean"
config "github.com/satit13/bc_api/config"
"log"
"fmt"
//	"github.com/revel/modules/db/app"
)

var dbtest *sqlx.DB
//var rs *Response

func init() {
	// Read configuration file from "config.json"
	//dsn := GetConfig("./model/config.json") // เปิดใช้งานจริงเมื่อ Docker Container run --link ตรงเข้า mariadb เท่านั้น

	dsn := config.LoadDSN("config.json",1)
	fmt.Println("ctrl.init.go -> start connect in init()")
	//1 = MsSql server
	dbtest = db.NewDB(dsn,1)
	err := dbtest.Ping()

	if err != nil {
		fmt.Println("Error connect database : ",err.Error())
	}

	fmt.Println("dsn : " ,dsn)

	log.Println("Connected db: ", dbtest)
}

