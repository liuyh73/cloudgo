package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/liuyh73/cloudgo/models"
)

type DetailController struct {
	beego.Controller
}

func (c *DetailController) Get() {
	username := c.Ctx.Input.Param(":name")
	fmt.Println(username)
	c.Data["user"] = models.User{
		Username:  username,
		Password:  "123",
		Email:     "sd",
		Telephone: "15989067460",
		Id:        "12345678",
	}
	c.TplName = "detail.jade"
}
