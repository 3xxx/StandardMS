package models

import (
	// "errors"
	// "strconv"
	"time"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//用户表
type Tag struct {
	Id            int64     `PK`
	Tagname       string    `orm:"unique"`
	Lastlogintime time.Time `orm:"null;type(datetime)" form:"-"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add" `
	Roles         []*Role   `orm:"rel(m2m)"`
}

func init() {
	orm.RegisterModel(new(Tag))
}

func SaveTag(Tag Tag) (uid int64, err error) {
	o := orm.NewOrm()
	//判断是否有重名
	// var spider Spider //下面这个filter放在topic=&Topic{后面用返回one(topic)则查询出错！
	//只有编号和主机都不同才写入。
	err = o.QueryTable("Tag").Filter("Tagname", Tag.Tagname).One(&Tag, "Id")
	// err = o.QueryTable("topic").Filter("categoryid", cid).Filter("tnumber", tnumber).One(&topic, "Id")

	if err == orm.ErrNoRows { //Filter("tnumber", tnumber).One(topic, "Id")==nil则无法建立
		// 没有找到记录
		// spider1 := &Spider{
		// 	Number:   number,
		// 	Name:     name,
		// 	Link:     link,
		// 	TagName: Tagname,
		// 	TagIp:   Tagip,
		// 	Created:  time.Now(),
		// 	Updated:  time.Now(),
		// }
		uid, err = o.Insert(&Tag)
		if err != nil {
			return 0, err //如果文章编号相同，则唯一性检查错误，返回id吗？
		}
	}
	return uid, err

	// 原来的代码orm := orm.NewOrm()
	// // fmt.Println(Tag)
	// uid, err = orm.Insert(&Tag) //_, err = o.Insert(reply)
	// return uid, err
}

func GetTag(tag Tag) ([]*Tag, error) {
	orm := orm.NewOrm()
	Tags := make([]*Tag, 0)
	qs := orm.QueryTable("Tag")
	_, err := qs.Filter("Tagname__contains", tag.Tagname).All(&Tags)
	if err != nil {
		return Tags, err
	}
	return Tags, err
}

/************************************************************/

//get Tag list
func GetTaglist(page int64, page_size int64, sort string) (Tags []orm.Params, count int64) {
	o := orm.NewOrm()
	Tag := new(Tag)
	qs := o.QueryTable(Tag)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&Tags)
	count, _ = qs.Count()
	return Tags, count
}

//添加用户
func AddTag(u *Tag) (int64, error) {

	o := orm.NewOrm()
	Tag := new(Tag)
	Tag.Tagname = u.Tagname
	id, err := o.Insert(Tag)
	return id, err
}

func UpdateTag(Tagid, tagname string) error {
	// if err := checkTag(&u); err != nil {
	// 	return err
	// }
	// id, _ := strconv.ParseInt(Tagid, 10, 64)
	// o := orm.NewOrm()
	// Tag := Tag{Id: id}
	// qs := o.QueryTable("Tag") //不知道主键就用这个过滤操作
	return nil
}

func DelTagById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Tag{Id: Id})
	return status, err
}

func GetTagByTagname(Tagname string) (Tag Tag) {

	o := orm.NewOrm()
	qs := o.QueryTable("Tag") //不知道主键就用这个过滤操作
	//进行编号唯一性检查
	qs.Filter("Tagname", Tagname).One(&Tag)

	// Tag = Tag{Tagname: Tagname} //指定字段查询，这样也行
	// o := orm.NewOrm()
	// o.Read(&Tag,"Tagname")
	return Tag
}

func GetTagByTagId(Tagid int64) (tag Tag) {
	tag = Tag{Id: Tagid}
	o := orm.NewOrm()
	o.Read(&tag) //这里是默认主键查询。=(&Tag,"Id")
	return tag
}

func GetRoleByTagId(Tagid int64) (roles []*Role, count int64) { //*Topic, []*Attachment, error
	roles = make([]*Role, 0)
	o := orm.NewOrm()
	// role := new(Role)
	count, _ = o.QueryTable("role").Filter("Tags__Tag__Id", Tagid).All(&roles)
	return roles, count
	// 通过 post title 查询这个 post 有哪些 tag
	// var tags []*Tag
	// num, err := dORM.QueryTable("tag").Filter("Posts__Post__Title", "Introduce Beego ORM").All(&tags)

}

func GetRoleByTagname(Tagname string) (roles []*Role, count int64) { //*Topic, []*Attachment, error
	roles = make([]*Role, 0)
	o := orm.NewOrm()
	// role := new(Role)
	count, _ = o.QueryTable("role").Filter("Tags__Tag__Tagname", Tagname).All(&roles)
	return roles, count
	// 通过 post title 查询这个 post 有哪些 tag
	// var tags []*Tag
	// num, err := dORM.QueryTable("tag").Filter("Posts__Post__Title", "Introduce Beego ORM").All(&tags)

}
