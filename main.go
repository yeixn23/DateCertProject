package main

import (
	"DataCertProject/blockchain"
	"DataCertProject/db_mysql"
	"DataCertProject/models"
	_ "DataCertProject/routers"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	user := models.User{
		Id:       0,
		Phone:    "11",
		Password: "123",
	}
	fmt.Println(user)
	user1,_:=json.Marshal(user)
	fmt.Println(string(user1))
	var user3 models.User
	json.Unmarshal(user1,&user3)
	fmt.Println(user3)
	return

	//1.生成第一个区块
	block := blockchain.NewBlock(0,[]byte{},[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})

	fmt.Println(block)
	fmt.Printf("区块链hsah值:%x",block.Hash)
	return


	//1.链接数据库
	db_mysql.ConnectDB()

	//2、静态资源路径设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	//3、允许
	beego.Run()//启动端口监听：阻塞
}