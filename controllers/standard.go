package controllers

import (
	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
	"os"
	"quick/models"
	"time"
)

type StandardController struct {
	beego.Controller
}

type Standardmore struct {
	Id            int64
	Number        string //`orm:"unique"`
	Title         string
	Uname         string //换成用户名
	CategoryName  int64  //换成规范类别GB、DL……
	Content       string `orm:"sie(5000)"`
	Route         string
	Created       time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"index","auto_now;type(datetime)"`
	Views         int64     `orm:"index"`
	LibraryNumber string    //规范有效版本库中的编号
	LibraryTitle  string
	LiNumber      string //完整编号
}

func (c *StandardController) Get() { //
	c.Data["IsStandard"] = true //这里修改到ListAllPosts()
	c.TplName = "standard.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		c.Data["Uname"] = v.(string)
	}
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err == nil {
	// 	c.Data["Uname"] = ck.Value
	// } else {
	// 	beego.Error(err)
	// }

	// beego.Info(ck.Value)
	// uname := ck.Value

	topics, err := models.GetAllTopics("", false) //这里传入空字符串
	if err != nil {
		beego.Error(err.Error)
	} else {
		c.Data["Topics"] = topics
		c.Data["Length"] = len(topics)
	}
	//var err error
	//	c.Data["Topic"], err = models.GetAllTopics()
	//	if err != nil {
	//		beego.Error(err)
	//	}
}

func (c *StandardController) Index() { //
	c.Data["IsStandard"] = true //
	c.TplName = "standard.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		c.Data["Uname"] = v.(string)
	}

	topics, err := models.GetAllTopics("", false) //这里传入空字符串
	if err != nil {
		beego.Error(err.Error)
	} else {
		c.Data["Topics"] = topics
		c.Data["Length"] = len(topics)
	}
}

//搜索规范或者图集的名称或编号
func (c *StandardController) Search() { //search用的是post方法
	name := c.Input().Get("name")
	c.Data["IsStandard"] = true //
	c.TplName = "standard.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		c.Data["Uname"] = v.(string)
	}
	//搜索名称
	Results1, err := models.SearchStandardsName(name, false)
	if err != nil {
		beego.Error(err.Error)
	}
	// beego.Info(Results1[0].Title)
	// beego.Info(Results1[1].Title)
	//搜索编号
	Results2, err := models.SearchStandardsNumber(name, false)
	if err != nil {
		beego.Error(err.Error)
	}
	// Standards := make([]*Standard, 0)
	Results1 = append(Results1, Results2...)

	//由categoryid查categoryname

	aa := make([]Standardmore, len(Results1))
	//由standardnumber查librarynumber
	for i, v := range Results1 {
		//由userid查username
		user := models.GetUserByUserId(v.Uid)
		//由standardnumber正则得到编号50268和分类GB
		Category, Number := SplitStandardFileNumber(v.Number)
		//由分类和编号查有效版本库中的编号
		library, err := models.SearchLiabraryNumber(Category, Number)
		if err != nil {
			beego.Error(err.Error)
		}
		aa[i].Id = v.Id
		aa[i].Number = v.Number //`orm:"unique"`
		aa[i].Title = v.Title
		aa[i].Uname = user.Username //换成用户名
		// CategoryName   //换成规范类别GB、DL……
		// Content
		aa[i].Route = v.Route
		aa[i].Created = v.Created
		aa[i].Updated = v.Updated
		aa[i].Views = v.Views
		if library != nil {
			aa[i].LibraryNumber = library.Number //规范有效版本库中的编号
			aa[i].LibraryTitle = library.Title
			aa[i].LiNumber = library.LiNumber //完整编号
		} else {
			aa[i].LibraryNumber = "no LibraryNumber find!"
		}
	}

	c.Data["json"] = aa //这里必须要是c.Data["json"]，其他c.Data["Data"]不行
	c.ServeJSON()
}

//上传excel文件，导入到数据库
//引用来自category的查看成果类型里的成果
func (c *StandardController) ImportExcel() {
	//获取上传的文件
	_, h, err := c.GetFile("excel")
	if err != nil {
		beego.Error(err)
	}
	// var attachment string
	var path string
	// var filesize int64
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
	}
	if err != nil {
		beego.Error(err)
	}
	var standard models.Standard
	//读出excel内容写入数据库
	// excelFileName := path                    //"/home/tealeg/foo.xlsx"
	xlFile, err := xlsx.OpenFile(path) //excelFileName
	if err != nil {
		beego.Error(err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			// 这里要判断单元格列数，如果超过单元格使用范围的列数，则出错for j := 2; j < 7; j += 5 {
			j := 0
			standard.Number, err = row.Cells[j].String()
			standard.Title, err = row.Cells[j+1].String()
			// Uname, err := row.Cells[j+2].String()
			// user := models.GetUserByUsername(Uname)
			// standard.Uid = user.Id
			// Category, err := row.Cells[j+3].String()
			// category, _ := models.GetCategoryName(Category)
			// standard.CategoryId = category.Id
			standard.Created = time.Now()
			standard.Updated = time.Now()
			standard.Content, err = row.Cells[j+4].String()
			standard.Route, err = row.Cells[j+5].String()
			_, err = models.SaveStandard(standard)

			if err != nil {
				beego.Error(err)
			}
			// }
			// for _, cell := range row.Cells {这里要继续循环cells，不能为空，即超出单元格使用范围
			// 	fmt.Printf("%s\n", cell.String())

			// }
		}
	}
	c.TplName = "standard.tpl"
	c.Redirect("/standard", 302)
}

//上传excel文件，导入到数据库
//引用来自category的查看成果类型里的成果
func (c *StandardController) ImportLibrary() {
	//获取上传的文件
	_, h, err := c.GetFile("excel")
	if err != nil {
		beego.Error(err)
	}
	// var attachment string
	var path string
	// var filesize int64
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
	}
	if err != nil {
		beego.Error(err)
	}
	var library models.Library
	//读出excel内容写入数据库
	// excelFileName := path                    //"/home/tealeg/foo.xlsx"
	xlFile, err := xlsx.OpenFile(path) //excelFileName
	if err != nil {
		beego.Error(err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			// 这里要判断单元格列数，如果超过单元格使用范围的列数，则出错for j := 2; j < 7; j += 5 {
			j := 0
			library.Number, err = row.Cells[j].String()
			library.Title, err = row.Cells[j+1].String()
			library.Category, err = row.Cells[j+2].String()
			library.LiNumber, err = row.Cells[j+3].String()
			library.Created = time.Now()
			library.Updated = time.Now()
			_, err = models.SaveLibrary(library)
			if err != nil {
				beego.Error(err)
			}
			// }
			// for _, cell := range row.Cells {这里要继续循环cells，不能为空，即超出单元格使用范围
			// 	fmt.Printf("%s\n", cell.String())
			// }
		}
	}
	c.TplName = "standard.tpl"
	c.Redirect("/standard", 302)
}

func (c *StandardController) Standard_one_addbaidu() { //一对一模式
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	_, fileNumber, fileName, _, category := SplitStandardName(h.Filename)
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		if category != "" {
			err := os.MkdirAll(".\\attachment\\Standard\\"+category, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
			if err != nil {
				beego.Error(err)
			}
		}
		path = ".\\attachment\\Standard\\" + category + "\\" + h.Filename
		// path := c.Input().Get("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()   // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
	}
	if fileName == "" {
		fileName = fileNumber
	}
	//获取用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	// v := sess.Get("uname")
	// uname := v.(string)

	var standard models.Standard

	standard.Number = fileNumber
	standard.Title = fileName

	//这里增加用户id和CategoryId

	standard.Created = time.Now()
	standard.Updated = time.Now()
	standard.Route = "/attachment/standard/" + category + "/" + h.Filename
	_, err = models.SaveStandard(standard)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": "111", "original": "demo.jpg", "url": standard.Route}
		c.ServeJSON()
	}
}
