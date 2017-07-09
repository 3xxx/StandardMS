package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strconv"
	"strings"
	"time"
)

type Category struct {
	Id              int64 `form:"-"`
	ParentId        int64
	Uid             int64
	Title           string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Number          string    `orm:"unique",form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"`
	Content         string    `orm:"sie(5000)"` //项目简介
	Cover           string    `orm:"sie(5000)"` //封面文字介绍
	Route           string    //封面图片的链接地址
	Created         time.Time `orm:"auto_now_add;type(datetime)"`
	Updated         time.Time `orm:"auto_now_add;type(datetime)"`
	Views           int64     `form:"-"`
	Author          string    //这个改成uid代替
	TopicCount      int64     //`form:"-"`
	TopicLastUserId int64     //`form:"-"`
	Isshow          bool
	Graphicmode     bool     //true表示图文模式
	Isuserdefined   bool     //是否自定义
	Label           []*Label `orm:"reverse(many)"` // 设置一对多的反向关系
	// DiskDirectory   string    `orm:"null"` //各级目录的物理文件夹地址
	// Url             string    `orm:"null"` //对应各级目录的链接地址
	//Type            string `orm:"null"` //项目类型：供水、枢纽、提防、河道、船闸、电站、水闸
	//这种类型是一对多关系
}

func init() {
	orm.RegisterModel(new(Category)) //, new(Article)
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/hydrocms.db", 10)
}

func AddCategory(name, number, label, content, cover, path, route, uname string) (id int64, err error) {
	o := orm.NewOrm()
	cate := &Category{
		Title:  name,
		Number: number,
		// Content: content,
		// Cover:         cover,
		Author: uname,
		// Route:         route,
		Created: time.Now(),
		Updated: time.Now(),
		// DiskDirectory: ".\\attachment\\" + number + name + "\\",
		// Url:           "/attachment/" + number + name + "/",
		Isshow: true,
	}
	qs := o.QueryTable("category") //不知道主键就用这个过滤操作
	//进行编号唯一性检查
	err = qs.Filter("number", number).One(cate)
	if err == nil {
		return 0, err
	}
	id, err = o.Insert(cate)
	if err != nil {
		return 0, err
	}
	//这里循环添加标签
	if label != "" {
		labelarray := strings.Split(label, ",")
		for _, labeltitle := range labelarray {
			_, err = AddLabel(labeltitle, id)
		}
	}
	// var id int64
	// err = qs.Filter("title", name).One(id, "Id")
	// id, err := strconv.ParseInt(tid, 10, 64)
	// var cate Category
	// cates := make([]*Category, 0)
	// id, err := qs.Filter("title", name).All(&cates, "Id") //只是返回查询个数
	var post Category
	err = o.QueryTable("Category").Filter("number", number).One(&post, "Id", "Title")
	if err != nil {
		return 0, err
	}
	// 	var post Post
	// o.QueryTable("post").Filter("Content__istartswith", "prefix string").One(&post, "Id", "Title")
	// //进行目录的添加和parentid的设置
	array := strings.Split(path, ",") //字符串切割 [a b c d e]
	// var j int
	// //将path存成3个数组
	// var JieDuan [10]string
	// var ZhuanYe [10]string
	// var ChengGuo [10]string
	// for i, v := range array {
	// 	switch v {
	// 	case "ghj", "xj", "ky":
	// 		// 先定义一个数组

	// 		JieDuan[i] = v
	// 		j = i
	// 	case "gh", "sg", "shg":
	// 		// 先定义一个数组

	// 		ZhuanYe[i-j] = v
	// 		j = i
	// 	case "dwg", "doc", "xls":
	// 		// 先定义一个数组

	// 		ChengGuo[i-j] = v
	// 	}
	// }
	//将数组存入数据库
	// for _, v := range JieDuan {
	// 	// if v == "" {
	// 	// 	break JLoop
	// 	// }
	// 	cate = &Category{
	// 		Title:    v,
	// 		ParentId: post.Id, //这里存入项目的id
	// 		Created:  time.Now(),
	// 		Updated:  time.Now(),
	// 	}
	// 	// JLoop:
	// 	_, err = o.Insert(cate)

	// 	var posts []Category //详见beego手册的All的示例
	// 	_, err = o.QueryTable("Category").Filter("parentid", post.Id).All(&posts, "Id")
	// 	for i, z := range ZhuanYe {
	// 		cate := &Category{
	// 			Title:    z,
	// 			ParentId: posts[i].Id,
	// 			Created:  time.Now(),
	// 			Updated:  time.Now(),
	// 		}
	// 		_, err = o.Insert(cate)

	// 	}
	// }
	var jieduan string
	for _, v := range array {
		switch v {
		case "A":
			jieduan = "规划"
		case "B":
			jieduan = "项目建议书"
		case "C":
			jieduan = "可行性研究"
		case "D":
			jieduan = "初步设计"
		case "E":
			jieduan = "招标设计"
		case "F":
			jieduan = "施工图设计"
		case "G":
			jieduan = "竣工图"
		case "L":
			jieduan = "专题"
		}
		switch v {
		case "A", "B", "C", "D", "E", "F", "G", "L":
			cate = &Category{
				Title:    jieduan,
				ParentId: post.Id, //这里存入项目的id
				Created:  time.Now(),
				Updated:  time.Now(),
				Author:   uname,
				// filepath := ".\\attachment\\" + ProNumber + category.Title + "\\" + ProJieduan
				// DiskDirectory: ".\\attachment\\" + number + name + "\\" + v + "\\",
				// Url:           "/attachment/" + number + name + "/" + v + "/",
				Isshow: true,
				// Style:         style,
			}
			_, err = o.Insert(cate)
			if err != nil {
				return 0, err
			}
			//建立目录——注意，models中无法建立目录，必须在controllers中才行
		}
	}

	var leixing string
	for _, v := range array {
		switch v {
		case "FB":
			leixing = "技术报告"
		case "FD":
			leixing = "设计大纲"
		case "FG":
			leixing = "设计/修改通知单"
		case "FT":
			leixing = "工程图纸"
		case "FJ":
			leixing = "计算书"
		case "FP":
			leixing = "PDF文件"
		case "Fdiary":
			leixing = "文章/设代日记"
		}
		switch v {
		case "FB", "FD", "FG", "FT", "FJ", "FP", "Fdiary": //文件类型
			//查到阶段的parentid，符合这个项目的，得出阶段id，作为文件类型parentid
			var posts []Category                                                            //详见beego手册的All的示例
			_, err = o.QueryTable("Category").Filter("parentid", post.Id).All(&posts, "Id") //Id没什么用？
			for _, w := range posts {
				cate := &Category{
					Title:    leixing, //v,
					ParentId: w.Id,
					Created:  time.Now(),
					Updated:  time.Now(),
					Author:   uname,
					// DiskDirectory: diskdirectory + w.DiskDirectory + v + "\\",
					// Url:           url + w.Url + v + "/",
					Isshow: true,
					// Style:         style,
				}
				_, err = o.Insert(cate)
				if err != nil {
					return 0, err
				}
			}
		}
	}
	var zhuanye string
	for _, v := range array {
		switch v {
		case "1":
			zhuanye = "综合"
		case "2":
			zhuanye = "规划(含水文、经评)"
		case "3":
			zhuanye = "测量"
		case "4":
			zhuanye = "地质(含钻探)"
		case "5":
			zhuanye = "水工(含公路、安全监测)"
		case "6":
			zhuanye = "建筑"
		case "7":
			zhuanye = "机电"
		case "8":
			zhuanye = "征地、环保、水保"
		case "9":
			zhuanye = "施工、工程造价"
		}
		switch v {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9": //专业
			//查到文件类型的parentid，符合这个阶段的，得出文件分类id，作为专业分类的parentid
			var posts []Category                                                                     //详见beego手册的All的示例
			_, err = o.QueryTable("Category").Filter("parentid", post.Id).All(&posts, "Id", "Title") //这里只用到Id？
			for _, w := range posts {
				var postss []Category //详见beego手册的All的示例
				_, err = o.QueryTable("Category").Filter("parentid", w.Id).All(&postss, "Id")
				for _, t := range postss {
					cate := &Category{
						Title:    zhuanye, //v,
						ParentId: t.Id,
						Created:  time.Now(),
						Updated:  time.Now(),
						Author:   uname,
						// DiskDirectory: diskdirectory + t.DiskDirectory + v + "\\",
						// Url:           url + t.Url + v + "/",
						Isshow: true,
						// Style:         style,
					}
					_, err = o.Insert(cate)
					if err != nil {
						return 0, err
					}
				}
			}
		}
	}
	return id, nil
}

