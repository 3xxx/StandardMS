//category的标签，比如供水工程，堤防工程，海堤工程
//label和category是多对一的关系
// Post 和 User 是 ManyToOne 关系，也就是 ForeignKey 为 User

// type Post struct {
//     Id    int
//     Title string
//     User  *User  `orm:"rel(fk)"`
//     Tags  []*Tag `orm:"rel(m2m)"`
// }
// var posts []*Post
// num, err := o.QueryTable("post").Filter("User", 1).RelatedSel().All(&posts)
// if err == nil {
//     fmt.Printf("%d posts read\n", num)
//     for _, post := range posts {
//         fmt.Printf("Id: %d, UserName: %d, Title: %s\n", post.Id, post.User.UserName, post.Title)
//     }
// }
// 根据 Post.Title 查询对应的 User：

// RegisterModel 时，ORM也会自动建立 User 中 Post 的反向关系，所以可以直接进行查询

// var user User
// err := o.QueryTable("user").Filter("Post__Title", "The Title").Limit(1).One(&user)
// if err == nil {
//     fmt.Printf(user)
//
package models

import (
	// "errors"
	// "strconv"
	// "fmt"
	// "log"
	// "time"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
	// "github.com/astaxie/beego/validation"
	// . "github.com/beego/admin/src/lib"
)

//用户表
// func AddRoleUser(roleid int64, userid int64) (int64, error) {
// 	o := orm.NewOrm()
// 	role := Role{Id: roleid}
// 	user := User{Id: userid}
// 	m2m := o.QueryM2M(&user, "Roles")
// 	num, err := m2m.Add(&role)
// 	return num, err
// }
//标签表
type Label struct {
	Id       int
	Title    string
	Category *Category `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Label)) //建立表后，label表中会出现categoryid字段，但这个字段不能直接赋值
}

//一对一和多对一关系不建立中间表吧？？也无操作方法
//果然，只是建立了一个categoryid在表label中
func AddLabel(labeltitle string, categoryid int64) (int64, error) {
	o := orm.NewOrm()
	category := &Category{Id: categoryid}
	label := &Label{
		Title:    labeltitle,
		Category: category, //虽然label表中是categoryid，但是这里不能用categoryid
	}
	id, err := o.Insert(label)
	if err != nil {
		return id, err
	}
	return id, err
}

//更新label
func UpdateLabel(labeltitle string, categoryid int64) error {
	o := orm.NewOrm()
	var err error
	// 依据当前查询条件，进行批量删除操作
	_, err = o.QueryTable("label").Filter("Category", categoryid).RelatedSel().Delete()

	category := &Category{Id: categoryid}
	labelarray := strings.Split(labeltitle, ",")
	for _, labeltitle1 := range labelarray {
		label := &Label{
			Title:    labeltitle1,
			Category: category, //虽然label表中是categoryid，但是这里不能用categoryid
		}
		_, err := o.Insert(label)
		if err != nil {
			return err
		}
	}
	return err
}
