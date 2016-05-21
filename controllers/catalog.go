package controllers

import (
	// 	// "fmt"
	"github.com/astaxie/beego"
	// 	// "github.com/astaxie/beego/utils/pagination"
	"github.com/tealeg/xlsx"
	// 	// "os"
	// 	"path"
	// 	// "path/filepath"
	// "quick/models"
	m "quick/models"
	"time"
	// 	// "regexp"
	"strconv"
	// 	"strings"
)

type CatalogController struct {
	beego.Controller
}

// 显示所有目录
func (c *CatalogController) Get() {
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, _, _ := checkRoleread(c.Ctx) //login里的
	c.Data["Uname"] = uname
	// rolename, _ = strconv.Atoi(role)
	c.Data["IsCatalog"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)

	c.TplName = "catalog.tpl"
	// cid := c.Input().Get("cid")
	catalogs, err := m.GetAllCatalogs("0") //这里传入空字符串
	if err != nil {
		beego.Error(err.Error)
	} else {
		c.Data["Catalogs"] = catalogs
		c.Data["Length"] = len(catalogs)
	}
}

//显示一个类型Id下的目录
func (c *CatalogController) View() {
	uname, _, _ := checkRoleread(c.Ctx) //login里的
	c.Data["Uname"] = uname
	c.Data["IsCatalog"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)

	c.TplName = "catalog_view.tpl"
	cid := c.Input().Get("id")
	// beego.Info(cid)
	// cid, _ := strconv.ParseInt(c.Input().Get("id"), 10, 64)
	catalogs, err := m.GetAllCatalogs(cid)
	if err != nil {
		beego.Error(err.Error)
	} else {
		c.Data["Catalogs"] = catalogs
		c.Data["Length"] = len(catalogs)
		c.Data["CategoryId"] = cid
	}
}

//更新目录
func (c *CatalogController) Update() {

}

//添加目录

//添加目录的页面
func (c *CatalogController) Add() { //这个作废，用上面的get
	// if !checkAccount(c.Ctx) {
	// 	c.Redirect("/login", 302)
	// 	return
	// }
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	// c.Data["IsCatalog"] = true
	// id := c.Input().Get("id")
	// mid := c.Input().Get("mid")

	c.TplName = "catalog.tpl"

	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
	// c.Data["Category"] = category
	// c.Data["Id"] = id
}

//添加目录，即插入一条目录
func (c *CatalogController) Post() {
	uname, _, _ := checkRolewrite(c.Ctx) //login里的
	c.Data["Uname"] = uname
	// rolename, _ = strconv.Atoi(role)
	// 	if rolename > 2 {
	// 		port := strconv.Itoa(c.Ctx.Input.Port())
	// 			route = c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
	// 			c.Data["Url"] = route
	// 			c.Redirect("/roleerr?url="+route, 302)
	// 			return
	// 	}
	var catalog m.Catalog
	catalog.Tnumber = c.Input().Get("Tnumber")
	catalog.Name = c.Input().Get("Name")
	catalog.Drawn = c.Input().Get("Drawn")
	catalog.Designd = c.Input().Get("Designd")
	catalog.Checked = c.Input().Get("Checked")
	catalog.Emamined = c.Input().Get("Emamined")
	catalog.Verified = c.Input().Get("Verified")
	catalog.Approved = c.Input().Get("Approved")
	catalog.Data = c.Input().Get("Data")
	catalog.DesignStage = c.Input().Get("DesignStage")
	catalog.Section = c.Input().Get("Section")
	catalog.Projec = c.Input().Get("Projec")
	pid := c.Input().Get("ParentId")
	parentid, _ := strconv.ParseInt(pid, 10, 64)
	catalog.ParentId = parentid
	catalog.Created = time.Now()
	catalog.Updated = time.Now()
	catalog.Views = 0
	catalog.Author = uname
	_, err := m.SaveCatalog(catalog)
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok!"
		c.Ctx.WriteString(data)
	}

	// err := models.ModifyCatalog(tid, title, tnumber)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// c.Redirect("/catalog", 302)
}

// func (c *CatalogController) AddCatalog() {
// 	if !checkAccount(c.Ctx) {
// 		c.Redirect("/login", 302)
// 		return
// 	}
// 	//解析表单
// 	tid := c.Input().Get("tid") //教程里漏了这句，导致修改总是变成添加文章
// 	title := c.Input().Get("title")
// 	tnumber := c.Input().Get("title")
// 	// content := c.Input().Get("content")
// 	// category := c.Input().Get("category")
// 	// categoryid := c.Input().Get("categoryid")

