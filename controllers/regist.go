package controllers

import (
	"crypto/md5"
	"encoding/hex"
	// "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"quick/models"
	"time"
)

type RegistController struct {
	beego.Controller
}

func (this *RegistController) Get() {
	// this.Data["IsLogin"] = true
	this.TplName = "regist.tpl"
}

func (this *RegistController) RegistErr() {
	// this.Data["IsLogin"] = true
	this.TplName = "registerr.tpl"
}

func (this *RegistController) CheckUname() {
	var user models.User //这里修改
	inputs := this.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("uname")
	err := models.CheckUname(user) //这里修改
	if err == nil {
		this.Ctx.WriteString("false")
		// return false
	} else {
		this.Ctx.WriteString("true")
		// return true
	}
	// return
}

func (this *RegistController) Post() {
	var user models.User //这里修改
	inputs := this.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("uname")
	Pwd1 := inputs.Get("pwd")

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	// fmt.Print(cipherStr)
	// fmt.Print("\n")
	// fmt.Print(hex.EncodeToString(cipherStr))

	user.Password = hex.EncodeToString(cipherStr)
	user.Lastlogintime = time.Now()
	uid, err := models.SaveUser(user) //这里修改

	_, err = models.AddRoleUser(4, uid)
	if err == nil {
		this.TplName = "success.tpl"
	} else {
		fmt.Println(err)
		this.TplName = "registerr.tpl"
	}
}

func (this *RegistController) GetUname() {
	var user models.User //这里修改[]*models.User(uname string)
	inputs := this.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("uname")
	// beego.Info(user.Username)
	uname1, err := models.GetUname(user) //这里修改
	//转换成json数据？
	// beego.Info(uname1[0].Username)
	// b, err := json.Marshal(uname1)
	if err == nil {
		// this.Ctx.WriteString(string(b))
		this.Data["json"] = uname1 //string(b)
		this.ServeJSON()
	}
	// 	this.Ctx.WriteString(uname1[1].Username)
	// 	// return uname1[0].Username
	// }
	// return uname1[0].Username
}
