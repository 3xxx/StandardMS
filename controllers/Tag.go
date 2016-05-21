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

func (c *TagController) Index() {
	var rolename int
	var uname string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, role, _ := checkRoleread(c.Ctx) //login里的
	if role != "0" {
		rolename, _ = strconv.Atoi(role)
		c.Data["Uname"] = uname
	} else {
		port := strconv.Itoa(c.Ctx.Input.Port())
		route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		return
	}
	if rolename > 1 { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// }

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//2.取得客户端用户名
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// } else {
	// 	c.Data["Uname"] = ck.Value
	// }
	users, count := m.Getuserlist(1, 2000, "Id")
	if c.IsAjax() {
		c.Data["json"] = &map[string]interface{}{"total": count, "rows": &users}
		c.ServeJSON()
		return
	} else {
		c.Data["Users"] = &users
		c.TplName = "user.tpl"
	}
}

func (c *TagController) View() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	// var rolename int
	var uname string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, role, _ := checkRoleread(c.Ctx) //login里的
	if role != "0" {
		// rolename, _ = strconv.Atoi(role)
		c.Data["Uname"] = uname
	} else {
		port := strconv.Itoa(c.Ctx.Input.Port())
		route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		return
	}
	userid, _ := strconv.ParseInt(c.Input().Get("useid"), 10, 64)
	user := m.GetUserByUserId(userid)
	list, _ := m.GetRoleByUserId(userid)
	c.Data["User"] = user
	c.Data["Role"] = list
	c.TplName = "admin_user_view.tpl"
}

func (c *TagController) AddUser() {
	u := m.User{}
	if err := c.ParseForm(&u); err != nil {
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
func (c *TagController) UpdateUser() {
	userid := c.Input().Get("userid")
	nickname := c.Input().Get("nickname")
	email := c.Input().Get("email")
	Pwd1 := c.Input().Get("password")
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
		roleid := c.Input().Get("roletitle1")
		if roleid != "" {
			roleid1, _ := strconv.ParseInt(roleid, 10, 64)
			roleid2, _ := strconv.ParseInt(c.Input().Get("roletitle2"), 10, 64)
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
		roleid := c.Input().Get("roletitle1")
		if roleid != "" {
			roleid1, _ := strconv.ParseInt(roleid, 10, 64)
			roleid2, _ := strconv.ParseInt(c.Input().Get("roletitle2"), 10, 64)
			userid1, _ := strconv.ParseInt(userid, 10, 64)
			_, err = m.UpdateRoleUser(roleid1, roleid2, userid1)
			if err != nil {
				beego.Error(err)
			}
		}
	}
	c.TplName = "user_view.tpl"
}

func (c *TagController) DelUser() {
	Id, _ := c.GetInt64("userid")
	status, err := m.DelUserById(Id)
	if err == nil && status > 0 {
		// c.Rsp(true, "Success")
		c.Redirect("/user/index", 302)
		return
	} else {
		// c.Rsp(false, err.Error())
		beego.Error(err.Error)
		return
	}
}

func (c *TagController) GetUserByUsername() {
	// 	c.Data["IsCategory"] = true
	// c.TplName = "category.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	// var rolename int
	var uname string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, role, _ := checkRoleread(c.Ctx) //login里的
	if role != "0" {
		// rolename, _ = strconv.Atoi(role)
		c.Data["Uname"] = uname
	} else {
		port := strconv.Itoa(c.Ctx.Input.Port())
		route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		return
	}
	username := c.Input().Get("username")
	// beego.Info(userid)
	user := m.GetUserByUsername(username)
	list, _, _ := m.GetRoleByUsername(username)
	c.Data["User"] = user
	c.Data["Role"] = list
	c.TplName = "user_view.tpl"
}

//上传excel文件，导入到数据库
//引用来自category的查看成果类型里的成果
func (c *TagController) ImportExcel() {
	//解析表单
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	// id := c.Input().Get("id")
	// path := c.Input().Get("path")
	// filename := c.Input().Get("filename")

	//获取上传的文件
	_, h, err := c.GetFile("excel")
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
		err = c.SaveToFile("excel", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
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
			user.Username, _ = row.Cells[j].String()
			Pwd1, _ := row.Cells[j+1].String()
			md5Ctx := md5.New()
			md5Ctx.Write([]byte(Pwd1))
			cipherStr := md5Ctx.Sum(nil)
			user.Password = hex.EncodeToString(cipherStr)
			user.Email, _ = row.Cells[j+2].String()
			user.Nickname, _ = row.Cells[j+3].String()
			user.Lastlogintime = time.Now()
			uid, err := m.SaveUser(user)
			role, _ := row.Cells[j+4].String()
			roleid, _ := strconv.ParseInt(role, 10, 64)
			_, err = m.AddRoleUser(roleid, uid)
			if err != nil {
				beego.Error(err)
			}
			// for _, cell := range row.Cells {
			// 	fmt.Printf("%s\n", cell.String())
		}
	}
}
