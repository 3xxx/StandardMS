package models

import (
	//"database/sql"
	//"github.com/astaxie/beedb"
	//_ "github.com/ziutek/mymysql/godrv"
	//"time"
	// "fmt"
	// "os"
	// "path"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"time"
	//"github.com/Unknwon/com
	// "errors"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	_ "github.com/mattn/go-sqlite3"
)

//const (
//	_DB_NAME        = "database/orm_test.db"
//	_SQLITE3_DRIVER = "sqlite3"
//)

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
}
type Topic struct {
	Id         int64
	Uid        int64
	Title      string
	Tnumber    string //`orm:"unique"`
	Category   string
	CategoryId int64
	Content    string `orm:"sie(5000)"`
	Attachment string
	// Attachments     []*Attachment `orm:"reverse(many)"` // fk 的反向关系
	Created         time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated         time.Time `orm:"index","auto_now;type(datetime)"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

//附件,attachment 和 topic 是 ManyToOne 关系，也就是 ForeignKey 为 topic
type Attachment struct {
	Id            int64
	Uid           int64
	FileName      string
	FileSize      string
	Downloads     int64
	DiskDirectory string
	Route         string
	Content       string    `orm:"sie(200)"`
	TopicId       int64     //*Topic    `orm:"rel(fk)"`
	Created       time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"index","auto_now;type(datetime)"`
	Views         int64     `orm:"index"`
	Author        string
}

//评论
type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

// type User struct {
// 	Id       int    `PK`
// 	Username string `orm:"unique"`
// 	Pwd      string
// Id            int64
// Username      string    `orm:"unique;size(32)" form:"Username"  valid:"Required;MaxSize(20);MinSize(6)"`
// Password      string    `orm:"size(32)" form:"Password" valid:"Required;MaxSize(20);MinSize(6)"`
// 	Repassword    string    `orm:"-" form:"Repassword" valid:"Required"`
// 	Nickname      string    `orm:"unique;size(32)" form:"Nickname" valid:"Required;MaxSize(20);MinSize(2)"`
// 	Email         string    `orm:"size(32)" form:"Email" valid:"Email"`
// 	Remark        string    `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
// 	Status        int       `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
// 	Lastlogintime time.Time `orm:"null;type(datetime)" form:"-"`
// 	Createtime    time.Time `orm:"type(datetime);auto_now_add" `
// 	Role          []*Role   `orm:"rel(m2m)"`
// }

// type Article struct {
// 	Id     int    `form:"-"`
// 	Name   string `form:"name,text,name:" valid:"MinSize(5);MaxSize(20)"`
// 	Client string `form:"client,text,client:"`
// 	Url    string `form:"url,text,url:"`
// }

