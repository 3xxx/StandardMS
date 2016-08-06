package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils/pagination"
	"image/png"
	"os"
	"quick/models"
	"strconv"
	"strings"
	"time"
)

type Uploadimage struct {
	url     string
	message string
	success int
}

type Categorydefine struct { //在category基础上加上label
	Id              int64
	ParentId        int64
	Uid             int64
	Title           string
	Number          string
	Content         string
	Cover           string
	Route           string //封面图片的链接地址
	Created         time.Time
	Updated         time.Time
	Views           int64
	Author          string //这个改成uid代替
	TopicCount      int64  //`form:"-"`
	TopicLastUserId int64  //`form:"-"`
	Isshow          bool
	Graphicmode     bool   //true表示图文模式
	Isuserdefined   bool   //是否自定义
	Label           string //[]*models.Label
}

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "add": //这个作废了，用后面的Post()代替
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.Data["IsCategory"] = true
		name := c.Input().Get("name")
		number := c.Input().Get("number")
		label := c.Input().Get("label")
		content := c.Input().Get("content")
		image := c.Input().Get("image")
		path := c.Input().Get("tempString")
		if len(name) == 0 {
			break
		}

		//建立目录
		array := strings.Split(path, ",") //字符串切割 [a b c d e]
		// for _, v := range array {
		// 	switch v {
		// 	case "ghj", "xj", "ky", "cs", "zb", "sgt", "jgt":
		// 		//建立目录——注意，models中无法建立目录，必须在controllers中才行
		// 		err := os.MkdirAll(".\\attachment\\"+v, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
		// 		// err := os.Mkdir(".\\attachment\\ghj", 0777)
		// 		if err != nil {
		// 			beego.Error(err)
		// 		}
		// 	}
		// }

		// for _, v := range array {
		// 	switch v {
		// 	case "gh", "sg", "jd", "shg", "dz", "ys", "zh": //专业
		// 		//查到所有阶段
		// 		for _, w := range array {
		// 			switch w {
		// 			case "ghj", "xj", "ky", "cs", "zb", "sgt", "jgt":
		// 				err := os.MkdirAll(".\\attachment\\"+w+"\\"+v, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
		// 				// err := os.Mkdir(".\\attachment\\ghj", 0777)
		// 				if err != nil {
		// 					beego.Error(err)
		// 				}
		// 			}
		// 		}
		// 	}
		// }
		for _, v := range array {
			switch v {
			case "ghj", "xj", "ky", "cs", "zb", "sgt", "jgt": //阶段//查到所有阶段
				for _, w := range array {
					switch w {
					case "gh", "sg", "jd", "shg", "dz", "ys", "zh": //专业
						for _, t := range array {
							switch t {
							case "dwg", "doc", "xls", "pdf", "jpg", "tif", "diary": //成果分类
								err := os.MkdirAll(".\\attachment\\"+number+""+name+"\\"+v+"\\"+w+"\\"+t, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
								if err != nil {
									beego.Error(err)
								}
							}
						}
					}
				}
			}
		}
		// String temp1[] = tempString.split(",");// 截取字符串，获得各个checkBox的值
		// 这个数组temp1[]里面的值就是要获取的各个复选框里取到的值
		// name = c.Input().Get("tempString")
		// array := strings.Split(name, ",") //字符串切割 [a b c d e]
		// if len(name) == 0 {
		// 	break
		// }
		// for _, v := range array { //i
		// 	// fmt.Println("Array element[", i, "]=", v)
		// 	// err := models.AddCategory(name)
		// 	err := models.AddCategory(v, "") //这个err后面的冒号为何没提示错误？
		// 	if err != nil {
		// 		beego.Error(err)
		// 	}
		// }
		// diskdirectory := ".\\attachment\\" + number + name + "\\"
		// url := "/attachment/" + number + name + "/"
		//保存上传的图片
		//获取上传的文件
		_, h, err := c.GetFile(image)
		// beego.Info(h)
		if err != nil {
			beego.Error(err)
		}
		// var attachment string
		// var path string
		var filesize int64
		if h != nil {
			//保存附件
			// attachment = h.Filename
			// beego.Info(attachment)
			path = ".\\attachment\\" + number + " " + name + "\\" + h.Filename
			err = c.SaveToFile("image", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
			if err != nil {
				beego.Error(err)
			}
			filesize, _ = FileSize(path)
			filesize = filesize / 1000.0
		}
		ck, err := c.Ctx.Request.Cookie("uname")
		if err != nil {
			beego.Error(err)
		}
		uname := ck.Value
		route := "/attachment/" + number + " " + name + "/" + h.Filename
		if err != nil {
			beego.Error(err)
		}
		//存入数据库
		id, err := models.AddCategory(name, number, label, content, "", path, route, uname)
		if err != nil {
			beego.Error(err)
		}
		id1 := strconv.FormatInt(id, 10)
		c.Redirect("/category?op=view&id="+id1, 301)
		return
	// case "del":
	// 	var rolename int
	// 	var uname string
	// 	uname, role, _ := checkRolewrite(c.Ctx) //login里的
	// 	rolename, _ = strconv.Atoi(role)
	// 	c.Data["Uname"] = uname
	// 	id := c.Input().Get("id")
	// 	if len(id) == 0 {
	// 		break
	// 	}
	// 	if rolename > 1 { //&& uname != category.Author
	// 		route := c.Ctx.Request.URL.String()
	// 		c.Data["Url"] = route
	// 		c.Redirect("/roleerr?url="+route, 302)
	// 		return
	// 	}
	// 	err := models.DelCategory(id)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	c.Redirect("/category", 301)
	// 	return
	//http://127.0.0.1/category?op=view&id=2726这个view是整个项目查看
	case "view":
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.Data["IsCategory"] = true
		c.TplName = "category_view.html"

		var uname string
		//2.如果登录或ip在允许范围内，进行访问权限检查
		uname, _, _ = checkRoleread(c.Ctx) //login里的
		c.Data["Uname"] = uname

		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		category, label, err := models.GetCategory(id) //由分类id取出本身（项目名称等）
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(category.Title)
		topics, err := models.GetAllTopics(category.Title, false)
		categorychengguo, _ := models.GetCategoryChengguo(id)
		categoryzhuanye, _ := models.GetCategoryLeixing(id)
		categoryjieduan, _ := models.GetCategoryJieduan(id)
		// if err != nil {
		// 	beego.Error(err)
		// 	c.Redirect("/", 302)
		// 	return
		// }
		c.Data["Category"] = category
		c.Data["Label"] = label
		c.Data["Categorychengguo"] = categorychengguo
		c.Data["Categoryzhuanye"] = categoryzhuanye
		c.Data["Categoryjieduan"] = categoryjieduan
		c.Data["Tid"] = id //教程中用的是圆括号，导致错误topic.go:52: cannot call non-function c.Controller.Ctx.Input.Params (type map[string]string)
		//教程第8章开头有修改

		//下面引自index
		c.Data["Id"] = c.Ctx.Input.Param(":id")
		// topics, err := models.GetAllTopics(c.Input().Get("cate"), true)
		// if err != nil {
		// 	beego.Error(err)
		// }
		c.Data["Topics"] = topics
		categories, err := models.GetAllCategories() //这个没有用吧
		if err != nil {
			beego.Error(err)
		}
		c.Data["Categories"] = categories
		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/test.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "ViewCategory" + " " + category.Title)
		logs.Close()

	case "view_b":
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.Data["IsCategoryb"] = true
		uname, _, _ := checkRoleread(c.Ctx) //login里的
		// rolename, _ = strconv.Atoi(role)
		c.Data["Uname"] = uname
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}

		category, label, _ := models.GetCategory(id) //由成果id取出成果
		c.TplName = "category_view_b.html"
		categorychengguo, _ := models.GetCategoryChengguo(id)
		categoryzhuanye, _ := models.GetCategoryLeixing(id)
		categoryjieduan, _ := models.GetCategoryJieduan(id)
		// if err != nil {
		// 	beego.Error(err)
		// 	c.Redirect("/", 302)
		// 	return
		// }
		c.Data["Category"] = category
		c.Data["Label"] = label
		c.Data["Categorychengguo"] = categorychengguo
		c.Data["Categoryzhuanye"] = categoryzhuanye
		c.Data["Categoryjieduan"] = categoryjieduan
		c.Data["Tid"] = id //教程中用的是圆括号，导致错误topic.go:52: cannot call non-function c.Controller.Ctx.Input.Params (type map[string]string)
		//教程第8章开头有修改

		//下面引自index
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

	case "viewlabel": //按label查看
		c.Data["IsCategory"] = true
		c.TplName = "category_label.tpl"
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		uname, _, _ := checkRoleread(c.Ctx) //login里的
		// rolename, _ = strconv.Atoi(role)
		c.Data["Uname"] = uname
		label := c.Input().Get("label")
		// beego.Info(label)
		if len(label) == 0 {
			break
		}
		categories, err := models.GetCategoriesbylabel(label) //由分类id取出本身（项目名称等）
		if err != nil {
			beego.Error(err)
		}

		// count := len(categories)
		// count1 := strconv.Itoa(count)
		// count2, err := strconv.ParseInt(count1, 10, 64)
		// if err != nil {
		// 	beego.Error(err)
		// }
		// c.Data["Length"] = len(categories)

		// sets this.Data["paginator"] with the current offset (from the url query param)
		// categoriesPerPage := 20
		// paginator := pagination.SetPaginator(c.Ctx, categoriesPerPage, count2)
		// // beego.Info(c.Ctx)
		// beego.Info(paginator.Offset())   0
		// p := pagination.NewPaginator(c.Ctx.Request, 10, 9)
		// beego.Info(p.Offset())   0
		// fetch the next 20 posts
		// categories, err = models.ListCategoriesByOffsetAndLimit(paginator.Offset(), categoriesPerPage)
		// if err != nil {
		// 	beego.Error(err)
		// }
		c.Data["Category"] = categories
		c.Data["Label"] = label
		// c.Data["paginator"] = paginator
		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/test.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "ViewCategory")
		logs.Close()
	default: //即http://127.0.0.1/category
		// c.EnableRender = false
		c.Data["IsCategory"] = true
		c.TplName = "category.tpl"
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		uname, _, _ := checkRoleread(c.Ctx) //login里的
		// rolename, _ = strconv.Atoi(role)
		c.Data["Uname"] = uname
		categories, err := models.GetAllCategories()
		if err != nil {
			beego.Error(err)
		}
		count := len(categories)
		count1 := strconv.Itoa(count)
		count2, err := strconv.ParseInt(count1, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		// c.Data["Category"] = categories
		c.Data["Length"] = len(categories)

		// sets this.Data["paginator"] with the current offset (from the url query param)
		categoriesPerPage := 20
		paginator := pagination.SetPaginator(c.Ctx, categoriesPerPage, count2)
		// beego.Info(c.Ctx)
		// beego.Info(paginator.Offset())   0
		// p := pagination.NewPaginator(c.Ctx.Request, 10, 9)
		// beego.Info(p.Offset())   0
		// fetch the next 20 posts
		categories, labels, err := models.ListCategoriesByOffsetAndLimit(paginator.Offset(), categoriesPerPage)
		if err != nil {
			beego.Error(err)
		}

		//循环获取每个category的label——这个方法太慢，取消
		// var label1 string
		// slice1 := make([]Categorydefine, 0)
		// for i1, _ := range categories {
		// 	cid := strconv.FormatInt(categories[i1].Id, 10)
		// 	_, label, err := models.GetCategory(cid) //由成果id取出成果
		// 	// _, numbers, marks, err := models.GetMeritTopic(0, categories[i1].Id)
		// 	if err != nil {
		// 		beego.Error(err)
		// 	}
		// 	//3.由cid查询数据库中的项目名
		// 	for i, label2 := range label {
		// 		if i == 0 {
		// 			label1 = label2.Title
		// 		} else {
		// 			label1 = label1 + "," + label2.Title
		// 		}
		// 	}
		// 	aa := make([]Categorydefine, 1)
		// 	aa[0].Id = categories[i1].Id //这里用for i1,v1,然后用v1.Id一样的意思
		// 	aa[0].ParentId = categories[i1].ParentId
		// 	aa[0].Uid = categories[i1].Uid
		// 	aa[0].Title = categories[i1].Title
		// 	aa[0].Number = categories[i1].Number
		// 	aa[0].Content = categories[i1].Content
		// 	aa[0].Cover = categories[i1].Cover
		// 	aa[0].Route = categories[i1].Route
		// 	aa[0].Created = categories[i1].Created
		// 	aa[0].Updated = categories[i1].Updated
		// 	aa[0].Views = categories[i1].Views
		// 	aa[0].Author = categories[i1].Author
		// 	aa[0].TopicCount = categories[i1].TopicCount
		// 	aa[0].TopicLastUserId = categories[i1].TopicLastUserId
		// 	aa[0].Isshow = categories[i1].Isshow
		// 	aa[0].Graphicmode = categories[i1].Graphicmode
		// 	aa[0].Isuserdefined = categories[i1].Isuserdefined
		// 	aa[0].Label = label1
		// 	slice1 = append(slice1, aa...)
		// }
		// c.Data["Category"] = slice1
		c.Data["Category"] = categories
		c.Data["Label"] = labels
		c.Data["paginator"] = paginator

		// c.EnableRender = false
		// c.Data["json"] = categories
		// c.Data["json1"] = labels //换名不行，只支持一组
		// c.ServeJSON()
		// c.Ctx.WriteString("hello")

		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/test.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "ViewCategory")
		logs.Close()
	}
	var err error
	if err != nil {
		beego.Error(err)
	}
}

