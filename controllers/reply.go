package controllers

import (
	"github.com/astaxie/beego"
	"quick/models"
)

type ReplyController struct {
	beego.Controller
}

func (c *ReplyController) Add() {
	// if !checkAccount(c.Ctx) {
	// 	c.Redirect("/login", 302)
	// 	return
	// }
	// //由文章Id获取文章author
	// role, _ := checkRole(c.Ctx)
	// if role != "1" { //or（||或） 文章的作者不等于user
	// 	c.Redirect("/roleerr", 302)
	// 	return
	// }
	tid := c.Input().Get("tid") //tid是文章id
	err := models.AddReply(tid, c.Input().Get("nickname"), c.Input().Get("content"))
	if err != nil {
		beego.Error(err)
	}
	op := c.Input().Get("op")
	if op == "b" {
		c.Redirect("/topic/view_b/"+tid, 302)
	} else {
		c.Redirect("/topic/view/"+tid, 302)
	}
}
func (c *ReplyController) Delete() {
	if !checkAccount(c.Ctx) {
		return
	}
	tid := c.Input().Get("tid")
	err := models.DeleteReply(c.Input().Get("rid"))
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic/view/"+tid, 302)
}
