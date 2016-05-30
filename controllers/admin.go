package controllers

import (
	// m "github.com/beego/admin/src/models"
	// "github.com/astaxie/beego/orm"
	// "bufio"
	// "crypto/md5"
	// "encoding/hex"
	"github.com/astaxie/beego"
	// "github.com/tealeg/xlsx"
	// "io"
	// "fmt"
	// "net"
	// "os"
	// m "quick/models"
	// "regexp"
	"strconv"
	// "strings"
	// "time"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Admin() {
	var rolename int
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//2.如果用户权限在允许范围内，进行访问权限检查
	uname, role, _ := checkRolewrite(c.Ctx) //login里的
	rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	//2.取得Id
	// id := c.Input().Get("cid")
	// if filetype != "pdf" && filetype != "jpg" && filetype != "diary" {
	if rolename > 1 { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	c.TplName = "AdminLTE/calendar.html"
}

func (c *AdminController) Ui() {
	var rolename int
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//2.如果用户权限在允许范围内，进行访问权限检查
	uname, role, _ := checkRolewrite(c.Ctx) //login里的
	rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	//2.取得Id
	// id := c.Input().Get("cid")
	// if filetype != "pdf" && filetype != "jpg" && filetype != "diary" {
	if rolename > 1 { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}

	// beego.Info(c.Ctx.Input.Param("0"))
	// beego.Info(c.Ctx.Input.Param("1"))
	c.TplName = "AdminLTE/pages/UI/" + c.Ctx.Input.Param("0")
}

func (c *AdminController) Mailbox() {
	var rolename int
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//2.如果用户权限在允许范围内，进行访问权限检查
	uname, role, _ := checkRolewrite(c.Ctx) //login里的
	rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	//2.取得Id
	// id := c.Input().Get("cid")
	// if filetype != "pdf" && filetype != "jpg" && filetype != "diary" {
	if rolename > 1 { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}

	// beego.Info(c.Ctx.Input.Param("0"))
	// beego.Info(c.Ctx.Input.Param("1"))
	c.TplName = "AdminLTE/pages/mailbox/" + c.Ctx.Input.Param("0")
}

func (c *AdminController) Get() {
	var rolename int
	var uname, role string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, role, _ = checkRoleread(c.Ctx) //login里的
	rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname

	//4.判断权限,如果用户访问的是doc/dwg/xls类，则需要注册和权限大于等于3.
	//或者在管理员设定的用户组内也可以，后续增加
	if rolename > 1 { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
}
