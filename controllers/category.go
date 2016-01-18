package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"image/png"
	"os"
	"quick/models"
	"strconv"
	"strings"
)

type Uploadimage struct {
	url     string
	message string
	success int
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
								err := os.MkdirAll(".\\attachment\\"+number+" "+name+"\\"+v+"\\"+w+"\\"+t, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
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
		diskdirectory := ".\\attachment\\" + number + name + "\\"
		url := "/attachment/" + number + name + "/"
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
		id, err := models.AddCategory(name, number, content, path, route, uname, diskdirectory, url)
		if err != nil {
			beego.Error(err)
		}
		id1 := strconv.FormatInt(id, 10)
		c.Redirect("/category?op=view&id="+id1, 301)
		return
	case "del":
		//1.首先判断是否注册
		if !checkAccount(c.Ctx) {
			// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/login?url="+route, 302)
			// c.Redirect("/login", 302)
			return
		}
		//2.取得Id
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		//3.由Id查询数据库中的用户名
		// category, err := models.GetCategory(id)
		// beego.Info(username)
		//4.取得客户端用户名
		// ck, err := c.Ctx.Request.Cookie("uname")
		// if err != nil {
		// 	beego.Error(err)
		// }
		// uname := ck.Value
		//5.取出用户的权限等级
		role, _ := checkRole(c.Ctx) //login里的
		// beego.Info(role)
		//6.进行逻辑分析：
		rolename, _ := strconv.ParseInt(role, 10, 64)
		// if filetype != "pdf" && filetype != "jpg" && filetype != "diary" {
		if rolename > 1 { //&& uname != category.Author
			// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			// c.Redirect("/roleerr", 302)
			return
		}
		// }
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		//删除目录
		// func RemoveAll(path string) errorRemoveAll删除path指定的文件，或目录及它包含的任何下级对象。它会尝试删除所有东西，除非遇到错误并返回。如果path指定的对象不存在，RemoveAll会返回nil而不返回错误。

		c.Redirect("/category", 301)
		return

	case "view":
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.Data["IsCategory"] = true
		c.TplNames = "category_view.html"
		//2.取得客户端用户名
		ck, err := c.Ctx.Request.Cookie("uname")
		if err == nil {
			c.Data["Uname"] = ck.Value
		} else {
			beego.Error(err)
		}
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}

		category, _ := models.GetCategory(id) //由分类id取出本身（项目名称等）
		topics, err := models.GetAllTopics(category.Title, false)
		categorychengguo, _ := models.GetCategoryChengguo(id)
		categoryzhuanye, _ := models.GetCategoryZhuanye(id)
		categoryjieduan, _ := models.GetCategoryJieduan(id)
		// if err != nil {
		// 	beego.Error(err)
		// 	c.Redirect("/", 302)
		// 	return
		// }
		c.Data["Category"] = category
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
	case "view_b":
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.Data["IsCategoryb"] = true
		//2.取得客户端用户名
		ck, err := c.Ctx.Request.Cookie("uname")
		if err == nil {
			c.Data["Uname"] = ck.Value
		} else {
			beego.Error(err)
		}
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}

		category, _ := models.GetCategory(id) //由成果id取出成果
		c.TplNames = "category_view_b.html"
		categorychengguo, _ := models.GetCategoryChengguo(id)
		categoryzhuanye, _ := models.GetCategoryZhuanye(id)
		categoryjieduan, _ := models.GetCategoryJieduan(id)
		// if err != nil {
		// 	beego.Error(err)
		// 	c.Redirect("/", 302)
		// 	return
		// }
		c.Data["Category"] = category
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
	default:
		c.Data["IsCategory"] = true
		c.TplNames = "category.tpl"
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		//2.取得客户端用户名
		ck, err := c.Ctx.Request.Cookie("uname")
		if err == nil {
			c.Data["Uname"] = ck.Value
		} else {
			beego.Error(err)
		}
		// var err error
		categories, err := models.GetAllCategories()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Category"] = categories
		c.Data["Length"] = len(categories)
	}
	var err error
	if err != nil {
		beego.Error(err)
	}
}

