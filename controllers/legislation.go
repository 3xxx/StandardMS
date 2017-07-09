package controllers

import (
	// "bufio"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/tealeg/xlsx"
	// "io"
	"hydrocms/models"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type LegislationController struct {
	beego.Controller
}

type Legislationmore struct {
	Id            int64
	Number        string //`orm:"unique"`
	Title         string //原法规名称
	LibraryNumber string //规范有效版本库中的编号
	LibraryTitle  string
	Execute       string //执行时间
}

func (c *LegislationController) Get() { //这个没用到
	c.Data["IsLegislation"] = true //这里修改到ListAllPosts()
	c.TplName = "legislation.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
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

func (c *LegislationController) Index() { //
	c.Data["IsLegislation"] = true //
	c.TplName = "legislation.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	legislations, err := models.GetAllLegislations() //这里传入空字符串
	if err != nil {
		beego.Error(err.Error)
	} else {
		c.Data["Legislations"] = legislations
		c.Data["Length"] = len(legislations) //得到总记录数
	}

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/test.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP())
	logs.Close()
}

//搜索规范或者图集的名称或编号
func (c *LegislationController) Checklist() { //checklist用的是post方法
	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/test.log"}`)
	logs.EnableFuncCallDepth(true)

	name := c.Input().Get("name")
	// beego.Info(name)
	array := strings.Split(name, "\n")
	aa := make([]Legislationmore, len(array))
	for i, v := range array {
		// beego.Info(v)
		//1、正则取到名称
		if v != "" { //空行的处理
			reg := regexp.MustCompile(`[《].*[》]`) //(`^\\<.*\\>`)
			text2 := reg.FindAllString(v, -1)
			// beego.Info(text2)
			if text2 != nil { //无书名号的处理。因为text2是数组，所以要用nil进行判断，而不能用""。
				text3 := SubString(text2[0], 1, len([]rune(text2[0]))-2)
				//2、根据名称搜索标准版本库，取得名称和版本号
				library, err := models.SearchLiabraryName(text3)
				// beego.Info(library)
				if err != nil {
					beego.Error(err.Error)
				}
				text4 := strconv.Itoa(i + 1)
				Id1, err := strconv.ParseInt(text4, 10, 64)
				if err != nil {
					beego.Error(err.Error)
				}
				aa[i].Id = Id1

				if len(library) != 0 { //library != nil这样不行，空数组不是nil
					beego.Info(library)
					//3、构造struct
					for j, w := range library {
						// beego.Info(w)
						if j == 0 {
							aa[i].LibraryNumber = w.Category + " " + w.Number + "-" + w.Year //规范有效版本库中的完整编号
							aa[i].LibraryTitle = w.Title
							aa[i].Execute = w.Execute //执行日期
						} else {
							aa[i].LibraryNumber = aa[i].LibraryNumber + "," + w.Category + " " + w.Number + "-" + w.Year //规范有效版本库中的完整编号
							aa[i].LibraryTitle = w.Title
							aa[i].Execute = aa[i].Execute + "," + w.Execute //执行日期
						}
					}
				} else {
					// beego.Info(library)
					// aa[i].Number = library.Number //`orm:"unique"`
					// aa[i].Title = text3
					aa[i].LibraryNumber = "No LibraryNumber Match Find!"
					aa[i].LibraryTitle = text3
					aa[i].Execute = ""
					logs.Info(c.Ctx.Input.IP() + " " + "No LibraryNumber:" + text3)
					// beego.Info(aa[i])
				}
			}
		}
	}
	c.Data["IsLegislation"] = true
	c.TplName = "legislation.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	//逐行读取
	// br := bufio.NewReader(name)
	// for {
	// 	a, _, c := br.ReadLine()
	// 	if c == io.EOF {
	// 		break
	// 	}
	// 	beego.Info(string(a))
	// }

	// bfRd := bufio.NewReader(f)
	// for {
	// 	line, err := bfRd.ReadBytes('\n')
	// 	hookfn(line)    //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
	// 	if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
	// 		if err == io.EOF {
	// 			return nil
	// 		}
	// 		return err
	// 	}
	// }

	// buf := bufio.NewReader(f)
	// for {
	// 	line, err := buf.ReadString('\n')
	// 	line = strings.TrimSpace(line)
	// 	handler(line)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			return nil
	// 		}
	// 		return err
	// 	}
	// }

	//由categoryid查categoryname
	// aa := make([]Legislationmore, len(Results1))
	// //由legislationnumber查librarynumber
	// for i, v := range Results1 {
	// 	//由userid查username
	// 	user := models.GetUserByUserId(v.Uid)
	// 	//由分类和编号查有效版本库中的编号
	// 	library, err := models.SearchLiabraryNumber(name, "Number")
	// 	if err != nil {
	// 		beego.Error(err.Error)
	// 	}
	// 	aa[i].Id = v.Id
	// 	aa[i].Number = v.Number //`orm:"unique"`
	// 	aa[i].Title = v.Title
	// 	aa[i].Uname = user.Username //换成用户名
	// 	aa[i].Route = v.Route
	// 	aa[i].Created = v.Created
	// 	aa[i].Updated = v.Updated
	// 	aa[i].Views = v.Views
	// 	if library != nil {
	// 		aa[i].LibraryNumber = library.Number //规范有效版本库中的编号
	// 		aa[i].LibraryTitle = library.Title
	// 		aa[i].LiNumber = library.LiNumber //完整编号
	// 	} else {
	// 		aa[i].LiNumber = "No LibraryNumber Match Find!"
	// 		aa[i].LibraryTitle = ""
	// 		aa[i].LibraryNumber = ""
	// 	}
	// }
	c.Data["json"] = aa //这里必须要是c.Data["json"]，其他c.Data["Data"]不行
	c.ServeJSON()

	logs.Info(c.Ctx.Input.IP() + " " + "SearchLegislationsName:" + name)
	logs.Close()
}

//上传excel文件，导入到数据库
//引用来自category的查看成果类型里的成果
func (c *LegislationController) ImportExcel() {
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
	var legislation models.Legislation
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
			legislation.Number, err = row.Cells[j].String()
			legislation.Title, err = row.Cells[j+1].String()
			// Uname, err := row.Cells[j+2].String()
			// user := models.GetUserByUsername(Uname)
			// Legislation.Uid = user.Id
			// Category, err := row.Cells[j+3].String()
			// category, _ := models.GetCategoryName(Category)
			// Legislation.CategoryId = category.Id
			legislation.Created = time.Now()
			legislation.Updated = time.Now()
			legislation.Content, err = row.Cells[j+4].String()
			legislation.Route, err = row.Cells[j+5].String()
			_, err = models.SaveLegislation(legislation)

			if err != nil {
				beego.Error(err)
			}
			// }
			// for _, cell := range row.Cells {这里要继续循环cells，不能为空，即超出单元格使用范围
			// 	fmt.Printf("%s\n", cell.String())

			// }
		}
	}
	c.TplName = "legislation.tpl"
	c.Redirect("/legislation", 302)
}

//上传excel文件，导入到数据库
//引用来自category的查看成果类型里的成果
func (c *LegislationController) ImportLibrary() {
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
	c.TplName = "legislation.tpl"
	c.Redirect("/legislation", 302)
}
