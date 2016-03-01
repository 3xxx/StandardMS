package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/tealeg/xlsx"
	"os"
	"path"
	"path/filepath"
	"quick/models"
	// "regexp"
	"strconv"
	"strings"
)

type Uploadimages struct {
	url     string
	message string
	success string
}
type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() { //这个给爬虫用。而为了配合pagenate，用后面的listall()
	c.Data["IsTopic"] = true //这里修改到ListAllPosts()
	c.TplName = "topic.tpl"
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

//根据用户名查看项目
func (c *TopicController) Viewbyuname() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsTopic"] = true
	c.TplName = "topic_uname.tpl"
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
	uname := c.Input().Get("uname")
	topic, _ := models.Gettopicsbyuname(uname) //由uname取出项目
	c.Data["Topics"] = topic
}

func (c *TopicController) Add() { //参考下面的 modify,这个add是topic/add

	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
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
		c.Data["Uname"] = v.(string)
	}
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err == nil {
	// 	c.Data["Uname"] = ck.Value
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
	role, _ := checkRole(c.Ctx) //login里的
	// beego.Info(role)
	//4.进行逻辑分析：
	rolename, _ := strconv.ParseInt(role, 10, 64)
	if rolename > 4 { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}

	id := c.Input().Get("id")
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsTopic"] = true

	mid := c.Input().Get("mid")
	// if mid == "1" {
	// } else {
	// 	c.TplName = "topic_add2.html"
	// }
	switch mid {
	case "1":
		c.TplName = "topic_one_add.html"
	case "2":
		c.TplName = "topic_many_add.html"
	case "3": //添加设代日记
		c.TplName = "diary_add.html"
	case "4": //自定义一对一模式
		c.TplName = "topic_user_one_add.html"
	case "5": //自定义一对多模式
		c.TplName = "topic_user_many_add.html"
	// default:
	// fmt.Printf("Default")
	case "6": //用百度的插件上传
		c.TplName = "topic_one_addbaidu.html"
	case "7": //用自动识别图号上传
		c.TplName = "topic_one_addstandard.html"
	case "8": //用自动识别图号上传
		c.TplName = "topic_many_addbaidu.html"
	}
	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
	if id != "" {
		categoryproj, err := models.GetCategoryProj(id)
		if err != nil {
			beego.Error(err)
			// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
			return
		}
		categoryphase, err := models.GetCategoryPhase(id)
		if err != nil {
			beego.Error(err)
			// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
			return
		}
		categoryspec, err := models.GetCategorySpec(id)
		if err != nil {
			beego.Error(err)
			// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
			return
		}
		category, err := models.GetCategory(id)
		if err != nil {
			beego.Error(err)
			// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
			return
		}

		c.Data["CategoryProj"] = categoryproj
		c.Data["CategoryPhase"] = categoryphase
		c.Data["CategorySpec"] = categoryspec
		c.Data["Category"] = category
		c.Data["Id"] = id
	}
}

func (c *TopicController) Post() { //这个post属于topic_modify.html提交修改。
	// if !checkAccount(c.Ctx) { //这里应该不需要判断
	// 	c.Redirect("/login", 302)
	// 	return
	// }
	//解析表单
	tid := c.Input().Get("tid") //教程里漏了这句，导致修改总是变成添加文章
	// beego.Info(tid)
	title := c.Input().Get("title")
	tnumber := c.Input().Get("tnumber")
	// beego.Info(tnumber)
	content := c.Input().Get("content")
	category := c.Input().Get("category")
	categoryid := c.Input().Get("categoryid")
	// categoryid, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	return
	// }
	//获取附件
	// _, fh, err := c.GetFile("attachment")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// var attachment string
	// if fh != nil {
	//保存附件
	// attachment = fh.Filename
	// beego.Info(attachment)
	// err = c.SaveToFile("attachment", path.Join("attachment", attachment))
	//err = models.AddTopic(title, category, content, attachment) //这句有用，但仍然没存进去category
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// }
	//将附件的编号和名称写入数据库
	// filename1, filename2 := SubStrings(attachment)
	// title = filename2
	// tnumber = filename1

	//var err error
	//var tid string //这里是增加的，不知为何教程没有
	// if len(tid) == 0 {
	// _, err = models.AddTopic(title, tnumber, category, categoryid, content, attachment)
	// } else {
	err := models.ModifyTopic(tid, title, tnumber, category, categoryid, content)
	// }
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic/view_b/"+tid, 302) //回到修改后的文章
}

func (c *TopicController) AddTopic() { //这个是否作废？？
	if !checkAccount(c.Ctx) { //这里应该不需要
		c.Redirect("/login", 302)
		return
	}
	//解析表单
	tid := c.Input().Get("tid") //教程里漏了这句，导致修改总是变成添加文章
	title := c.Input().Get("title")
	tnumber := c.Input().Get("title")
	content := c.Input().Get("content")
	category := c.Input().Get("category")
	categoryid := c.Input().Get("categoryid")

	_, h, err := c.GetFile("image") //获取上传的文件
	if err != nil {
		beego.Error(err)
	}
	var attachment string
	if h != nil {
		//保存附件
		attachment = h.Filename
		// beego.Info(attachment)

		// err = c.SaveToFile("attachment", path.Join("attachment", attachment))
		// path := c.Input().Get("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()   // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("image", path.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
	}

	if title == "" || tnumber == "" {
		//将附件的编号和名称写入数据库
		filename1, filename2 := SubStrings(attachment)
		title = filename2
		tnumber = filename1
	}
	// var err error
	// var tid string //这里是增加的，不知为何教程没有
	//获取用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	// if v != nil {
	uname := v.(string)
	// }
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value

	if len(tid) == 0 {
		_, err = models.AddTopicOne(title, tnumber, category, categoryid, content, uname, attachment)
		// beego.Info(attachment)
	} else {
		err = models.ModifyTopic(tid, title, tnumber, category, categoryid, content)
	}
	if err != nil {
		beego.Error(err)
	}
	// c.Redirect("/topic", 302)
}

