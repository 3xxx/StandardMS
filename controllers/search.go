package controllers

import (
	"github.com/astaxie/beego"
	// "path"
	"quick/models"
)

type SearchController struct {
	beego.Controller
}

//导航栏搜索
func (c *SearchController) Get() { //search用的是get方法
	tid := c.Input().Get("tuming")
	c.Data["IsSearch"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "search.tpl"
	Searchs, err := models.SearchTopics(tid, false)
	if err != nil {
		beego.Error(err.Error)
	} else {
		c.Data["Searchs"] = Searchs
	}
	//var err error
	//	c.Data["Search"], err = models.GetAllSearchs()
	//	if err != nil {
	//		beego.Error(err)
	//	}
}

//水利设计院本地搜索
func (c *SearchController) Searchlocal() { //search用的是get方法
	tid := c.Input().Get("name")
	c.Data["IsSearch"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	// c.TplName = "search.tpl"
	Searchs, err := models.SearchTopics(tid, false)
	if err != nil {
		beego.Error(err.Error)
	} else {
		// c.Data["Searchs"] = Searchs
		c.Data["json"] = Searchs //这里必须要是c.Data["json"]，其他c.Data["Data"]不行
		c.ServeJSON()
	}

	//var err error
	//	c.Data["Search"], err = models.GetAllSearchs()
	//	if err != nil {
	//		beego.Error(err)
	//	}
}

// func (c *SearchController) Post() {
// 	// if !checkAccount(c.Ctx) {
// 	// 	c.Redirect("/login", 302)
// 	// 	return
// 	// }
// 	//解析表单
// 	tid := c.Input().Get("tuming") //教程里漏了这句，导致修改总是变成添加文章
// 	// title := c.Input().Get("title")
// 	// content := c.Input().Get("content")
// 	// category := c.Input().Get("category")
// 	c.Data["IsSearch"] = true
// 	c.Data["IsLogin"] = checkAccount(c.Ctx)
// 	c.TplName = "search.tpl"
// 	Searchs, err := models.SearchTopics(tid, false) //这里传入空字符串
// 	if err != nil {
// 		beego.Error(err.Error)
// 	} else {
// 		c.Data["Searchs"] = Searchs
// 	}
// 	//获取附件
// 	_, fh, err := c.GetFile("attachment")
// 	if err != nil {
// 		beego.Error(err)
// 	}

// 	var attachment string
// 	if fh != nil {
// 		//保存附件
// 		attachment = fh.Filename
// 		beego.Info(attachment)
// 		err = c.SaveToFile("attachment", path.Join("attachment", attachment))
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 	}

// 	c.Redirect("/search", 302)
// }