// 	_, h, err := c.GetFile("image") //获取上传的文件
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	var attachment string
// 	if h != nil {
// 		//保存附件
// 		attachment = h.Filename
// 		beego.Info(attachment)

// 		// err = c.SaveToFile("attachment", path.Join("attachment", attachment))
// 		// path := c.Input().Get("url")  //存文件的路径
// 		// path = path[3:]
// 		// path = "./attachment" + "/" + h.Filename
// 		// f.Close()   // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
// 		err = c.SaveToFile("image", path.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 	}

// 	if title == "" || tnumber == "" {
// 		//将附件的编号和名称写入数据库
// 		filename1, filename2 := SubStrings(attachment)
// 		title = filename2
// 		tnumber = filename1
// 	}
// 	// var err error
// 	// var tid string //这里是增加的，不知为何教程没有
// 	// ck, err := c.Ctx.Request.Cookie("uname")
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	// uname := ck.Value

// 	if len(tid) == 0 {
// 		_, err = models.AddCatalog(title, tnumber)
// 		beego.Info(attachment)
// 	} else {
// 		err = models.ModifyCatalog(tid, title, tnumber)
// 	}
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	// c.Redirect("/catalog", 302)
// }

// //查看一条目录
// func (c *CatalogController) View() {
// 	c.Data["IsLogin"] = checkAccount(c.Ctx)
// 	c.Data["IsCatalog"] = true

// 	//这里是通过文章的id获得文章及上级目录情况
// 	// catalogproj, err := models.GetCatalogProj(c.Ctx.Input.Param("0"))
// 	// catalogphase, err := models.GetCatalogPhase(c.Ctx.Input.Param("0"))
// 	// catalogspec, err := models.GetCatalogSpec(c.Ctx.Input.Param("0"))
// 	// _, catalogchengguo, err := models.GetCatalogChengguo(c.Ctx.Input.Param("0"))
// 	// if catalogchengguo.Title == "diary" {
// 	// 	c.TplName = "diary_view1.html"
// 	// } else {
// 	// 	c.TplName = "catalog_view.html"
// 	// }

// 	// catalog, attachment, err := models.GetCatalog(c.Ctx.Input.Param("0"))
// 	//articleId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
// 	//id, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
// 	// if err != nil {
// 	// 	beego.Error(err)
// 	// 	c.Redirect("/", 302)
// 	// 	return
// 	// }
// 	// c.Data["CatalogProj"] = catalogproj
// 	// c.Data["CatalogPhase"] = catalogphase
// 	// c.Data["CatalogSpec"] = catalogspec
// 	// c.Data["CatalogChengguo"] = Catalogchengguo
// 	// c.Data["Catalog"] = catalog
// 	// c.Data["Attachment"] = attachment
// 	c.Data["Tid"] = c.Ctx.Input.Param("0") //教程中用的是圆括号，导致错误Catalog.go:52: cannot call non-function c.Controller.Ctx.Input.Params (type map[string]string)
// 	//教程第8章开头有修改
// 	replies, err := models.GetAllReplies(c.Ctx.Input.Param("0"))
// 	if err != nil {
// 		beego.Error(err)
// 		return
// 	}
// 	c.Data["Replies"] = replies
// 	c.Data["IsLogin"] = checkAccount(c.Ctx)
// }

//修改一条目录
func (c *CatalogController) ModifyCatalog() {
	uname, _, _ := checkRolewrite(c.Ctx) //login里的
	c.Data["Uname"] = uname
	var catalog m.Catalog
	catalog.Tnumber = c.Input().Get("Tnumber")
	catalog.Name = c.Input().Get("Name")
	catalog.Drawn = c.Input().Get("Drawn")
	catalog.Designd = c.Input().Get("Designd")
	catalog.Checked = c.Input().Get("Checked")
	catalog.Emamined = c.Input().Get("Emamined")
	catalog.Verified = c.Input().Get("Verified")
	catalog.Approved = c.Input().Get("Approved")
	catalog.Data = c.Input().Get("Data")
	catalog.DesignStage = c.Input().Get("DesignStage")
	catalog.Section = c.Input().Get("Section")
	catalog.Projec = c.Input().Get("Projec")
	cid := c.Input().Get("CatalogId")
	// catalogid, _ := strconv.ParseInt(cid, 10, 64)
	var id string
	if cid != "" {
		id = string(cid[3:len(cid)])
		// beego.Info(id)
	}
	catalog.Created = time.Now()
	catalog.Updated = time.Now()
	catalog.Views = 0
	catalog.Author = uname
	err := m.ModifyCatalog(id, catalog)
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok!"
		c.Ctx.WriteString(data)
	}
}