func (c *TopicController) Topic_many_add() { //一对多模式
	//解析表单
	tid := c.Input().Get("tid") //教程里漏了这句，导致修改总是变成添加文章
	title := c.Input().Get("title")
	tnumber := c.Input().Get("tnumber")
	content := c.Input().Get("content")
	category := c.Input().Get("category")
	categoryid := c.Input().Get("categoryid")

	//获取文件保存路径，有了categoryid可以求出整个路径
	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
	// categoryproj, err := models.GetCategoryProj(categoryid)
	// categoryphase, err := models.GetCategoryPhase(categoryid)
	// categoryspec, err := models.GetCategorySpec(categoryid)
	category1, err := models.GetCategory(categoryid)
	if err != nil {
		beego.Error(err)
		// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
		return
	}
	//获取上传的文件
	_, h, err := c.GetFile("image")
	if err != nil {
		beego.Error(err)
	}
	var attachment string
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		attachment = h.Filename
		// beego.Info(attachment)
		// path = ".\\attachment\\" + categoryproj.Number + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename
		path = category1.DiskDirectory + h.Filename
		// path := c.Input().Get("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("image", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
	}
	if title == "" || tnumber == "" {
		//将附件的编号和名称写入数据库
		filename1, filename2 := SubStrings(attachment)
		if filename1 == "" {
			filename1 = filename2 //如果编号为空，则用文件名代替，否则多个编号为空导致存入数据库唯一性检查错误
		}
		title = filename2
		tnumber = filename1
	}
	// var err error
	// var tid string //这里是增加的，不知为何教程没有
	// path := ".\\attachment\\" + categoryproj.Number + " " + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename
	//获取用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	// if v != nil {
	uname := v.(string)
	// }
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value

	// route := "/attachment/" + categoryproj.Number + categoryproj.Title + "/" + categoryphase.Title + "/" + categoryspec.Title + "/" + category + "/" + h.Filename
	route := category1.Url + h.Filename
	//topicid := c.Input().Get("topicid")
	var topicid int64
	if len(tid) == 0 {
		topicid, err = models.AddTopicMany(title, tnumber, category, categoryid, uname, content, attachment)
		//这里返回topicid，并存入attachment表中
		if err != nil { //如果发生错误，返回错误，并获取该文章的topicid
			beego.Error(err)
		}
		if topicid == 0 { //这个已经不用了。
			topicid, err = models.GetTopicIdbytnumber(tnumber)
		}
		cid := strconv.FormatInt(topicid, 10)
		filesize := strconv.FormatInt(filesize, 10)
		err = models.AddAttachment(attachment, filesize, path, route, cid, uname)
		// beego.Info(attachment)
	} else {
		err = models.ModifyTopic(tid, title, tnumber, category, categoryid, content)
	}
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "topic_many_add.tpl" //不加这句上传出错，虽然可以成功上传
	// c.Redirect("/topic", 302)
}

func (c *TopicController) Topic_one_add() { //一对一模式
	//解析表单
	tid := c.Input().Get("tid") //教程里漏了这句，导致修改总是变成添加文章
	title := c.Input().Get("title")
	tnumber := c.Input().Get("tnumber")
	content := c.Input().Get("content")
	category := c.Input().Get("category")
	categoryid := c.Input().Get("categoryid")

	//获取文件保存路径，有了categoryid可以求出整个路径
	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
	// categoryproj, err := models.GetCategoryProj(categoryid)
	// categoryphase, err := models.GetCategoryPhase(categoryid)
	// categoryspec, err := models.GetCategorySpec(categoryid)
	category1, err := models.GetCategory(categoryid)
	if err != nil {
		beego.Error(err)
		// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
		return
	}
	//获取上传的文件
	_, h, err := c.GetFile("image")
	beego.Info(h.Filename)
	if err != nil {
		beego.Error(err)
	}
	var attachment string
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		attachment = h.Filename
		// beego.Info(attachment)
		// path = ".\\attachment\\" + categoryproj.Number + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename
		path = category1.DiskDirectory + h.Filename
		// path := c.Input().Get("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("image", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
	}
	if title == "" || tnumber == "" {
		//将附件的编号和名称写入数据库
		filename1, filename2 := SubStrings(attachment)
		//当2个文件都取不到filename1的时候，数据库里的tnumber的唯一性检查出错。
		// beego.Info(filename1)
		// beego.Info(filename2)
		if filename1 == "" {
			filename1 = filename2 //如果编号为空，则用文件名代替，否则多个编号为空导致存入数据库唯一性检查错误
		}
		tnumber = filename1
		title = filename2
	}
	//获取用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	// if v != nil {
	uname := v.(string)
	// }
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value
	// var err error
	// var tid string //这里是增加的，不知为何教程没有
	// path := ".\\attachment\\" + categoryproj.Number + " " + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename
	// route := "/attachment/{{.TopicProj.Number}} {{.TopicProj.Title}}/{{.TopicPhase.Title}}/{{.TopicSpec.Title}}/{{.TopicChengguo.Title}}/{{.Topic.Attachment}}"
	// route := "/attachment/" + categoryproj.Number + categoryproj.Title + "/" + categoryphase.Title + "/" + categoryspec.Title + "/" + category + "/" + h.Filename
	route := category1.Url + h.Filename
	// topicid := c.Input().Get("topicid")
	var topicid int64
	if len(tid) == 0 {
		topicid, err = models.AddTopicOne(title, tnumber, category, categoryid, uname, content, attachment)
		if err != nil {
			beego.Error(err)
		}
		cid := strconv.FormatInt(topicid, 10)
		filesize := strconv.FormatInt(filesize, 10)
		err = models.AddAttachment(attachment, filesize, path, route, cid, uname)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(attachment)
	} else {
		err = models.ModifyTopic(tid, title, tnumber, category, categoryid, content)
	}
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "topic_one_add.tpl" //不加这句上传出错，虽然可以成功上传
	// c.Redirect("/topic", 302)
}