func init() {
	orm.RegisterModel(new(Category), new(Topic), new(Attachment), new(Comment), new(Catalog)) //, new(Article)
	orm.RegisterDriver("sqlite", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db", 10)
}

// 多字段唯一键
// func (c *Category) TableUnique() [][]string {
// 	return [][]string{
// 		[]string{"Title", "Number"},
// 	}
// }
// func getLink() beedb.Model {
// 	db, err := sql.Open("mysql", "root:root@tcp(192.168.1.81:3306)/test_my?charset=utf8")
// 	if err != nil {
// 		panic(err)
// 	}
// 	orm := beedb.New(db)
// 	return orm
// }

// func SaveUser(user User) error {
// 	orm := getLink()
// 	fmt.Println(user)
// 	err := orm.Save(&user)
// 	return err
// }

// func ValidateUser(user User) error {
// 	orm := getLink()
// 	var u User
// 	orm.Where("username=? and pwd=?", user.Username, user.Pwd).Find(&u)
// 	if u.Username == "" {
// 		return errors.New("用户名或密码错误！")
// 	}
// 	return nil
// }
// func SaveUser(user User) error {
// 	orm := orm.NewOrm()
// 	fmt.Println(user)
// 	_, err := orm.Insert(&user) //_, err = o.Insert(reply)
// 	return err
// }

// func ValidateUser(user User) error {
// 	orm := orm.NewOrm()
// 	var u User

// 	// user = new(User)
// 	qs := orm.QueryTable("user")
// 	err := qs.Filter("username", user.Username).Filter("pwd", user.Pwd).One(&u)
// 	if err != nil {
// 		return err
// 	}

// 	// orm.Where("username=? and pwd=?", user.Username, user.Pwd).Find(&u)
// 	if u.Username == "" {
// 		return errors.New("用户名或密码错误！")
// 	}
// 	return nil
// }

// func CheckUname(user User) error {
// 	orm := orm.NewOrm()
// 	var u User
// 	// user = new(User)
// 	qs := orm.QueryTable("user")
// 	err := qs.Filter("username", user.Username).One(&u)
// 	if err != nil {
// 		return err
// 	}
// 	// orm.Where("username=? and pwd=?", user.Username, user.Pwd).Find(&u)
// 	// if u.Username == "" {
// 	// 	return errors.New("用户名或密码错误！")
// 	// }
// 	return nil
// }

func DeleteReply(rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	var tidNum int64
	reply := &Comment{Id: ridNum}
	if o.Read(reply) == nil {
		tidNum = reply.Tid
		_, err = o.Delete(reply)
		if err != nil {
			return err
		}
	}
	replies := make([]*Comment, 0) //slice,将所有文章取出来
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).OrderBy("-created").All(&replies)
	if err != nil {
		return err
	}
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil { //如果错误为空
		if len(replies) != 0 { //如果回复不为空，则……
			topic.ReplyTime = replies[0].Created
			topic.ReplyCount = int64(len(replies))
			_, err = o.Update(topic)
		}
	}
	return err
}

func AddReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	reply := &Comment{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}
	o := orm.NewOrm()
	_, err = o.Insert(reply)
	if err != nil {
		return err
	}
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.ReplyTime = time.Now()
		topic.ReplyCount++
		_, err = o.Update(topic)
	}
	return err
}

func GetAllReplies(tid string) (replies []*Comment, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	replies = make([]*Comment, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).All(&replies)
	return replies, err

}

// func (a *Article) TableName() string {
// 	return "articles"
// }

// func GetArticleById(id int) (v *Article, err error) {
// 	o := orm.NewOrm()
// 	v = &Article{Id: id}
// 	if err = o.Read(v); err == nil {
// 		return v, nil
// 	}
// 	return nil, err
// }

//一对一模式
func AddTopicOne(title, tnumber, category, categoryid, uname, content, attachment string) (id int64, err error) {
	cid, err := strconv.ParseInt(categoryid, 10, 64)
	o := orm.NewOrm()
	//判断在这个类型id下，是否有重号
	var topic Topic //下面这个filter放在topic=&Topic{后面用返回one(topic)则查询出错！
	err = o.QueryTable("topic").Filter("categoryid", cid).Filter("tnumber", tnumber).One(&topic, "Id")
	if err == orm.ErrNoRows { //Filter("tnumber", tnumber).One(topic, "Id")==nil则无法建立
		// 没有找到记录
		topic1 := &Topic{
			Title:      title,
			Tnumber:    tnumber,
			Category:   category,
			CategoryId: cid,
			Content:    content,
			Attachment: attachment,
			Author:     uname,
			Created:    time.Now(),
			Updated:    time.Now(),
			ReplyTime:  time.Now(),
		}
		// err = o.QueryTable("topic").Filter("category_id", cid).One(&topic, "id")
		// if err != nil {.Filter("tnumber", tnumber)
		// 	return 0, err
		// }
		// if err == orm.ErrMultiRows {
		// 	// 多条的时候报错
		// 	// fmt.Printf("Returned Multi Rows Not One")
		// }
		// if err == orm.ErrNoRows {
		// 没有找到记录
		// fmt.Printf("Not row found")

		//	qs := o.QueryTable("category") //不知道主键就用这个过滤操作
		//	err := qs.Filter("title", name).One(cate)
		//	if err == nil {
		//		return err
		//	}

		id, err = o.Insert(topic1)
		if err != nil {
			return id, err //如果文章编号相同，则唯一性检查错误，返回id吗？
		}
	}
	// if id == 0 {//这个是为一对多模式用的。
	// 	var topic Topic
	// 	err = o.QueryTable("topic").Filter("tnumber", tnumber).One(&topic, "Id")
	// 	id = topic.Id
	// }
	//更新分类统计
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("id", categoryid).One(cate)
	// beego.Info(categoryid)
	if err == nil {
		//如果不存在，简单地忽略更新
		//如果想精确地统计，用获取所有文章就调用getalltopic方法
		cate.TopicCount++
		// beego.Info(cate.TopicCount)
		_, err = o.Update(cate)
	}
	// }
	return id, err
}

