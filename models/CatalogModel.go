package models

import (
	//"database/sql"
	//"github.com/astaxie/beedb"
	//_ "github.com/ziutek/mymysql/godrv"
	//"time"
	// "fmt"
	// "os"
	// "path"
	"strconv"
	// "strings"
	"time"
	//"github.com/Unknwon/com
	// "errors"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	_ "github.com/mattn/go-sqlite3"
)

type Catalog struct {
	Id          int64
	Name        string
	ParentId    int64
	Tnumber     string
	Drawn       string
	Designd     string
	Checked     string
	Emamined    string
	Verified    string
	Approved    string
	Data        string
	DesignStage string
	Section     string
	Projec      string
	Created     time.Time `orm:"index"`
	Updated     time.Time `orm:"index"`
	Views       int64     `form:"-",orm:"index"`
	Author      string
	Exist       bool
	TopicId     int64
}

// func init() {
// 	orm.RegisterModel(new(Catalog)) //, new(Article)
// 	// orm.RegisterDriver("sqlite", orm.DR_Sqlite)
// 	// orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db", 10)
// }

func SaveCatalog(catalog Catalog) (cid int64, err error) {
	orm := orm.NewOrm()
	// fmt.Println(user)
	cid, err = orm.Insert(&catalog) //_, err = o.Insert(reply)
	return cid, err
}

func GetAllCatalogs(cid string) (catalogs []*Catalog, err error) {
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return nil, err
	}
	catalogs = make([]*Catalog, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("catalog")
	_, err = qs.Filter("parentid", cidNum).All(&catalogs)

	//读出这个分类id下所有成果
	var topics []Topic
	qs1 := o.QueryTable("topic")
	_, err = qs1.Filter("categoryid", cidNum).All(&topics)

	// cate := &Category{Id: cidNum}
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
	// if o.Read(cate) == nil {
	// 	if route == "" { //如果没有更新图片，则不更新图片地址
	// 		cate.Title = name
	// 		cate.Number = number
	// 		cate.Content = content
	// 		cate.Author = uname
	// 		// cate.Route = route
	// 		// cate.Created: time.Now(),
	// 		cate.Updated = time.Now()
	// 		_, err = o.Update(cate)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	} else {
	// 		cate.Title = name
	// 		cate.Number = number
	// 		cate.Content = content
	// 		cate.Author = uname
	// 		cate.Route = route
	// 		// cate.Created: time.Now(),
	// 		cate.Updated = time.Now()
	// 		_, err = o.Update(cate)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	//进行图纸目录更新——一次搞定，欣喜。
	for _, k := range topics {
		for _, v := range catalogs {
			if v.Name == k.Title && v.Tnumber == k.Tnumber {
				v.Updated = time.Now()
				v.Exist = true
				v.TopicId = k.Id
				_, err = o.Update(v)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return catalogs, err
}