func (c *TopicController) Topic_one_addstandard() { //一对一上传，自动识别
	//获取文件名——解析出 项目编号、阶段、专业、类型、扩展名 和成果名称，查询数据库是否有此项目
	//没有项目则建立？？
	//ueditor中的附件如何处理呢？
	content := c.Input().Get("content")
	beego.Info(content)
	//获取上传的文件
	_, h, err := c.GetFile("file") //ueditor用upfile
	if err != nil {
		beego.Error(err)
	}
	//Suffix
	// beego.Info(h.Filename)
	_, FileNumber, FileName, ProNumber, ProJieduan, ProLeixing, ProZhuanye := Record(h.Filename)
	// beego.Info(FileNumber)
	// beego.Info(FileName)
	// beego.Info(ProNumber)
	// beego.Info(ProJieduan)
	// beego.Info(ProLeixing)
	// beego.Info(ProZhuanye)
	//由项目号查出项目名称
	category, err := models.GetCategoryTitle(ProNumber)
	if err != nil {
		beego.Error(err)
		// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
		return
	}
	//由项目编号、阶段、文件类型、专业查出专业Id.
	Id, err := models.GetCategoryzhuanye(ProNumber, ProJieduan, ProLeixing, ProZhuanye)
	ProLeixingId := strconv.FormatInt(Id, 10)
	// var attachment string
	// var path string
	// var filesize int64
	// if h != nil {
	//保存附件
	attachment := h.Filename
	// ".\\attachment\\" + categoryproj.Number + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\"
	filepath := ".\\attachment\\" + ProNumber + category.Title + "\\" + ProJieduan + "\\" + ProLeixing + "\\" + ProZhuanye + "\\" + h.Filename
	// 关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = c.SaveToFile("file", filepath) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
	if err != nil {
		beego.Error(err)
	}
	filesize, _ := FileSize(filepath)
	filesize = filesize / 1000.0
	// }
	// if title == "" || tnumber == "" {
	// 	//将附件的编号和名称写入数据库
	// 	filename1, filename2 := SubStrings(attachment)
	// 	//当2个文件都取不到filename1的时候，数据库里的tnumber的唯一性检查出错。
	// 	if filename1 == "" {
	// 		filename1 = filename2 //如果编号为空，则用文件名代替，否则多个编号为空导致存入数据库唯一性检查错误
	// 	}
	// 	tnumber = filename1
	// 	title = filename2
	// }
	//获取用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	// if v != nil {
	uname := v.(string)
	// }
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value
	route := "/attachment/" + ProNumber + category.Title + "/" + ProJieduan + "/" + ProLeixing + "/" + ProZhuanye + "/" + h.Filename
	var topicid int64
	// if len(tid) == 0 {
	topicid, err = models.AddTopicOne(FileName, FileNumber, ProLeixing, ProLeixingId, uname, content, attachment)
	if err != nil {
		beego.Error(err)
	}
	cid := strconv.FormatInt(topicid, 10)
	filesize1 := strconv.FormatInt(filesize, 10)
	err = models.AddAttachment(attachment, filesize1, filepath, route, cid, uname)
	if err != nil {
		beego.Error(err)
	}
	// } else {
	// 	err = models.ModifyTopic(tid, title, tnumber, category, categoryid, content)
	// }
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = map[string]interface{}{
			"state":    "SUCCESS",
			"title":    "111",
			"original": "demo.jpg",
			"url":      route,
		}
		c.ServeJSON()
	}
}

func (c *TopicController) Topic_many_addbaidu() { //一对多模式
	//解析表单
	tid := c.Input().Get("tid") //这句没用。教程里漏了这句，导致修改总是变成添加文章
	title := c.Input().Get("title")
	tnumber := c.Input().Get("tnumber")
	content := c.Input().Get("content")
	category := c.Input().Get("category")
	categoryid := c.Input().Get("categoryid")

	//获取文件保存路径，有了categoryid可以求出整个路径
	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
	// categoryproj, err := models.GetCategoryProj(categoryid)
	// categoryphase, err := models.GetCategoryPhase(categoryid)
	// categoryspec, err := models.GetCategorySpec(categoryid)
	category1, err := models.GetCategory(categoryid)
	if err != nil {
		beego.Error(err)
		// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
		return
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	var attachment string
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		attachment = h.Filename
		// beego.Info(attachment)
		// path = ".\\attachment\\" + categoryproj.Number + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename
		path = category1.DiskDirectory + h.Filename
		// path := c.Input().Get("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
	}
	if title == "" || tnumber == "" {
		//将附件的编号和名称写入数据库
		filename1, filename2 := SubStrings(attachment)
		if filename1 == "" {
			filename1 = filename2 //如果编号为空，则用文件名代替，否则多个编号为空导致存入数据库唯一性检查错误
		}
		title = filename2
		tnumber = filename1
	}
	// var err error
	// var tid string //这里是增加的，不知为何教程没有
	// path := ".\\attachment\\" + categoryproj.Number + " " + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename
	//获取用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	// if v != nil {
	uname := v.(string)
	// }
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value

	// route := "/attachment/" + categoryproj.Number + categoryproj.Title + "/" + categoryphase.Title + "/" + categoryspec.Title + "/" + category + "/" + h.Filename
	route := category1.Url + h.Filename
	//topicid := c.Input().Get("topicid")
	var topicid int64
	if len(tid) == 0 {
		topicid, err = models.AddTopicMany(title, tnumber, category, categoryid, uname, content, attachment)
		//这里返回topicid，并存入attachment表中
		if err != nil { //如果发生错误，返回错误，并获取该文章的topicid
			beego.Error(err)
		}
		if topicid == 0 { //这个已经不用了。
			topicid, err = models.GetTopicIdbytnumber(tnumber)
		}
		cid := strconv.FormatInt(topicid, 10)
		filesize := strconv.FormatInt(filesize, 10)
		err = models.AddAttachment(attachment, filesize, path, route, cid, uname)
		// beego.Info(attachment)
	} else {
		err = models.ModifyTopic(tid, title, tnumber, category, categoryid, content)
	}
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "topic_many_add.tpl" //不加这句上传出错，虽然可以成功上传
	// c.Redirect("/topic", 302)
}

func (c *TopicController) Topic_one_addbaidu() { //一对一模式
	//解析表单
	tid := c.Input().Get("tid") //教程里漏了这句，导致修改总是变成添加文章
	title := c.Input().Get("title")
	tnumber := c.Input().Get("tnumber")
	content := c.Input().Get("content")
	category := c.Input().Get("category")
	categoryid := c.Input().Get("categoryid")

	//获取文件保存路径，有了categoryid可以求出整个路径
	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
	// categoryproj, err := models.GetCategoryProj(categoryid)
	// categoryphase, err := models.GetCategoryPhase(categoryid)
	// categoryspec, err := models.GetCategorySpec(categoryid)
	category1, err := models.GetCategory(categoryid)
	if err != nil {
		beego.Error(err)
		// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
		return
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	var attachment string
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		attachment = h.Filename
		// beego.Info(attachment)
		// path = ".\\attachment\\" + categoryproj.Number + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename
		path = category1.DiskDirectory + h.Filename
		// path := c.Input().Get("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
	}
	if title == "" || tnumber == "" {
		//将附件的编号和名称写入数据库
		filename1, filename2 := SubStrings(attachment)
		//当2个文件都取不到filename1的时候，数据库里的tnumber的唯一性检查出错。
		// beego.Info(filename1)
		// beego.Info(filename2)
		if filename1 == "" {
			filename1 = filename2 //如果编号为空，则用文件名代替，否则多个编号为空导致存入数据库唯一性检查错误
		}
		tnumber = filename1
		title = filename2
	}
	//获取用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	// if v != nil {
	uname := v.(string)
	// }

	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value
	// var err error
	// var tid string //这里是增加的，不知为何教程没有
	// path := ".\\attachment\\" + categoryproj.Number + " " + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename
	// route := "/attachment/{{.TopicProj.Number}} {{.TopicProj.Title}}/{{.TopicPhase.Title}}/{{.TopicSpec.Title}}/{{.TopicChengguo.Title}}/{{.Topic.Attachment}}"
	// route := "/attachment/" + categoryproj.Number + categoryproj.Title + "/" + categoryphase.Title + "/" + categoryspec.Title + "/" + category + "/" + h.Filename
	route := category1.Url + h.Filename
	// topicid := c.Input().Get("topicid")
	var topicid int64
	if len(tid) == 0 {
		topicid, err = models.AddTopicOne(title, tnumber, category, categoryid, uname, content, attachment)
		if err != nil {
			beego.Error(err)
		}
		cid := strconv.FormatInt(topicid, 10)
		filesize := strconv.FormatInt(filesize, 10)
		err = models.AddAttachment(attachment, filesize, path, route, cid, uname)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(attachment)
	} else {
		err = models.ModifyTopic(tid, title, tnumber, category, categoryid, content)
	}
	if err != nil {
		beego.Error(err)
	} else {
		// f := Uploadimages{
		// 	url:     route,
		// 	success: "1",
		// 	message: "ok",
		// }
		// c.Data["json"] = f
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": "111", "original": "demo.jpg", "url": route}
		c.ServeJSON()
	}
	// c.TplName = "topic_one_add.tpl" //不加这句上传出错，虽然可以成功上传
	// c.Redirect("/topic", 302)
	// success : 0 | 1,           // 0 表示上传失败，1 表示上传成功
	//    message : "提示的信息，上传成功或上传失败及错误信息等。",
	//    url     : "图片地址"        // 上传成功时才返回
}