//一对多模式
func AddTopicMany(title, tnumber, category, categoryid, uname, content, attachment string) (id int64, err error) {
	cid, err := strconv.ParseInt(categoryid, 10, 64)
	o := orm.NewOrm()
	//判断在这个类型id下，是否有重号
	var topic Topic //下面这个filter放在topic=&Topic{后面用返回one(topic)则查询出错！
	err = o.QueryTable("topic").Filter("categoryid", cid).Filter("tnumber", tnumber).One(&topic, "Id")
	if err == orm.ErrNoRows { //Filter("tnumber", tnumber).One(topic, "Id")==nil则无法建立
		// 没有找到记录
		topic1 := &Topic{
			Title:      title,
			Tnumber:    tnumber,
			Category:   category,
			CategoryId: cid,
			Content:    content,
			Attachment: attachment,
			Author:     uname,
			Created:    time.Now(),
			Updated:    time.Now(),
			ReplyTime:  time.Now(),
		}
		// if err == orm.ErrMultiRows {
		// 	// 多条的时候报错
		// 	// fmt.Printf("Returned Multi Rows Not One")
		// }
		// if err == orm.ErrNoRows {
		// 没有找到记录
		// fmt.Printf("Not row found")

		//	qs := o.QueryTable("category") //不知道主键就用这个过滤操作
		//	err := qs.Filter("title", name).One(cate)
		//	if err == nil {
		//		return err
		//	}
		id, err = o.Insert(topic1)
		if err != nil {
			return id, err //如果文章编号相同，则唯一性检查错误，返回id吗？
		}
		// if id == 0 {//这个是为一对多模式用的。
		// 	var topic Topic
		// 	err = o.QueryTable("topic").Filter("tnumber", tnumber).One(&topic, "Id")
		// 	id = topic.Id
		// }
		//更新分类统计
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("id", categoryid).One(cate)
		// beego.Info(categoryid)
		if err == nil {
			//如果不存在，简单地忽略更新
			//如果想精确地统计，用获取所有文章就调用getalltopic方法
			cate.TopicCount++
			// beego.Info(cate.TopicCount)
			_, err = o.Update(cate)
		}
	} else {
		id = topic.Id
	}
	return id, err
}

func GetTopicIdbytnumber(tnumber string) (id int64, err error) { //想把这个函数与addtopic合并，可怎么也不成功。
	o := orm.NewOrm()
	var topic Topic
	err = o.QueryTable("topic").Filter("tnumber", tnumber).One(&topic, "Id")
	id = topic.Id
	return id, err
}

