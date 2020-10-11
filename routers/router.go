package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)
/*
*router.go文件的作用：路由功能。用于接收并分发接收到的浏览器的请求
 */
func init() {
    beego.Router("/", &controllers.MainController{})
    //用户注册的接口请求
    beego.Router("/user_register",&controllers.RegisterController{})
    beego.Router("/login.html",&controllers.LoginController{})
    beego.Router("home.html",&controllers.HomeController{})
}
