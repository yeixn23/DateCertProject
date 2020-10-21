package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)
/*
*router.go文件的作用：路由功能。用于接收并分发接收到的浏览器的请求
 */
func init() {
	//注册页面
    beego.Router("/", &controllers.MainController{})
    //用户注册的接口请求
    beego.Router("/user_register",&controllers.RegisterController{})
    //直接登录的页面请求接口
    beego.Router("/login.html",&controllers.LoginController{})
    beego.Router("/login",&controllers.LoginController{})
	beego.Router("/upload", &controllers.UploadController{})
}