//添加自定义目录
func AdduserdefinedCategory(name, number, label, content, cover string, path2, path3, path4 []string, radio, route, uname string) (id int64, err error) {
	o := orm.NewOrm()
	cate := &Category{
		Title:  name,
		Number: number,
		// Content:       content,
		// Cover:         cover,
		Author: uname,
		// Route:         route,
		Created: time.Now(),
		Updated: time.Now(),
		// DiskDirectory: ".\\attachment\\" + number + name + "\\",
		// Url:           "/attachment/" + number + name + "/",
		Isshow:        true,
		Isuserdefined: true,
	}
	qs := o.QueryTable("category") //不知道主键就用这个过滤操作
	// 进行编号唯一性检查
	err = qs.Filter("number", number).One(cate)
	if err == nil {
		return 0, err
	}
	id, err = o.Insert(cate)
	if err != nil {
		return 0, err
	}
	//这里循环添加标签
	if label != "" {
		labelarray := strings.Split(label, ",")
		for _, labeltitle := range labelarray {
			_, err = AddLabel(labeltitle, id)
		}
	}

	var post Category //取出number项目编号的Id
	err = o.QueryTable("Category").Filter("number", number).One(&post, "Id", "Title")
	if err != nil {
		return 0, err
	}
	// array := strings.Split(path, ",") //字符串切割 [a b c d e]
	//Graphicmode图文模式
	radioarray := strings.Split(radio, ",")
	beego.Info(radioarray)
	beego.Info(radioarray[0])
	beego.Info(radioarray[1])
	if len(path2) > 0 {
		for _, v := range path2 {
			if v != "" {
				cate = &Category{
					Title:    v,
					ParentId: post.Id, //这里存入项目的id
					Created:  time.Now(),
					Updated:  time.Now(),
					Author:   uname,
					// DiskDirectory: diskdirectory + v + "\\",
					// Url:           url + v + "/",
					Isshow:        true,
					Isuserdefined: true,
				}
				_, err = o.Insert(cate)
				if err != nil {
					return 0, err
				}
			}
		}
	}
	if len(path3) > 0 {
		for _, v := range path3 {
			if v != "" {
				//查到阶段的parentid，符合这个项目的，得出阶段id，作为专业parentid
				var posts []Category //详见beego手册的All的示例
				_, err = o.QueryTable("Category").Filter("parentid", post.Id).All(&posts, "Id", "Title")
				for _, w := range posts {
					cate := &Category{
						Title:    v,
						ParentId: w.Id,
						Created:  time.Now(),
						Updated:  time.Now(),
						Author:   uname,
						// DiskDirectory: diskdirectory + w.Title + "\\" + v + "\\",
						// Url:           url + w.Title + "/" + v + "/",
						Isshow:        true,
						Isuserdefined: true,
					}
					_, err = o.Insert(cate)
					if err != nil {
						return 0, err
					}
				}
			}

		}
	}

	if len(path4) > 0 {
		for i, v := range path4 {
			if v != "" {
				//查到专业的parentid，符合这个阶段的，得出专业id，作为成果分类的parentid
				var posts []Category //详见beego手册的All的示例
				_, err = o.QueryTable("Category").Filter("parentid", post.Id).All(&posts, "Id", "Title")
				if err != nil {
					return 0, err
				}
				for _, w := range posts {
					var postss []Category //详见beego手册的All的示例
					_, err = o.QueryTable("Category").Filter("parentid", w.Id).All(&postss, "Id", "Title")
					if err != nil {
						return 0, err
					}
					for _, t := range postss {
						//如果是图文模式Graphicmode为true
						if radioarray[i] == "Fdiary" {
							cate := &Category{
								Title:    v,
								ParentId: t.Id,
								Created:  time.Now(),
								Updated:  time.Now(),
								Author:   uname,
								// DiskDirectory: diskdirectory + w.Title + "\\" + t.Title + "\\" + v + "\\",
								// Url:           url + w.Title + "/" + t.Title + "/" + v + "/",
								Isshow:        true,
								Graphicmode:   true,
								Isuserdefined: true,
							}
							_, err = o.Insert(cate)
							if err != nil {
								return 0, err
							}
						} else {
							cate := &Category{
								Title:    v,
								ParentId: t.Id,
								Created:  time.Now(),
								Updated:  time.Now(),
								Author:   uname,
								// DiskDirectory: diskdirectory + w.Title + "\\" + t.Title + "\\" + v + "\\",
								// Url:           url + w.Title + "/" + t.Title + "/" + v + "/",
								Isshow:        true,
								Graphicmode:   false,
								Isuserdefined: true,
							}
							_, err = o.Insert(cate)
							if err != nil {
								return 0, err
							}
						}

					}
				}
			}
		}
	}
	return id, nil
}