func (c *TopicController) Uploadimagesmd() {
	//解析表单
	tid := c.Input().Get("tid") //教程里漏了这句，导致修改总是变成添加文章
	title := c.Input().Get("title")
	tnumber := c.Input().Get("tnumber")
	content := c.Input().Get("content")
	category := c.Input().Get("category")
	categoryid := c.Input().Get("categoryid")

	//获取文件保存路径，有了categoryid可以求出整个路径
	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
	// categoryproj, err := models.GetCategoryProj(categoryid)
	// categoryphase, err := models.GetCategoryPhase(categoryid)
	// categoryspec, err := models.GetCategorySpec(categoryid)
	category1, err := models.GetCategory(categoryid)
	if err != nil {
		beego.Error(err)
		// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
		return
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	var attachment string
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		attachment = h.Filename
		// beego.Info(attachment)
		// path = ".\\attachment\\" + categoryproj.Number + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename
		path = category1.DiskDirectory + h.Filename
		// path := c.Input().Get("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
	}
	if title == "" || tnumber == "" {
		//将附件的编号和名称写入数据库
		filename1, filename2 := SubStrings(attachment)
		//当2个文件都取不到filename1的时候，数据库里的tnumber的唯一性检查出错。
		// beego.Info(filename1)
		// beego.Info(filename2)
		if filename1 == "" {
			filename1 = filename2 //如果编号为空，则用文件名代替，否则多个编号为空导致存入数据库唯一性检查错误
		}
		tnumber = filename1
		title = filename2
	}

	//获取用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	// if v != nil {
	uname := v.(string)
	// }
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value
	// var err error
	// var tid string //这里是增加的，不知为何教程没有
	// path := ".\\attachment\\" + categoryproj.Number + " " + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename
	// route := "/attachment/{{.TopicProj.Number}} {{.TopicProj.Title}}/{{.TopicPhase.Title}}/{{.TopicSpec.Title}}/{{.TopicChengguo.Title}}/{{.Topic.Attachment}}"
	// route := "/attachment/" + categoryproj.Number + categoryproj.Title + "/" + categoryphase.Title + "/" + categoryspec.Title + "/" + category + "/" + h.Filename
	route := category1.Url + h.Filename
	// topicid := c.Input().Get("topicid")
	var topicid int64
	if len(tid) == 0 {
		topicid, err = models.AddTopicOne(title, tnumber, category, categoryid, uname, content, attachment)
		if err != nil {
			beego.Error(err)
		}
		cid := strconv.FormatInt(topicid, 10)
		filesize := strconv.FormatInt(filesize, 10)
		err = models.AddAttachment(attachment, filesize, path, route, cid, uname)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(attachment)
	} else {
		err = models.ModifyTopic(tid, title, tnumber, category, categoryid, content)
	}
	if err != nil {
		beego.Error(err)
	} else {
		f := Uploadimages{
			url:     route,
			success: "1",
			message: "ok",
		}
		c.Data["json"] = f
		c.ServeJSON()
	}
	// c.TplName = "topic_one_add.tpl" //不加这句上传出错，虽然可以成功上传
	// c.Redirect("/topic", 302)
	// success : 0 | 1,           // 0 表示上传失败，1 表示上传成功
	//    message : "提示的信息，上传成功或上传失败及错误信息等。",
	//    url     : "图片地址"        // 上传成功时才返回
}

func (c *TopicController) Diary_add() { //文章添加
	//解析表单
	// tid := c.Input().Get("tid") //这个没用到。教程里漏了这句，导致修改总是变成添加文章
	title := c.Input().Get("title")
	tnumber := c.Input().Get("tnumber")
	content := c.Input().Get("editorValue") //content
	category := c.Input().Get("category")
	categoryid := c.Input().Get("categoryid")

	//获取文件保存路径，有了categoryid可以求出整个路径
	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
	// categoryproj, err := models.GetCategoryProj(categoryid)
	// categoryphase, err := models.GetCategoryPhase(categoryid)
	// categoryspec, err := models.GetCategorySpec(categoryid)
	// category1, err := models.GetCategory(categoryid)
	// if err != nil {
	// 	beego.Error(err)
	// 	// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
	// 	return
	// }
	//获取用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	uname := v.(string)

	//存入数据库
	// id, err := models.AddCategory(name, number, content, path, "", uname, "", "")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// id1 := strconv.FormatInt(id, 10)
	// c.Redirect("/category?op=view&id="+id1, 301)
	// return //???

	// var topicid int64
	id, err := models.AddTopicMany(title, tnumber, category, categoryid, uname, content, "")
	if err != nil { //如果发生错误，返回错误，并获取该文章的topicid
		beego.Error(err)
	}
	c.TplName = "diary_add.tpl" //不加这句上传出错，虽然可以成功上传
	id1 := strconv.FormatInt(id, 10)
	// c.Redirect("/category?op=view&id="+id1, 301)
	c.Redirect("/topic/view_b/"+id1, 301)
	return //???
}

