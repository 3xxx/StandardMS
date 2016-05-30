package models

import (
	"errors"
	"strconv"
	// "fmt"
	"log"
	"time"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	. "github.com/beego/admin/src/lib"
)

//用户表
type User struct {
	Id            int64  `PK`
	Username      string `orm:"unique"`
	Password      string
	Repassword    string    `orm:"-" form:"Repassword" valid:"Required"`
	Nickname      string    //`orm:"unique;size(32)" form:"Nickname" valid:"Required;MaxSize(20);MinSize(2)"`
	Email         string    `orm:"size(32)" form:"Email" valid:"Email"`
	Department    string    //分院
	Secoffice     string    //科室
	Remark        string    `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
	Status        int       `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
	Lastlogintime time.Time `orm:"null;type(datetime)" form:"-"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add" `
	Roles         []*Role   `orm:"rel(m2m)"` //用户和权限是多对一的关系。
}

// Id            int64
// Username      string    `orm:"unique;size(32)" form:"Username"  valid:"Required;MaxSize(20);MinSize(6)"`
// Password      string    `orm:"size(32)" form:"Password" valid:"Required;MaxSize(20);MinSize(6)"`

func init() {
	orm.RegisterModel(new(User))
}

func SaveUser(user User) (uid int64, err error) {
	o := orm.NewOrm()
	//判断是否有重名
	// var spider Spider //下面这个filter放在topic=&Topic{后面用返回one(topic)则查询出错！
	//只有编号和主机都不同才写入。
	err = o.QueryTable("user").Filter("username", user.Username).One(&user, "Id")
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
		uid, err = o.Insert(&user)
		if err != nil {
			return 0, err //如果文章编号相同，则唯一性检查错误，返回id吗？
		}
	}
	return uid, err

	// 原来的代码orm := orm.NewOrm()
	// // fmt.Println(user)
	// uid, err = orm.Insert(&user) //_, err = o.Insert(reply)
	// return uid, err
}

func ValidateUser(user User) error {
	orm := orm.NewOrm()
	var u User

	// user = new(User)
	qs := orm.QueryTable("user")
	err := qs.Filter("username", user.Username).Filter("password", user.Password).One(&u)
	if err != nil {
		return err
	}

	// orm.Where("username=? and pwd=?", user.Username, user.Pwd).Find(&u)
	if u.Username == "" {
		return errors.New("用户名或密码错误！")
	}
	return nil
}

func CheckUname(user User) error {
	orm := orm.NewOrm()
	var u User
	// user = new(User)
	qs := orm.QueryTable("user")
	err := qs.Filter("username", user.Username).One(&u)
	if err != nil {
		return err
	}
	// orm.Where("username=? and pwd=?", user.Username, user.Pwd).Find(&u)
	// if u.Username == "" {
	// 	return errors.New("用户名或密码错误！")
	// }
	return nil
}

func GetUname(user User) ([]*User, error) {
	orm := orm.NewOrm()
	users := make([]*User, 0)
	qs := orm.QueryTable("user")
	_, err := qs.Filter("Username__contains", user.Username).All(&users)
	if err != nil {
		return users, err
	}
	return users, err
}

// func SearchTopics(tuming string, isDesc bool) ([]*Topic, error) {
// 	o := orm.NewOrm()
// 	topics := make([]*Topic, 0)
// 	qs := o.QueryTable("topic")
// 	var err error
// 	if isDesc {
// 		if len(tuming) > 0 {
// 			qs = qs.Filter("Title__contains", tuming) //这里取回
// 		}
// 		_, err = qs.OrderBy("-created").All(&topics)
// 	} else {
// 		_, err = qs.Filter("Title__contains", tuming).OrderBy("-created").All(&topics)
// 		//o.QueryTable("user").Filter("name", "slene").All(&users)
// 	}
// 	return topics, err
// }

// func (u *User) TableName() string {
// 	return beego.AppConfig.String("rbac_user_table")
// }

func (u *User) Valid(v *validation.Validation) {
	if u.Password != u.Repassword {
		v.SetError("Repassword", "两次输入的密码不一样")
	}
}

//验证用户信息
func checkUser(u *User) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&u)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

/************************************************************/

//get user list——没有取到权限，为何？
func Getuserlist(page int64, page_size int64, sort string) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&users)
	//顺带把权限也取出来，不具备这个功能。只能一个个查权限
	//用户和权限是多对一的关系。
	count, _ = qs.Count()
	return users, count
}

//添加用户
func AddUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := new(User)
	user.Username = u.Username
	user.Password = Strtomd5(u.Password)
	user.Nickname = u.Nickname
	user.Email = u.Email
	user.Remark = u.Remark
	user.Status = u.Status

	id, err := o.Insert(user)
	return id, err
}