//添加一个自定义目录
func AdduserdefinedCategoryOne(name, pid, radio, uname string) (id int64, err error) {
	o := orm.NewOrm()
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		return 0, err
	}
	//如果是图文模式Graphicmode为true
	if radio == "Fdiary" {
		cate := &Category{
			Title:         name,
			Author:        uname,
			ParentId:      pidNum,
			Created:       time.Now(),
			Updated:       time.Now(),
			Isshow:        true,
			Graphicmode:   true,
			Isuserdefined: true,
		}
		_, err = o.Insert(cate)
		if err != nil {
			return 0, err
		}
	} else {
		cate := &Category{
			Title:         name,
			Author:        uname,
			ParentId:      pidNum,
			Created:       time.Now(),
			Updated:       time.Now(),
			Isshow:        true,
			Graphicmode:   false,
			Isuserdefined: true,
		}
		_, err = o.Insert(cate)
		if err != nil {
			return 0, err
		}
	}
	return id, err
}

//修改项目——也是第二步添加封面、简介的提交方法，共用
func ModifyCategory(cid, name, number, label, content, cover, path, route, uname string) error {
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	// cid, err := strconv.ParseInt(categoryid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cidNum}
	// if o.Read(cate) == nil {
	// 	// oldAttach = topic.Attachment
	// 	// oldCate = topic.Category
	// 	cate.Title = name
	// 	topic.Tnumber = tnumber
	// 	topic.Category = category
	// 	topic.CategoryId = cid
	// 	topic.Content = content
	// 	// topic.Attachment = attachment
	// 	topic.Updated = time.Now()
	// 	_, err = o.Update(topic)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	if o.Read(cate) == nil {
		if name == "" { //如果名称为空，则是添加封面
			// cate.Title = name
			// cate.Number = number
			// cate.Route = route
			cate.Content = content
			cate.Cover = cover
			// cate.Author = uname
			cate.Route = route
			// cate.Created: time.Now(),
			cate.Updated = time.Now()
			_, err = o.Update(cate)
			if err != nil {
				return err
			}
		} else { //修改
			cate.Title = name
			cate.Number = number
			//这里循环添加标签
			// if label != "" {
			// labelarray := strings.Split(label, ",")
			// for _, labeltitle := range labelarray {
			err = UpdateLabel(label, cidNum) //这里是修改，不是添加
			if err != nil {
				return err
			}
			// }
			// }
			cate.Content = content
			cate.Cover = cover
			cate.Author = uname
			if route != "" { //如果用户未更新封面图片，所取到的封面图片地址为空，则保留原数据库中的封面图片地址
				cate.Route = route
			}
			// cate.Created: time.Now(),
			cate.Updated = time.Now()
			_, err = o.Update(cate)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//仅仅删除项目数据库
func DelCategory(id string) error { //应该显示警告
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	//先删成果专业
	var posts []Category                                                        //详见beego手册的All的示例
	_, err = o.QueryTable("Category").Filter("parentid", cid).All(&posts, "Id") //获取阶段
	for _, v := range posts {
		var postss []Category //详见beego手册的All的示例
		_, err = o.QueryTable("Category").Filter("parentid", v.Id).All(&postss, "Id")
		for _, w := range postss {
			var postsss []Category //详见beego手册的All的示例
			_, err = o.QueryTable("Category").Filter("parentid", w.Id).All(&postsss, "Id")
			for _, z := range postsss {
				cate := &Category{Id: z.Id}
				_, err = o.Delete(cate)
			}
		}
	}

	//再删文档类型
	// var posts []Category //详见beego手册的All的示例
	_, err = o.QueryTable("Category").Filter("parentid", cid).All(&posts, "Id")
	for _, w := range posts {
		var postss []Category //详见beego手册的All的示例
		_, err = o.QueryTable("Category").Filter("parentid", w.Id).All(&postss, "Id")
		for _, z := range postss {
			cate := &Category{Id: z.Id}
			_, err = o.Delete(cate)
		}
	}

	//再删阶段分类
	// var posts []Category //详见beego手册的All的示例
	_, err = o.QueryTable("Category").Filter("parentid", cid).All(&posts, "Id")
	for _, w := range posts {
		cate := &Category{Id: w.Id}
		_, err = o.Delete(cate)
	}
	//最后删除项目
	cate := &Category{Id: cid} //这种必须是主键，非主键用下面的方式
	// user := User{Name: "slene"}
	// err = o.Read(&user, "Name")
	_, err = o.Delete(cate)
	return err
}

//删除category部分目录结构
func DeleteCategory(id int64) error { //应该在controllers中显示警告
	var tid, cid, cid1, cid2 string
	o := orm.NewOrm()
	// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	// user := User{Name: "slene"}
	// err = o.Read(&user, "Name")
	category := Category{Id: id}
	if o.Read(&category) == nil {
		_, err := o.Delete(&category) //删除本级（最上级是阶段）
		if err != nil {
			return err
		}
	}

	//查询下级，如果有下级
	var categories []Category //categories是分类
	_, err := o.QueryTable("Category").Filter("parentid", id).All(&categories, "Id")
	if err != nil {
		//应该在这里删除本级
		//再删除成果
		cid = strconv.FormatInt(id, 10)
		topics, _ := GetTopicsbyparentid(cid, true)
		for _, w := range topics {
			tid = strconv.FormatInt(w.Id, 10)
			err = DeletTopic(tid)
			if err != nil {
				beego.Error(err)
			}
		}
		//再删除物理目录
		_, diskdirectory, err := GetCategoryUrl(cid)
		if err != nil {
			beego.Error(err)
		}
		err = os.RemoveAll(diskdirectory)
		if err != nil {
			beego.Error(err)
		}

		return err
	} else { //如果有下级（文档分类）
		_, err = o.QueryTable("Category").Filter("parentid", id).Delete() //删除类型
		// _, err := o.Delete(&categories)
		if err != nil {
			return err
		}
		for _, v := range categories { //循环分类
			var categories1 []Category
			_, err = o.QueryTable("Category").Filter("parentid", v.Id).All(&categories1, "Id")
			if err != nil { //如果没有下级专业了，则说明是分类，进行删除成果
				//再删除成果
				cid1 = strconv.FormatInt(v.Id, 10)
				topics, _ := GetTopicsbyparentid(cid1, true)
				for _, w := range topics {
					tid = strconv.FormatInt(w.Id, 10)
					err = DeletTopic(tid)
					if err != nil {
						beego.Error(err)
					}
				}
				//再删除物理目录
				_, diskdirectory, err := GetCategoryUrl(cid1)
				if err != nil {
					beego.Error(err)
				}
				err = os.RemoveAll(diskdirectory)
				if err != nil {
					beego.Error(err)
				}
				return err
			} else { //如果有下级（专业）
				_, err = o.QueryTable("Category").Filter("parentid", v.Id).Delete() //删除专业
				// _, err := o.Delete(&categories1)
				if err != nil {
					return err
				}
				for _, w := range categories1 {
					//再删除成果
					cid2 = strconv.FormatInt(w.Id, 10)
					topics, _ := GetTopicsbyparentid(cid2, true)
					for _, ww := range topics {
						tid = strconv.FormatInt(ww.Id, 10)
						err = DeletTopic(tid)
						if err != nil {
							beego.Error(err)
						}
					}
					//再删除物理目录
					_, diskdirectory, err := GetCategoryUrl(cid2)
					if err != nil {
						beego.Error(err)
					}
					err = os.RemoveAll(diskdirectory)
					if err != nil {
						beego.Error(err)
					}
					// var categories2 []Category
					// _, err = o.QueryTable("Category").Filter("parentid", w.Id).All(&categories2, "Id")
					// if err != nil {
					// 	return err
					// } else {
					// 	_, err = o.QueryTable("Category").Filter("parentid", w.Id).Delete() //删除价值内容
					// 	if err != nil {
					// 		return err
					// 	}
					// }
				}
			}
		}
	}
	// 依据当前查询条件，进行批量删除操作
	// num, err := o.QueryTable("user").Filter("name", "slene").Delete()
	// fmt.Printf("Affected Num: %s, %s", num, err)
	// // DELETE FROM user WHERE name = "slene"
	return err
}

//显示category部分目录结构
func ShowCategory(id int64) error { //应该在controllers中显示警告
	o := orm.NewOrm()
	// cate := &Category{Id: cidNum}
	// if o.Read(cate) == nil {
	// 		cate.Isshow = true
	// 		cate.Updated = time.Now()
	// 		_, err = o.Update(cate)
	// 		if err != nil {
	// 			return err
	// 		}
	// }
	// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	// user := User{Name: "slene"}
	// err = o.Read(&user, "Name")
	category := &Category{Id: id}
	if o.Read(category) == nil {
		category.Isshow = true
		category.Updated = time.Now()
		_, err := o.Update(category)
		if err != nil {
			return err
		}
	}
	//查询下级
	var categories []Category
	_, err := o.QueryTable("Category").Filter("parentid", id).All(&categories, "Id")
	if err != nil {
		return err
	} else {
		// _, err = o.QueryTable("Category").Filter("parentid", id).Delete() //删除类型
		// _, err := o.Delete(&categories)
		// if err != nil {
		// 	return err
		// }
		// 依据当前查询条件，进行批量更新操作
		_, err := o.QueryTable("Category").Filter("parentid", id).Update(orm.Params{
			"isshow": true, "updated": time.Now(),
		})
		if err != nil {
			return err
		}
		for _, v := range categories {
			// v.Isshow = true
			// v.Updated = time.Now()
			// _, err := o.Update(v) //这里不对
			// if err != nil {
			// 	return err
			// }
			var categories1 []Category
			_, err = o.QueryTable("Category").Filter("parentid", v.Id).All(&categories1, "Id")
			if err != nil {
				return err
			} else {
				// 依据当前查询条件，进行批量更新操作
				_, err := o.QueryTable("Category").Filter("parentid", v.Id).Update(orm.Params{
					"isshow": true, "updated": time.Now(),
				})
				// _, err = o.QueryTable("Category").Filter("parentid", v.Id).Delete() //删除专业
				// _, err := o.Delete(&categories1)
				if err != nil {
					return err
				}
				// for _, w := range categories1 {
				// 	w.Isshow = true
				// 	w.Updated = time.Now()
				// 	_, err := o.Update(w)
				// 	if err != nil {
				// 		return err
				// 	}
				// 	var categories2 []Category
				// 	_, err = o.QueryTable("Category").Filter("parentid", w.Id).All(&categories2, "Id")
				// 	if err != nil {
				// 		return err
				// 	} else {
				// 		_, err = o.QueryTable("Category").Filter("parentid", w.Id).Delete() //删除价值内容
				// 		if err != nil {
				// 			return err
				// 		}
				// 	}
				// }
			}
		}
	}
	// 依据当前查询条件，进行批量删除操作
	// num, err := o.QueryTable("user").Filter("name", "slene").Delete()
	// fmt.Printf("Affected Num: %s, %s", num, err)
	// // DELETE FROM user WHERE name = "slene"
	return err
}

//隐藏category部分目录结构
func HideCategory(id int64) error { //应该在controllers中显示警告
	o := orm.NewOrm()
	// cate := &Category{Id: cidNum}
	// if o.Read(cate) == nil {
	// 		cate.Isshow = true
	// 		cate.Updated = time.Now()
	// 		_, err = o.Update(cate)
	// 		if err != nil {
	// 			return err
	// 		}
	// }
	// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	// user := User{Name: "slene"}
	// err = o.Read(&user, "Name")
	category := &Category{Id: id}
	if o.Read(category) == nil {
		category.Isshow = false
		category.Updated = time.Now()
		_, err := o.Update(category)
		if err != nil {
			return err
		}
	}
	//查询下级
	var categories []Category
	_, err := o.QueryTable("Category").Filter("parentid", id).All(&categories, "Id")
	if err != nil {
		return err
	} else {
		// _, err = o.QueryTable("Category").Filter("parentid", id).Delete() //删除类型
		// _, err := o.Delete(&categories)
		// if err != nil {
		// 	return err
		// }
		// 依据当前查询条件，进行批量更新操作
		_, err := o.QueryTable("Category").Filter("parentid", id).Update(orm.Params{
			"isshow": false, "updated": time.Now(),
		})
		if err != nil {
			return err
		}
		for _, v := range categories {
			// v.Isshow = true
			// v.Updated = time.Now()
			// _, err := o.Update(v) //这里不对
			// if err != nil {
			// 	return err
			// }
			var categories1 []Category
			_, err = o.QueryTable("Category").Filter("parentid", v.Id).All(&categories1, "Id")
			if err != nil {
				return err
			} else {
				// 依据当前查询条件，进行批量更新操作
				_, err := o.QueryTable("Category").Filter("parentid", v.Id).Update(orm.Params{
					"isshow": false, "updated": time.Now(),
				})
				// _, err = o.QueryTable("Category").Filter("parentid", v.Id).Delete() //删除专业
				// _, err := o.Delete(&categories1)
				if err != nil {
					return err
				}
				// for _, w := range categories1 {
				// 	w.Isshow = true
				// 	w.Updated = time.Now()
				// 	_, err := o.Update(w)
				// 	if err != nil {
				// 		return err
				// 	}
				// 	var categories2 []Category
				// 	_, err = o.QueryTable("Category").Filter("parentid", w.Id).All(&categories2, "Id")
				// 	if err != nil {
				// 		return err
				// 	} else {
				// 		_, err = o.QueryTable("Category").Filter("parentid", w.Id).Delete() //删除价值内容
				// 		if err != nil {
				// 			return err
				// 		}
				// 	}
				// }
			}
		}
	}
	// 依据当前查询条件，进行批量删除操作
	// num, err := o.QueryTable("user").Filter("name", "slene").Delete()
	// fmt.Printf("Affected Num: %s, %s", num, err)
	// // DELETE FROM user WHERE name = "slene"
	return err
}

//修改“自定义”目录名称以及下级名称；url，diskdirectory名称通过GetCategoryUrl获得
func ModifyCategoryTitle(id int64, name string) error {
	o := orm.NewOrm()
	var err error
	//先把数据库中的旧路径取出来
	cid := strconv.FormatInt(id, 10)
	// beego.Info(cid)
	_, diskdirectory, err := GetCategoryUrl(cid)
	if err != nil {
		beego.Error(err)
	}
	//修改数据库中的路径
	category := &Category{Id: id}
	if o.Read(category) == nil {
		category.Title = name
		_, err = o.Update(category)
		if err != nil {
			return err
		}
	}

	// beego.Info(diskdirectory)
	Length1 := len(diskdirectory)                //汉字占2位
	Disk := string(diskdirectory[0 : Length1-1]) //汉字占2位
	index := strings.LastIndex(Disk, "\\")       //汉字占2位

	Disk1 := string(Disk[0 : index+1])
	// beego.Info(Disk1)
	err = os.Rename(diskdirectory, Disk1+name+"\\")
	if err != nil {
		return err
	}
	return err
}

//修改“自定义”目录名称以及下级名称；url，diskdirectory名称通过GetCategoryUrl获得
// func ModifyCategoryTitleback(id int64, name string) error {
// 	var Length1, Length2, Length3, Length4, Length5, Length6 int
// 	var Disk1, Url1 string
// 	o := orm.NewOrm()
// 	category := &Category{Id: id}
// 	if o.Read(category) == nil {
// 		category.Title = name
// 		// beego.Info(category.DiskDirectory)                     //.\attachment\SL2016测试添加成果\A\Fdiary\1\
// 		Length1 = len(category.DiskDirectory)                 //汉字占2位
// 		Disk := string(category.DiskDirectory[0 : Length1-1]) //汉字占2位
// 		beego.Info(Disk)                                      //.\attachment\SL2016测试添加成果\A\Fdiary\1
// 		index := strings.LastIndex(Disk, "\\")                //汉字占2位
// 		// beego.Info(index)
// 		Disk1 = string(Disk[0 : index+1]) //不能用SubString(Disk, 0, index)汉字占1位
// 		beego.Info(Disk1)                 //.\attachment\SL2016测试添加成果\A\Fdiary\1
// 		err := os.Rename(category.DiskDirectory, Disk1+name+"\\")
// 		if err != nil {
// 			return err
// 		}
// 		// beego.Info(category.DiskDirectory)
// 		Length2 := len(category.Url)
// 		Url := string(category.Url[0 : Length2-1])
// 		// beego.Info(Url)
// 		index = strings.LastIndex(Url, "/")
// 		Url1 = string(Url[0 : index+1])
// 		// beego.Info(Url1)

// 		// category.DiskDirectory = Disk1 + name + "\\"
// 		// category.Url = Url1 + name + "/"
// 		// beego.Info(category.Url) //attachment/SL2016测试添加成果/A/Fdiary/1/综合1\
// 		_, err = o.Update(category)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	//查询下级
// 	var categories []Category
// 	_, err := o.QueryTable("Category").Filter("parentid", id).All(&categories)
// 	if err != nil {
// 		return err
// 	} else {
// 		// 依据当前查询条件，进行批量更新操作
// 		// _, err := o.QueryTable("Category").Filter("parentid", id).Update(orm.Params{
// 		// 	"title": name,
// 		// })
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		for _, v := range categories {
// 			//这里分别对数据库中disk和url的修改
// 			Length3 = len(v.DiskDirectory)
// 			beego.Info(Length3) //
// 			beego.Info(v.DiskDirectory)
// 			beego.Info(Length1)
// 			Disk2 := string(v.DiskDirectory[Length1:Length3])

// 			Length4 = len(v.Url)
// 			Url2 := string(v.Url[Length2:Length4])

// 			category1 := &Category{Id: v.Id}
// 			if o.Read(category1) == nil {
// 				// category1.DiskDirectory = Disk1 + name + "\\" + Disk2 + "\\"
// 				// category1.Url = Url1 + name + "/" + Url2 + "/"
// 				_, err = o.Update(category1)
// 				if err != nil {
// 					return err
// 				}
// 			}

// 			var categories1 []Category
// 			_, err = o.QueryTable("Category").Filter("parentid", v.Id).All(&categories1)
// 			if err != nil {
// 				return err
// 			} else {
// 				// 依据当前查询条件，进行批量更新操作
// 				// _, err := o.QueryTable("Category").Filter("parentid", v.Id).Update(orm.Params{
// 				// 	"title": name,
// 				// })
// 				// if err != nil {
// 				// 	return err
// 				// }
// 				for _, w := range categories1 {
// 					//这里进行物理目录修改，分别对数据库中disk和url的修改
// 					//这里分别对数据库中disk和url的修改
// 					Length5 = len(w.DiskDirectory)                    //汉字占1位
// 					Disk3 := string(w.DiskDirectory[Length1:Length5]) //汉字占1位

// 					Length6 = len(w.Url)
// 					Url3 := string(w.Url[Length2:Length6])

// 					category2 := &Category{Id: w.Id}
// 					if o.Read(category2) == nil {
// 						// category2.DiskDirectory = Disk1 + name + "\\" + Disk3 + "\\"
// 						// category2.Url = Url1 + name + "/" + Url3 + "/"
// 						_, err = o.Update(category2)
// 						if err != nil {
// 							return err
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return err
// }

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	var err error
	//这里进行过滤，parentid为空的才显示
	qs = qs.Filter("ParentId", 0)
	_, err = qs.OrderBy("-created").All(&cates)

	// _, err := qs.All(&cates)
	return cates, err
}

//由label取得categories
//搞错了，应该是多对多的关系
func GetCategoriesbylabel(title string) ([]*Category, error) {
	// var user User
	// err := o.QueryTable("user").Filter("Post__Title", "The Title").Limit(1).One(&user)
	// if err == nil {
	//     fmt.Printf(user)
	o := orm.NewOrm()
	// cate := new(Category)
	cate := make([]*Category, 0)
	cates := make([]*Category, 0)
	labels := make([]*Label, 0)
	qs1 := o.QueryTable("category")
	qs := o.QueryTable("label")
	var err error
	_, err = qs.Filter("Title", title).All(&labels)
	for _, v := range labels {
		// beego.Info(v.Category.Id)
		err = qs1.Filter("id", v.Category.Id).One(&cate)
		cates = append(cates, cate...)
	}

	// _, err := qs.All(&cates)
	return cates, err
}

//取出分页的项目
func ListCategoriesByOffsetAndLimit(set, categoriesPerPage int) ([]*Category, []*Label, error) {
	o := orm.NewOrm()
	categories := make([]*Category, 0)
	labels := make([]*Label, 0)
	qs := o.QueryTable("category")
	var err error
	_, err = qs.Filter("ParentId", 0).Limit(categoriesPerPage, set).OrderBy("-created").All(&categories)
	qs1 := o.QueryTable("label")
	_, err = qs1.All(&labels)
	return categories, labels, err
}

//由用户名取得分类：项目
func GetCategoriesbyuname(uname string) ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	var err error
	//这里进行过滤，parentid为空的才显示
	qs = qs.Filter("ParentId", 0).Filter("Author", uname)
	_, err = qs.OrderBy("-created").All(&cates)

	// _, err := qs.All(&cates)
	return cates, err
}

//由分类ID取得分类本身
func GetCategory(id string) (category *Category, labels []*Label, err error) {
	o := orm.NewOrm()
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, nil, err
	}
	// o := orm.NewOrm()
	category = new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("id", idNum).One(category)
	if err != nil {
		return nil, nil, err
	}
	//再查出label
	// var labels []*Label
	_, err = o.QueryTable("label").Filter("Category", idNum).RelatedSel().All(&labels)
	if err != nil {
		// for i, label := range labels {
		// 	if i == 0 {
		// 		label1 = label.Title
		// 	} else {
		// 		label1 = label1 + "," + label.Title
		// 	}
		// }
		return nil, nil, err
	}

	category.Views++
	_, err = o.Update(category)
	return category, labels, err
}

