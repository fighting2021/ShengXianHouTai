package controllers

import (
	"ShengXianHouTai/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

//显示登录界面
func (c *UserController)ShowLogin() {
	userName := c.Ctx.GetCookie("userName")
	if userName != "" {
		c.Data["checked"] = "checked"
	}
	c.Data["userName"] = userName
	c.TplName = "login.html"
}

//处理登录请求
func (c *UserController)HandleLogin() {
	//获取请求参数
	userName := c.GetString("userName")
	userPass := c.GetString("password")
	//查询用户
	o := orm.NewOrm() //创建ORM对象
	var user models.User
	user.Name = userName //根据名字查询
	err := o.Read(&user, "Name") //执行查询
	if err != nil {
		c.Data["errMsg"] = "用户名不存在"
		c.TplName = "login.html"
		return
	}
	if user.PassWord != userPass || user.Power != 1 {
		c.Data["errMsg"] = "用户名或密码不正确"
		c.TplName = "login.html"
		return
	}

	//记住用户名
	rem := c.GetString("remember")
	if rem == "on" {
		c.Ctx.SetCookie("userName", userName, 60 * 60 * 24 * 7)
		c.Data["userName"] = userName
		c.Data["checked"] = "checked"
	} else {
		c.Ctx.SetCookie("userName", "", 0)
		c.Data["userName"] = ""
		c.Data["checked"] = ""
	}
	//把用户信息记录在Session中
	c.SetSession("user", user)
	c.Redirect("/index", 302)
}

//处理用户退出
func (c *UserController)HandleLogout() {
	c.DelSession("user") //删除Session
	c.Redirect("/login", 302)
}