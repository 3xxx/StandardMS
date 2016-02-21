package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
	m "quick/models"
	"strconv"
	"time"
)

type TagController struct {
	beego.Controller
}

func (this *TagController) Index() {
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
		this.Data["Users"] = &users
		this.TplName = "user.tpl"
	}
}

func (this *TagController) View() {
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
	userid, _ := strconv.ParseInt(this.Input().Get("useid"), 10, 64)
	user := m.GetUserByUserId(userid)
	list, _ := m.GetRoleByUserId(userid)
	this.Data["User"] = user
	this.Data["Role"] = list
	this.TplName = "admin_user_view.tpl"
}

func (this *TagController) AddUser() {
	u := m.User{}
	if err := this.ParseForm(&u); err != nil {
		beego.Error(err.Error)
		return
	}
	id, err := m.AddUser(&u)
	if err == nil && id > 0 {
		return
	} else {
		beego.Error(err.Error)
		return
	}

}
func (this *TagController) UpdateUser() {
	userid := this.Input().Get("userid")
	nickname := this.Input().Get("nickname")
	email := this.Input().Get("email")
	Pwd1 := this.Input().Get("password")
	if Pwd1 != "" {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(Pwd1))
		cipherStr := md5Ctx.Sum(nil)
		password := hex.EncodeToString(cipherStr)

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

func (this *TagController) DelUser() {
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

func (this *TagController) GetUserByUsername() {
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
func (this *TagController) ImportExcel() {
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

	if err != nil {
		beego.Error(err)
	}

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
			// for _, cell := range row.Cells {
			// 	fmt.Printf("%s\n", cell.String())
		}
	}
}