//由分类number（项目编号）取得分类本身
func GetCategoryTitle(number string) (*Category, error) {
	o := orm.NewOrm()
	// Num, err := strconv.ParseInt(number, 10, 64)
	// if err != nil {
	// 	return nil, err
	// }
	// o := orm.NewOrm()
	category := new(Category)
	qs := o.QueryTable("category")
	err := qs.Filter("number", number).One(category, "Title")
	// beego.Info(number)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		beego.Info("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		beego.Info("Not row found")
	}
	return category, err
}

//由分类name（项目名称）取得分类本身
func GetCategoryName(title string) (*Category, error) {
	o := orm.NewOrm()
	// Num, err := strconv.ParseInt(number, 10, 64)
	// if err != nil {
	// 	return nil, err
	// }
	// o := orm.NewOrm()
	category := new(Category)
	qs := o.QueryTable("category")
	err := qs.Filter("title", title).One(category, "Id")
	// beego.Info(number)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		beego.Info("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		beego.Info("Not row found")
	}
	return category, err
}

//由成果类型id取出父一级的专业
func GetCategorySpec(id string) (*Category, error) {
	o := orm.NewOrm()
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	// o := orm.NewOrm()
	category := new(Category)
	qs := o.QueryTable("category")
	// err = qs.Filter("id", idNum).One(category, "ParentId") //返回父级专业id,这句不行
	// err = qs.Filter("id", category).One(category)           //由专业id返回专业
	err = qs.Filter("id", idNum).One(category)             //由id取到成果类型的struct
	err = qs.Filter("id", category.ParentId).One(category) //再由成果类型的父id取得专业struct

	if err != nil {
		return nil, err
	}
	// category.Views++
	// _, err = o.Update(category)
	return category, err
}