//这个给UI用
func (c *CategoryController) CategoryUi() {
	c.EnableRender = false
	c.Data["IsCategory"] = true
	// c.TplName = "category.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	uname, _, _ := checkRoleread(c.Ctx) //login里的
	c.Data["Uname"] = uname
	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	count := len(categories)
	count1 := strconv.Itoa(count)
	count2, err := strconv.ParseInt(count1, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Length"] = len(categories)

	categoriesPerPage := 20
	paginator := pagination.SetPaginator(c.Ctx, categoriesPerPage, count2)

	categories, labels, err := models.ListCategoriesByOffsetAndLimit(paginator.Offset(), categoriesPerPage)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Category"] = categories
	c.Data["Label"] = labels
	c.Data["paginator"] = paginator

	// c.EnableRender = false
	c.Data["json"] = categories
	// c.Data["json1"] = labels //换名不行，只支持一组
	c.ServeJSON()
	// c.Ctx.WriteString("hello")

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/test.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "ViewCategory")
	logs.Close()
}

//这个给UI用
func (c *CategoryController) CategoryViewUi() {
	c.EnableRender = false
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	// c.TplName = "category_view.html"

	var uname string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, _, _ = checkRoleread(c.Ctx) //login里的
	c.Data["Uname"] = uname

	id := c.Input().Get("id")
	// if len(id) == 0 {
	// 	break
	// }
	category, label, err := models.GetCategory(id) //由分类id取出本身（项目名称等）
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(category.Title)
	topics, err := models.GetAllTopics(category.Title, false)
	categorychengguo, _ := models.GetCategoryChengguo(id)
	categoryzhuanye, _ := models.GetCategoryLeixing(id)
	categoryjieduan, _ := models.GetCategoryJieduan(id)
	// if err != nil {
	// 	beego.Error(err)
	// 	c.Redirect("/", 302)
	// 	return
	// }
	c.Data["Category"] = category
	c.Data["Label"] = label
	c.Data["Categorychengguo"] = categorychengguo
	c.Data["Categoryzhuanye"] = categoryzhuanye
	c.Data["Categoryjieduan"] = categoryjieduan
	c.Data["Tid"] = id //教程中用的是圆括号，导致错误topic.go:52: cannot call non-function c.Controller.Ctx.Input.Params (type map[string]string)
	//教程第8章开头有修改
	c.Data["json"] = categoryjieduan
	// c.Data["json1"] = labels //换名不行，只支持一组
	c.ServeJSON()
	//下面引自index
	c.Data["Id"] = c.Ctx.Input.Param(":id")
	c.Data["Topics"] = topics
	categories, err := models.GetAllCategories() //这个没有用吧
	if err != nil {
		beego.Error(err)
	}
	c.Data["Categories"] = categories
	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/test.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "ViewCategory" + " " + category.Title)
	logs.Close()
}

