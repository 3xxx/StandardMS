package controllers

import (
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
	"path"
	"quick/models"
	"strconv"
	"strings"
)

type AttachController struct {
	beego.Controller
}

func (c *AttachController) Get() {
	//1.url处理中文字符路径，[1:]截掉路径前面的/斜杠
	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:])
	// beego.Info(filePath)
	// fileext := path.Ext(filePath) //取得文件扩展名
	route, err := url.QueryUnescape(c.Ctx.Request.RequestURI)
	// beego.Info(route)
	//2.取到成果类型
	filename := path.Dir(filePath) //这一步取得去除文件名后，前面的路径
	// beego.Info(filename)
	index := strings.LastIndex(filename, "/") //取得最后一个斜杠的索引值
	// fmt.Println(index)
	//	fmt.Println(len([]rune(filePath)))
	// fmt.Println(len(filePath))
	// fmt.Println(filePath[index+1 : len(filePath)])
	filetype := filename[index+1 : len(filename)] //取得文件类型字符
	// beego.Info(filetype)

	if filetype == "doc" || filetype == "dwg" || filetype == "xls" {
		//判断用户是否注册，注册后才能取得cookie中的用户，否则出错
		//取到用户后才能去除用户权限。其他不需要权限的，就跳过这些
		//1.首先判断是否注册
		if !checkAccount(c.Ctx) {
			// route, _ := url.QueryUnescape(c.Ctx.Request.RequestURI)
			port := strconv.Itoa(c.Ctx.Input.Port())
			route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
			c.Data["Url"] = route
			c.Redirect("/login?url="+route, 302)
			return
		}
		//2.取得客户端用户名
		sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
		defer sess.SessionRelease(c.Ctx.ResponseWriter)
		v := sess.Get("uname")
		// if v != nil {
		uname := v.(string)
		// }
		// ck, err := c.Ctx.Request.Cookie("uname")
		// if err != nil {
		// 	beego.Error(err)
		// }
		// uname := ck.Value
		// beego.Info(uname)
		//3.取出用户的权限等级
		role, _ := checkRole(c.Ctx) //login里的
		// beego.Info(role)
		//3.由路径查询数据库中的用户名
		username, err := models.GetattatchAuthor(route)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(username)
		//4.判断权限,如果用户访问的是doc/dwg/xls类，则需要注册和权限大于等于3.
		rolename, _ := strconv.ParseInt(role, 10, 64)
		if rolename > 3 && uname != username {
			port := strconv.Itoa(c.Ctx.Input.Port())
			route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
			c.Data["Url"] = route
			// c.Redirect("/login?url="+route, 302)
			c.Redirect("/roleerr?url="+route, 302)
			return
		}
	} else if filetype == "pdf" || filetype == "diary" || filetype == "jpg" || filetype == "tif" {
		//如果用户访问pdf或diary……，需要注册和权限4级
		//1.首先判断是否注册
		if !checkAccount(c.Ctx) {
			//这里获取登录前的url
			// route, _ := url.QueryUnescape(c.Ctx.Request.RequestURI)
			// beego.Info(route)
			port := strconv.Itoa(c.Ctx.Input.Port())
			route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
			// beego.Info(c.Ctx.Input.Site()) //http://localhost
			// beego.Info(c.Ctx.Input.Port())
			// beego.Info(c.Ctx.Input.URI())
			// beego.Info(c.Ctx.Input.URL()) // /attachment/DL65987 注册后用户名登陆测
			c.Data["Url"] = route
			c.Redirect("/login?url="+route, 302)
			return
		}
		//2.取得客户端用户名
		sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
		defer sess.SessionRelease(c.Ctx.ResponseWriter)
		v := sess.Get("uname")
		// if v != nil {
		uname := v.(string)
		// }
		// ck, err := c.Ctx.Request.Cookie("uname")
		// if err != nil {
		// 	beego.Error(err)
		// }
		// uname := ck.Value
		// beego.Info(uname)
		//3.取出用户的权限等级
		role, _ := checkRole(c.Ctx) //login里的
		// beego.Info(role)
		//3.由路径查询数据库中的用户名
		username, err := models.GetattatchAuthor(route)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(username)
		//4.判断权限,如果用户访问的是pdf/jpg/diary类，则需要注册和权限4.
		rolename, _ := strconv.ParseInt(role, 10, 64)
		if rolename > 4 && uname != username {
			port := strconv.Itoa(c.Ctx.Input.Port())
			route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
			c.Data["Url"] = route
			// c.Redirect("/login?url="+route, 302)
			c.Redirect("/roleerr?url="+route, 302)
			return
		}
	}

	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	f, err := os.Open(filePath)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	defer f.Close()
	_, err = io.Copy(c.Ctx.ResponseWriter, f)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
}
