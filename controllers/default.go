package controllers

import (
	"github.com/astaxie/beego"
	"quick/models"
	//（1）导入session包
	// "encoding/json"
	// "fmt"
	"github.com/astaxie/beego/session"
	// "github.com/bitly/go-simplejson" // for json get
)

//（2）建立一个全局session mananger对象
var globalSessions *session.Manager

//（3）在初始化“全局session mananger对象”
func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	go globalSessions.GC()
	// globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	// go globalSessions.GC()
}

type MainController struct {
	beego.Controller
}

// func (c *MainController) Hello() {
// 	c.Data["Website"] = "My Website"
// 	c.Data["Email"] = "your.email.address@example.com"
// 	c.Data["EmailName"] = "Your Name"
// 	c.Data["Id"] = c.Ctx.Input.Param(":id")
// 	c.TplName = "index.tpl" //"default/hello.tpl"

// }

func (c *MainController) Help() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsHelp"] = true
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {

		// } else {
		// }
		//2.取得客户端用户名
		// ck, err := c.Ctx.Request.Cookie("uname")
		// if err != nil {
		// beego.Error(err)
		// } else {
		c.Data["Uname"] = v.(string) //ck.Value
	}
	c.TplName = "help.html"

}
func (c *MainController) Test() {
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	// c.Data["IsHelp"] = true
	c.TplName = "test.tpl"

}
func (c *MainController) Test1() {
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	// c.Data["IsHelp"] = true
	c.TplName = "test1.tpl"

}
func (c *MainController) Test2() {
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	// c.Data["IsHelp"] = true
	c.TplName = "test2.tpl"

}
func (c *MainController) Get() {
	// w http.ResponseWriter, r *http.Request
	// sess := globalSessions.SessionStart(w, r)
	// defer sess.SessionRelease()
	// username := sess.Get("username")
	// if r.Method == "GET" {
	// 	t, _ := template.ParseFiles("login.gtpl")
	// 	t.Execute(w, nil)
	// } else {
	// 	sess.Set("username", r.Form["username"])
	// }
	//（4）获取当前的请求会话，并返回当前请求会话的对象
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	//（5）根据当前请求对象，设置一个session
	sess.Set("mySession", "qq504284")

	c.Data["Website"] = "广东省水利电力勘测设计研究院■☆●施工预算分院"
	//（6）从session中读取值
	c.Data["Email"] = sess.Get("mySession")

	// c.Data["Website"] = "127.0.0.1:8080/hello"
	// c.Data["Email"] = "astaxie@gmail.com"
	beego.Info(c.Ctx.Input.IP())
	c.Data["IsHome"] = true
	c.TplName = "index.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx) //大小写害死人！IsLogin
	//2.取得客户端用户名
	// sess, _ = globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		c.Data["Uname"] = v.(string) //ck.Value
	}
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// } else {
	// 	c.Data["Uname"] = ck.Value
	// }
	c.Data["Id"] = c.Ctx.Input.Param(":id")
	topics, err := models.GetAllTopics(c.Input().Get("cate"), true)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topics"] = topics
	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Categories"] = categories
	//c.Ctx.Output.Download("database/1.txt", "1.txt")
	//试验控制器数据赋值
	// ss := []string{"a", "b", "c"}
	// c.Data["s"] = ss

	//转换成json数据？
	// beego.Info(uname1[0].Username)
	// b, err := json.Marshal(categories)
	// if err == nil {
	// 	c.Data["json"] = categories
	// 	c.ServeJSON()
	// }

}
func (c *MainController) Post() {

	c.Data["IsHome"] = true
	c.TplName = "index.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx) //大小写害死人！IsLogin
	//2.取得客户端用户名
	// 	type Controller
	// type Controller struct {
	//     // context data
	//     Ctx  *context.Context
	//     Data map[interface{}]interface{}
	//下面的Ctx是因为beego的Controller方法里写好了Ctx  *context.Context

	// 	type Context

	// type Context struct {
	//     Input          *BeegoInput
	//     Output         *BeegoOutput
	//     Request        *http.Request
	//     ResponseWriter *Response
	//     // contains filtered or unexported fields
	// }
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		c.Data["Uname"] = v.(string) //ck.Value
	}
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// } else {
	// 	c.Data["Uname"] = ck.Value
	// }
	c.Data["Id"] = c.Ctx.Input.Param(":id")
	topics, err := models.GetAllTopics(c.Input().Get("cate"), true)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topics"] = topics
	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	// c.Data["Categories"] = categories
	//c.Ctx.Output.Download("database/1.txt", "1.txt")
	//试验控制器数据赋值
	// ss := []string{"a", "b", "c"}
	// c.Data["s"] = ss

	//转换成json数据
	// beego.Info(uname1[0].Username)
	// b, err := json.Marshal(categories)
	if err == nil {
		// c.Ctx.WriteString(string(b))
		c.Data["json"] = categories
		c.ServeJSON()
	}
	// 你自己判断比如
	// if isMobile {
	// 	c.ServeJSON()
	// } else {
	// 	c.TplName = "xxx"
	// }

	// js, err := simplejson.NewJson(b)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(js)

}

//下面这个例题也过时了
// func login(w http.ResponseWriter, r *http.Request) {
// 	sess := globalSessions.SessionStart(w, r)
// 	defer sess.SessionRelease()
// 	username := sess.Get("username")
// 	if r.Method == "GET" {
// 		t, _ := template.ParseFiles("login.gtpl")
// 		t.Execute(w, nil)
// 	} else {
// 		sess.Set("username", r.Form["username"])
// 	}
// }
// 包引用
// import (
//     "encoding/json"
//     "github.com/bitly/go-simplejson" // for json get
// )
// 用于存放数据的结构体
// type MyData struct {
//     Name   string    `json:"item"`
//     Other  float32   `json:"amount"`
// }
// 这里需要注意的就是后面单引号中的内容。

// `json:"item"`
// 这个的作用，就是Name字段在从结构体实例编码到JSON数据格式的时候，使用item作为名字。算是一种重命名的方式吧。

// 编码JSON
// var detail MyData

// detail.Name  = "1"
// detail.Other = "2"

// body, err := json.Marshal(detail)
// if err != nil {
//     panic(err.Error())
// }
// 我们使用Golang自带的encoding/json包对结构体进行编码到JSON数据。

// json.Marshal(...)
// JSON解码
// 由于Golang自带的json包处理解码的过程较为复杂，所以这里使用一个第三方的包simplejson进行json数据的解码操作。

// js, err := simplejson.NewJson(body)
// if err != nil {
//     panic(err.Error())
// }

// fmt.Println(js)
