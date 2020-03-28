package controllers

import (
	"ShengXianHouTai/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	user := c.GetSession("user").(models.User)
	c.Data["userName"] = user.Name
	c.Layout = "layout.html"
	c.TplName = "index.html"
}