//ui文档类型——下文虽然是专业，实际上是类型
func (c *CategoryController) CategoryView1Ui() {
	c.EnableRender = false
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	// c.TplName = "category_view.html"

	var uname string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, _, _ = checkRoleread(c.Ctx) //login里的
	c.Data["Uname"] = uname

	id := c.Input().Get("id")

	categoryzhuanye, _ := models.GetCategoryLeixing(id)

	c.Data["json"] = categoryzhuanye
	c.ServeJSON()

}

//ui专业——下文虽然是chengguo，但是实际上是专业
func (c *CategoryController) CategoryView2Ui() {
	c.EnableRender = false
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true

	var uname string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, _, _ = checkRoleread(c.Ctx) //login里的
	c.Data["Uname"] = uname

	id := c.Input().Get("id")
	categorychengguo, _ := models.GetCategoryChengguo(id)
	c.Data["json"] = categorychengguo
	c.ServeJSON()
}

//删除项目数据库——删除项目中的成果(删除附件)——删除物理目录
func (c *CategoryController) Delete() {
	cid := c.Input().Get("cid")
	var rolename int
	var uname string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, role, _ := checkRolewrite(c.Ctx) //login里的
	// beego.Info(uname) = nil
	rolename, _ = strconv.Atoi(role)
	// beego.Info(rolename)=5
	c.Data["Uname"] = uname
	//3.由cid查询数据库中的项目名
	var label1 string
	category, label, err := models.GetCategory(cid)
	// beego.Info(category.Author)=admin
	for i, label2 := range label {
		if i == 0 {
			label1 = label2.Title
		} else {
			label1 = label1 + "," + label2.Title
		}
	}
	if rolename > 1 && uname != category.Author { //要么管理员，要么作者自己可以修改项目
		// port := strconv.Itoa(c.Ctx.Input.Port())
		route := c.Ctx.Request.URL.String() //c.Ctx.Input.Site() + ":" + port +
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	url := c.Input().Get("url")
	var tid string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	// 2016-7-31既未登录，IP也不在范围内，可能出现错误。
	//删除数据库中的成果（和附件）
	topics, err := models.GetAllTopics(category.Title, false)
	//这个放删除项目后面不行，因为是用的是title，models里还需要转换成categoryid
	for _, w := range topics {
		// beego.Info(w)
		// beego.Info(w.Id)
		tid = strconv.FormatInt(w.Id, 10)
		err = models.DeletTopic(tid)
		if err != nil {
			beego.Error(err)
		}
	}
	//删除物理目录
	// func RemoveAll(path string) errorRemoveAll删除path指定的文件，或目录及它包含的任何下级对象。它会尝试删除所有东西，除非遇到错误并返回。如果path指定的对象不存在，RemoveAll会返回nil而不返回错误。
	err = os.RemoveAll(".\\attachment\\" + category.Number + category.Title + "\\")
	if err != nil {
		beego.Error(err)
	}
	// 删除数据库中的项目
	err = models.DelCategory(cid)
	if err != nil {
		beego.Error(err)
		// data := "权限不够，请登录！"
		// c.Ctx.WriteString(data)
	} else {
		data := "ok!"
		c.Ctx.WriteString(data)
	}
	// err :=os.MkdirAll(".\\attachment\\"+number+" "+name+"\\"
	// return               //Handler crashed with error can't find templatefile in the path:topiccontroller/delete.tpl
	c.Redirect(url, 302) //这里增加topic
	// c.Redirect("/category", 301)
	// return
}

func (c *CategoryController) Get_b() { //项目B显示控制
	c.Data["IsCategoryb"] = true
	c.TplName = "category_b.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		c.Data["Uname"] = v.(string)
	}
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// } else {
	// 	c.Data["Uname"] = ck.Value
	// }
	// var err error
	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Category"] = categories
	c.Data["Length"] = len(categories)

	// var err error
	if err != nil {
		beego.Error(err)
	}
}