func (c *TopicController) Diary_add_back() { //这个作废！！日记上传图片——进行压缩
	//解析表单
	tid := c.Input().Get("tid") //这个没用到。教程里漏了这句，导致修改总是变成添加文章
	title := c.Input().Get("title")
	tnumber := c.Input().Get("tnumber")
	content := c.Input().Get("content")
	category := c.Input().Get("category")
	categoryid := c.Input().Get("categoryid")

	//获取文件保存路径，有了categoryid可以求出整个路径
	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
	// categoryproj, err := models.GetCategoryProj(categoryid)
	// categoryphase, err := models.GetCategoryPhase(categoryid)
	// categoryspec, err := models.GetCategorySpec(categoryid)
	category1, err := models.GetCategory(categoryid)
	if err != nil {
		beego.Error(err)
		// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
		return
	}
	//获取上传的文件
	_, h, err := c.GetFile("image")
	if err != nil {
		beego.Error(err)
	}
	var attachment string
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		attachment = h.Filename
		// beego.Info(attachment)
		// path = ".\\attachment\\" + categoryproj.Number + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename
		path = category1.DiskDirectory + h.Filename
		// path := c.Input().Get("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("image", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		//如果扩展名为jpg
		// if strings.ToLower(path.Ext(h.Filename)) == ".jpg" {

		// }
		//如果包含jpg，则进行压缩
		if strings.Contains(strings.ToLower(h.Filename), ".jpg") { //ToLower转成小写
			// 随机名称
			// to := path + random_name() + ".jpg"
			origin := path //path + file.Name()
			fmt.Println("正在处理" + origin + ">>>" + origin)
			cmd_resize(origin, 2048, 0, origin)
			//				defer os.Remove(origin)//删除原文件
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
	}
	if title == "" || tnumber == "" {
		//将附件的编号和名称写入数据库
		filename1, filename2 := SubStrings(attachment)
		if filename1 == "" {
			filename1 = filename2 //如果编号为空，则用文件名代替，否则多个编号为空导致存入数据库唯一性检查错误
		}
		title = filename2
		tnumber = filename1
	}
	// var err error
	// var tid string //这里是增加的，不知为何教程没有
	// path := ".\\attachment\\" + categoryproj.Number + " " + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename
	ck, err := c.Ctx.Request.Cookie("uname")
	if err != nil {
		beego.Error(err)
	}
	uname := ck.Value

	// route := "/attachment/" + categoryproj.Number + categoryproj.Title + "/" + categoryphase.Title + "/" + categoryspec.Title + "/" + category + "/" + h.Filename
	route := category1.Url + h.Filename
	//topicid := c.Input().Get("topicid")
	var topicid int64
	if len(tid) == 0 {
		topicid, err = models.AddTopicMany(title, tnumber, category, categoryid, uname, content, attachment)
		//这里返回topicid，并存入attachment表中
		if err != nil { //如果发生错误，返回错误，并获取该文章的topicid
			beego.Error(err)
		}
		if topicid == 0 {
			topicid, err = models.GetTopicIdbytnumber(tnumber)
		}
		cid := strconv.FormatInt(topicid, 10)

		//把这个cid返回给uploader
		c.Ctx.WriteString(cid)

		filesize := strconv.FormatInt(filesize, 10)
		err = models.AddAttachment(attachment, filesize, path, route, cid, uname)
		// beego.Info(attachment)
	} else {
		err = models.ModifyTopic(tid, title, tnumber, category, categoryid, content)
	}
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "diary_add.tpl" //不加这句上传出错，虽然可以成功上传
	// c.Redirect("/topic", 302)
}

func (c *TopicController) Diary_second_add() { //二次存入设代日记中图片的说明
	//解析表单
	// jsoninfo := c.Ctx.Input.Query("aid")
	// beego.Info(jsoninfo)
	tid := c.Input().Get("tid")
	aid := c.GetStrings("aid")
	// beego.Info(aid[0])
	content := c.GetStrings("content")
	for i, _ := range aid {
		//由图片附件的id，存入图片的content
		err := models.ModifyAttachment(aid[i], content[i])
		if err != nil {
			beego.Error(err)
		}
	}
	// c.TplName = "addtopic3.tpl"
	c.Redirect("/topic/view_b/"+tid, 302) //这里应该是跳到当前日记
}

func (c *TopicController) View() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsTopic"] = true
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
	//这里是通过文章的id获得文章及上级目录情况
	topicproj, err := models.GetTopicProj(c.Ctx.Input.Param("0"))
	topicphase, err := models.GetTopicPhase(c.Ctx.Input.Param("0"))
	topicspec, err := models.GetTopicSpec(c.Ctx.Input.Param("0"))
	_, topicchengguo, err := models.GetTopicChengguo(c.Ctx.Input.Param("0"))
	if topicchengguo.Title == "diary" {
		c.TplName = "diary_view1.html"
	} else {
		c.TplName = "topic_view.html"
	}

	topic, attachment, err := models.GetTopic(c.Ctx.Input.Param("0"))
	//articleId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
	//id, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["TopicProj"] = topicproj
	c.Data["TopicPhase"] = topicphase
	c.Data["TopicSpec"] = topicspec
	c.Data["TopicChengguo"] = topicchengguo
	c.Data["Topic"] = topic
	c.Data["Attachment"] = attachment
	c.Data["Tid"] = c.Ctx.Input.Param("0") //教程中用的是圆括号，导致错误topic.go:52: cannot call non-function c.Controller.Ctx.Input.Params (type map[string]string)
	//教程第8章开头有修改
	replies, err := models.GetAllReplies(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		return
	}
	c.Data["Replies"] = replies
	c.Data["IsLogin"] = checkAccount(c.Ctx)
}

//设代日记查看，全页模式
func (c *TopicController) View_b() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsTopic"] = true
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
	//这里是通过文章的id获得文章及上级目录情况
	topicproj, err := models.GetTopicProj(c.Ctx.Input.Param("0"))
	topicphase, err := models.GetTopicPhase(c.Ctx.Input.Param("0"))
	topicspec, err := models.GetTopicSpec(c.Ctx.Input.Param("0"))
	_, topicchengguo, err := models.GetTopicChengguo(c.Ctx.Input.Param("0"))
	if topicchengguo.Title == "diary" {
		c.TplName = "diary_view1_b.html"
	} else {
		c.TplName = "topic_view_b.html"
	}

	topic, attachment, err := models.GetTopic(c.Ctx.Input.Param("0"))
	//articleId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
	//id, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["TopicProj"] = topicproj
	c.Data["TopicPhase"] = topicphase
	c.Data["TopicSpec"] = topicspec
	c.Data["TopicChengguo"] = topicchengguo
	c.Data["Topic"] = topic
	c.Data["Attachment"] = attachment
	c.Data["Tid"] = c.Ctx.Input.Param("0") //教程中用的是圆括号，导致错误topic.go:52: cannot call non-function c.Controller.Ctx.Input.Params (type map[string]string)
	//教程第8章开头有修改
	replies, err := models.GetAllReplies(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		return
	}
	c.Data["Replies"] = replies
	c.Data["IsLogin"] = checkAccount(c.Ctx)
}

