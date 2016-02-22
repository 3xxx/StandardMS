package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"strings"
	"time"
)

type Category struct {
	Id              int64 `form:"-"`
	ParentId        int64
	Uid             int64
	Title           string `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Number          string `orm:"unique",form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"`
	Content         string `orm:"sie(5000)"`
	Route           string
	Created         time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated         time.Time `orm:"index","auto_now_add;type(datetime)"`
	Views           int64     `form:"-",orm:"index"`
	Author          string
	TopicCount      int64  //`form:"-"`
	TopicLastUserId int64  //`form:"-"`
	DiskDirectory   string `orm:"null"`
	Url             string `orm:"null"`
	// Type            string `orm:"null"` //项目类型：供水、枢纽、提防、河道、船闸、电站、水闸
}

func init() {
	orm.RegisterModel(new(Category)) //, new(Article)
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/hydrocms.db", 10)
}

func AddCategory(name, number, content, path, route, uname, diskdirectory, url string) (id int64, err error) {
	o := orm.NewOrm()
	cate := &Category{
		Title:         name,
		Number:        number,
		Content:       content,
		Author:        uname,
		Route:         route,
		Created:       time.Now(),
		Updated:       time.Now(),
		DiskDirectory: diskdirectory,
		Url:           url,
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
				DiskDirectory: ".\\attachment\\" + number + name + "\\" + v + "\\",
				Url:           "/attachment/" + number + name + "/" + v + "/",
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
			var posts []Category                                                                                    //详见beego手册的All的示例
			_, err = o.QueryTable("Category").Filter("parentid", post.Id).All(&posts, "Id", "DiskDirectory", "Url") //Id没什么用？
			for _, w := range posts {
				cate := &Category{
					Title:         leixing, //v,
					ParentId:      w.Id,
					Created:       time.Now(),
					Updated:       time.Now(),
					Author:        uname,
					DiskDirectory: diskdirectory + w.DiskDirectory + v + "\\",
					Url:           url + w.Url + v + "/",
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
				_, err = o.QueryTable("Category").Filter("parentid", w.Id).All(&postss, "Id", "DiskDirectory", "Url")
				for _, t := range postss {
					cate := &Category{
						Title:         zhuanye, //v,
						ParentId:      t.Id,
						Created:       time.Now(),
						Updated:       time.Now(),
						Author:        uname,
						DiskDirectory: diskdirectory + t.DiskDirectory + v + "\\",
						Url:           url + t.Url + v + "/",
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
func AdduserdefinedCategory(name, number, content string, path2, path3, path4 []string, route, uname, diskdirectory, url string) (id int64, err error) {
	o := orm.NewOrm()
	cate := &Category{
		Title:         name,
		Number:        number,
		Content:       content,
		Author:        uname,
		Route:         route,
		Created:       time.Now(),
		Updated:       time.Now(),
		DiskDirectory: diskdirectory,
		Url:           url,
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
	var post Category //取出number项目编号的Id
	err = o.QueryTable("Category").Filter("number", number).One(&post, "Id", "Title")
	if err != nil {
		return 0, err
	}
	// array := strings.Split(path, ",") //字符串切割 [a b c d e]
	if len(path2) > 0 {
		for _, v := range path2 {
			cate = &Category{
				Title:         v,
				ParentId:      post.Id, //这里存入项目的id
				Created:       time.Now(),
				Updated:       time.Now(),
				Author:        uname,
				DiskDirectory: diskdirectory + v + "\\",
				Url:           url + v + "/",
			}
			_, err = o.Insert(cate)
			if err != nil {
				return 0, err
			}
		}
	}
	if len(path3) > 0 {
		for _, v := range path3 {
			//查到阶段的parentid，符合这个项目的，得出阶段id，作为专业parentid
			var posts []Category //详见beego手册的All的示例
			_, err = o.QueryTable("Category").Filter("parentid", post.Id).All(&posts, "Id", "Title")
			for _, w := range posts {
				cate := &Category{
					Title:         v,
					ParentId:      w.Id,
					Created:       time.Now(),
					Updated:       time.Now(),
					Author:        uname,
					DiskDirectory: diskdirectory + w.Title + "\\" + v + "\\",
					Url:           url + w.Title + "/" + v + "/",
				}
				_, err = o.Insert(cate)
				if err != nil {
					return 0, err
				}
			}

		}
	}

	if len(path4) > 0 {
		for _, v := range path4 {
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
					cate := &Category{
						Title:         v,
						ParentId:      t.Id,
						Created:       time.Now(),
						Updated:       time.Now(),
						Author:        uname,
						DiskDirectory: diskdirectory + w.Title + "\\" + t.Title + "\\" + v + "\\",
						Url:           url + w.Title + "/" + t.Title + "/" + v + "/",
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

func ModifyCategory(cid, name, number, content, path, route, uname string) error {
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
		if route == "" { //如果没有更新图片，则不更新图片地址
			cate.Title = name
			cate.Number = number
			cate.Content = content
			cate.Author = uname
			// cate.Route = route
			// cate.Created: time.Now(),
			cate.Updated = time.Now()
			_, err = o.Update(cate)
			if err != nil {
				return err
			}
		} else {
			cate.Title = name
			cate.Number = number
			cate.Content = content
			cate.Author = uname
			cate.Route = route
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

func DelCategory(id string) error { //应该显示警告
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	//先删成果分类
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

	//再删专业分类
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

//取出分页的项目
func ListCategoriesByOffsetAndLimit(set, categoriesPerPage int) ([]*Category, error) {
	o := orm.NewOrm()
	categories := make([]*Category, 0)
	qs := o.QueryTable("category")
	var err error
	_, err = qs.Filter("ParentId", 0).Limit(categoriesPerPage, set).OrderBy("-created").All(&categories)
	return categories, err
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
func GetCategory(id string) (*Category, error) {
	o := orm.NewOrm()
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	// o := orm.NewOrm()
	category := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("id", idNum).One(category)
	if err != nil {
		return nil, err
	}
	category.Views++
	_, err = o.Update(category)
	return category, err
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
	beego.Info(number)
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
	//取出成果分类
	var posts []Category                                                          //详见beego手册的All的示例
	_, err = o.QueryTable("Category").Filter("parentid", idNum).All(&posts, "Id") //获取阶段
	for _, v := range posts {
		var postss []Category //详见beego手册的All的示例
		_, err = o.QueryTable("Category").Filter("parentid", v.Id).All(&postss, "Id")
		for _, w := range postss {
			// var cates []Category //这句导致无法赋值
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

//由项目id取出所有文件类型
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
