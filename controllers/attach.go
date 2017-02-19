package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/httplib"
	"hydrocms/models"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	// "strings"
)

type AttachController struct {
	beego.Controller
}

func (c *AttachController) Get1() {
	var rolename int
	var uname, role string
	var route string
	//1.url处理中文字符路径，[1:]截掉路径前面的/斜杠
	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:]) //  attachment/SL2016测试添加成果/A/FB/1/Your First Meteor Application.pdf
	fileext := path.Ext(filePath)                                    //取得文件扩展名.pdf
	// route, err := url.QueryUnescape(c.Ctx.Request.RequestURI)
	// beego.Info(route) //   /attachment/SL2016测试添加成果/A/FB/1/Your First Meteor Application.pdf
	//2.取到成果类型
	// filename := path.Dir(filePath) //这一步取得去除文件名后，前面的路径
	// beego.Info(filename)                      //attachment/SL2016测试添加成果/A/FB/1
	// index := strings.LastIndex(filename, "/") //取得最后一个斜杠的索引值
	// fmt.Println(index)                        //40
	//	fmt.Println(len([]rune(filePath)))
	// fmt.Println(len(filePath))
	// fmt.Println(filePath[index+1 : len(filePath)]) //  1/Your First Meteor Application.pdf
	// filetype := filename[index+1 : len(filename)]  //取得文件类型字符
	// beego.Info(filetype)                           // 1
	//即取得上级目录的类型，如果是fdiary，或扩展名是jpg等图片，或自定义的图文模式等，均不用权限
	switch fileext {
	case ".doc", ".docx", ".dwg", ".xls", ".xlsx", ".dgn", ".rar", ".zip", ".tar", ".gz", ".7z", ".bz2", ".cab", ".iso", ".ppt", ".pptx":
		//判断用户是否登录，登录后才能取得cookie/session中的用户，否则出错
		//1.首先判断是否登录或ip在预设允许范围内
		//这一步可以不用，直接用第2步代替
		// if !checkAccount(c.Ctx) {
		// 	// route, _ := url.QueryUnescape(c.Ctx.Request.RequestURI)
		// 	port := strconv.Itoa(c.Ctx.Input.Port())
		// 	route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
		// 	c.Data["Url"] = route
		// 	c.Redirect("/login?url="+route, 302)
		// 	return
		// }
		//2.如果登录或ip在允许范围内，进行访问权限检查
		uname, role, _ = checkRoleread(c.Ctx) //login里的
		rolename, _ = strconv.Atoi(role)
		c.Data["Uname"] = uname

		//3.由路径查询数据库中的用户名
		username, err := models.GetattachAuthor(route)
		if err != nil {
			beego.Error(err)
		}
		//4.判断权限,如果用户访问的是doc/dwg/xls类，则需要注册和权限大于等于3.
		//或者在管理员设定的用户组内也可以，后续增加
		if rolename > 3 && uname != username {
			port := strconv.Itoa(c.Ctx.Input.Port())
			route = c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			return
		}
	case ".pdf", ".mp4":
		//2.如果登录或ip在允许范围内，进行访问权限检查
		uname, role, _ := checkRoleread(c.Ctx) //login里的
		rolename, _ = strconv.Atoi(role)
		c.Data["Uname"] = uname
		//3.由路径查询数据库中的用户名
		username, err := models.GetattachAuthor(route)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(username)
		//4.判断权限,如果用户访问的是pdf/jpg/diary类，则需要注册和权限4.
		if rolename > 4 && uname != username {
			port := strconv.Itoa(c.Ctx.Input.Port())
			route = c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			return
		}
		// case "jpg" || "png" || "tif" || "gif"jpeg",  "bmp"
	}
	// if filetype == "doc" || filetype == "dwg" || filetype == "xls" {
	// } else if filetype == "pdf" || filetype == "diary" || filetype == "jpg" || filetype == "tif" {
	// }
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

func (c *AttachController) DownloadAttachment() {
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	//4.取得客户端用户名
	// var uname string
	// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
	// v := sess.Get("uname")
	// var role, userrole int
	// if v != nil {
	// 	uname = v.(string)
	// 	c.Data["Uname"] = v.(string)
	// 	user, err := models.GetUserByUsername(uname)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	userrole = user.Role
	// } else {
	// 	userrole = 5
	// }

	// iprole := Getiprole(c.Ctx.Input.IP())
	// if iprole <= userrole {
	// 	role = iprole
	// } else {
	// 	role = userrole
	// }

	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:]) //  attachment/SL2016测试添加成果/A/FB/1/Your First Meteor Application.pdf
	if err != nil {
		beego.Error(err)
	}
	fileext := path.Ext(filePath)
	switch fileext {
	case ".pdf", ".PDF":
		// if role <= 3 {
		http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
		// } else {
		// 	route := c.Ctx.Request.URL.String()
		// 	c.Data["Url"] = route
		// 	c.Redirect("/roleerr?url="+route, 302)
		// 	return
		// }
	default:
		// if role <= 2 {
		http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
		// } else {
		// 	route := c.Ctx.Request.URL.String()
		// 	c.Data["Url"] = route
		// 	c.Redirect("/roleerr?url="+route, 302)
		// 	return
		// }
	}
}
