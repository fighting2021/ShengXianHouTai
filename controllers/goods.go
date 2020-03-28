package controllers

import (
	"ShengXianHouTai/models"
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/weilaihui/fdfs_client"
	"io/ioutil"
	"path"
)

type GoodsController struct {
	beego.Controller
}

//显示添加分类界面
func (c *GoodsController)ShowAddType() {
	o := orm.NewOrm()
	var types []models.GoodsType
	o.QueryTable("GoodsType").All(&types)
	c.Data["types"] = types
	c.Layout = "layout.html"
	c.TplName = "addType.html"
}

//添加分类
func (c *GoodsController)AddType() {
	typeName := c.GetString("typeName")
	logo := c.GetString("logo")
	//上传类型图片
	imagePath, errMsg := UploadFile(&c.Controller, "uploadTypeImage")
	fmt.Printf("imagePath = %s, errMsg = %s", imagePath, errMsg)
	if errMsg != "" {
		c.Data["errMsg"] = errMsg
		c.ShowAddType()
	}
	if typeName == "" || logo == "" || imagePath == "noImg" {
		c.Data["errMsg"] = "填写信息不完整"
		c.ShowAddType()
		return
	}
	var goodsType models.GoodsType
	goodsType.Name = typeName
	goodsType.Logo = logo
	goodsType.Image = imagePath
	o := orm.NewOrm()
	_, err := o.Insert(&goodsType)
	if err != nil {
		c.Data["errMsg"] = "添加失败"
		c.ShowAddType()
		return
	}
	c.Redirect("/goods/addType",302)
}

//功能：实现FDFS的文件上传
//参数：第一个参数代表当前的控制器对象，第二个参数代表上传字段名
//返回值：第一个返回值代表上传文件ID，第二个返回值代表上传失败信息
func UploadFile(c *beego.Controller, name string) (fileId string, errMsg string) {
	file, header, err := c.GetFile(name)
	if err != nil{
		fileId = "NoImg"
		return
	}
	defer file.Close()
	//验证文件格式
	fileExt := path.Ext(header.Filename) //获取文件的后缀名
	if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".jpeg" {
		errMsg = "上传文件格式有误"
		return
	}
	//验证文件大小
	if header.Size > (100 * 1024) {
		errMsg = "上传图片太大"
		return
	}

	//如果上传文件验证通过，则把文件上传到FDFS服务器。
	client, err := fdfs_client.NewFdfsClient("../github.com/weilaihui/fdfs_client/client.conf")
	if err != nil {
		errMsg = "无法访问FDFS服务器"
		return
	}
	//获取文件读取到Reader中
	reader := bufio.NewReader(file)
	//从Reader中读取文件内容，并保存到[]byte中。
	fileBuffer, err := ioutil.ReadAll(reader)
	//应为fileExt包含.，所以fileExt[1:]把.去掉
	resp, err := client.UploadByBuffer(fileBuffer, fileExt[1:])
	if err != nil {
		errMsg = "上传文件失败"
		return
	}
	fileId = resp.RemoteFileId
	return
}