//由成果类型id取出父一级的专业和爷一级的阶段
func GetCategoryPhase(id string) (*Category, error) {
	o := orm.NewOrm()
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	// o := orm.NewOrm()
	category := new(Category)
	qs := o.QueryTable("category")
	// err = qs.Filter("id", idNum).One(category, "ParentId") //返回父级专业id,这句不行
	// err = qs.Filter("id", category).One(category)           //由专业id返回专业
	err = qs.Filter("id", idNum).One(category)             //由id取到成果类型的struct
	err = qs.Filter("id", category.ParentId).One(category) //再由成果类型的父id取得专业struct
	err = qs.Filter("id", category.ParentId).One(category) //再由专业的父id取得阶段struct

	if err != nil {
		return nil, err
	}
	// category.Views++
	// _, err = o.Update(category)
	return category, err
}

//由成果类型id取出父一级的专业和爷一级的阶段及祖父一级的项目名称
func GetCategoryProj(id string) (*Category, error) {
	o := orm.NewOrm()
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	// o := orm.NewOrm()
	category := new(Category)
	qs := o.QueryTable("category")
	// err = qs.Filter("id", idNum).One(category, "ParentId") //返回父级专业id,这句不行
	// err = qs.Filter("id", category).One(category)           //由专业id返回专业
	err = qs.Filter("id", idNum).One(category)             //由id取到成果类型的struct
	err = qs.Filter("id", category.ParentId).One(category) //再由成果类型的父id取得专业struct
	err = qs.Filter("id", category.ParentId).One(category) //再由专业的父id取得阶段struct
	err = qs.Filter("id", category.ParentId).One(category)
	if err != nil {
		return nil, err
	}
	// category.Views++
	// _, err = o.Update(category)
	return category, err
}

