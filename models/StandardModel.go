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
	// "strconv"
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

type Standard struct {
	Id         int64
	Number     string //`orm:"unique"`
	Title      string
	Uid        int64
	CategoryId int64
	Content    string `orm:"sie(5000)"`
	Route      string
	// AttachmentId int64
	// Attachments     []*Attachment `orm:"reverse(many)"` // fk 的反向关系
	Created time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated time.Time `orm:"index","auto_now;type(datetime)"`
	Views   int64     `orm:"index"`
}

type Library struct {
	Id       int64
	Number   string //`orm:"unique"`
	Title    string
	LiNumber string //完整编号
	Category string
	Content  string    `orm:"sie(5000)"`
	Created  time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"index","auto_now;type(datetime)"`
}

//附件,attachment 和 Standard 是 ManyToOne 关系，也就是 ForeignKey 为 Standard
// type Attachment struct {
// 	Id            int64
// 	Uid           int64
// 	FileName      string
// 	FileSize      string
// 	Downloads     int64
// 	DiskDirectory string
// 	Route         string
// 	Content       string    `orm:"sie(200)"`
// 	StandardId    int64     //*Standard    `orm:"rel(fk)"`
// 	Created       time.Time `orm:"index","auto_now_add;type(datetime)"`
// 	Updated       time.Time `orm:"index","auto_now;type(datetime)"`
// 	Views         int64     `orm:"index"`
// 	Author        string
// }

func init() {
	orm.RegisterModel(new(Standard), new(Library)) //, new(Attachment), new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db", 10)
}

//标准存入数据库
func SaveStandard(standard Standard) (sid int64, err error) {
	o := orm.NewOrm()
	//判断是否有重名
	// var spider Spider //下面这个filter放在topic=&Topic{后面用返回one(topic)则查询出错！
	//只有编号和主机都不同才写入。
	err = o.QueryTable("standard").Filter("title", standard.Title).One(&standard, "Id")
	// err = o.QueryTable("topic").Filter("categoryid", cid).Filter("tnumber", tnumber).One(&topic, "Id")
	if err == orm.ErrNoRows { //Filter("tnumber", tnumber).One(topic, "Id")==nil则无法建立
		// 没有找到记录
		// spider1 := &Spider{
		// 	Number:   number,
		// 	Name:     name,
		// 	Link:     link,
		// 	UserName: username,
		// 	UserIp:   userip,
		// 	Created:  time.Now(),
		// 	Updated:  time.Now(),
		// }
		sid, err = o.Insert(&standard)
		if err != nil {
			return 0, err //如果文章编号相同，则唯一性检查错误，返回id吗？
		}
	}
	return sid, err
	// 原来的代码orm := orm.NewOrm()
	// // fmt.Println(user)
	// uid, err = orm.Insert(&user) //_, err = o.Insert(reply)
	// return uid, err
}

//有效版本库存入数据库
func SaveLibrary(library Library) (lid int64, err error) {
	o := orm.NewOrm()
	//判断是否有重名
	// var spider Spider //下面这个filter放在topic=&Topic{后面用返回one(topic)则查询出错！
	//只有编号和主机都不同才写入。
	err = o.QueryTable("library").Filter("title", library.Title).One(&library, "Id")
	// err = o.QueryTable("topic").Filter("categoryid", cid).Filter("tnumber", tnumber).One(&topic, "Id")
	if err == orm.ErrNoRows { //Filter("tnumber", tnumber).One(topic, "Id")==nil则无法建立
		// 没有找到记录
		// spider1 := &Spider{
		// 	Number:   number,
		// 	Name:     name,
		// 	Link:     link,
		// 	UserName: username,
		// 	UserIp:   userip,
		// 	Created:  time.Now(),
		// 	Updated:  time.Now(),
		// }
		lid, err = o.Insert(&library)
		if err != nil {
			return 0, err //如果文章编号相同，则唯一性检查错误，返回id吗？
		}
	}
	return lid, err
	// 原来的代码orm := orm.NewOrm()
	// // fmt.Println(user)
	// uid, err = orm.Insert(&user) //_, err = o.Insert(reply)
	// return uid, err
}

//由名字模糊搜索
func SearchStandardsName(name string, isDesc bool) ([]*Standard, error) {
	o := orm.NewOrm()
	Standards := make([]*Standard, 0)
	qs := o.QueryTable("Standard")
	var err error
	if isDesc {
		if len(name) > 0 {
			qs = qs.Filter("Title__contains", name) //这里取回
		}
		_, err = qs.OrderBy("-created").All(&Standards)
	} else {
		_, err = qs.Filter("Title__contains", name).OrderBy("-created").All(&Standards)
		//o.QueryTable("user").Filter("name", "slene").All(&users)
	}
	return Standards, err
}

//由编号模糊搜索
func SearchStandardsNumber(number string, isDesc bool) ([]*Standard, error) {
	o := orm.NewOrm()
	Standards := make([]*Standard, 0)
	qs := o.QueryTable("Standard")
	var err error
	if isDesc {
		if len(number) > 0 {
			qs = qs.Filter("Number__contains", number) //这里取回
		}
		_, err = qs.OrderBy("-created").All(&Standards)
	} else {
		_, err = qs.Filter("Number__contains", number).OrderBy("-created").All(&Standards)
		//o.QueryTable("user").Filter("name", "slene").All(&users)
	}
	return Standards, err
}

//由分类和编号搜索有效版本库
func SearchLiabraryNumber(Category, Number string) (*Library, error) {
	o := orm.NewOrm()
	library := new(Library)
	qs := o.QueryTable("library")
	err := qs.Filter("category", Category).Filter("number", Number).One(library)
	if err != nil {
		return nil, err
	}
	return library, err
}
