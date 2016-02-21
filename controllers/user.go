package controllers

import (
	// m "github.com/beego/admin/src/models"
	// "github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
	m "quick/models"
	"strconv"
	"time"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	// page, _ := this.GetInt64("page")
	// page_size, _ := this.GetInt64("rows")
	// sort := this.GetString("sort")
	// order := this.GetString("order")
	// if len(order) > 0 {
	// 	if order == "desc" {
	// 		sort = "-" + sort
	// 	}
	// } else {
	// 	sort = "Id"
	// }
	// 	c.Data["IsCategory"] = true
	// c.TplName = "category.tpl"
	//1.首先判断是否注册
	if !checkAccount(this.Ctx) {
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := this.Ctx.Request.URL.String()
		this.Data["Url"] = route
		this.Redirect("/login?url="+route, 302)
		// c.Redirect("/login", 302)
		return
	}
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		this.Data["Uname"] = v.(string)
	}
	// ck, err := this.Ctx.Request.Cookie("uname")
	// if err == nil {
	// 	this.Data["Uname"] = ck.Value
	// } else {
	// 	beego.Error(err)
	// }
	//2.取得客户端用户名
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value
	//3.取出用户的权限等级
	// category, err := models.GetCategory(id)
	// beego.Info(username)
	//4.取得客户端用户名
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value
	//5.取出用户的权限等级
	role, _ := checkRole(this.Ctx) //login里的
	// beego.Info(role)
	//6.进行逻辑分析：
	rolename, _ := strconv.ParseInt(role, 10, 64)
	// if filetype != "pdf" && filetype != "jpg" && filetype != "diary" {
	if rolename > 1 { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := this.Ctx.Request.URL.String()
		this.Data["Url"] = route
		this.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// }

	this.Data["IsLogin"] = checkAccount(this.Ctx)
	//2.取得客户端用户名
	// ck, err := this.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// } else {
	// 	this.Data["Uname"] = ck.Value
	// }

	users, count := m.Getuserlist(1, 2000, "Id")
	if this.IsAjax() {
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &users}
		this.ServeJSON()
		return
	} else {
		// tree := this.GetTree()
		// this.Data["tree"] = &tree
		this.Data["Users"] = &users
		this.TplName = "user.tpl"
		// if this.GetTemplatetype() != "easyui" {
		// this.Layout = this.GetTemplatetype() + "/public/layout.tpl"
		// }
		// this.TplName = this.GetTemplatetype() + "/rbac/user.tpl"
	}

}

func (this *UserController) View() {
	// c.Data["IsCategory"] = true
	// c.TplName = "category.tpl"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		this.Data["Uname"] = v.(string)
	}
	// ck, err := this.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// } else {
	// 	this.Data["Uname"] = ck.Value
	// }
	// userid, _ := this.GetInt64("Id")
	// id := this.Ctx.Input.Param("0")这里为何无效？？？？这个需要routers中设置AutoRouter
	// beego.Info(id)
	// userid, _ := strconv.ParseInt(id, 10, 64)

	userid, _ := strconv.ParseInt(this.Input().Get("useid"), 10, 64)
	// beego.Info(userid)
	user := m.GetUserByUserId(userid)
	// if this.IsAjax() {
	// users, _ := m.Getuserlist(1, 1000, "Id")
	list, _ := m.GetRoleByUserId(userid)
	// if err != nil {
	// 	beego.Error(err)
	// 	c.Redirect("/", 302)
	// 	return
	// }
	// beego.Info(list[1])
	// for i := 0; i < len(users); i++ {
	// 	for x := 0; x < len(list); x++ {
	// 		if users[i]["Id"] == list[x]["Id"] {
	// 			users[i]["checked"] = 1
	// 		}
	// 	}
	// }
	// if len(users) < 1 {
	// 	users = []orm.Params{}
	// }
	// this.Data["json"] = &map[string]interface{}{"total": count, "rows": &users}
	// this.ServeJSON()
	// return
	// } else {
	this.Data["User"] = user
	this.Data["Role"] = list
	// this.Data["Users"] = &users
	this.TplName = "admin_user_view.tpl"
	// this.TplName = this.GetTemplatetype() + "/rbac/roletouserlist.tpl"
	// }
}

func (this *UserController) AddUser() {
	u := m.User{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		// this.Rsp(false, err.Error())
		beego.Error(err.Error)
		return
	}
	id, err := m.AddUser(&u)
	if err == nil && id > 0 {
		// this.Rsp(true, "Success")
		return
	} else {
		// this.Rsp(false, err.Error())
		beego.Error(err.Error)
		return
	}

}

// func (this *UserController) UpdateUser() {
// 	u := m.User{}
// 	if err := this.ParseForm(&u); err != nil {
// 		//handle error
// 		// this.Rsp(false, err.Error())
// 		beego.Error(err.Error)
// 		return
// 	}
// 	id, err := m.UpdateUser(&u)
// 	if err == nil && id > 0 {
// 		// this.Rsp(true, "Success")
// 		return
// 	} else {
// 		// this.Rsp(false, err.Error())
// 		beego.Error(err.Error)
// 		return
// 	}

