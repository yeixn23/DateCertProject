package db_mysql

import (
	"fmt"
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"

)

var Db *sql.DB

/*
*在初始函数中连接数据库
 */
func ConnectDB() {
	fmt.Println("连接mysql数据库")
	config := beego.AppConfig
	dbDriver := config.String("driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")

	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	db, err := sql.Open(dbDriver, connUrl)
	if err != nil {
		panic("连接数据库失败")
	}
	//为全局赋值
	Db = db
}