//设代日记查看，轮播模式
func (c *TopicController) Carousel() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsTopic"] = true
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
	//这里是通过文章的id获得文章及上级目录情况
	topicproj, err := models.GetTopicProj(c.Ctx.Input.Param("0"))
	topicphase, err := models.GetTopicPhase(c.Ctx.Input.Param("0"))
	topicspec, err := models.GetTopicSpec(c.Ctx.Input.Param("0"))
	_, topicchengguo, err := models.GetTopicChengguo(c.Ctx.Input.Param("0"))
	if topicchengguo.Title == "diary" {
		c.TplName = "carousel.html"
	} else {
		c.TplName = "topic_view_b.html"
	}

	topic, attachment, err := models.GetTopic(c.Ctx.Input.Param("0"))
	//articleId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
	//id, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["TopicProj"] = topicproj
	c.Data["TopicPhase"] = topicphase
	c.Data["TopicSpec"] = topicspec
	c.Data["TopicChengguo"] = topicchengguo
	c.Data["Topic"] = topic
	c.Data["Attachment"] = attachment
	c.Data["Tid"] = c.Ctx.Input.Param("0") //教程中用的是圆括号，导致错误topic.go:52: cannot call non-function c.Controller.Ctx.Input.Params (type map[string]string)
	//教程第8章开头有修改
	replies, err := models.GetAllReplies(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		return
	}
	c.Data["Replies"] = replies
	c.Data["IsLogin"] = checkAccount(c.Ctx)
}

//添加好日记图片后开始写图片说明
func (c *TopicController) ViewDiary() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsTopic"] = true
	c.TplName = "diary_view.html"
	//这里是通过日记的编号获取日记id,由日记id获取日记及上级目录情况
	tid := c.Input().Get("tid")
	topicid, _ := strconv.ParseInt(tid, 10, 64)
	// tnumber := c.Input().Get("tnumber")
	// topicid, err := models.GetTopicIdbytnumber(tnumber)
	// if err != nil {
	// 	beego.Error(err)
	// 	return
	// }
	cid := strconv.FormatInt(topicid, 10)
	topicproj, err := models.GetTopicProj(cid)
	topicphase, err := models.GetTopicPhase(cid)
	topicspec, err := models.GetTopicSpec(cid)
	_, topicchengguo, err := models.GetTopicChengguo(cid)
	topic, attachment, err := models.GetTopic(cid)
	//articleId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
	//id, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["TopicProj"] = topicproj
	c.Data["TopicPhase"] = topicphase
	c.Data["TopicSpec"] = topicspec
	c.Data["TopicChengguo"] = topicchengguo
	c.Data["Topic"] = topic
	c.Data["Attachment"] = attachment
	c.Data["Tid"] = c.Ctx.Input.Param("0") //教程中用的是圆括号，导致错误topic.go:52: cannot call non-function c.Controller.Ctx.Input.Params (type map[string]string)
	//教程第8章开头有修改
	replies, err := models.GetAllReplies(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		return
	}
	c.Data["Replies"] = replies
	c.Data["IsLogin"] = checkAccount(c.Ctx)
}

func (c *TopicController) ViewDiary1() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsTopic"] = true
	c.TplName = "diary_view1.html"
	topicproj, err := models.GetTopicProj(c.Ctx.Input.Param("0"))
	topicphase, err := models.GetTopicPhase(c.Ctx.Input.Param("0"))
	topicspec, err := models.GetTopicSpec(c.Ctx.Input.Param("0"))
	_, topicchengguo, err := models.GetTopicChengguo(c.Ctx.Input.Param("0"))
	topic, attachment, err := models.GetTopic(c.Ctx.Input.Param("0"))
	//articleId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
	//id, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["TopicProj"] = topicproj
	c.Data["TopicPhase"] = topicphase
	c.Data["TopicSpec"] = topicspec
	c.Data["TopicChengguo"] = topicchengguo
	c.Data["Topic"] = topic
	c.Data["Attachment"] = attachment
	c.Data["Tid"] = c.Ctx.Input.Param("0") //教程中用的是圆括号，导致错误topic.go:52: cannot call non-function c.Controller.Ctx.Input.Params (type map[string]string)
	//教程第8章开头有修改
	replies, err := models.GetAllReplies(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		return
	}
	c.Data["Replies"] = replies
	c.Data["IsLogin"] = checkAccount(c.Ctx)
}

func (c *TopicController) Modify() { //这个也要登陆验证
	tid := c.Input().Get("tid")
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		// c.Redirect("/login", 302)
		return
	}
	//2.取得文章的作者
	topic, attachment, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	//3.取得客户端用户名
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
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	uname := v.(string) //ck.Value
	//4.取出用户的权限等级
	role, _ := checkRole(c.Ctx) //login里的
	// beego.Info(role)
	//5.进行逻辑分析：
	rolename, _ := strconv.ParseInt(role, 10, 64)
	if rolename > 2 && uname != topic.Author { //
		// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "topic_modify.html"

	c.Data["Topic"] = topic
	c.Data["Attachment"] = attachment
	c.Data["Tid"] = tid

	c.Data["IsTopic"] = true

	topicproj, err := models.GetTopicProj(tid)
	topicphase, err := models.GetTopicPhase(tid)
	topicspec, err := models.GetTopicSpec(tid)
	_, topicchengguo, err := models.GetTopicChengguo(tid)

	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["TopicProj"] = topicproj
	c.Data["TopicPhase"] = topicphase
	c.Data["TopicSpec"] = topicspec
	c.Data["TopicChengguo"] = topicchengguo
}

func (c *TopicController) ModifyTopic() { //一对多模式,向文章中追加附件
	//解析表单
	tid := c.Input().Get("tid") //教程里漏了这句，导致修改总是变成添加文章
	// title := c.Input().Get("title")
	// tnumber := c.Input().Get("tnumber")
	// content := c.Input().Get("content")
	category := c.Input().Get("category")
	categoryid := c.Input().Get("categoryid")

	//获取文件保存路径，有了categoryid可以求出整个路径
	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
	categoryproj, err := models.GetCategoryProj(categoryid)
	categoryphase, err := models.GetCategoryPhase(categoryid)
	categoryspec, err := models.GetCategorySpec(categoryid)
	// category, err := models.GetCategory(categoryid)
	if err != nil {
		beego.Error(err)
		// c.Redirect("/", 302)//这里注释掉，否则在图纸页面无法进入添加页面，因为传入的id为空，导致err发生
		return
	}
	//获取上传的文件
	_, h, err := c.GetFile("image")
	if err != nil {
		beego.Error(err)
	}
	var attachment string
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		attachment = h.Filename
		// beego.Info(attachment)
		path = ".\\attachment\\" + categoryproj.Number + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename

		// path := c.Input().Get("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("image", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
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

	route := "/attachment/" + categoryproj.Number + categoryproj.Title + "/" + categoryphase.Title + "/" + categoryspec.Title + "/" + category + "/" + h.Filename
	//topicid := c.Input().Get("topicid")
	// var topicid int64
	// if len(tid) == 0 {
	// 	topicid, err = models.AddTopic(title, tnumber, category, categoryid, content, attachment)
	// 	//这里返回topicid，并存入attachment表中
	// 	if err != nil { //如果发生错误，返回错误，并获取该文章的topicid
	// 		beego.Error(err)
	// 	}
	// 	if topicid == 0 {
	// 		topicid, err = models.GetTopicIdbytnumber(tnumber)
	// 	}
	// cid := strconv.FormatInt(tid, 10)
	//获取用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	// if v != nil {
	uname := v.(string)
	// }
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value

	size := strconv.FormatInt(filesize, 10)
	err = models.AddAttachment(attachment, size, path, route, tid, uname)
	// beego.Info(attachment)
	// } else {
	// err = models.ModifyTopic(tid, title, tnumber, category, categoryid, content, attachment)

	// }
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "modifytopic.tpl" //不加这句上传出错，虽然可以成功上传
	// c.Redirect("/topic", 302)
}