func (c *CategoryController) Get_b() { //项目B显示控制
	c.Data["IsCategoryb"] = true
	c.TplNames = "category_b.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//2.取得客户端用户名
	ck, err := c.Ctx.Request.Cookie("uname")
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Uname"] = ck.Value
	}
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
	content := c.Input().Get("test-editormd-html-code")
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
		case "ghj", "xj", "ky", "cs", "zb", "sgt", "jgt": //阶段//查到所有阶段
			for _, w := range array {
				switch w {
				case "gh", "sg", "jd", "shg", "dz", "ys", "zh": //专业
					for _, t := range array {
						switch t {
						case "dwg", "doc", "xls", "pdf", "jpg", "tif", "diary": //成果分类
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
	diskdirectory := ".\\attachment\\" + number + name + "\\"
	url := "/attachment/" + number + name + "/"
	//保存上传的图片
	//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
	_, h, err := c.GetFile("image")
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
		// attachment = h.Filename
		// beego.Info(attachment)
		path1 := ".\\attachment\\" + number + name + "\\" + h.Filename
		err = c.SaveToFile("image", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
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
			origin := path1 //path + file.Name()
			fmt.Println("正在处理" + origin + ">>>" + origin)
			cmd_resize(origin, 2048, 0, origin)
			//				defer os.Remove(origin)//删除原文件
		}
		filesize, _ = FileSize(path1)
		filesize = filesize / 1000.0
		route = "/attachment/" + number + name + "/" + h.Filename
	} else {
		img := CreateRandomAvatar([]byte(number + name))
		fi, _ := os.Create("./attachment/" + number + name + "/u1.png")
		png.Encode(fi, img)
		fi.Close()
		route = "/attachment/" + number + name + "/u1.png"
	}
	ck, err := c.Ctx.Request.Cookie("uname")
	if err != nil {
		beego.Error(err)
	}
	uname := ck.Value

	//存入数据库
	id, err := models.AddCategory(name, number, content, path, route, uname, diskdirectory, url)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Uname"] = ck.Value
	id1 := strconv.FormatInt(id, 10)
	c.Redirect("/category?op=view&id="+id1, 301)
	return //???
}

func (c *CategoryController) Uploadimagesct() {

	name := "111"    //c.Input().Get("name")
	number := "222"  //c.Input().Get("number")
	content := "333" //c.Input().Get("test-editormd-html-code")
	path := "c"      //c.Input().Get("tempString")

	diskdirectory := ".\\attachment\\" + "test" + "\\"
	url := "/attachment/" + "test" + "/"
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
	_, err = models.AddCategory(name, number, content, path, route, uname, diskdirectory, url)
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
		c.ServeJson()
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

func (c *CategoryController) UserdefinedPost() {
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
	name := c.Input().Get("name")
	number := c.Input().Get("number")
	content := c.Input().Get("test-editormd-html-code")
	// image := c.Input().Get("image")
	// path1 := c.Input().Get("category2")
	// beego.Info(path1) //只能取到一个值 [I] 2-1
	// path2 := c.Input().Get("category3")
	// path3 := c.Input().Get("category4")
	var path2, path3, path4 []string
	var diskdirectory, url string
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
	category3 := c.GetStrings("category3")
	category4 := c.GetStrings("category4")
	// beego.Info(category2) //[2-1 2-2]
	// path := make([]string, 0)
	if len(category2) > 0 { //如果有2级目录，则建立2级，没有则建立1级
		for _, v := range category2 {
			if len(category3) > 0 { //如果有3级目录，则建立3级，没有则建立2级
				for _, w := range category3 {
					if len(category4) > 0 { //如果有4级目录，则建立4级，没有则建立3级
						for _, t := range category4 {
							err := os.MkdirAll(".\\attachment\\"+number+name+"\\"+v+"\\"+w+"\\"+t, 0777)
							path4 = category4
							path3 = category3
							path2 = category2
							// diskdirectory = ".\\attachment\\" + number + name + "\\" + v + "\\" + w + "\\" + t + "\\"
							// url = "/attachment/" + number + name + "/" + v + "/" + w + "/" + t + "/"
							// path = append(category2, category3...)
							// path = append(path, category4...)
							// path = path1 + "," + path2 + "," + path3
							// beego.Info(path)
							if err != nil {
								beego.Error(err)
							}
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
	diskdirectory = ".\\attachment\\" + number + name + "\\"
	url = "/attachment/" + number + name + "/"
	//保存上传的图片
	//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
	_, h, err := c.GetFile("image")
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
		// attachment = h.Filename
		// beego.Info(attachment)
		path1 := ".\\attachment\\" + number + name + "\\" + h.Filename
		err = c.SaveToFile("image", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
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
			origin := path1 //path + file.Name()
			fmt.Println("正在处理" + origin + ">>>" + origin)
			cmd_resize(origin, 2048, 0, origin)
			//defer os.Remove(origin)//删除原文件
		}
		filesize, _ = FileSize(path1)
		filesize = filesize / 1000.0
		route = "/attachment/" + number + name + "/" + h.Filename
	} else {
		route = ""
	}
	ck, err := c.Ctx.Request.Cookie("uname")
	if err != nil {
		beego.Error(err)
	}
	uname := ck.Value

	//存入数据库
	var id int64
	id, err = models.AdduserdefinedCategory(name, number, content, path2, path3, path4, route, uname, diskdirectory, url)
	if err != nil {
		beego.Error(err)
	}
	id1 := strconv.FormatInt(id, 10)
	c.Redirect("/category?op=view&id="+id1, 301)
	return
	//2.取得客户端用户名
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	c.Data["Uname"] = ck.Value
}

func (c *CategoryController) Add() {
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
	ck, err := c.Ctx.Request.Cookie("uname")
	if err == nil {
		c.Data["Uname"] = ck.Value
	} else {
		beego.Error(err)
	}
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
	c.TplNames = "category_add.tpl"
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

func (c *CategoryController) Add_b() {
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
	ck, err := c.Ctx.Request.Cookie("uname")
	if err == nil {
		c.Data["Uname"] = ck.Value
	} else {
		beego.Error(err)
	}

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
	c.TplNames = "category_add_b.tpl"
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

//根据用户名查看项目
func (c *CategoryController) Viewbyuname() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	c.TplNames = "category_uname.tpl"
	//2.取得客户端用户名
	ck, err := c.Ctx.Request.Cookie("uname")
	if err == nil {
		c.Data["Uname"] = ck.Value
	} else {
		beego.Error(err)
	}
	uname := c.Input().Get("uname")
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

//查看成果类型里的成果
func (c *CategoryController) View() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	//2.取得客户端用户名
	ck, err := c.Ctx.Request.Cookie("uname")
	if err == nil {
		c.Data["Uname"] = ck.Value
	} else {
		beego.Error(err)
	}

	id := c.Input().Get("id")
	// if len(id) == 0 {
	// 	break
	// }
	category, _ := models.GetCategory(id)
	if category.Title == "diary" {
		c.TplNames = "proddiary_view.tpl"
	} else {
		c.TplNames = "prod_view.tpl"
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
	c.Data["Chengguo"] = chengguo
	c.Data["Length"] = len(chengguo)

	//
	cid := strconv.FormatInt(categoryproj.Id, 10)
	categorycelan, _ := models.GetCategory(cid) //由分类id取出本身（项目名称等）
	// topics, err := models.GetAllTopics(category.Title, false)
	categorychengguo, _ := models.GetCategoryChengguo(cid)
	categoryzhuanye, _ := models.GetCategoryZhuanye(cid)
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

//查看专业里或第3级目录中的成果
func (c *CategoryController) View_3() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	//2.取得客户端用户名
	ck, err := c.Ctx.Request.Cookie("uname")
	if err == nil {
		c.Data["Uname"] = ck.Value
	} else {
		beego.Error(err)
	}

	id := c.Input().Get("id")
	// if len(id) == 0 {
	// 	break
	// }
	category, _ := models.GetCategory(id)
	if category.Title == "diary" {
		c.TplNames = "proddiary_view.tpl"
	} else {
		c.TplNames = "prod_3_view.tpl"
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
	//2.取得客户端用户名
	ck, err := c.Ctx.Request.Cookie("uname")
	if err == nil {
		c.Data["Uname"] = ck.Value
	} else {
		beego.Error(err)
	}
	id := c.Input().Get("id")
	// if len(id) == 0 {
	// 	break
	// }
	category, _ := models.GetCategory(id)
	if category.Title == "diary" {
		c.TplNames = "proddiary_view_b.tpl"
	} else {
		c.TplNames = "prod_view_b.tpl"
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
	//2.取得客户端用户名
	ck, err := c.Ctx.Request.Cookie("uname")
	if err == nil {
		c.Data["Uname"] = ck.Value
	} else {
		beego.Error(err)
	}
	id := c.Input().Get("id")
	// beego.Info(id)
	title := c.Input().Get("title")
	// beego.Info(title)
	// if len(id) == 0 {
	// 	break
	// }
	topics, err := models.GetAllTopics(title, false)

	c.TplNames = "category_prod_view.tpl"

	//由项目id获取所有成果
	// chengguo, _ := models.GetTopicsbyparentid(id, true)
	//取得成果类型id的专业parentid以及阶段parentid以及项目parentid才行
	category, _ := models.GetCategory(id)
	categoryproj, _ := models.GetCategoryProj(id)
	categoryphase, _ := models.GetCategoryPhase(id)
	categoryspec, _ := models.GetCategorySpec(id)

	c.Data["CategoryProj"] = categoryproj
	c.Data["CategoryPhase"] = categoryphase
	c.Data["CategorySpec"] = categoryspec
	c.Data["Category"] = category
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

func (c *CategoryController) Modify() {
	cid := c.Input().Get("cid")
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		// port := strconv.Itoa(c.Ctx.Input.Port())
		route := c.Ctx.Request.URL.String() //c.Ctx.Input.Site() + ":" + port +
		// beego.Info(c.Ctx.Input.Url())

		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		// c.Redirect("/login", 302)
		return
	}
	//2.取得客户端用户名
	ck, err := c.Ctx.Request.Cookie("uname")
	if err == nil {
		c.Data["Uname"] = ck.Value
	} else {
		beego.Error(err)
	}
	uname := ck.Value
	//3.由cid查询数据库中的用户名
	category, err := models.GetCategory(cid)
	//4.取出用户的权限等级
	role, _ := checkRole(c.Ctx) //login里的
	// beego.Info(role)
	//5.进行逻辑分析：
	rolename, _ := strconv.ParseInt(role, 10, 64)
	if rolename > 1 && uname != category.Author { //要么管理员，要么作者自己可以修改项目
		// port := strconv.Itoa(c.Ctx.Input.Port())
		route := c.Ctx.Request.URL.String() //c.Ctx.Input.Site() + ":" + port +
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategoryb"] = true
	c.TplNames = "category_modify.tpl"

	categorychengguo, err := models.GetCategoryChengguo(cid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	categoryzhuanye, err := models.GetCategoryZhuanye(cid)
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
		case "ghj":
			c.Data["Ghj"] = true
		case "xj":
			c.Data["Xj"] = true
		case "ky":
			c.Data["Ky"] = true
		case "cs":
			c.Data["Cs"] = true
		case "zb":
			c.Data["Zb"] = true
		case "sgt":
			c.Data["Sgt"] = true
		case "jgt":
			c.Data["Jgt"] = true
		}
	}
	for _, w := range categoryzhuanye {
		switch w.Title {
		case "gh":
			c.Data["Gh"] = true
		case "sg":
			c.Data["Sg"] = true
		case "jd":
			c.Data["Jd"] = true
		case "shg":
			c.Data["Shg"] = true
		case "dz":
			c.Data["Dz"] = true
		case "ys":
			c.Data["Ys"] = true
		case "zh":
			c.Data["Zh"] = true
		}
	}
	for _, x := range categorychengguo {
		switch x.Title {
		case "dwg":
			c.Data["Dwg"] = true
		case "doc":
			c.Data["Doc"] = true
		case "xls":
			c.Data["Xls"] = true
		case "pdf":
			c.Data["Pdf"] = true
		case "jpg":
			c.Data["Jpg"] = true
		case "Tif":
			c.Data["Tif"] = true
		case "diary":
			c.Data["Diary"] = true
		}
	}
	// c.Data["Id"] = cid

}

func (c *CategoryController) ModifyCategory() {
	// if !checkAccount(c.Ctx) {
	// 	port := strconv.Itoa(c.Ctx.Input.Port())
	// 	route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.Url()
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
	// content := c.Input().Get("content")
	content := c.Input().Get("test-editormd-html-code")
	// image := c.Input().Get("image")
	path := c.Input().Get("tempString")

	ck, err := c.Ctx.Request.Cookie("uname")
	if err != nil {
		beego.Error(err)
	}
	uname := ck.Value

	//保存上传的图片
	//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
	_, h, err := c.GetFile("image")
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
		// attachment = h.Filename
		// beego.Info(attachment)
		path1 := ".\\attachment\\" + number + name + "\\" + h.Filename
		err = c.SaveToFile("image", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		//如果包含jpg，则进行压缩
		if strings.Contains(strings.ToLower(h.Filename), ".jpg") {
			// 随机名称
			// to := path + random_name() + ".jpg"
			origin := path1 //path + file.Name()
			fmt.Println("正在处理" + origin + ">>>" + origin)
			cmd_resize(origin, 2048, 0, origin)
			//				defer os.Remove(origin)//删除原文件
		}
		filesize, _ = FileSize(path1)
		filesize = filesize / 1000.0
		route = "/attachment/" + number + name + "/" + h.Filename
		//存入数据库ModifyCategory(cid, name, number, content, path, route, uname string)
		err = models.ModifyCategory(cid, name, number, content, path, route, uname)
		if err != nil {
			beego.Error(err)
		}
	} else {
		//如果没有更新图片，则图片地址不存入
		err = models.ModifyCategory(cid, name, number, content, path, "", uname)
		if err != nil {
			beego.Error(err)
		}
	}
	c.Data["Uname"] = ck.Value
	c.Redirect("/category", 301)
	return
}

// func (c *TopicController) View() {
// 	c.TplNames = "topic_view.html"
// 	topic, err := models.GetTopic(c.Ctx.Input.Params["0"])
// 	//articleId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
// 	//id, _ := strconv.Atoi(manage.Ctx.Input.Params[":id"])
// 	if err != nil {
// 		beego.Error(err)
// 		c.Redirect("/", 302)
// 		return
// 	}
// 	c.Data["Topic"] = topic
// 	c.Data["Tid"] = c.Ctx.Input.Params["0"] //教程中用的是圆括号，导致错误topic.go:52: cannot call non-function c.Controller.Ctx.Input.Params (type map[string]string)
// 	//教程第8章开头有修改
// 	replies, err := models.GetAllReplies(c.Ctx.Input.Params["0"])
// 	if err != nil {
// 		beego.Error(err)
// 		return
// 	}
// 	c.Data["Replies"] = replies
// 	c.Data["IsLogin"] = checkAccount(c.Ctx)
// }