//删除一条目录
func (c *CatalogController) Delete() {
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, _, _ := checkRolewrite(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	//取得用户名

	// if rolename > 2 && uname != username {
	cid := c.Input().Get("CatalogId")
	var id string
	if cid != "" {
		id = string(cid[3:len(cid)])
		beego.Info(id)
	}
	err := m.DeletCatalog(id)
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok!"
		c.Ctx.WriteString(data)
	}
}

// func (c *CatalogController) DeleteAll() {
// 	if !checkAccount(c.Ctx) {
// 		c.Redirect("/login", 302)
// 		return
// 	}
// 	//解析表单
// 	c.Data["IsLogin"] = checkAccount(c.Ctx)
// 	cid := c.Input().Get("cid")
// 	Catalogid := c.Input().Get("tempstring")
// 	array := strings.Split(Catalogid, ",") //字符串切割 [a b c d e]
// 	for _, v := range array {
// 		err := models.DeletCatalog(v)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 	}
// 	c.Redirect("/category/view_b?id="+cid, 302) //这里增加Catalog
// }

//上传excel文件，导入到数据库
//引用来自category的查看成果类型里的成果
func (c *CatalogController) Import_Xls_Catalog() {
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
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	var uname string
	if v != nil {
		uname = v.(string)
	} else {
		beego.Error(err)
	}
	// ck, err := c.Ctx.Request.Cookie("uname")
	//
	// if err == nil {
	// 	uname = ck.Value
	// }
	// route := "/attachment/" + categoryproj.Number + " " + categoryproj.Title + "/" + categoryphase.Title + "/" + categoryspec.Title + "/" + category + "/" + h.Filename
	//Catalogid := c.Input().Get("Catalogid")
	var catalog m.Catalog
	id1 := c.Input().Get("id")
	cid, _ := strconv.ParseInt(id1, 10, 64)
	catalog.ParentId = cid
	//读出excel内容写入数据库
	// excelFileName := path                    //"/home/tealeg/foo.xlsx"
	xlFile, err := xlsx.OpenFile(path) //excelFileName
	if err != nil {
		beego.Error(err)
	}
	j := 0
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows { //行数,第一行从0开始
			if i != 0 { //忽略第一行
				if len(row.Cells) >= 2 { //总列数，从1开始
					catalog.Tnumber, err = row.Cells[j+1].String() //第一列从0开始
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 3 {
					catalog.Name, err = row.Cells[j+2].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 4 {
					catalog.Drawn, _ = row.Cells[j+3].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 5 {
					catalog.Designd, _ = row.Cells[j+4].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 6 {
					catalog.Checked, _ = row.Cells[j+5].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 7 {
					catalog.Emamined, _ = row.Cells[j+6].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 8 {
					catalog.Verified, _ = row.Cells[j+7].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 9 {
					catalog.Approved, _ = row.Cells[j+8].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 10 {
					catalog.Data, _ = row.Cells[j+9].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 11 {
					catalog.DesignStage, _ = row.Cells[j+10].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 12 {
					catalog.Section, _ = row.Cells[j+11].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 13 {
					catalog.Projec, _ = row.Cells[j+12].String()
					if err != nil {
						beego.Error(err)
					}
				}
				catalog.Created = time.Now()
				catalog.Updated = time.Now()
				catalog.Views = 0
				catalog.Author = uname
				_, err := m.SaveCatalog(catalog)
				if err != nil {
					beego.Error(err)
				}
				// roleid, _ := strconv.ParseInt(row.Cells[j+4].String(), 10, 64)
				// _, err = m.AddRoleUser(roleid, uid)
				// if err != nil {
				// 	beego.Error(err)
				// }
				// }
				// for _, cell := range row.Cells {
				// 	fmt.Printf("%s\n", cell.String())
			}
		}
	}
	c.TplName = "catalog.tpl"
	// c.Redirect("/catalog/view?id="+id1, 302)
}

// PrepareInsert
// 用于一次 prepare 多次 insert 插入，以提高批量插入的速度。
// var users []*User
// ...
// qs := o.QueryTable("user")
// i, _ := qs.PrepareInsert()
// for _, user := range users {
//     id, err := i.Insert(user)
//     if err == nil {
//         ...
//     }
// }
// PREPARE INSERT INTO user (`name`, ...) VALUES (?, ...)
// EXECUTE INSERT INTO user (`name`, ...) VALUES ("slene", ...)
// EXECUTE ...
// ...
// i.Close() // 别忘记关闭 statement