//删除文章
func (c *TopicController) Delete() { //应该显示警告
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		// c.Redirect("/login", 302)
		return
	}
	//2.取得文章的作者
	topic, _, err := models.GetTopic(c.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	//3.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	// if v != nil {
	uname := v.(string)
	// }
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value
	//4.取出用户的权限等级
	role, _ := checkRole(c.Ctx) //login里的
	// beego.Info(role)
	//5.进行逻辑分析：
	rolename, _ := strconv.ParseInt(role, 10, 64)
	if rolename > 2 && uname != topic.Author { //
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	err = models.DeletTopic(c.Input().Get("tid")) //(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic", 302) //这里增加topic
}

//删除文章中的附件，保持页面不跳转怎么办？
func (c *TopicController) DeleteAttachment() { //应该显示警告
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		// c.Redirect("/login", 302)
		return
	}
	//2.取得文章的作者
	topic, _, err := models.GetTopic(c.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	//3.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	// if v != nil {
	uname := v.(string)
	// }
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value
	//4.取出用户的权限等级
	role, _ := checkRole(c.Ctx) //login里的
	// beego.Info(role)
	//5.进行逻辑分析：
	rolename, _ := strconv.ParseInt(role, 10, 64)
	if rolename > 2 && uname != topic.Author { //
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// Tid := c.Ctx.Input.Param("0")
	Tid := c.Input().Get("tid")
	// beego.Info(Tid)
	err = models.DeletAttachment(c.Input().Get("aid")) //(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
	}
	op := c.Input().Get("op")
	switch op {
	case "modify":
		c.Redirect("/topic/modify?tid="+Tid, 302)
	default:
		c.Redirect("/topic/view/"+Tid, 302) //这里增加topic
	}
}

func FileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

//导出到excel
//引用来自category的查看成果类型里的成果
func (c *TopicController) ExportToExcel() {
	//解析表单
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	id := c.Input().Get("id")
	// beego.Info(id)
	path := c.Input().Get("path")
	// beego.Info(path)
	filename := c.Input().Get("filename")
	// beego.Info(filename)
	// if len(id) == 0 {
	// 	break
	// }
	// category, _ := models.GetCategory(id)
	// if category.Title == "diary" {
	// 	c.TplName = "proddiary_view_b.tpl"
	// } else {
	// 	c.TplName = "prod_view_b.tpl"
	// }
	chengguo, err := models.GetTopicsbyparentid(id, true)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(chengguo)
	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
	// categoryproj, _ := models.GetCategoryProj(id)
	// categoryphase, _ := models.GetCategoryPhase(id)
	// categoryspec, _ := models.GetCategorySpec(id)

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	// var err error

	file = xlsx.NewFile()
	sheet = file.AddSheet("Sheet1")
	row = sheet.AddRow() //增加行

	// for j := 2; j < 5; j++ {
	cell = row.AddCell() //增加列
	cell.Value = "#"

	cell = row.AddCell() //增加列
	cell.Value = "成果编号"

	cell = row.AddCell() //增加列
	cell.Value = "成果名称"

	cell = row.AddCell() //增加列
	cell.Value = "成果上传时间"

	cell = row.AddCell() //增加列
	cell.Value = "回复数量"
	for _, z := range chengguo { //行
		// cate := &Category{
		// 	Title:    v,
		// 	ParentId: z.Id,
		// 	Created:  time.Now(),
		// 	Updated:  time.Now(),
		// }
		// _, err = o.Insert(cate)
		// }
		row = sheet.AddRow() //增加行

		// for j := 2; j < 5; j++ {
		cell = row.AddCell() //增加列
		Num1 := strconv.FormatInt(z.Id, 10)
		cell.Value = Num1

		cell = row.AddCell() //增加列
		cell.Value = z.Tnumber

		cell = row.AddCell() //增加列
		cell.Value = z.Title

		cell = row.AddCell() //增加列
		time := z.Updated.String()
		cell.Value = time

		cell = row.AddCell() //增加列
		Num2 := strconv.FormatInt(z.Views, 10)
		cell.Value = Num2
		// }
	}
	path = path + filename + ".xlsx" //categoryproj.Number + " " + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + category + ".xlsx"
	// beego.Info(path)
	err = file.Save(path) //(".\\attachment\\MyXLSXFile.xlsx")

	// path := c.Input().Get("url")  //存文件的路径
	// path = path[3:]
	// path = "./attachment" + "/" + h.Filename
	// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
	// err = c.SaveToFile("image", path)
	if err != nil {
		fmt.Printf(err.Error())
	}
	c.Ctx.Output.Download(path) //("./attachment/MyXLSXFile.xlsx", "123.xlsx")
	// c.Data["CategoryProj"] = categoryproj
	// c.Data["CategoryPhase"] = categoryphase
	// c.Data["CategorySpec"] = categoryspec
	// c.Data["Category"] = category
	// c.Data["Chengguo"] = chengguo
}

func (c *TopicController) DeleteAll() {
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		// c.Redirect("/login", 302)
		return
	}
	//2.取得文章的作者
	// topic, _, err := models.GetTopic(c.Input().Get("tid"))
	// if err != nil {
	// 	beego.Error(err)
	// 	c.Redirect("/", 302)
	// 	return
	// }
	//3.取得客户端用户名
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value
	//4.取出用户的权限等级
	role, _ := checkRole(c.Ctx) //login里的
	// beego.Info(role)
	//5.进行逻辑分析：只有管理员可以
	rolename, _ := strconv.ParseInt(role, 10, 64)
	if rolename > 2 { //&& uname != topic.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//port + c.Ctx.Input.Site() + ":" +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	//解析表单
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	cid := c.Input().Get("cid")
	topicid := c.Input().Get("tempstring")
	array := strings.Split(topicid, ",") //字符串切割 [a b c d e]
	for _, v := range array {
		err := models.DeletTopic(v)
		if err != nil {
			beego.Error(err)
		}
	}
	c.Redirect("/category/view?id="+cid, 302) //这里增加topic
}

func (c *TopicController) DownloadAll() {
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		// c.Redirect("/login", 302)
		return
	}
	//2.取得文章的作者
	// topic, _, err := models.GetTopic(c.Input().Get("tid"))
	// if err != nil {
	// 	beego.Error(err)
	// 	c.Redirect("/", 302)
	// 	return
	// }
	//3.取得客户端用户名
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value
	//4.取出用户的权限等级
	role, _ := checkRole(c.Ctx) //login里的
	// beego.Info(role)
	//5.进行逻辑分析：管理员或该项目的负责人，后者没有完善
	rolename, _ := strconv.ParseInt(role, 10, 64)
	if rolename > 2 { //&& uname != topic.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	//解析表单
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	// cid := c.Input().Get("cid")
	// beego.Info(cid)
	topicid := c.Input().Get("tempstring1")
	// beego.Info(topicid)
	array := strings.Split(topicid, ",") //字符串切割 [a b c d e]

	//建立随机目录
	var to string
	to = "./" + random_name()
	err := os.Mkdir(to, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range array {
		_, attachment, err := models.GetTopic(v)
		//articleId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
		//id, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
		if err != nil {
			beego.Error(err)
			c.Redirect("/", 302)
			return
		}
		for _, v1 := range attachment {
			// beego.Info(v1)
			filename := filepath.Base(v1.DiskDirectory)
			// path := filepath.Dir(v1.DiskDirectory)
			w, err := CopyFile(to+"/"+filename, v1.DiskDirectory) //temp后面的斜杠必须要
			//targetfile,sourcefile
			fmt.Println(w)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	to1 := "./" + random_name() //保存压缩文件的文件夹
	err = os.Mkdir(to1, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}
	zippath(to, to1+"/"+to1+".tar.gz")
	// TarGzPath(to, to1+"/"+to1+".tar.gz") //压缩temp后面的斜杠可要可不要
	//UnTarGz("./temp/ty4z2008/1.tar.gz", "/temp/ty4z2008")     //解压
	c.Ctx.Output.Download(to1 + "/" + to1 + ".tar.gz")
	// 打开文件夹
	// dir, err := os.Open(to + "/")
	// if err != nil {
	// 	panic(nil)
	// }
	// defer dir.Close()
	// 读取文件列表
	// fis, err := dir.Readdir(0)
	// if err != nil {
	// 	panic(err)
	// }
	// 遍历文件列表
	// for _, fi := range fis {
	// os.RemoveAll(to + "/" + fi.Name())
	// }
	os.RemoveAll(to + "/")
	os.RemoveAll(to1 + "/")
	// c.Redirect("/category/view_b?id="+cid, 302) //这句多余，因为做不到。会出现http: multiple response.WriteHeader calls
}

func (c *TopicController) ListAllPosts() {
	c.Data["IsTopic"] = true
	c.TplName = "topic.tpl"
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

	// Directory:github.com/astaxie/beego/context     Pakage in Source:context
	// func (input *BeegoInput) IP() string {}
	//c是TopicController,TopicController是beego.controller，即beego.controller.ctx.input.ip
	// beego.Info(c.Ctx.Input.IP())
	topics, err := models.GetAllTopics("", false)
	if err != nil {
		beego.Error(err)
	}
	count := len(topics)
	count1 := strconv.Itoa(count)
	count2, err := strconv.ParseInt(count1, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	// sets this.Data["paginator"] with the current offset (from the url query param)
	postsPerPage := 20
	paginator := pagination.SetPaginator(c.Ctx, postsPerPage, count2)
	// beego.Info(c.Ctx)
	// beego.Info(paginator.Offset())   0
	// p := pagination.NewPaginator(c.Ctx.Request, 10, 9)
	// beego.Info(p.Offset())   0
	// fetch the next 20 posts
	topics, err = models.ListPostsByOffsetAndLimit(paginator.Offset(), postsPerPage)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topics"] = topics
	c.Data["paginator"] = paginator
	// count, _ := models.M("logoperation").Alias(`op`).Field(`count(op.id) as count`).Where(where).Count()
	// if count > 0 {
	// 	pagesize := 10
	// 	p := tools.NewPaginator(this.Ctx.Request, pagesize, count)
	// 	log, _ := models.M("logoperation").Alias(`op`).Where(where).Limit(strconv.Itoa(p.Offset()), strconv.Itoa(pagesize)).Order(`op.id desc`).Select()
	// 	this.Data["data"] = log
	// 	this.Data["paginator"] = p
	// }
}

// chengguo, _ := models.GetTopicsbyparentid(id, true)
//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
// categoryproj, _ := models.GetCategoryProj(id)
// categoryphase, _ := models.GetCategoryPhase(id)
// categoryspec, _ := models.GetCategorySpec(id)

// path = ".\\attachment\\" + path + filename + ".xlsx" //categoryproj.Number + " " + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + category + ".xlsx"
// beego.Info(path)
// err = file.Save(path) //(".\\attachment\\MyXLSXFile.xlsx")

// path := c.Input().Get("url")  //存文件的路径
// path = path[3:]
// path = "./attachment" + "/" + h.Filename
// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
// err = c.SaveToFile("image", path)
// if err != nil {
// 	fmt.Printf(err.Error())
// }
// c.Ctx.Output.Download(path) //("./attachment/MyXLSXFile.xlsx", "123.xlsx")

// func ToExcel() {
// 	var file *xlsx.File
// 	var sheet *xlsx.Sheet
// 	var row *xlsx.Row
// 	var cell *xlsx.Cell
// 	var err error

// 	file = xlsx.NewFile()
// 	sheet = file.AddSheet("Sheet1")
// 	//    for _, sheet := range xlFile.Sheets {
// 	for i := 0; i < 10; i++ { //行
// 						for _, z := range postss {
// 					cate := &Category{
// 						Title:    v,
// 						ParentId: z.Id,
// 						Created:  time.Now(),
// 						Updated:  time.Now(),
// 					}
// 					_, err = o.Insert(cate)
// 				}
// 		row = sheet.AddRow() //增加行

// 		for j := 11; j < 20; j++ {
// 			cell = row.AddCell() //增加列
// 			Num1 := strconv.Itoa(j)
// 			cell.Value = Num1
// 		}
// 	}
// 	err = file.Save("MyXLSXFile.xlsx")
// 	if err != nil {
// 		fmt.Printf(err.Error())
// 	}
// }