//由项目id取出所有成果分类.......再由成果分类取出所有成果GetTopicsbyparentid
func GetCategoryChengguo(id string) ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	cates2 := make([]*Category, 0)
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	//这里循环取出子目录
	//取出阶段
	var posts []Category                                                          //详见beego手册的All的示例
	_, err = o.QueryTable("Category").Filter("parentid", idNum).All(&posts, "Id") //获取阶段
	for _, v := range posts {
		var postss []Category //详见beego手册的All的示例——取出文档类型
		_, err = o.QueryTable("Category").Filter("parentid", v.Id).All(&postss, "Id")
		for _, w := range postss {
			// var cates []Category //这句导致无法赋值——取出专业
			_, err = o.QueryTable("Category").Filter("parentid", w.Id).All(&cates2)
			// 给mySlice后面添加另一个数组切片
			cates = append(cates, cates2...)
			// for _, z := range postsss {
			// 	cate := &Category{Id: z.Id}
			// 	_, err = o.Delete(cate)
			// }
		}
	}
	return cates, err
}

//由项目id取出所有文档类型
func GetCategoryLeixing(id string) ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0) //创建一个初始元素个数为0的数组切片，元素初始值为0
	cates2 := make([]*Category, 0)
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	//再取出文件类型
	var posts []Category                                                          //详见beego手册的All的示例
	_, err = o.QueryTable("Category").Filter("parentid", idNum).All(&posts, "Id") //只取出阶段ID
	for _, w := range posts {
		_, err = o.QueryTable("Category").Filter("parentid", w.Id).All(&cates2) //取出每个阶段的文件类型
		// 给mySlice后面添加另一个数组切片
		cates = append(cates, cates2...)
	}
	return cates, err
}