//添加项目第一步视图
func (c *CategoryController) Add() {
	var rolename int
	var uname string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, role, _ := checkRolewrite(c.Ctx) //login里的
	rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	if rolename > 2 { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	c.TplName = "category_add.tpl"
	// id := c.Input().Get("id")
	// category, err := models.GetCategory(id)
	// if err != nil {
	// beego.Error(err)
	// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
	// return
	// }
	// c.Data["Category"] = category
	// c.Data["Id"] = id
}

//添加项目第一步方法提交
func (c *CategoryController) Post() {
	//也可以先c.Input().Get("category2")再切割strings.Split(category2, ",")
	// category2 := c.GetStrings("category2")
	// category3 := c.GetStrings("category3")
	// category4 := c.GetStrings("category4")
	// for _, v := range category2 {
	//   if category3!=""{}如果有3级目录
	//   for _, w := range category3 {
	//     if category4!=""{}如果有4级目录
	//     for _, t := range category4 {
	// 	err :=os.MkdirAll(".\\attachment\\"+number+" "+name+"\\"+v+"\\"+w+"\\"+t, 0777)
	//
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// }
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	//2.取得客户端用户名
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }

	name := c.Input().Get("name")
	number := c.Input().Get("number")
	label := c.Input().Get("label")
	// content := c.Input().Get("editorValue")
	// image := c.Input().Get("image")
	path := c.Input().Get("tempString")
	// if len(name) == 0 {
	// 	break
	// }

	//建立目录
	array := strings.Split(path, ",") //字符串切割 [a b c d e]
	// for _, v := range array {
	// 	switch v {
	// 	case "ghj", "xj", "ky", "cs", "zb", "sgt", "jgt":
	// 		//建立目录——注意，models中无法建立目录，必须在controllers中才行
	// 		err := os.MkdirAll(".\\attachment\\"+v, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	// 		// err := os.Mkdir(".\\attachment\\ghj", 0777)
	// 		if err != nil {
	// 			beego.Error(err)
	// 		}
	// 	}
	// }

	// for _, v := range array {
	// 	switch v {
	// 	case "gh", "sg", "jd", "shg", "dz", "ys", "zh": //专业
	// 		//查到所有阶段
	// 		for _, w := range array {
	// 			switch w {
	// 			case "ghj", "xj", "ky", "cs", "zb", "sgt", "jgt":
	// 				err := os.MkdirAll(".\\attachment\\"+w+"\\"+v, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	// 				// err := os.Mkdir(".\\attachment\\ghj", 0777)
	// 				if err != nil {
	// 					beego.Error(err)
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	for _, v := range array {
		switch v {
		case "A", "B", "C", "D", "E", "F", "G", "L": //阶段
			for _, w := range array {
				switch w {
				case "FB", "FD", "FG", "FT", "FJ", "FP", "Fdiary": //文件类型
					for _, t := range array {
						switch t {
						case "1", "2", "3", "4", "5", "6", "7", "8", "9": //专业
							err := os.MkdirAll(".\\attachment\\"+number+name+"\\"+v+"\\"+w+"\\"+t, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
							if err != nil {
								beego.Error(err)
							}
						}
					}
				}
			}
		}
	}
	// String temp1[] = tempString.split(",");// 截取字符串，获得各个checkBox的值
	// 这个数组temp1[]里面的值就是要获取的各个复选框里取到的值
	// name = c.Input().Get("tempString")
	// array := strings.Split(name, ",") //字符串切割 [a b c d e]
	// if len(name) == 0 {
	// 	break
	// }
	// for _, v := range array { //i
	// 	// fmt.Println("Array element[", i, "]=", v)
	// 	// err := models.AddCategory(name)
	// 	err := models.AddCategory(v, "") //这个err后面的冒号为何没提示错误？
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// }
	// diskdirectory := ".\\attachment\\" + number + name + "\\"
	// url := "/attachment/" + number + name + "/"
	//保存上传的图片
	//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
	// _, h, err := c.GetFile("image")
	// beego.Info(h)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// var attachment string
	// var path string
	// var filesize int64
	// var route string
	// if h != nil {
	// 	//保存附件
	// 	// attachment = h.Filename
	// 	// beego.Info(attachment)
	// 	path1 := ".\\attachment\\" + number + name + "\\" + h.Filename
	// 	err = c.SaveToFile("image", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	//如果扩展名为jpg
	// 	// if strings.ToLower(path.Ext(h.Filename)) == ".jpg" {

	// 	// }
	// 	//如果包含jpg，则进行压缩
	// 	if strings.Contains(strings.ToLower(h.Filename), ".jpg") { //ToLower转成小写
	// 		// 随机名称
	// 		// to := path + random_name() + ".jpg"
	// 		origin := path1 //path + file.Name()
	// 		fmt.Println("正在处理" + origin + ">>>" + origin)
	// 		cmd_resize(origin, 2048, 0, origin)
	// 		//				defer os.Remove(origin)//删除原文件
	// 	}
	// 	filesize, _ = FileSize(path1)
	// 	filesize = filesize / 1000.0
	// 	route = "/attachment/" + number + name + "/" + h.Filename
	// } else {//如果没有图片就自动生成一个
	// 	img := CreateRandomAvatar([]byte(number + name))
	// 	fi, _ := os.Create("./attachment/" + number + name + "/u1.png")
	// 	png.Encode(fi, img)
	// 	fi.Close()
	// 	route = "/attachment/" + number + name + "/u1.png"
	// }
	uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	//存入数据库
	id, err := models.AddCategory(name, number, label, "", "", path, "", uname)
	if err != nil {
		beego.Error(err)
	}
	id1 := strconv.FormatInt(id, 10)
	// c.Redirect("/category?op=view&id="+id1, 301)
	c.Redirect("/category/add2?id="+id1, 301) //跳转到添加封面图片、封面说明和项目简介
	return                                    //???
}

//添加项目第二步视图
func (c *CategoryController) Add2() {
	var rolename int
	var uname string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, role, _ := checkRolewrite(c.Ctx) //login里的
	rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	if rolename > 2 { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	c.TplName = "category_add2.tpl"

	id := c.Input().Get("id") //这个get也能获取地址栏中的id啊
	category, label, err := models.GetCategory(id)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302) //这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
		return
	}
	c.Data["Category"] = category
	c.Data["Label"] = label
	c.Data["Id"] = id
}

//添加项目第二步方法提交
func (c *CategoryController) Post2() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true

	cid := c.Input().Get("categoryid")
	//取得上传之封面图片
	route := c.Input().Get("route")
	// beego.Info(cid)
	// number := c.Input().Get("number")
	cover := c.Input().Get("editor_cover")
	content := c.Input().Get("editor_property")

	// image := c.Input().Get("image")
	// path := c.Input().Get("tempString")
	// if len(name) == 0 {
	// 	break
	// }

	//更新数据库
	err := models.ModifyCategory(cid, "", "", "", content, cover, "", route, "")
	// id, err := models.AddCategory("", "", content, "", "", "", "", "")
	if err != nil {
		beego.Error(err)
	}
	// id1 := strconv.FormatInt(id, 10)
	c.Redirect("/category?op=view&id="+cid, 301)
	// c.Redirect("/category/add2?id="+id1, 301) //跳转到添加封面图片、封面说明和项目简介
	// return //???
}

//项目第二步添加封面图片
func (c *CategoryController) AddCoverPhoto() {
	//解析表单
	cid := c.Input().Get("categoryid")
	// beego.Info(cid)
	// category1, _, err := models.GetCategory(cid)
	// if err != nil {
	// 	beego.Error(err)
	// 	return
	// }
	url, diskdirectory, err := models.GetCategoryUrl(cid)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(diskdirectory)
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	// var attachment string
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		// attachment := h.Filename
		path = diskdirectory + h.Filename
		beego.Info(h.Filename)
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
	}

	route := url + h.Filename

	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": "111", "original": "demo.jpg", "url": route}
		c.ServeJSON()
	}
	// c.TplName = "topic_one_add.tpl" //不加这句上传出错，虽然可以成功上传
	// c.Redirect("/topic", 302)
	// success : 0 | 1,           // 0 表示上传失败，1 表示上传成功
	//    message : "提示的信息，上传成功或上传失败及错误信息等。",
	//    url     : "图片地址"        // 上传成功时才返回
}

