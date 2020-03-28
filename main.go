package main

import (
	_ "ShengXianHouTai/routers"
	"github.com/astaxie/beego"
)

func main() {
	//定义视图函数
	//第一个代表视图函数名，第二个参数代表Go函数
	beego.AddFuncMap("showPrePage", showPrePage)
	beego.AddFuncMap("showNextPage", showNextPage)
	beego.Run()
}

//定义视图函数，返回上一页
func showPrePage(pageIndex int) int {
	if pageIndex == 1 {
		return 1
	}
	return pageIndex - 1
}

//定义视图函数，返回下一页
func showNextPage(pageIndex int, pageCount int) int {
	if pageIndex == pageCount {
		return pageCount
	}
	return pageIndex + 1
}
