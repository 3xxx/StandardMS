package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"time"
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

func init() {
	orm.RegisterModel(new(Catalog)) //, new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db", 10)
}

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
