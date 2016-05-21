package models

import (
	//"database/sql"
	//"github.com/astaxie/beedb"
	//_ "github.com/ziutek/mymysql/godrv"
	//"time"
	// "fmt"
	// "os"
	// "path"
	// "github.com/astaxie/beego"
	"strconv"
	// "strings"
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

type Topic struct {
	Id                int64
	Uid               int64
	Title             string
	Tnumber           string //`orm:"unique"`
	Category          string
	CategoryId        int64
	Content           string `orm:"sie(5000)"`
	Attachment        string
	Created           time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated           time.Time `orm:"index","auto_now;type(datetime)"`
	Views             int64     `orm:"index"`
	Author            string
	ReplyTime         time.Time `orm:"index"`
	ReplyCount        int64
	ReplyLastUserName string
	// Attachments     []*Attachment `orm:"reverse(many)"` // fk 的反向关系
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

func init() {
	orm.RegisterModel(new(Topic), new(Attachment)) //, new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db", 10)
}

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
	if len(cate) > 0 {
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