//由项目id取出所有阶段
func GetCategoryJieduan(id string) ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	//再取出阶段分类
	_, err = o.QueryTable("Category").Filter("parentid", idNum).All(&cates)
	return cates, err
}

//由项目编号、阶段、成果类型、专业名取出专业Id
func GetCategoryzhuanye(pronumber, projieduan, proleixing, prozhuanye string) (Id int64, err error) {
	o := orm.NewOrm()
	// cates := make([]*Category, 0)
	cate := Category{Number: pronumber}
	err = o.Read(&cate, "Number") //查询字段Number
	//由阶段parentid和projieduan名查出阶段id
	// jieduan:=Category{Number: cate.}
	var jieduan string
	switch projieduan {
	case "A":
		jieduan = "规划"
	case "B":
		jieduan = "项目建议书"
	case "C":
		jieduan = "可行性研究"
	case "D":
		jieduan = "初步设计"
	case "E":
		jieduan = "招标设计"
	case "F":
		jieduan = "施工图设计"
	case "G":
		jieduan = "竣工图"
	case "L":
		jieduan = "专题"
	}
	var jieduan1 Category
	err = o.QueryTable("category").Filter("title", jieduan).Filter("parentid", cate.Id).One(&jieduan1, "Id")
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		beego.Info("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		beego.Info("Not row found")
	}
	//由阶段id和文件类型名称，查出文件类型id
	var leixing string
	switch proleixing {
	case "FB":
		leixing = "技术报告"
	case "FD":
		leixing = "设计大纲"
	case "FG":
		leixing = "设计/修改通知单"
	case "FT":
		leixing = "工程图纸"
	case "FJ":
		leixing = "计算书"
	case "FP":
		leixing = "PDF文件"
	case "Fdiary":
		leixing = "文章/设代日记"
	}
	var leixing1 Category
	err = o.QueryTable("category").Filter("title", leixing).Filter("parentid", jieduan1.Id).One(&leixing1, "Id")
	//由专业id和类型名称，查出类型id
	var zhuanye string
	switch prozhuanye {
	case "1":
		zhuanye = "综合"
	case "2":
		zhuanye = "规划(含水文、经评)"
	case "3":
		zhuanye = "测量"
	case "4":
		zhuanye = "地质(含钻探)"
	case "5":
		zhuanye = "水工(含公路、安全监测)"
	case "6":
		zhuanye = "建筑"
	case "7":
		zhuanye = "机电"
	case "8":
		zhuanye = "征地、环保、水保"
	case "9":
		zhuanye = "施工、工程造价"
	}
	var zhuanye1 Category
	err = o.QueryTable("category").Filter("title", zhuanye).Filter("parentid", leixing1.Id).One(&zhuanye1, "Id")

	return zhuanye1.Id, err
}

//汉字占1位，begin从0开始，length从1开始，（string.index从0开始）
func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)
	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	// 返回子串
	return string(rs[begin:end])
}