//更新用户
// func UpdateUser(u *User) (int64, error) {
// 	if err := checkUser(u); err != nil {
// 		return 0, err
// 	}
// 	o := orm.NewOrm()
// 	user := make(orm.Params)
// 	if len(u.Username) > 0 {
// 		user["Username"] = u.Username
// 	}
// 	if len(u.Nickname) > 0 {
// 		user["Nickname"] = u.Nickname
// 	}
// 	if len(u.Email) > 0 {
// 		user["Email"] = u.Email
// 	}
// 	if len(u.Remark) > 0 {
// 		user["Remark"] = u.Remark
// 	}
// 	if len(u.Password) > 0 {
// 		user["Password"] = Strtomd5(u.Password)
// 	}
// 	if u.Status != 0 {
// 		user["Status"] = u.Status
// 	}
// 	if len(user) == 0 {
// 		return 0, errors.New("update field is empty")
// 	}
// 	var table User
// 	num, err := o.QueryTable(table).Filter("Id", u.Id).Update(user)
// 	return num, err
// }
func UpdateUser(userid, nickname, email, password string) error {
	// if err := checkUser(&u); err != nil {
	// 	return err
	// }
	id, err := strconv.ParseInt(userid, 10, 64)
	o := orm.NewOrm()
	// qs := o.QueryTable("user") //不知道主键就用这个过滤操作
	user := User{Id: id}
	// err := qs.Filter("Username", u.Username).One(&u)
	// if err == nil {
	// 	return err
	// }
	// user := User{Username: u.Username}
	// if err := o.Read(&user); err == nil {
	user.Nickname = nickname
	user.Email = email
	if password != "" {
		user.Password = password
		// user1 := make(orm.Params)
		// var table User
		_, err = o.Update(&user, "password", "nickname", "email")
		if err != nil {
			return err
		}
	} else {
		_, err = o.Update(&user, "nickname", "email")
		if err != nil {
			return err
		}
	}
	// } else {
	// return err
	// }
	return nil
}

//更新用户登陆时间
func UpdateUserlastlogintime(username string) error {
	o := orm.NewOrm()
	user := make(orm.Params)
	if len(username) > 0 {
		user["Lastlogintime"] = time.Now()
	}

	if len(username) == 0 {
		return errors.New("update field is empty")
	}
	var table User
	_, err := o.QueryTable(table).Filter("Username", username).Update(user)
	return err
}

func DelUserById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&User{Id: Id})
	return status, err
}

func GetUserByUsername(username string) (user User) {

	o := orm.NewOrm()
	qs := o.QueryTable("user") //不知道主键就用这个过滤操作
	//进行编号唯一性检查
	qs.Filter("username", username).One(&user)

	// user = User{Username: username} //指定字段查询，这样也行
	// o := orm.NewOrm()
	// o.Read(&user,"Username")
	return user
}

func GetUserByUserId(userid int64) (user User) {
	user = User{Id: userid}
	o := orm.NewOrm()
	o.Read(&user) //这里是默认主键查询。=(&user,"Id")
	return user
}

// func GetAllReplies(tid string) (replies []*Comment, err error) {
// 	tidNum, err := strconv.ParseInt(tid, 10, 64)
// 	if err != nil {
// 		return nil, err
// 	}
// 	replies = make([]*Comment, 0)

// 	o := orm.NewOrm()
// 	qs := o.QueryTable("comment")
// 	_, err = qs.Filter("tid", tidNum).All(&replies)
// 	return replies, err

// }
func GetRoleByUserId(userid int64) (roles []*Role, count int64) { //*Topic, []*Attachment, error
	roles = make([]*Role, 0)
	o := orm.NewOrm()
	// role := new(Role)
	count, _ = o.QueryTable("role").Filter("Users__User__Id", userid).All(&roles)
	return roles, count
	// 通过 post title 查询这个 post 有哪些 tag
	// var tags []*Tag
	// num, err := dORM.QueryTable("tag").Filter("Posts__Post__Title", "Introduce Beego ORM").All(&tags)

}

func GetRoleByUsername(username string) (roles []*Role, count int64, err error) { //*Topic, []*Attachment, error
	roles = make([]*Role, 0)
	o := orm.NewOrm()
	// role := new(Role)
	count, err = o.QueryTable("role").Filter("Users__User__Username", username).All(&roles)
	return roles, count, err
	// 通过 post title 查询这个 post 有哪些 tag
	// var tags []*Tag
	// num, err := dORM.QueryTable("tag").Filter("Posts__Post__Title", "Introduce Beego ORM").All(&tags)

}
