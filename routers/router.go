package routers

import (
	"ShengXianHouTai/controllers"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
)


func init() {
	//添加过滤器
	beego.InsertFilter("/index", beego.BeforeExec, loginFilter)
	beego.InsertFilter("/user/*", beego.BeforeExec, loginFilter)
	beego.InsertFilter("/goods/*", beego.BeforeExec, loginFilter)
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/index", &controllers.MainController{})
	beego.Router("/goods/addType", &controllers.GoodsController{}, "get:ShowAddType;post:AddType")
	beego.Router("/user/logout", &controllers.UserController{}, "get:HandleLogout")
}

//登录过滤器
var loginFilter = func(ctx *context.Context) {
	user := ctx.Input.Session("user")
	if user == nil {
		ctx.Redirect(302, "/login")
	}
}