// }
func (this *UserController) UpdateUser() {

	userid := this.Input().Get("userid")
	// username := this.Input().Get("username")
	nickname := this.Input().Get("nickname")
	email := this.Input().Get("email")
	Pwd1 := this.Input().Get("password")
	if Pwd1 != "" {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(Pwd1))
		cipherStr := md5Ctx.Sum(nil)
		// fmt.Print(cipherStr)
		// fmt.Print("\n")
		// fmt.Print(hex.EncodeToString(cipherStr))

		password := hex.EncodeToString(cipherStr)
		// user.Lastlogintime = time.Now()
		err := m.UpdateUser(userid, nickname, email, password) //这里修改
		if err != nil {
			beego.Error(err)
		}
		//更新role
		roleid := this.Input().Get("roletitle1")
		if roleid != "" {
			roleid1, _ := strconv.ParseInt(roleid, 10, 64)
			roleid2, _ := strconv.ParseInt(this.Input().Get("roletitle2"), 10, 64)
			userid1, _ := strconv.ParseInt(userid, 10, 64)
			_, err = m.UpdateRoleUser(roleid1, roleid2, userid1)
			if err != nil {
				beego.Error(err)
			}
		}
	} else {
		err := m.UpdateUser(userid, nickname, email, "") //这里修改
		if err != nil {
			beego.Error(err)
		}
		//更新role
		roleid := this.Input().Get("roletitle1")
		if roleid != "" {
			roleid1, _ := strconv.ParseInt(roleid, 10, 64)
			roleid2, _ := strconv.ParseInt(this.Input().Get("roletitle2"), 10, 64)
			userid1, _ := strconv.ParseInt(userid, 10, 64)
			_, err = m.UpdateRoleUser(roleid1, roleid2, userid1)
			if err != nil {
				beego.Error(err)
			}
		}
	}
	this.TplName = "user_view.tpl"
}

func (this *UserController) DelUser() {
	Id, _ := this.GetInt64("userid")
	status, err := m.DelUserById(Id)
	if err == nil && status > 0 {
		// this.Rsp(true, "Success")
		this.Redirect("/user/index", 302)
		return
	} else {
		// this.Rsp(false, err.Error())
		beego.Error(err.Error)
		return
	}
}

func (this *UserController) GetUserByUsername() {
	// 	c.Data["IsCategory"] = true
	// c.TplName = "category.tpl"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		this.Data["Uname"] = v.(string)
	}
	// ck, err := this.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// } else {
	// 	this.Data["Uname"] = ck.Value
	// }
	username := this.Input().Get("username")
	// beego.Info(userid)
	user := m.GetUserByUsername(username)
	list, _, _ := m.GetRoleByUsername(username)
	this.Data["User"] = user
	this.Data["Role"] = list
	this.TplName = "user_view.tpl"
}

//上传excel文件，导入到数据库
//引用来自category的查看成果类型里的成果
func (this *UserController) ImportExcel() {
	//解析表单
	// this.Data["IsLogin"] = checkAccount(this.Ctx)
	// id := c.Input().Get("id")
	// path := c.Input().Get("path")
	// filename := c.Input().Get("filename")

	//获取上传的文件
	_, h, err := this.GetFile("excel")
	if err != nil {
		beego.Error(err)
	}
	// var attachment string
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		// attachment = h.Filename
		// beego.Info(attachment)
		path = ".\\attachment\\" + h.Filename

		// path := c.Input().Get("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = this.SaveToFile("excel", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
	}
	// if title == "" || tnumber == "" {
	// 	//将附件的编号和名称写入数据库
	// 	filename1, filename2 := SubStrings(attachment)
	// 	if filename1 == "" {
	// 		filename1 = filename2 //如果编号为空，则用文件名代替，否则多个编号为空导致存入数据库唯一性检查错误
	// 	}
	// 	title = filename2
	// 	tnumber = filename1
	// }
	// var err error
	// var tid string //这里是增加的，不知为何教程没有
	// path := ".\\attachment\\" + categoryproj.Number + " " + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename
	// ck, err := c.Ctx.Request.Cookie("uname")
	if err != nil {
		beego.Error(err)
	}
	// uname := ck.Value

	// route := "/attachment/" + categoryproj.Number + " " + categoryproj.Title + "/" + categoryphase.Title + "/" + categoryspec.Title + "/" + category + "/" + h.Filename
	//Catalogid := c.Input().Get("Catalogid")
	var user m.User

	//读出excel内容写入数据库
	// excelFileName := path                    //"/home/tealeg/foo.xlsx"
	xlFile, err := xlsx.OpenFile(h.Filename) //excelFileName
	if err != nil {
		beego.Error(err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			// for j := 2; j < 7; j += 5 {
			j := 1
			user.Username = row.Cells[j].String()
			Pwd1 := row.Cells[j+1].String()
			md5Ctx := md5.New()
			md5Ctx.Write([]byte(Pwd1))
			cipherStr := md5Ctx.Sum(nil)
			user.Password = hex.EncodeToString(cipherStr)
			user.Email = row.Cells[j+2].String()
			user.Nickname = row.Cells[j+3].String()
			user.Lastlogintime = time.Now()
			uid, err := m.SaveUser(user)
			roleid, _ := strconv.ParseInt(row.Cells[j+4].String(), 10, 64)
			_, err = m.AddRoleUser(roleid, uid)
			if err != nil {
				beego.Error(err)
			}
			// }
			// for _, cell := range row.Cells {
			// 	fmt.Printf("%s\n", cell.String())

			// }
		}
	}
}