//删除项目结构中的项目——删除下级——删除下级中的成果——删除硬盘目录
func (c *CategoryController) DeleteCategory() {
	cid := c.Input().Get("cid") //项目id
	id := c.Input().Get("id")   //要删除的id
	idNum, err := strconv.ParseInt(id, 10, 64)
	err = models.DeleteCategory(idNum) //删除本级和下级
	if err != nil {
		beego.Error(err)
	}
	// err = models.DelCategory(id)
	// if err != nil {
	// 	beego.Error(err)
	// } else {
	// 	data := "ok!"
	// 	c.Ctx.WriteString(data)
	// }
	c.Redirect("/category/modifyfrm?cid="+cid, 301)
}

//显示项目结构中的项目
func (c *CategoryController) ShowCategory() {
	cid := c.Input().Get("cid") //项目id
	id := c.Input().Get("id")   //要删除的id
	idNum, err := strconv.ParseInt(id, 10, 64)
	err = models.ShowCategory(idNum)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/category/modifyfrm?cid="+cid, 301)
}

//隐藏项目结构中的项目
func (c *CategoryController) HideCategory() {
	cid := c.Input().Get("cid") //项目id
	id := c.Input().Get("id")   //要删除的id
	idNum, err := strconv.ParseInt(id, 10, 64)
	err = models.HideCategory(idNum)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/category/modifyfrm?cid="+cid, 301)
}

//这个是测试用的
func (c *CategoryController) Uploadimagesct() {
	name := "111"    //c.Input().Get("name")
	number := "222"  //c.Input().Get("number")
	content := "333" //c.Input().Get("test-editormd-html-code")
	path := "c"      //c.Input().Get("tempString")
	label := c.Input().Get("label")
	// diskdirectory := ".\\attachment\\" + "test" + "\\"
	// url := "/attachment/" + "test" + "/"
	//保存上传的图片
	//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
	_, h, err := c.GetFile("editormd-image-file") //editormd-image-file
	// beego.Info(h)
	if err != nil {
		beego.Error(err)
	}
	// var attachment string
	// var path string
	var filesize int64
	var route string
	if h != nil {
		//保存附件
		path1 := ".\\attachment\\" + "test" + "\\" + h.Filename
		err = c.SaveToFile("editormd-image-file", path1) //editormd-image-file  .Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}

		if strings.Contains(strings.ToLower(h.Filename), ".jpg") { //ToLower转成小写
			// 随机名称
			// to := path + random_name() + ".jpg"
			origin := path1 //path + file.Name()
			fmt.Println("正在处理" + origin + ">>>" + origin)
			cmd_resize(origin, 2048, 0, origin)
			//				defer os.Remove(origin)//删除原文件
		}
		filesize, _ = FileSize(path1)
		filesize = filesize / 1000.0
		route = "/attachment/" + "test" + "/" + h.Filename
	} else {
		img := CreateRandomAvatar([]byte(number + name))
		fi, _ := os.Create("./attachment/" + "test" + "/u1.png")
		png.Encode(fi, img)
		fi.Close()
		route = "/attachment/" + "test" + "/u1.png"
	}

	uname := "4"

	//存入数据库
	_, err = models.AddCategory(name, number, label, content, "", path, route, uname)
	if err != nil {
		beego.Error(err)
	} else {
		// f := Uploadimage{
		// 	url:     route,
		// 	success: 1,
		// 	message: "ok",
		// }
		// beego.Info(f)2016/01/17 01:40:03 [category.go:549] [I] {/attachment/test/u1.png ok 1}
		c.Data["json"] = map[string]interface{}{"success": 1, "message": "111", "url": route}
		// c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": "111", "original": "demo.jpg", "url": route}
		// c.Data["json"] = f
		c.ServeJSON()
		// beego.Info(c.Data["json"])
		// 2016/01/17 01:42:00 [category.go:554] [I] map[success:1 message:111 url:/attachm
		// ent/test/u1.png]
		// 		{
		//     "state": "SUCCESS",
		//     "url": "upload/demo.jpg",
		//     "title": "demo.jpg",
		//     "original": "demo.jpg"
		//      }
	}

	// c.Data["Uname"] = ck.Value
	// id1 := strconv.FormatInt(id, 10)
	// c.Redirect("/category?op=view&id="+id1, 301)
	return //???
}

//添加自定义项目第一步视图
func (c *CategoryController) Add_b() {
	var rolename int
	var uname string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, role, _ := checkRolewrite(c.Ctx) //login里的
	rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	if rolename > 2 { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	c.TplName = "category_add_b.tpl"
}

//添加自定义项目第一步提交方法
func (c *CategoryController) UserdefinedPost() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	name := c.Input().Get("name")
	number := c.Input().Get("number")
	label := c.Input().Get("label")
	// content := c.Input().Get("editorValue") //editorValue  test-editormd-html-code
	// image := c.Input().Get("image")
	// path1 := c.Input().Get("category2")
	// beego.Info(path1) //只能取到一个值 [I] 2-1
	// path2 := c.Input().Get("category3")
	// path3 := c.Input().Get("category4")
	var path2, path3, path4 []string
	// var diskdirectory, url string
	//单选按钮值的字符串，用,号隔开；；单选按钮的字符；；单选按钮传值的名字button0~buttoni
	var radio, radio1, radiostring string
	// path := c.Input()通过把这个map传到models里，models取得map，通过键取得值，就可以实现无数层的目录建立了
	// slice := []int{10, 20, 30, 40, 50}
	// arrSlice := []int{1, 2, 3, 4, 5}
	// m := map[string][]int{}定义map
	// m["key1"] = arrSlice//map赋值
	// fmt.Println(m)
	// beego.Info(path4)               // map[category4:[4-1 4-2] name:[……
	// beego.Info(path4["catelogry4"]) //这句无结果，按道理应该是[4-1 4-2]
	// if len(name) == 0 {
	// 	break
	// }

	//建立目录
	// array := strings.Split(path, ",") //字符串切割 [a b c d e]
	category2 := c.GetStrings("category2") //func (c *Controller) GetStrings(key string) []string
	// beego.Info(category2)
	category3 := c.GetStrings("category3")
	category4 := c.GetStrings("category4")
	// beego.Info(category2) //[2-1 2-2]
	// path := make([]string, 0)
	if len(category2) > 0 { //如果有2级目录，则建立2级，没有则建立1级
		for _, v := range category2 {
			if len(category3) > 0 { //如果有3级目录，则建立3级，没有则建立2级
				for _, w := range category3 {
					if len(category4) > 0 { //如果有4级目录，则建立4级，没有则建立3级
						path4 = category4
						path3 = category3
						path2 = category2
						for i, t := range category4 {
							err := os.MkdirAll(".\\attachment\\"+number+name+"\\"+v+"\\"+w+"\\"+t, 0777)
							if err != nil {
								beego.Error(err)
							}
							ii := strconv.Itoa(i)
							radiostring = "radiobutton" + ii
							radio1 = c.Input().Get(radiostring)
							if i == 0 {
								radio = radio1
							} else {
								radio = radio + "," + radio1
							}
							// diskdirectory = ".\\attachment\\" + number + name + "\\" + v + "\\" + w + "\\" + t + "\\"
							// url = "/attachment/" + number + name + "/" + v + "/" + w + "/" + t + "/"
							// path = append(category2, category3...)
							// path = append(path, category4...)
							// path = path1 + "," + path2 + "," + path3
							// beego.Info(path)

						}
					} else { //如果没有4级目录，则建立3级目录
						err := os.MkdirAll(".\\attachment\\"+number+name+"\\"+v+"\\"+w, 0777)
						// path = append(category2, category3...)
						// path = path1 + "," + path2
						path3 = category3
						path2 = category2
						// diskdirectory = ".\\attachment\\" + number + name + "\\" + v + "\\" + w + "\\"
						// url = "/attachment/" + number + name + "/" + v + "/" + w + "/"
						if err != nil {
							beego.Error(err)
						}
					}
				}
			} else { //如果没有3级目录，则建立2级目录
				err := os.MkdirAll(".\\attachment\\"+number+name+"\\"+v, 0777)
				// path := category2
				// path = path1
				path2 = category2
				// diskdirectory = ".\\attachment\\" + number + name + "\\" + v + "\\"
				// url = "/attachment/" + number + name + "/" + v + "/"
				if err != nil {
					beego.Error(err)
				}
			}
		}
	} else { //没有2级目录则建立1级目录
		err := os.MkdirAll(".\\attachment\\"+number+name, 0777)
		// diskdirectory = ".\\attachment\\" + number + name + "\\"
		// url = "/attachment/" + number + name + "/"
		if err != nil {
			beego.Error(err)
		}
	}
	// diskdirectory = ".\\attachment\\" + number + name + "\\"
	// url = "/attachment/" + number + name + "/"
	uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	//存入数据库
	// var id int64
	id, err := models.AdduserdefinedCategory(name, number, label, "", "", path2, path3, path4, radio, "", uname)
	if err != nil {
		beego.Error(err)
	}
	id1 := strconv.FormatInt(id, 10)
	c.Redirect("/category/add2?id="+id1, 301) //跳转到添加封面图片、封面说明和项目简介
	// beego.Info(radio)
	return
	//2.取得客户端用户名
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
}

