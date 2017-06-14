package Model


import (
//"database/sql"
"log"
"github.com/jmoiron/sqlx"
 _ "github.com/denisenkom/go-mssqldb"

)

func NewDB(dsn string,dbType int) (db *sqlx.DB ){
	//mySql
	//db := *sqlx.DB{}
	if dbType==0{

		db = sqlx.MustConnect("mysql", dsn)

	}
	//sql server
	if dbType ==1 {
		db = sqlx.MustConnect("mssql", dsn)

	}


	log.Println("db = ", db)
	return db
}

