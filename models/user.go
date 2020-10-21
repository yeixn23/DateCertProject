package models

import (
	"BeegoHello/db_mysql"
	"DataCertProject/unti"
)

type User struct {
	Id       int       `from:"id"`
	Phone    string		`form:"phone"`
	Password string		`from:"password"`
}

/*
*保存用户信息的方法：保存用户信息到数据库中
*/
func (u User) SaveUser()(int64,error){
	//1,密码脱敏处理
	//md5Hash := md5.New()
	//	//md5Hash.Write([]byte(u.Password))
	//	//bytes := md5Hash.Sum(nil)
	u.Password = unti.Md5Hash(u.Password)
	//2，执行数据库操作
	row ,err := db_mysql.Db.Exec("insert into user(phone,password)" + "values(?,?)", u.Phone, u.Password)
	if err != nil { //保存数据时遇到错误
		return -1, err
	}
	id, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id,nil
}
func (u User) QueryUser()(*User,error){
	//md5Hash := md5.New()
	//md5Hash.Write([]byte(u.Password))
	//bytes := md5Hash.Sum(nil)
	u.Password = unti.Md5Hash(u.Password)
	row := db_mysql.Db.QueryRow("select phone from user where phone = ? and password =?",u.Phone,u.Password)
	err := row.Scan(&u.Phone)
	if err != nil{
		return nil,err
	}
	return &u,err
}