//根据用户名查看项目
func (c *CategoryController) Viewbyuname() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	c.TplName = "category_uname.tpl"
	uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	uname = c.Input().Get("uname")
	// if len(uname) == 0 {
	// 	break
	// }

	category, _ := models.GetCategoriesbyuname(uname) //由uname取出项目
	// categorychengguo, _ := models.GetCategoryChengguo(id)
	// categoryzhuanye, _ := models.GetCategoryZhuanye(id)
	// categoryjieduan, _ := models.GetCategoryJieduan(id)
	// if err != nil {
	// 	beego.Error(err)
	// 	c.Redirect("/", 302)
	// 	return
	// }
	c.Data["Category"] = category
	// c.Data["Categorychengguo"] = categorychengguo
	// c.Data["Categoryzhuanye"] = categoryzhuanye
	// c.Data["Categoryjieduan"] = categoryjieduan
	// c.Data["Tid"] = id //教程中用的是圆括号，导致错误topic.go:52: cannot call non-function c.Controller.Ctx.Input.Params (type map[string]string)
	//教程第8章开头有修改

	//下面引自index
	// c.Data["Id"] = c.Ctx.Input.Param(":id")
	// topics, err := models.GetAllTopics(c.Input().Get("cate"), true)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// c.Data["Topics"] = topics
	// categories, err := models.GetAllCategories()
	// if err != nil {
	// 	beego.Error(err)
	// }
	// c.Data["Categories"] = categories
}

//查看成果类型里的成果——这个view是点击侧栏后显示的页面
func (c *CategoryController) View() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	id := c.Input().Get("id")
	// if len(id) == 0 {
	// 	break
	// }
	//判断父级title是否是“文章/设代日记”或者图文模式为真
	category, _ := models.GetCategorySpec(id) //这是父一级的，所以下面判断是否图文模式还不行
	category1, label, _ := models.GetCategory(id)
	// beego.Info(category1.Graphicmode)
	if category.Title == "文章/设代日记" || category1.Graphicmode {
		c.TplName = "proddiary_view.tpl"
	} else {
		c.TplName = "prod_view.tpl"
	}
	// category, _ := models.GetCategory(id)
	// if category.Title == "diary" {
	// 	c.TplName = "proddiary_view.tpl"
	// } else {
	// 	c.TplName = "prod_view.tpl"
	// }
	chengguo, _ := models.GetTopicsbyparentid(id, true)
	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
	categoryproj, _ := models.GetCategoryProj(id)
	categoryphase, _ := models.GetCategoryPhase(id)
	categoryspec, _ := models.GetCategorySpec(id)

	c.Data["CategoryProj"] = categoryproj
	c.Data["CategoryPhase"] = categoryphase
	c.Data["CategorySpec"] = categoryspec
	c.Data["Category"] = category1
	c.Data["Label"] = label
	c.Data["Chengguo"] = chengguo
	c.Data["Length"] = len(chengguo)

	//
	cid := strconv.FormatInt(categoryproj.Id, 10)
	categorycelan, _, _ := models.GetCategory(cid) //由分类id取出本身（项目名称等）
	// topics, err := models.GetAllTopics(category.Title, false)
	categorychengguo, _ := models.GetCategoryChengguo(cid)
	categoryzhuanye, _ := models.GetCategoryLeixing(cid)
	categoryjieduan, _ := models.GetCategoryJieduan(cid)
	// if err != nil {
	// 	beego.Error(err)
	// 	c.Redirect("/", 302)
	// 	return
	// }
	c.Data["Categorycelan"] = categorycelan
	c.Data["Categorychengguo"] = categorychengguo
	c.Data["Categoryzhuanye"] = categoryzhuanye
	c.Data["Categoryjieduan"] = categoryjieduan
	c.Data["Tid"] = id
	// if err != nil {
	// 	beego.Error(err)
	// 	c.Redirect("/", 302)
	// 	return
	// }
	// id := c.Input().Get("id")
	// category, err := models.GetCategory(id)
	// if err != nil {
	// beego.Error(err)
	// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
	// return
	// }
	// c.Data["Category"] = category
	// c.Data["Id"] = id
}

//查看项目简介
func (c *CategoryController) ViewBrief() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	id := c.Input().Get("id")
	category, label, _ := models.GetCategory(id)
	c.Data["Category"] = category
	c.Data["Label"] = label
	// if category.Title == "diary" {
	c.TplName = "category_view_brief.tpl"
	// } else {
	// 	c.TplName = "prod_view_b.tpl"
	// }
}

//查看专业里或第3级目录中的成果
func (c *CategoryController) View_3() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	id := c.Input().Get("id")
	// if len(id) == 0 {
	// 	break
	// }
	category, label, _ := models.GetCategory(id)
	if category.Title == "diary" {
		c.TplName = "proddiary_view.tpl"
	} else {
		c.TplName = "prod_3_view.tpl"
	}
	chengguo, _ := models.GetTopicsbyparentid(id, true)
	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
	// categoryproj, _ := models.GetCategoryProj(id)
	categoryphase, _ := models.GetCategoryPhase(id) //这里实际上相当于取得项目，即一级目录
	categoryspec, _ := models.GetCategorySpec(id)   //这里相当于取得阶段，即二级目录

	// c.Data["CategoryProj"] = categoryproj
	c.Data["CategoryPhase"] = categoryphase
	c.Data["CategorySpec"] = categoryspec
	c.Data["Category"] = category
	c.Data["Label"] = label
	c.Data["Chengguo"] = chengguo
	c.Data["Length"] = len(chengguo)

	// if err != nil {
	// 	beego.Error(err)
	// 	c.Redirect("/", 302)
	// 	return
	// }
	// id := c.Input().Get("id")
	// category, err := models.GetCategory(id)
	// if err != nil {
	// beego.Error(err)
	// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
	// return
	// }
	// c.Data["Category"] = category
	// c.Data["Id"] = id
}

