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

type Legislation struct {
	Id       int64
	Number   string //`orm:"unique"`
	Title    string
	Uid      int64
	Category string
	Content  string `orm:"sie(5000)"`
	Route    string
	// AttachmentId int64
	// Attachments     []*Attachment `orm:"reverse(many)"` // fk 的反向关系
	Created time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated time.Time `orm:"index","auto_now;type(datetime)"`
	Views   int64     `orm:"index"`
}

//附件,attachment 和 Legislation 是 ManyToOne 关系，也就是 ForeignKey 为 Legislation
// type Attachment struct {
// 	Id            int64
// 	Uid           int64
// 	FileName      string
// 	FileSize      string
// 	Downloads     int64
// 	DiskDirectory string
// 	Route         string
// 	Content       string    `orm:"sie(200)"`
// 	LegislationId    int64     //*Legislation    `orm:"rel(fk)"`
// 	Created       time.Time `orm:"index","auto_now_add;type(datetime)"`
// 	Updated       time.Time `orm:"index","auto_now;type(datetime)"`
// 	Views         int64     `orm:"index"`
// 	Author        string
// }

func init() {
	orm.RegisterModel(new(Legislation)) //, new(Attachment), new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db", 10)
}

//标准存入数据库
func SaveLegislation(legislation Legislation) (sid int64, err error) {
	o := orm.NewOrm()
	//判断是否有重名
	// var spider Spider //下面这个filter放在topic=&Topic{后面用返回one(topic)则查询出错！
	//只有编号和主机都不同才写入。
	err = o.QueryTable("legislation").Filter("number", legislation.Number).Filter("title", legislation.Title).One(&legislation, "Id")
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
		sid, err = o.Insert(&legislation)
		if err != nil {
			return 0, err //如果文章编号相同，则唯一性检查错误，返回id吗？
		}
	} else { //进行更新
		// _, err = o.Update(cate)
		sid, err = o.Update(&legislation)
	}
	return sid, err
	// 原来的代码orm := orm.NewOrm()
	// // fmt.Println(user)
	// uid, err = orm.Insert(&user) //_, err = o.Insert(reply)
	// return uid, err
}

//有效版本库存入数据库

//由名字模糊搜索
func SearchLegislationsName(name string, isDesc bool) ([]*Legislation, error) {
	o := orm.NewOrm()
	Legislations := make([]*Legislation, 0)
	qs := o.QueryTable("Legislation")
	var err error
	if isDesc {
		if len(name) > 0 {
			qs = qs.Filter("Title__contains", name) //这里取回
		}
		_, err = qs.OrderBy("-created").All(&Legislations)
	} else {
		_, err = qs.Filter("Title__contains", name).OrderBy("-created").All(&Legislations)
		//o.QueryTable("user").Filter("name", "slene").All(&users)
	}
	return Legislations, err
}

//由编号模糊搜索
func SearchLegislationsNumber(number string, isDesc bool) ([]*Legislation, error) {
	o := orm.NewOrm()
	Legislations := make([]*Legislation, 0)
	qs := o.QueryTable("Legislation")
	var err error
	if isDesc {
		if len(number) > 0 {
			qs = qs.Filter("Number__contains", number) //这里取回
		}
		_, err = qs.OrderBy("-created").All(&Legislations)
	} else {
		_, err = qs.Filter("Number__contains", number).OrderBy("-created").All(&Legislations)
		//o.QueryTable("user").Filter("name", "slene").All(&users)
	}
	return Legislations, err
}

func GetAllLegislations() ([]*Legislation, error) {
	o := orm.NewOrm()
	legislations := make([]*Legislation, 0)
	qs := o.QueryTable("legislation")
	var err error
	//这里进行过滤，parentid为空的才显示
	// qs = qs.Filter("ParentId", 0)
	_, err = qs.OrderBy("-created").All(&legislations)
	// _, err := qs.All(&cates)
	return legislations, err
}