//由category.id返回url和disk
func GetCategoryUrl(id string) (url, diskdirectory string, err error) {
	var jieduan, leixing, zhuanye string
	o := orm.NewOrm()
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return "", "", err
	}
	// o := orm.NewOrm()
	category := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("id", idNum).One(category) //由id取到成果类型的struct
	if err != nil {
		return "", "", err
	}
	if category.Isuserdefined == true { //如果是自定义目录
		if category.ParentId != 0 {
			category1 := new(Category)
			err = qs.Filter("id", category.ParentId).One(category1) //再由成果类型的父id取得专业struct
			if err != nil {
				return "", "", err
			}
			if category1.ParentId != 0 {
				category2 := new(Category)
				err = qs.Filter("id", category1.ParentId).One(category2) //再由专业的父id取得阶段struct
				if err != nil {
					return "", "", err
				}
				if category2.ParentId != 0 {
					category3 := new(Category)
					err = qs.Filter("id", category2.ParentId).One(category3)
					if err != nil {
						return "", "", err
					}
					url = "/attachment/" + category3.Number + category3.Title + "/" + category2.Title + "/" + category1.Title + "/" + category.Title + "/"
					diskdirectory = ".\\attachment\\" + category3.Number + category3.Title + "\\" + category2.Title + "\\" + category1.Title + "\\" + category.Title + "\\"
				} else {
					url = "/attachment/" + category2.Number + category2.Title + "/" + category1.Title + "/" + category.Title + "/"
					diskdirectory = ".\\attachment\\" + category2.Number + category2.Title + "\\" + category1.Title + "\\" + category.Title + "\\"
				}
			} else {
				url = "/attachment/" + category1.Number + category1.Title + "/" + category.Title + "/"
				diskdirectory = ".\\attachment\\" + category1.Number + category1.Title + "\\" + category.Title + "\\"
			}
		} else {
			url = "/attachment/" + category.Number + category.Title + "/"
			diskdirectory = ".\\attachment\\" + category.Number + category.Title + "\\"
		}
	} else { //如果是标准目录
		if category.ParentId != 0 {
			category1 := new(Category)
			err = qs.Filter("id", category.ParentId).One(category1) //再由成果类型的父id取得专业struct
			if err != nil {
				return "", "", err
			}
			if category1.ParentId != 0 {
				category2 := new(Category)
				err = qs.Filter("id", category1.ParentId).One(category2) //再由专业的父id取得阶段struct
				if err != nil {
					return "", "", err
				}
				if category2.ParentId != 0 {
					category3 := new(Category)
					err = qs.Filter("id", category2.ParentId).One(category3)
					if err != nil {
						return "", "", err
					}

					switch category2.Title {
					case "规划":
						jieduan = "A"
					case "项目建议书":
						jieduan = "B"
					case "可行性研究":
						jieduan = "C"
					case "初步设计":
						jieduan = "D"
					case "招标设计":
						jieduan = "E"
					case "施工图设计":
						jieduan = "F"
					case "竣工图":
						jieduan = "G"
					case "专题":
						jieduan = "L"
					}
					switch category1.Title {
					case "技术报告":
						leixing = "FB"
					case "设计大纲":
						leixing = "FD"
					case "设计/修改通知单":
						leixing = "FG"
					case "工程图纸":
						leixing = "FT"
					case "计算书":
						leixing = "FJ"
					case "PDF文件":
						leixing = "FP"
					case "文章/设代日记":
						leixing = "Fdiary"
					}
					switch category.Title {
					case "综合":
						zhuanye = "1"
					case "规划(含水文、经评)":
						zhuanye = "2"
					case "测量":
						zhuanye = "3"
					case "地质(含钻探)":
						zhuanye = "4"
					case "水工(含公路、安全监测)":
						zhuanye = "5"
					case "建筑":
						zhuanye = "6"
					case "机电":
						zhuanye = "7"
					case "征地、环保、水保":
						zhuanye = "8"
					case "施工、工程造价":
						zhuanye = "9"
					}

					url = "/attachment/" + category3.Number + category3.Title + "/" + jieduan + "/" + leixing + "/" + zhuanye + "/"
					diskdirectory = ".\\attachment\\" + category3.Number + category3.Title + "\\" + jieduan + "\\" + leixing + "\\" + zhuanye + "\\"
				} else {
					switch category1.Title {
					case "规划":
						jieduan = "A"
					case "项目建议书":
						jieduan = "B"
					case "可行性研究":
						jieduan = "C"
					case "初步设计":
						jieduan = "D"
					case "招标设计":
						jieduan = "E"
					case "施工图设计":
						jieduan = "F"
					case "竣工图":
						jieduan = "G"
					case "专题":
						jieduan = "L"
					}
					switch category.Title {
					case "技术报告":
						leixing = "FB"
					case "设计大纲":
						leixing = "FD"
					case "设计/修改通知单":
						leixing = "FG"
					case "工程图纸":
						leixing = "FT"
					case "计算书":
						leixing = "FJ"
					case "PDF文件":
						leixing = "FP"
					case "文章/设代日记":
						leixing = "Fdiary"
					}

					url = "/attachment/" + category2.Number + category2.Title + "/" + jieduan + "/" + leixing + "/"
					diskdirectory = ".\\attachment\\" + category2.Number + category2.Title + "\\" + jieduan + "\\" + leixing + "\\"
				}
			} else {
				switch category.Title {
				case "规划":
					jieduan = "A"
				case "项目建议书":
					jieduan = "B"
				case "可行性研究":
					jieduan = "C"
				case "初步设计":
					jieduan = "D"
				case "招标设计":
					jieduan = "E"
				case "施工图设计":
					jieduan = "F"
				case "竣工图":
					jieduan = "G"
				case "专题":
					jieduan = "L"
				}
				url = "/attachment/" + category1.Number + category1.Title + "/" + jieduan + "/"
				diskdirectory = ".\\attachment\\" + category1.Number + category1.Title + "\\" + jieduan + "\\"
			}
		} else {
			url = "/attachment/" + category.Number + category.Title + "/"
			diskdirectory = ".\\attachment\\" + category.Number + category.Title + "\\"
		}
	}
	return url, diskdirectory, err
}

func SearchCategories(categoryname string, isDesc bool) ([]*Category, []*Label, error) {
	o := orm.NewOrm()
	categories := make([]*Category, 0)
	qs := o.QueryTable("category")
	qs1 := o.QueryTable("label")
	var err error
	if isDesc {
		if len(categoryname) > 0 {
			qs = qs.Filter("Title__contains", categoryname).Filter("ParentId", 0) //这里取回
		}
		_, err = qs.OrderBy("-created").All(&categories)
	} else {
		_, err = qs.Filter("Title__contains", categoryname).Filter("ParentId", 0).OrderBy("-created").All(&categories)
		//o.QueryTable("user").Filter("name", "slene").All(&users)
	}
	labels := make([]*Label, 0)
	_, err = qs1.All(&labels)
	return categories, labels, err
}