//查看成果类型里的成果
func (c *CategoryController) View_b() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategoryb"] = true
	uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	id := c.Input().Get("id")
	// if len(id) == 0 {
	// 	break
	// }
	category, label, _ := models.GetCategory(id)
	if category.Title == "diary" {
		c.TplName = "proddiary_view_b.tpl"
	} else {
		c.TplName = "prod_view_b.tpl"
	}
	chengguo, _ := models.GetTopicsbyparentid(id, true)
	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
	categoryproj, _ := models.GetCategoryProj(id)
	categoryphase, _ := models.GetCategoryPhase(id)
	categoryspec, _ := models.GetCategorySpec(id)

	c.Data["CategoryProj"] = categoryproj
	c.Data["CategoryPhase"] = categoryphase
	c.Data["CategorySpec"] = categoryspec
	c.Data["Category"] = category
	c.Data["Label"] = label
	c.Data["Chengguo"] = chengguo
	c.Data["Length"] = len(chengguo)
	// if err != nil {
	// 	beego.Error(err)
	// 	c.Redirect("/", 302)
	// 	return
	// }
	// id := c.Input().Get("id")
	// category, err := models.GetCategory(id)
	// if err != nil {
	// beego.Error(err)
	// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
	// return
	// }
	// c.Data["Category"] = category
	// c.Data["Id"] = id
}

//查看成果类型里的成果
func (c *CategoryController) Category_prod_view() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	id := c.Input().Get("id")
	// beego.Info(id)
	title := c.Input().Get("title")
	// beego.Info(title)
	// if len(id) == 0 {
	// 	break
	// }
	topics, err := models.GetAllTopics(title, false)

	c.TplName = "category_prod_view.tpl"

	//由项目id获取所有成果
	// chengguo, _ := models.GetTopicsbyparentid(id, true)
	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
	category, label, _ := models.GetCategory(id)
	categoryproj, _ := models.GetCategoryProj(id)
	categoryphase, _ := models.GetCategoryPhase(id)
	categoryspec, _ := models.GetCategorySpec(id)

	c.Data["CategoryProj"] = categoryproj
	c.Data["CategoryPhase"] = categoryphase
	c.Data["CategorySpec"] = categoryspec
	c.Data["Category"] = category
	c.Data["Label"] = label
	// c.Data["Chengguo"] = chengguo
	c.Data["Chengguo"] = topics
	c.Data["Length"] = len(topics) //将这个结果写入数据库
	catid, err := strconv.ParseInt(id, 10, 64)
	// beego.Info(id)
	// beego.Info(catid)
	length := strconv.Itoa(len(topics))
	// beego.Info(length)
	length1, err := strconv.ParseInt(length, 10, 64)
	// beego.Info(length1)
	err = models.TopicCount(catid, length1)
	if err != nil {
		beego.Error(err)
		// c.Redirect("/", 302)
		// return
	}
}

