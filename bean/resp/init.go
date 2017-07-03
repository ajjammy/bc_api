package Resp



import (
	"github.com/jmoiron/sqlx"
	db "github.com/satit13/bc_api/bean"
	config "github.com/satit13/bc_api/config"
	"log"
	"fmt"
)

var dbtest *sqlx.DB
//var rs *Response

func init() {
	// Read configuration file from "config.json"
	//dsn := GetConfig("./model/config.json") // เปิดใช้งานจริงเมื่อ Docker Container run --link ตรงเข้า mariadb เท่านั้น

	dsn := config.LoadDSN("./config.json",1)

	//curDir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	//fmt.Println("curdir :",curDir)


	fmt.Println("resp.init.go -> start connect test in init()")
	//1 = MsSql server
	dbtest = db.NewDB(dsn,1)
	err := dbtest.Ping()

	if err != nil {
		fmt.Println("Error connect database : ",err.Error())
	}

	fmt.Println("dsn : " ,dsn)

	log.Println("Connected db: ", dbtest)
}

