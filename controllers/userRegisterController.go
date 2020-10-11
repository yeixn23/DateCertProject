package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post() {
	//1,解析请求
	var user models.User
	err := r.ParseForm(&user)
	if err !=nil {
		//返回错误信息给浏览器，提示用户
		r.Ctx.WriteString("抱歉解析数据错误，请重试！")
		return
	}
	//2,保存用户信息到数据库
	_, err1 := user.SaveUser()
	//3,返回前端结果(成功跳登录页面，失败弹出错误信息)
	if err1 != nil {
		r.Ctx.WriteString("抱歉用户注册失败,请重试！")
		return
	}
	//用户注册成功
	r.TplName = "login.html"
}