//显示修改项目简介、封面的界面
func (c *CategoryController) Modify() {
	cid := c.Input().Get("cid")
	var rolename int
	var uname string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, role, _ := checkRolewrite(c.Ctx) //login里的
	rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	//3.由cid查询数据库中的项目名
	var label1 string
	category, label, err := models.GetCategory(cid)
	for i, label2 := range label {
		if i == 0 {
			label1 = label2.Title
		} else {
			label1 = label1 + "," + label2.Title
		}
	}

	if rolename > 1 && uname != category.Author { //要么管理员，要么作者自己可以修改项目
		// port := strconv.Itoa(c.Ctx.Input.Port())
		route := c.Ctx.Request.URL.String() //c.Ctx.Input.Site() + ":" + port +
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	c.TplName = "category_modify.tpl"

	categorychengguo, err := models.GetCategoryChengguo(cid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	categoryzhuanye, err := models.GetCategoryLeixing(cid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	categoryjieduan, err := models.GetCategoryJieduan(cid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}

	// category, err := models.GetCategory(cid)
	// if err != nil {
	// 	beego.Error(err)
	// 	c.Redirect("/", 302)
	// 	return
	// }

	c.Data["Category"] = category
	c.Data["Label"] = label1
	//由route截取文件名
	Attachment := strings.Split(category.Route, "/")
	s := len(Attachment)
	name := Attachment[s-1]
	c.Data["Filename"] = name
	c.Data["Categoryjieduan"] = categoryjieduan
	c.Data["Categoryzhuanye"] = categoryzhuanye
	c.Data["Categorychengguo"] = categorychengguo

	for _, v := range categoryjieduan {
		switch v.Title {
		case "A":
			c.Data["A"] = true
		case "B":
			c.Data["B"] = true
		case "C":
			c.Data["C"] = true
		case "D":
			c.Data["D"] = true
		case "E":
			c.Data["E"] = true
		case "F":
			c.Data["F"] = true
		case "G":
			c.Data["G"] = true
		case "L":
			c.Data["L"] = true
		}
	}

	for _, x := range categorychengguo {
		switch x.Title {
		case "FB":
			c.Data["FB"] = true
		case "FD":
			c.Data["FD"] = true
		case "FG":
			c.Data["FG"] = true
		case "FT":
			c.Data["FT"] = true
		case "FJ":
			c.Data["FJ"] = true
		case "FP":
			c.Data["FP"] = true
		case "Fdiary":
			c.Data["Fdiary"] = true
		}
	}
	for _, w := range categoryzhuanye {
		switch w.Title {
		case "1":
			c.Data["1"] = true
		case "2":
			c.Data["2"] = true
		case "3":
			c.Data["3"] = true
		case "4":
			c.Data["4"] = true
		case "5":
			c.Data["5"] = true
		case "6":
			c.Data["6"] = true
		case "7":
			c.Data["7"] = true
		case "8":
			c.Data["8"] = true
		case "9":
			c.Data["9"] = true
		}
	}
	// c.Data["Id"] = cid
}

//显示修改项目目录结构的界面
func (c *CategoryController) ModifyFRM() {
	cid := c.Input().Get("cid")
	var rolename int
	var uname, role string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, role, _ = checkRolewrite(c.Ctx) //login里的
	rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	//3.由cid查询数据库中的用户名
	category, label, err := models.GetCategory(cid)
	if rolename > 1 && uname != category.Author { //要么管理员，要么作者自己可以修改项目
		// port := strconv.Itoa(c.Ctx.Input.Port())
		route := c.Ctx.Request.URL.String() //c.Ctx.Input.Site() + ":" + port +
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	c.TplName = "category_modifyfrm.tpl"

	categorychengguo, err := models.GetCategoryChengguo(cid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	categoryzhuanye, err := models.GetCategoryLeixing(cid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	categoryjieduan, err := models.GetCategoryJieduan(cid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}

	// category, err := models.GetCategory(cid)
	// if err != nil {
	// 	beego.Error(err)
	// 	c.Redirect("/", 302)
	// 	return
	// }

	c.Data["Category"] = category
	c.Data["Label"] = label
	//由route截取文件名
	Attachment := strings.Split(category.Route, "/")
	s := len(Attachment)
	name := Attachment[s-1]
	c.Data["Filename"] = name
	c.Data["Categoryjieduan"] = categoryjieduan
	c.Data["Categoryzhuanye"] = categoryzhuanye
	c.Data["Categorychengguo"] = categorychengguo

	for _, v := range categoryjieduan {
		switch v.Title {
		case "A":
			c.Data["A"] = true
		case "B":
			c.Data["B"] = true
		case "C":
			c.Data["C"] = true
		case "D":
			c.Data["D"] = true
		case "E":
			c.Data["E"] = true
		case "F":
			c.Data["F"] = true
		case "G":
			c.Data["G"] = true
		case "L":
			c.Data["L"] = true
		}
	}

	for _, x := range categorychengguo {
		switch x.Title {
		case "FB":
			c.Data["FB"] = true
		case "FD":
			c.Data["FD"] = true
		case "FG":
			c.Data["FG"] = true
		case "FT":
			c.Data["FT"] = true
		case "FJ":
			c.Data["FJ"] = true
		case "FP":
			c.Data["FP"] = true
		case "Fdiary":
			c.Data["Fdiary"] = true
		}
	}
	for _, w := range categoryzhuanye {
		switch w.Title {
		case "1":
			c.Data["1"] = true
		case "2":
			c.Data["2"] = true
		case "3":
			c.Data["3"] = true
		case "4":
			c.Data["4"] = true
		case "5":
			c.Data["5"] = true
		case "6":
			c.Data["6"] = true
		case "7":
			c.Data["7"] = true
		case "8":
			c.Data["8"] = true
		case "9":
			c.Data["9"] = true
		}
	}
	// c.Data["Id"] = cid

}

//修改项目提交方法
func (c *CategoryController) ModifyCategory() {
	// if !checkAccount(c.Ctx) {
	// 	port := strconv.Itoa(c.Ctx.Input.Port())
	// 	route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/login?url="+route, 302)
	// 	// c.Redirect("/login", 302)
	// 	return
	// }
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true

	cid := c.Input().Get("cid")
	name := c.Input().Get("name")
	number := c.Input().Get("number")
	label := c.Input().Get("label")
	// content := c.Input().Get("content")
	// content := c.Input().Get("editorValue") //test-editormd-html-code
	//取得上传之封面图片
	route := c.Input().Get("route")
	cover := c.Input().Get("editor_cover")
	content := c.Input().Get("editor_property")
	// image := c.Input().Get("image")
	// path := c.Input().Get("tempString")

	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value
	uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname

	//保存上传的图片
	//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
	// _, h, err := c.GetFile("image")
	// // beego.Info(h)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// // var attachment string
	// // var path string
	// var filesize int64
	// var route string
	// if h != nil {
	// 	//保存附件
	// 	// attachment = h.Filename
	// 	// beego.Info(attachment)
	// 	path1 := ".\\attachment\\" + number + name + "\\" + h.Filename
	// 	err = c.SaveToFile("image", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	//如果包含jpg，则进行压缩
	// 	if strings.Contains(strings.ToLower(h.Filename), ".jpg") {
	// 		// 随机名称
	// 		// to := path + random_name() + ".jpg"
	// 		origin := path1 //path + file.Name()
	// 		fmt.Println("正在处理" + origin + ">>>" + origin)
	// 		cmd_resize(origin, 2048, 0, origin)
	// 		//				defer os.Remove(origin)//删除原文件
	// 	}
	// 	filesize, _ = FileSize(path1)
	// 	filesize = filesize / 1000.0
	// 	route = "/attachment/" + number + name + "/" + h.Filename
	//存入数据库cid, name, number, content, cover, path, route, uname
	err := models.ModifyCategory(cid, name, number, label, content, cover, "", route, uname)
	if err != nil {
		beego.Error(err)
	}
	// } else {
	// 	//如果没有更新图片，则图片地址不存入
	// 	err = models.ModifyCategory(cid, name, number, content, path, "", uname)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// }
	// c.Data["Uname"] = ck.Value
	c.Redirect("/category?op=view&id="+cid, 301)
	return
}

//修改目录名称提交
func (c *CategoryController) ModifyCategoryTitle() {
	cid := c.Input().Get("cid") //项目id
	id := c.Input().Get("pid")  //要修改的id
	name := c.Input().Get("title")
	//修改物理目录，
	//修改id及其下级的title
	idNum, err := strconv.ParseInt(id, 10, 64)
	err = models.ModifyCategoryTitle(idNum, name)
	// if err != nil {
	// 	beego.Error(err)
	// }
	if err != nil {
		beego.Error(err)
	} else {
		// c.Data["json"] = map[string]interface{}{
		// 	"state":    "SUCCESS",
		// 	"data":     "111",
		// 	"original": "demo.jpg",
		// }
		// c.ServeJSON()
		//返回值给ajax的data
		data := name
		c.Ctx.WriteString(data)
	}
	c.Redirect("/category/modifyfrm?cid="+cid, 301) //这个有用
	// return加这个return就初夏下面这个错误
	// 2016/05/15 15:07:56 [CategoryModel.go:818][I] .\attachment\20……
	// 2016/05/15 15:07:56 http: multiple response.WriteHeader calls
}

//添加一级目录
func (c *CategoryController) UserdefinedPostOne() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	name := c.Input().Get("title")
	pid := c.Input().Get("pid") //要添加目录的父id
	cid := c.Input().Get("cid") //项目id
	radio := c.Input().Get("radiostring")
	//由cid查出diskdirectory，category
	_, diskdirectory, err := models.GetCategoryUrl(pid)
	if err != nil {
		beego.Error(err)
	}
	err = os.MkdirAll(diskdirectory+name, 0777)
	if err != nil {
		beego.Error(err)
	}

	uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	//存入数据库
	_, err = models.AdduserdefinedCategoryOne(name, pid, radio, uname)
	if err != nil {
		beego.Error(err)
	} else {
		data := "OK!"
		c.Ctx.WriteString(data)
	}
	// id1 := strconv.FormatInt(id, 10)
	//项目id
	c.Redirect("/category/modifyfrm?cid="+cid, 301)
	// return
}

// func (c *TopicController) View() {
// 	c.TplName = "topic_view.html"
// 	topic, err := models.GetTopic(c.Ctx.Input.Param("0"))
// 	//articleId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
// 	//id, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
// 	if err != nil {
// 		beego.Error(err)
// 		c.Redirect("/", 302)
// 		return
// 	}
// 	c.Data["Topic"] = topic
// 	c.Data["Tid"] = c.Ctx.Input.Param("0") //教程中用的是圆括号，导致错误topic.go:52: cannot call non-function c.Controller.Ctx.Input.Params (type map[string]string)
// 	//教程第8章开头有修改
// 	replies, err := models.GetAllReplies(c.Ctx.Input.Param("0"))
// 	if err != nil {
// 		beego.Error(err)
// 		return
// 	}
// 	c.Data["Replies"] = replies
// 	c.Data["IsLogin"] = checkAccount(c.Ctx)
// }
