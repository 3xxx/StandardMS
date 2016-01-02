//这个作废了。
package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"quick/models"
)

type IndexController struct {
	beego.Controller
}

func (index *IndexController) Get() {
	sess := index.StartSession()
	username := sess.Get("username")
	fmt.Println(username)
	if username == nil || username == "" {
		index.TplNames = "login.tpl"
	} else {
		index.TplNames = "success.tpl"
	}

}

func (index *IndexController) Post() {
	sess := index.StartSession()
	var user models.User
	inputs := index.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("username")

	Pwd1 := inputs.Get("pwd")

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	fmt.Print(cipherStr)
	fmt.Print("\n")
	fmt.Print(hex.EncodeToString(cipherStr))

	user.Password = hex.EncodeToString(cipherStr)
	err := models.ValidateUser(user)
	if err == nil {
		sess.Set("username", user.Username)
		fmt.Println("username:", sess.Get("username"))
		index.TplNames = "success.tpl"
	} else {
		fmt.Println(err)
		index.TplNames = "error.tpl"
	}
}
