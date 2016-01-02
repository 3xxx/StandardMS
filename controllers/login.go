package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	// "net/url"
	"quick/models"
	// "strconv"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	isExit := c.Input().Get("exit") == "true"
	url1 := c.Input().Get("url") //这里不支持这样的url，http://192.168.9.13/login?url=/topic/add?id=955&mid=3
	url2 := c.Input().Get("mid")
	var url string
	if url2 == "" {
		url = url1
	} else {
		url = url1 + "&mid=" + url2
	}

	c.Data["Url"] = url
	// beego.Info(isExit)
	if isExit {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Redirect("/", 301)
		return
	}

	//	c.Data["Website"] = "My Website"
	//	c.Data["Email"] = "your.email.address@example.com"
	//	c.Data["EmailName"] = "Your Name"
	//	c.Data["Id"] = c.Ctx.Input.Param(":id")
	c.TplNames = "login.html"
}

func (c *LoginController) Loginerr() {
	url := c.Input().Get("url")
	// port := strconv.Itoa(c.Ctx.Input.Port())
	// url := c.Ctx.Input.Site() + ":" + port + c.Ctx.Request.URL.String()
	c.Data["Url"] = url
	// beego.Info(url)
	c.TplNames = "loginerr.html"
}

func (c *LoginController) Post() {
	// uname := c.Input().Get("uname")
	url := c.Input().Get("returnUrl")
	// beego.Info(url)
	// pwd := c.Input().Get("pwd")
	// autoLogin := c.Input().Get("autoLogin") == "on"

	// if beego.AppConfig.String("uname") == uname &&
	// 	beego.AppConfig.String("pwd") == pwd {
	// 	maxAge := 0
	// 	if autoLogin {
	// 		maxAge = 1<<31 - 1
	// 	}
	// 	c.Ctx.SetCookie("uname", uname, maxAge, "/")
	// 	c.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	// 	c.Redirect("/", 301)
	// } else {
	// 	c.Redirect("/login", 302)
	// }
	// return
	var user models.User
	user.Username = c.Input().Get("uname")
	Pwd1 := c.Input().Get("pwd")
	autoLogin := c.Input().Get("autoLogin") == "on"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	fmt.Print(cipherStr)
	fmt.Print("\n")
	fmt.Print(hex.EncodeToString(cipherStr))

	user.Password = hex.EncodeToString(cipherStr)
	err := models.ValidateUser(user)
	if err == nil {

		// if beego.AppConfig.String("uname") == uname &&
		// 	beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		c.Ctx.SetCookie("uname", user.Username, maxAge, "/")
		c.Ctx.SetCookie("pwd", user.Password, maxAge, "/")
		//更新user表的lastlogintime
		models.UpdateUserlastlogintime(user.Username)
		if url != "" {
			c.Redirect(url, 301)
			// beego.Info(url)
		} else {
			c.Redirect("/", 301)
		}
	} else {
		// port := strconv.Itoa(c.Ctx.Input.Port())
		// route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.Url()
		// c.Data["Url"] = route
		// c.Redirect("/login?url="+route, 302)
		c.Redirect("/loginerr?url="+url, 302)
	}
	return

	// sess := index.StartSession()
	// var user models.User
	// inputs := index.Input()
	//fmt.Println(inputs)
	// user.Username = c.Input().Get("uname")
	// Pwd1 := c.Input().Get("pwd")

	// md5Ctx := md5.New()
	// md5Ctx.Write([]byte(Pwd1))
	// cipherStr := md5Ctx.Sum(nil)
	// fmt.Print(cipherStr)
	// fmt.Print("\n")
	// fmt.Print(hex.EncodeToString(cipherStr))

	// user.Pwd = hex.EncodeToString(cipherStr)
	// err := models.ValidateUser(user)
	// if err == nil {
	// 	sess.Set("username", user.Username)
	// 	fmt.Println("username:", sess.Get("username"))
	// 	index.TplNames = "success.tpl"
	// } else {
	// 	fmt.Println(err)
	// 	index.TplNames = "error.tpl"
	// }
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	var user models.User
	user.Username = ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	user.Password = ck.Value
	err = models.ValidateUser(user)
	if err == nil {
		return true
	} else {
		return false
	}
	// return beego.AppConfig.String("uname") == uname &&
	// 	beego.AppConfig.String("pwd") == pwd
}

func checkRole(ctx *context.Context) (role string, err error) { //这里返回用户的role
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return "", err
	}
	var user models.User
	user.Username = ck.Value
	var roles []*models.Role
	roles, _ = models.GetRoleByUsername(user.Username)
	if err == nil {
		return roles[0].Name, err //这里修改
	} else {
		return "", err
	}
}

// func checkRole(ctx *context.Context) (roles []*models.Role, err error) {
// 	ck, err := ctx.Request.Cookie("uname")
// 	if err != nil {
// 		return roles, err
// 	}
// 	var user models.User
// 	user.Username = ck.Value

// 	roles, _ = models.GetRoleByUsername(user.Username)
// 	if err == nil {
// 		return roles, err
// 	} else {
// 		return roles, err
// 	}
// }

// func GetRoleByUserId(userid int64) (roles []*Role, count int64) { //*Topic, []*Attachment, error
// 	roles = make([]*Role, 0)
// 	o := orm.NewOrm()
// 	// role := new(Role)
// 	count, _ = o.QueryTable("role").Filter("Users__User__Id", userid).All(&roles)
// 	return roles, count
// 	// 通过 post title 查询这个 post 有哪些 tag
// 	// var tags []*Tag
// 	// num, err := dORM.QueryTable("tag").Filter("Posts__Post__Title", "Introduce Beego ORM").All(&tags)

// }
