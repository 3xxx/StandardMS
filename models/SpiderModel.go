package models

import (
	// "errors"
	// "strconv"
	// "fmt"
	// "log"
	"time"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	// . "github.com/beego/admin/src/lib"
)

//成果表//这个是测试用，已经作废
type Spider struct {
	Id       int64
	Number   string
	Name     string
	Link     string
	UserName string
	UserIp   string
	Created  time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"index","auto_now_add;type(datetime)"`
}

//成果表
type Spidertopic struct {
	Id       int64
	Number   string
	Name     string
	Link     string
	UserName string
	UserIp   string
	Created  time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"index","auto_now_add;type(datetime)"`
}

//项目表
type Spidercategory struct {
	Id       int64
	Number   string
	Name     string
	Link     string
	UserName string
	UserIp   string
	Created  time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"index","auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Spider), new(Spidertopic), new(Spidercategory))
}

// func SaveSpider(spider Spider) (sid int64, err error) {
// 	orm := orm.NewOrm()
// 	// fmt.Println(user)
// 	sid, err = orm.Insert(&spider) //_, err = o.Insert(reply)
// 	return sid, err
// }
//这个是测试用，已经作废
func AddSpider(number, name, link, username, userip string) (id int64, err error) {
	o := orm.NewOrm()
	//判断在这个类型number下，是否有重号
	var spider Spider //下面这个filter放在topic=&Topic{后面用返回one(topic)则查询出错！
	//只有编号和主机都不同才写入。
	err = o.QueryTable("spider").Filter("number", number).Filter("userip", userip).One(&spider, "Id")
	// err = o.QueryTable("topic").Filter("categoryid", cid).Filter("tnumber", tnumber).One(&topic, "Id")

	if err == orm.ErrNoRows { //Filter("tnumber", tnumber).One(topic, "Id")==nil则无法建立
		// 没有找到记录
		spider1 := &Spider{
			Number:   number,
			Name:     name,
			Link:     link,
			UserName: username,
			UserIp:   userip,
			Created:  time.Now(),
			Updated:  time.Now(),
		}
		id, err = o.Insert(spider1)
		if err != nil {
			return id, err //如果文章编号相同，则唯一性检查错误，返回id吗？
		}
	}
	return id, err
}

//加到项目表
func AddSpiderCategory(number, name, link, username, userip string) (id int64, err error) {
	o := orm.NewOrm()
	//判断在这个类型number下，是否有重号
	var spidercategory Spidercategory //下面这个filter放在topic=&Topic{后面用返回one(topic)则查询出错！
	err = o.QueryTable("spidercategory").Filter("number", number).Filter("name", name).Filter("userip", userip).One(&spidercategory, "Id")
	if err == orm.ErrNoRows { //Filter("tnumber", tnumber).One(topic, "Id")==nil则无法建立
		// 没有找到记录
		spider1 := &Spidercategory{
			Number:   number,
			Name:     name,
			Link:     link,
			UserName: username,
			UserIp:   userip,
			Created:  time.Now(),
			Updated:  time.Now(),
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
		id, err = o.Insert(spider1)
		if err != nil {
			return id, err //如果文章编号相同，则唯一性检查错误，返回id吗？
		}
	}
	return id, err
}

//加到成果表
func AddSpiderTopic(number, name, link, username, userip string) (id int64, err error) {
	o := orm.NewOrm()
	//判断在这个类型number下，是否有重号
	var spidertopic Spidertopic //下面这个filter放在topic=&Topic{后面用返回one(topic)则查询出错！
	err = o.QueryTable("spidertopic").Filter("number", number).Filter("name", name).Filter("userip", userip).One(&spidertopic, "Id")
	if err == orm.ErrNoRows { //Filter("tnumber", tnumber).One(topic, "Id")==nil则无法建立
		// 没有找到记录
		spider1 := &Spidertopic{
			Number:   number,
			Name:     name,
			Link:     link,
			UserName: username,
			UserIp:   userip,
			Created:  time.Now(),
			Updated:  time.Now(),
		}
		id, err = o.Insert(spider1)
		if err != nil {
			return id, err //如果文章编号相同，则唯一性检查错误，返回id吗？
		}
	}
	return id, err
}

//缺少排序，获取所有成果
func GetSpiderTopic() ([]*Spidertopic, error) {
	o := orm.NewOrm()
	spidertopic := make([]*Spidertopic, 0)
	qs := o.QueryTable("spidertopic")
	var err error
	_, err = qs.OrderBy("-created").All(&spidertopic)
	return spidertopic, err
}

//获取所有项目
func GetSpiderCategory() ([]*Spidercategory, error) {
	o := orm.NewOrm()
	spidercategory := make([]*Spidercategory, 0)
	qs := o.QueryTable("spidercategory")
	var err error
	//这里进行过滤，parentid为空的才显示
	// qs = qs.Filter("ParentId", 0)
	_, err = qs.OrderBy("-created").All(&spidercategory)
	// _, err := qs.All(&cates)
	return spidercategory, err
}

//设计院首页全局搜索
func Searchspidertopics(title string, isDesc bool) ([]*Spidertopic, []*Spidercategory, error) {
	o := orm.NewOrm()
	spidertopics := make([]*Spidertopic, 0)
	spidercategories := make([]*Spidercategory, 0)
	// spidercategories := make([]*Spidercategory, 0)
	qs := o.QueryTable("spidertopic")
	var err error
	if isDesc {
		if len(title) > 0 {
			qs = qs.Filter("Name__contains", title) //这里取回
		}
		_, err = qs.OrderBy("-created").All(&spidertopics)
	} else {
		_, err = qs.Filter("Name__contains", title).OrderBy("-created").All(&spidertopics)
		//o.QueryTable("user").Filter("name", "slene").All(&users)
	}
	qs1 := o.QueryTable("spidercategory")
	if isDesc {
		if len(title) > 0 {
			qs1 = qs1.Filter("Name__contains", title) //这里取回
		}
		_, err = qs1.OrderBy("-created").All(&spidercategories)
	} else {
		_, err = qs1.Filter("Name__contains", title).OrderBy("-created").All(&spidercategories)
		//o.QueryTable("user").Filter("name", "slene").All(&users)
	}

	return spidertopics, spidercategories, err
}
