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
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		// c.Redirect("/login", 302)
		return
	}
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		c.Data["Uname"] = v.(string) //uname := v.(string)
	}
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err == nil {
	//
	// } else {
	// 	beego.Error(err)
	// }

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
	//1.首先判断是否注册
	// if !checkAccount(c.Ctx) {
	// 	port := strconv.Itoa(c.Ctx.Input.Port())
	// 	route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/login?url="+route, 302)
	// 	// c.Redirect("/login", 302)
	// 	return
	// }
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		c.Data["Uname"] = v.(string)
	}
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err == nil {
	// 	 = ck.Value
	// } else {
	// 	beego.Error(err)
	// }

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

// //添加目录，即插入一条目录
// func (c *CatalogController) Post() {
// 	if !checkAccount(c.Ctx) {
// 		c.Redirect("/login", 302)
// 		return
// 	}
// 	//解析表单
// 	tid := c.Input().Get("tid") //教程里漏了这句，导致修改总是变成添加文章
// 	beego.Info(tid)
// 	title := c.Input().Get("title")
// 	tnumber := c.Input().Get("tnumber")
// 	beego.Info(tnumber)
// 	// content := c.Input().Get("content")
// 	// category := c.Input().Get("category")
// 	// categoryid := c.Input().Get("categoryid")

// 	err := models.ModifyCatalog(tid, title, tnumber)

// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	c.Redirect("/catalog", 302)
// }

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

// //修改一条目录
// func (c *CatalogController) ModifyCatalog() { //一对多模式,向文章中追加附件
// 	//解析表单
// 	tid := c.Input().Get("tid") //教程里漏了这句，导致修改总是变成添加文章
// 	// title := c.Input().Get("title")
// 	// tnumber := c.Input().Get("tnumber")
// 	// content := c.Input().Get("content")
// 	category := c.Input().Get("category")
// 	categoryid := c.Input().Get("categoryid")

// 	//获取文件保存路径，有了categoryid可以求出整个路径
// 	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
// 	categoryproj, err := models.GetCategoryProj(categoryid)
// 	categoryphase, err := models.GetCategoryPhase(categoryid)
// 	categoryspec, err := models.GetCategorySpec(categoryid)
// 	// category, err := models.GetCategory(categoryid)
// 	if err != nil {
// 		beego.Error(err)
// 		// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
// 		return
// 	}
// 	//获取上传的文件
// 	_, h, err := c.GetFile("image")
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	var attachment string
// 	var path string
// 	var filesize int64
// 	if h != nil {
// 		//保存附件
// 		attachment = h.Filename
// 		beego.Info(attachment)
// 		path = ".\\attachment\\" + categoryproj.Number + " " + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename

// 		// path := c.Input().Get("url")  //存文件的路径
// 		// path = path[3:]
// 		// path = "./attachment" + "/" + h.Filename
// 		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
// 		err = c.SaveToFile("image", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		filesize, _ = FileSize(path)
// 		filesize = filesize / 1000.0
// 	}

// 	route := "/attachment/" + categoryproj.Number + " " + categoryproj.Title + "/" + categoryphase.Title + "/" + categoryspec.Title + "/" + category + "/" + h.Filename

// 	size := strconv.FormatInt(filesize, 10)
// 	err = models.AddAttachment(attachment, size, path, route, tid)
// 	beego.Info(attachment)
// 	// } else {
// 	// err = models.ModifyCatalog(tid, title, tnumber, category, categoryid, content, attachment)

// 	// }
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	// c.Redirect("/Catalog", 302)
// }

// //删除一条目录
// func (c *CatalogController) Delete() { //应该显示警告
// 	if !checkAccount(c.Ctx) {
// 		c.Redirect("/login", 302)
// 		return
// 	}
// 	err := models.DeletCatalog(c.Input().Get("tid")) //(c.Ctx.Input.Param("0"))
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	c.Redirect("/Catalog", 302) //这里增加Catalog
// }

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
	for _, sheet := range xlFile.Sheets {
		// for i := 1; i <= len(xlFile.Sheets); i++ {
		for _, row := range sheet.Rows {
			// for j := 2; j < 7; j += 5 {
			j := 1
			catalog.Tnumber, _ = row.Cells[j].String()
			catalog.Name, _ = row.Cells[j+1].String()
			catalog.Drawn, _ = row.Cells[j+2].String()
			catalog.Designd, _ = row.Cells[j+3].String()
			catalog.Checked, _ = row.Cells[j+4].String()
			catalog.Emamined, _ = row.Cells[j+5].String()
			catalog.Verified, _ = row.Cells[j+6].String()
			catalog.Approved, _ = row.Cells[j+7].String()
			catalog.Data, _ = row.Cells[j+8].String()
			catalog.DesignStage, _ = row.Cells[j+9].String()
			catalog.Section, _ = row.Cells[j+10].String()
			catalog.Projec, _ = row.Cells[j+11].String()
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

			// }
		}
	}
	c.TplName = "catalog.tpl"
	c.Redirect("/catalog/view?id="+id1, 302)
}
