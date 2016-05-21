package controllers

import (
	"crypto/md5"
	"encoding/hex"
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	// "net/url"
	"quick/models"
	"strconv"
	// "github.com/astaxie/beego/session"
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
	// logout user
	// func LogoutUser(ctx *context.Context) {
	// 	DeleteRememberCookie(ctx)
	// 	ctx.Input.CruSession.Delete("auth_user_id")
	// 	ctx.Input.CruSession.Flush()
	// 	beego.GlobalSessions.SessionDestroy(ctx.ResponseWriter, ctx.Request)
	// }
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	if isExit {
		// c.Ctx.SetCookie("uname", "", -1, "/")
		// c.Ctx.SetCookie("pwd", "", -1, "/")
		// c.DelSession("gosessionid")
		// c.DelSession("gosessionid")
		// c.DelSession("admin")
		// c.DelSession("uname")//这个不行
		// c.DestroySession()
		// c.Ctx.Input.CruSession.Delete("gosessionid")这句与上面一句重复
		// c.Ctx.Input.CruSession.Flush()
		// beego.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
		sess.Delete("uname") //这个可行。
		// m.DestroySession()
		// beego.Info(sess.SessionID())
		// sess.SessionDestroy(sess.SessionID())
		// c.DestroySession()
		// sess.Flush()//这个不灵
		c.Redirect("/", 301)
		return
	}
	// https://github.com/astaxie/beego/issues/1196
	// You can start session after user login , destory session after user logout.
	// e.g:
	// func (m *MainController) Login() {
	//     // Save user info after user login.
	//     m.SetSession("name", "ysqi")
	//     m.SetSession("email", "devysq@gmail.com")
	// }

	// func (m *MainController) Logout() {
	//     // Destory Session after user logout.
	//     m.DestroySession()
	// }
	// If you only use beego session model.
	// e.g:
	// func login(w http.ResponseWriter, r *http.Request) {
	//     sess, _ := globalSessions.SessionStart(w, r)
	//     defer sess.SessionRelease(w)
	//     username := sess.Get("username")
	//     if r.Method == "GET" {
	//         t, _ := template.ParseFiles("login.gtpl")
	//         t.Execute(w, nil)
	//     } else {
	//         sess.Set("username", r.Form["username"])
	//     }
	// }
	// func logout(w http.ResponseWriter, r *http.Request) {
	//     sess, _ := globalSessions.SessionStart(w, r)
	//     defer sess.SessionRelease(w)
	//     sess.SessionDestroy(w,r)
	// }

	//	c.Data["Website"] = "My Website"
	//	c.Data["Email"] = "your.email.address@example.com"
	//	c.Data["EmailName"] = "Your Name"
	//	c.Data["Id"] = c.Ctx.Input.Param(":id")
	c.TplName = "login.html"
}

func (c *LoginController) Loginerr() {
	// url := c.Input().Get("url")
	url1 := c.Input().Get("url") //这里不支持这样的url，http://192.168.9.13/login?url=/topic/add?id=955&mid=3
	url2 := c.Input().Get("mid")
	var url string
	if url2 == "" {
		url = url1
	} else {
		url = url1 + "&mid=" + url2
	}
	c.Data["Url"] = url
	// beego.Info(url)
	c.TplName = "loginerr.html"
}

func (c *LoginController) Post() {
	// uname := c.Input().Get("uname")
	url := c.Input().Get("returnUrl")

	//（4）获取当前的请求会话，并返回当前请求会话的对象
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	//（5）根据当前请求对象，设置一个session
	// sess.Set("mySession", "qq504284")
	// c.Data["Website"] = "广东省水利电力勘测设计研究院■☆●施工预算分院"
	//（6）从session中读取值
	// c.Data["Email"] = sess.Get("mySession")

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
	// autoLogin := c.Input().Get("autoLogin") == "on"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	// fmt.Print(cipherStr)
	// fmt.Print("\n")
	// fmt.Print(hex.EncodeToString(cipherStr))

	user.Password = hex.EncodeToString(cipherStr)
	err := models.ValidateUser(user)
	if err == nil {
		// if beego.AppConfig.String("uname") == uname &&
		// 	beego.AppConfig.String("pwd") == pwd {
		// maxAge := 0
		// if autoLogin {
		// 	maxAge = 1<<31 - 1
		// }
		// c.Ctx.SetCookie("uname", user.Username, maxAge, "/")
		sess.Set("uname", user.Username)
		sess.Set("pwd", user.Password)
		// beego.Info(sess.Get("uname"))
		// c.Ctx.SetCookie("pwd", user.Password, maxAge, "/")

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
		// route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
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
	// 	index.TplName = "success.tpl"
	// } else {
	// 	fmt.Println(err)
	// 	index.TplName = "error.tpl"
	// }
}

//检查是否登录或ip在预设允许范围内
func checkAccount(ctx *context.Context) bool {
	var user models.User
	//（4）获取当前的请求会话，并返回当前请求会话的对象
	//但是我还是建议大家采用 SetSession、GetSession、DelSession 三个方法来操作，避免自己在操作的过程中资源没释放的问题
	sess, _ := globalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	defer sess.SessionRelease(ctx.ResponseWriter)
	v := sess.Get("uname")
	if v == nil {
		role1 := Getiprole(ctx.Input.IP())
		if role1 != 0 {
			return true
		} else {
			return false
		}
	} else {
		user.Username = v.(string)
		v = sess.Get("pwd")        //没必要检查密码吧，因为只有登录了才产生session，才能获取用户名
		user.Password = v.(string) //ck.Value
		err := models.ValidateUser(user)
		if err == nil {
			return true
		} else {
			return false
		}
	}
}

//访问（读取）权限检查
func checkRoleread(ctx *context.Context) (uname, role string, err error) { //这里返回用户的role
	var user models.User
	var roles []*models.Role
	//（4）获取当前的请求会话，并返回当前请求会话的对象
	sess, _ := globalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	defer sess.SessionRelease(ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		user.Username = v.(string) //ck.Value
		roles, _, err = models.GetRoleByUsername(user.Username)
		if err == nil {
			return v.(string), roles[0].Title, err //这里修改Name改为title就对了
		} else {
			return v.(string), "5", err
		}
	} else {
		role1 := Getiprole(ctx.Input.IP())
		if role1 != 0 {
			return ctx.Input.IP(), strconv.Itoa(role1), nil
		} else {
			return ctx.Input.IP(), "5", nil
		}

	}
}

//写权限检查——只能是登录用户
func checkRolewrite(ctx *context.Context) (uname, role string, err error) { //这里返回用户的role
	var user models.User
	var roles []*models.Role
	//（4）获取当前的请求会话，并返回当前请求会话的对象
	sess, _ := globalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	defer sess.SessionRelease(ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		user.Username = v.(string) //ck.Value
		roles, _, err = models.GetRoleByUsername(user.Username)
		if err == nil {
			return v.(string), roles[0].Title, err //这里修改Name改为title就对了
		} else {
			return v.(string), "5", err
		}
	} else {
		return "", "5", err
	}
	// ck, err := ctx.Request.Cookie("uname")
	// if err != nil {
	// }
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
