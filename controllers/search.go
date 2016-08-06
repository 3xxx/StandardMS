package controllers

import (
	"github.com/astaxie/beego"
	// "path"
	"quick/models"
)

type SearchController struct {
	beego.Controller
}

//搜索项目
func (c *SearchController) SearchCategory() { //search用的是get方法
	tid := c.Input().Get("categoryname")
	if tid != "" {
		c.Data["IsCategory"] = true
		// c.Data["IsSearch"] = true
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.TplName = "searchcategory.tpl"
		Searchs, labels, err := models.SearchCategories(tid, false)
		if err != nil {
			beego.Error(err.Error)
		} else {
			c.Data["Searchs"] = Searchs
			c.Data["Label"] = labels
		}
	} else {
		c.Data["IsCategory"] = true
		// c.Data["IsSearch"] = true
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.TplName = "searchcategory.tpl"
	}
}

//搜索成果
func (c *SearchController) SearchProduction() { //search用的是get方法
	tid := c.Input().Get("topic")
	if tid != "" {
		c.Data["IsTopic"] = true
		// c.Data["IsSearch"] = true
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.TplName = "searchtopic.tpl"
		Searchs, err := models.SearchTopics(tid, false)
		if err != nil {
			beego.Error(err.Error)
		} else {
			c.Data["Searchs"] = Searchs
		}
	} else {
		c.Data["IsTopic"] = true
		// c.Data["IsSearch"] = true
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.TplName = "searchtopic.tpl"
	}
}

//搜索wiki
func (c *SearchController) SearchWiki() { //search用的是get方法
	tid := c.Input().Get("wikiname")
	if tid != "" {
		c.Data["IsWiki"] = true
		// c.Data["IsSearch"] = true
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.TplName = "searchwiki.tpl"
		Searchs, err := models.SearchWikis(tid, false)
		if err != nil {
			beego.Error(err.Error)
		} else {
			c.Data["Searchs"] = Searchs
		}
	} else {
		c.Data["IsWiki"] = true
		// c.Data["IsSearch"] = true
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.TplName = "searchwiki.tpl"
	}
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