func AddAttachment(filename, filesize, diskdirectory, route, topicid, author string) error {
	tid, err := strconv.ParseInt(topicid, 10, 64)
	o := orm.NewOrm()
	attachment := &Attachment{
		FileName: filename,
		FileSize: filesize,
		// Downloads     int64
		DiskDirectory: diskdirectory,
		Route:         route,
		TopicId:       tid,
		Created:       time.Now(),
		Updated:       time.Now(),
		// Views         int64     `orm:"index"`
		Author: author,
	}
	//	qs := o.QueryTable("category") //不知道主键就用这个过滤操作
	//	err := qs.Filter("title", name).One(cate)
	//	if err == nil {
	//		return err
	//	}
	_, err = o.Insert(attachment)
	if err != nil {
		return err
	}
	//更新分类统计
	// atta := new(Attachment)
	// qs := o.QueryTable("attachment")
	// err = qs.Filter("fileName", category).One(cate)
	// if err == nil {
	// 	//如果不存在，简单地忽略更新
	// 	//如果想精确地统计，用获取所有文章就调用getalltopic方法
	// 	cate.TopicCount++
	// 	_, err = o.Update(cate)
	// }
	return err
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

	for _, v := range array {
		switch v {
		case "ghj", "xj", "ky", "cs", "zb", "sgt", "jgt":
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
			//建立目录——注意，models中无法建立目录，必须在controllers中才行
		}
	}

	for _, v := range array {
		switch v {
		case "gh", "sg", "jd", "shg", "dz", "ys", "zh": //专业
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

	for _, v := range array {
		switch v {
		case "dwg", "doc", "xls", "pdf", "jpg", "tif", "diary": //成果分类
			//查到专业的parentid，符合这个阶段的，得出专业id，作为成果分类的parentid
			var posts []Category //详见beego手册的All的示例
			_, err = o.QueryTable("Category").Filter("parentid", post.Id).All(&posts, "Id", "Title")
			for _, w := range posts {
				var postss []Category //详见beego手册的All的示例
				_, err = o.QueryTable("Category").Filter("parentid", w.Id).All(&postss, "Id", "Title")
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

func DeletTopic(tid string) error { //应该在controllers中显示警告
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	var oldCateId int64
	o := orm.NewOrm()
	// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	// user := User{Name: "slene"}
	// err = o.Read(&user, "Name")

	topic := Topic{Id: tidNum}
	if o.Read(&topic) == nil {
		oldCateId = topic.CategoryId
		_, err = o.Delete(&topic)
		if err != nil {
			return err
		}
	}

	attachment := Attachment{TopicId: tidNum}
	if o.Read(&attachment, "TopicId") == nil {
		// oldCate = topic.Category
		_, err = o.Delete(&attachment)
		if err != nil {
			return err
		}
	}

	if oldCateId > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("id", oldCateId).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}
	_, err = o.Delete(&topic) //这句为何重复？
	return err
}

//缺少排序，由项目名称获取项目下所有成果，如果没有项目名称，则获取所有成果
func GetAllTopics(cate string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	if len(cate) > 0 { //这一半整个系统没有用到？？
		qs := o.QueryTable("Category")
		category := new(Category)
		err = qs.Filter("Title", cate).One(category) //由项目名称获得项目id
		if err != nil {
			return nil, err
		}
		id := strconv.FormatInt(category.Id, 10)
		categoryid, _ := GetCategoryChengguo(id) //由项目id获得所有成果分类
		//因为这里的categoryid有很多组，所以要用range才行
		for _, v := range categoryid {
			id = strconv.FormatInt(v.Id, 10)
			topics2, _ := GetTopicsbyparentid(id, true) //由成果分类获得所有成果
			topics = append(topics, topics2...)
		}
		// var err error
		// if isDesc {
		// 	// if len(cate) > 0 {
		// 	// 	qs = qs.Filter("Title", cate) //这里取回GetCategoryChengguo
		// 	// }
		// 	_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.OrderBy("-created").All(&topics)
	}
	return topics, err
}

//取出分页的文章
func ListPostsByOffsetAndLimit(set, postsPerPage int) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	_, err = qs.Limit(postsPerPage, set).OrderBy("-created").All(&topics)
	return topics, err
}

//由成果分类id取出所有成果
func GetTopicsbyparentid(id string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	// attachments := make([]*Attachment, 0)
	qs := o.QueryTable("topic")
	// qs1 := o.QueryTable("attachment")
	var err error
	if isDesc {
		if len(id) > 0 {
			idNum, _ := strconv.ParseInt(id, 10, 64)
			if err != nil {
				return nil, err
			}
			qs = qs.Filter("categoryid", idNum) //这里总是习惯性写成parentid造成错误
		}
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
	// var posts []Topic
	// _, err = o.QueryTable("Topic").Filter("parentid", idNum).All(&posts)
	// return posts, err

}

func SearchTopics(tuming string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	if isDesc {
		if len(tuming) > 0 {
			qs = qs.Filter("Title__contains", tuming) //这里取回
		}
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.Filter("Title__contains", tuming).OrderBy("-created").All(&topics)
		//o.QueryTable("user").Filter("name", "slene").All(&users)
	}
	return topics, err
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

//由项目id取出所有专业
func GetCategoryZhuanye(id string) ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0) //创建一个初始元素个数为0的数组切片，元素初始值为0
	cates2 := make([]*Category, 0)
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	//再取出专业分类
	var posts []Category                                                          //详见beego手册的All的示例
	_, err = o.QueryTable("Category").Filter("parentid", idNum).All(&posts, "Id") //只取出阶段ID
	for _, w := range posts {
		_, err = o.QueryTable("Category").Filter("parentid", w.Id).All(&cates2) //取出每个阶段的专业
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

//由项目编号、阶段、专业、成果类型取出成果类型Id
func GetCategoryleixing(pronumber, projieduan, proleixing, prozhuanye string) (Id int64, err error) {
	o := orm.NewOrm()
	// cates := make([]*Category, 0)
	cate := Category{Number: pronumber}
	err = o.Read(&cate, "Number") //查询字段Number
	//由阶段parentid和projieduan名查出阶段id
	// jieduan:=Category{Number: cate.}
	var jieduan Category
	err = o.QueryTable("category").Filter("name", projieduan).Filter("parentid", cate.Id).One(&jieduan, "Id")
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		beego.Info("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		beego.Info("Not row found")
	}
	//由阶段id和专业名称，查出专业id
	var zhuanye Category
	err = o.QueryTable("category").Filter("name", prozhuanye).Filter("parentid", jieduan.Id).One(&zhuanye, "Id")
	//由专业id和类型名称，查出类型id
	var leixing Category
	err = o.QueryTable("category").Filter("name", proleixing).Filter("parentid", zhuanye.Id).One(&leixing, "Id")

	return leixing.Id, err
}
func GetTopic(tid string) (*Topic, []*Attachment, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, nil, err
	}
	topic.Views++
	_, err = o.Update(topic)

	attachments := make([]*Attachment, 0)
	// attachment := new(Attachment)
	qs = o.QueryTable("attachment")
	_, err = qs.Filter("topicId", tidNum).OrderBy("FileName").All(&attachments)
	if err != nil {
		return nil, nil, err
	}
	return topic, attachments, err
}

//由用户名取得文章
func Gettopicsbyuname(uname string) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error

	qs = qs.Filter("Author", uname)
	_, err = qs.OrderBy("-created").All(&topics)

	// _, err := qs.All(&cates)
	return topics, err
}

//由文章id得到父级文章类型
func GetTopicChengguo(tid string) (*Topic, *Category, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, nil, err
	}

	// o = orm.NewOrm()
	category := new(Category)
	qs = o.QueryTable("category")
	err = qs.Filter("id", topic.CategoryId).One(category)
	if err != nil {
		return nil, nil, err
	}
	// topic.Views++
	// _, err = o.Update(topic)
	return topic, category, err
}

//由文章id得到父级文章类型和爷级专业
func GetTopicSpec(tid string) (*Category, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	// o = orm.NewOrm()
	category := new(Category)
	qs = o.QueryTable("category")
	err = qs.Filter("id", topic.CategoryId).One(category)
	if err != nil {
		return nil, err
	}
	err = qs.Filter("id", category.ParentId).One(category)
	if err != nil {
		return nil, err
	}
	// topic.Views++
	// _, err = o.Update(topic)
	return category, err
}

//由文章id得到父级文章类型和祖父级阶段
func GetTopicPhase(tid string) (*Category, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	// o = orm.NewOrm()
	category := new(Category)
	qs = o.QueryTable("category")
	err = qs.Filter("id", topic.CategoryId).One(category)
	if err != nil {
		return nil, err
	}
	err = qs.Filter("id", category.ParentId).One(category)
	if err != nil {
		return nil, err
	}
	err = qs.Filter("id", category.ParentId).One(category)
	if err != nil {
		return nil, err
	}
	// topic.Views++
	// _, err = o.Update(topic)
	return category, err
}

//由文章id得到父级文章类型和曾祖父级项目
func GetTopicProj(tid string) (*Category, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	// o = orm.NewOrm()
	category := new(Category)
	qs = o.QueryTable("category")
	err = qs.Filter("id", topic.CategoryId).One(category)
	if err != nil {
		return nil, err
	}
	err = qs.Filter("id", category.ParentId).One(category)
	if err != nil {
		return nil, err
	}
	err = qs.Filter("id", category.ParentId).One(category)
	if err != nil {
		return nil, err
	}
	err = qs.Filter("id", category.ParentId).One(category)
	if err != nil {
		return nil, err
	}
	// topic.Views++
	// _, err = o.Update(topic)
	return category, err
}

//只修改编号、名称和内容，不修改附件及分类树状目录
func ModifyTopic(tid, title, tnumber, category, categoryid, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	cid, err := strconv.ParseInt(categoryid, 10, 64)
	if err != nil {
		return err
	}

	// var oldAttach string //oldCate,
	var oldCate string
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		// oldAttach = topic.Attachment
		oldCate = topic.Category
		topic.Title = title
		topic.Tnumber = tnumber
		topic.Category = category
		topic.CategoryId = cid
		topic.Content = content
		// topic.Attachment = attachment
		topic.Updated = time.Now()
		_, err = o.Update(topic)
		if err != nil {
			return err
		}
	}

	//删除旧的附件
	// if len(oldAttach) > 0 {
	// 	os.Remove(path.Join("attachment", oldAttach))
	// }
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err := qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
	}
	return err
}

//设代日记：由图片附件的id，存入图片的content
func ModifyAttachment(aid, content string) error {
	aidNum, err := strconv.ParseInt(aid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	attachment := &Attachment{Id: aidNum}
	if o.Read(attachment) == nil {
		attachment.Content = content

		attachment.Updated = time.Now()
		_, err = o.Update(attachment)
		if err != nil {
			return err
		}
	}
	return err
}

//删除文章中的附件
func DeletAttachment(aid string) error { //应该显示警告
	aidNum, err := strconv.ParseInt(aid, 10, 64)
	if err != nil {
		return err
	}

	// var oldCate string
	o := orm.NewOrm()
	// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	// user := User{Name: "slene"}
	// err = o.Read(&user, "Name")

	// topic := Topic{Id: tidNum}
	// if o.Read(&topic) == nil {
	// 	oldCate = topic.Category
	// 	_, err = o.Delete(&topic)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	attachment := Attachment{Id: aidNum}
	if o.Read(&attachment) == nil {
		// oldCate = topic.Category
		_, err = o.Delete(&attachment)
		if err != nil {
			return err
		}
	}

	// if len(oldCate) > 0 {
	// 	cate := new(Category)
	// 	qs := o.QueryTable("category")
	// 	err = qs.Filter("title", oldCate).One(cate)
	// 	if err == nil {
	// 		cate.TopicCount--
	// 		_, err = o.Update(cate)
	// 	}
	// }
	// _, err = o.Delete(&topic) //这句为何重复？
	return err
}

func TopicCount(id, count int64) (err error) {
	o := orm.NewOrm()
	//	qs := o.QueryTable("category") //不知道主键就用这个过滤操作
	//	err := qs.Filter("title", name).One(cate)
	//	if err == nil {
	//		return err
	//	}

	//更新分类统计
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("id", id).One(cate)
	// beego.Info(categoryid)
	if err == nil {
		//如果不存在，简单地忽略更新
		//如果想精确地统计，用获取所有文章就调用getalltopic方法
		cate.TopicCount = count
		// beego.Info(cate.TopicCount)
		_, err = o.Update(cate)
	}
	return err
}

func AddCatalog(name, tnumber string) (id int64, err error) {
	// cid, err := strconv.ParseInt(categoryid, 10, 64)
	o := orm.NewOrm()
	catalog := &Catalog{
		Name:    name,
		Tnumber: tnumber,
		// Category:   category,
		// CategoryId: cid,
		// Content:    content,
		// Attachment: attachment,
		// Author:     uname,
		// Created:    time.Now(),
		// Updated:    time.Now(),
		// ReplyTime:  time.Now(),
	}
	//	qs := o.QueryTable("category") //不知道主键就用这个过滤操作
	//	err := qs.Filter("title", name).One(cate)
	//	if err == nil {
	//		return err
	//	}
	id, err = o.Insert(catalog)
	if err != nil {
		return id, err //如果文章编号相同，则唯一性检查错误，返回id吗？
	}
	if id == 0 {
		var catalog Catalog
		err = o.QueryTable("catalog").Filter("tnumber", tnumber).One(&catalog, "Id")
		id = catalog.Id
	}

	return id, err
}

func ModifyCatalog(tid, title, tnumber string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	_, err = strconv.ParseInt(title, 10, 64)
	if err != nil {
		return err
	}

	// var oldAttach string //oldCate,
	// var oldCatalog string
	o := orm.NewOrm()
	catalog := &Catalog{Id: tidNum}
	if o.Read(catalog) == nil {
		// oldAttach = topic.Attachment

		catalog.Name = title
		catalog.Tnumber = tnumber

		// topic.Attachment = attachment
		// catalog.Updated = time.Now()
		_, err = o.Update(catalog)
		if err != nil {
			return err
		}
	}
	return err
}

func DeletCatalog(tid string) error { //应该在controllers中显示警告
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	var oldCateId int64
	o := orm.NewOrm()
	// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	// user := User{Name: "slene"}
	// err = o.Read(&user, "Name")

	topic := Topic{Id: tidNum}
	if o.Read(&topic) == nil {
		oldCateId = topic.CategoryId
		_, err = o.Delete(&topic)
		if err != nil {
			return err
		}
	}

	attachment := Attachment{TopicId: tidNum}
	if o.Read(&attachment, "TopicId") == nil {
		// oldCate = topic.Category
		_, err = o.Delete(&attachment)
		if err != nil {
			return err
		}
	}

	if oldCateId > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("id", oldCateId).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}
	_, err = o.Delete(&topic) //这句为何重复？
	return err
}
func GetCatalog(tid string) (*Catalog, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	catalog := new(Catalog)
	qs := o.QueryTable("catalog")
	err = qs.Filter("id", tidNum).One(catalog)
	if err != nil {
		return nil, err
	}
	// catalog.Views++
	// _, err = o.Update(topic)

	// attachments := make([]*Attachment, 0)
	// attachment := new(Attachment)
	// qs = o.QueryTable("attachment")
	// _, err = qs.Filter("topicId", tidNum).OrderBy("FileName").All(&attachments)
	// if err != nil {
	// 	return nil, err
	// }
	return catalog, err
}

//根据路由获取category表的author
func GetattatchAuthor(route string) (uname string, err error) {
	o := orm.NewOrm()
	var attachment Attachment
	// category := new(Category) //第一种这种形式的category，这种不能返回one的author
	qs := o.QueryTable("attachment")
	err = qs.Filter("route", route).One(&attachment, "Author")
	if err != nil {
		return "", err
	}
	return attachment.Author, nil

	// o := orm.NewOrm()
	// topic := new(Topic)
	// qs := o.QueryTable("topic")
	// err = qs.Filter("id", tidNum).One(topic)
	// if err != nil {
	// 	return nil, err
	// }

	// var topic Topic第二种这种形式的topic
	// err = o.QueryTable("topic").Filter("tnumber", tnumber).One(&topic, "Id")
	// id = topic.Id

	// user = User{Id: userid}第三种这种形式的user
	// o := orm.NewOrm()
	// o.Read(&user, "Id")
	// return user
}

//type Blog struct {
//	Id      int `PK`
//	Title   string
//	Content string
//	Created time.Time
//}

//func GetLink() beedb.Model {
//	db, err := sql.Open("mymysql", "blog/astaxie/123456")
//	if err != nil {
//		panic(err)
//	}
//	orm := beedb.New(db)
//	return orm
//}
//func GetAll() (blogs []Blog) {
//	db := GetLink()
//	db.FindAll(&blogs)
//	return
//}

//func SaveBlog(blog Blog) (bg Blog) {
//	db := GetLink()
//	db.Save(&blog)
//	return bg
//}
//func DelBlog(blog Blog) {
//	db := GetLink()
//	db.Delete(&blog)
//	return
//}
