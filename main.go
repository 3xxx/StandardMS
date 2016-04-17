package main

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	// "github.com/beego/i18n"
	"os"
	"quick/controllers"
	_ "quick/routers"
	// "quick/models"
)

//func Init() {
//	models.RegisterDB()
//}
func main() {
	//开启orm调试模式
	orm.Debug = true
	//自动建表
	orm.RunSyncdb("default", false, true)

	//创建附件目录
	os.Mkdir("attachment", os.ModePerm)
	//作為靜態文件
	//beego.SetStaticPath("/attachment", "attachment")
	//作为单独一个控制器来处理
	// beego.Router("/attachment/:all", &controllers.AttachController{}) //这句：all什么意思
	// beego.AutoRouter(&controllers.AttachController{})

	// 需要先注册一个模板函数
	// beego.AddFuncMap("i18n", i18n.Tr)

	time := beego.AppConfig.String("spec") //"0/time * * * * *"
	// time1 := "0/" + time + " * * * * *"
	time1 := "0 0 */" + time + " * * *"
	tk1 := toolbox.NewTask("tk1", time1, func() error { controllers.TestJsonStartsWithArray(); return nil }) //func() error { fmt.Println("tk1"); return nil }
	toolbox.AddTask("tk1", tk1)
	toolbox.StartTask()
	defer toolbox.StopTask()

	//启动beeego
	beego.Run()
}

// spec 格式是参照 crontab 做的，详细的解释如下所示：

//前6个字段分别表示：
//       秒钟：0-59
//       分钟：0-59
//       小时：1-23
//       日期：1-31
//       月份：1-12
//       星期：0-6（0 表示周日）

//还可以用一些特殊符号：
//       *： 表示任何时刻
//       ,：　表示分割，如第三个字段里里：2,4，表示 2 点和 4 点执行
//　　    -：表示一个段，如第三个字段里里： 1-5，就表示 1 到 5 点
//       /n : 表示每个n的单位执行一次，如果在第三个字段里是*/1, 就表示每隔 1 个小时执行一次命令。也可以写成1-23/1.
/////////////////////////////////////////////////////////
//  0/30 * * * * *                        每 30 秒 执行
//  0 43 21 * * *                         21:43 执行
//  0 15 05 * * * 　　                     05:15 执行
//  0 0 17 * * *                          17:00 执行
//  0 0 17 * * 1                          每周一的 17:00 执行
//  0 0,10 17 * * 0,2,3                   每周日,周二,周三的 17:00和 17:10 执行
//  0 0-10 17 1 * *                       毎月1日从 17:00 到 7:10 毎隔 1 分钟 执行
//  0 0 0 1,15 * 1                        毎月1日和 15 日和 一日的 0:00 执行
//  0 42 4 1 * * 　 　                     毎月1日的 4:42 分 执行
//  0 0 21 * * 1-6　　                     周一到周六 21:00 执行
//  0 0,10,20,30,40,50 * * * *　           每隔 10 分 执行
//  0 */10 * * * * 　　　　　　              每隔 10 分 执行
//  0 * 1 * * *　　　　　　　　               从 1:0 到 1:59 每隔 1 分钟 执行
//  0 0 1 * * *　　　　　　　　               1:00 执行
//  0 0 */1 * * *　　　　　　　               毎时 0 分 每隔 1 小时 执行
//  0 0 * * * *　　　　　　　　               毎时 0 分 每隔 1 小时 执行
//  0 2 8-20/3 * * *　　　　　　             8:02,11:02,14:02,17:02,20:02 执行
//  0 30 5 1,15 * *　　　　　　              1 日 和 15 日的 5:30 执行

// 请问，beego CORS 跨域如何实现？
// beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
// 		AllowOrigins: []string{"*"},
// 		AllowMethods: []string{"*"},
// 		ExposeHeaders: []string{"Content-Length"},
// 		AllowCredentials: true,
// 	}))

// 请问一下 map[string]string{"a":"apple", "b":"banana"}  这个 value 是什么类型
// package main

// import (
// 	"encoding/json"
// 	"flag"
// 	"fmt"
// )

// type JsonHolder struct {
// 	Data map[string]string `json:"data"`
// }

// func main() {
// 	var m string
// 	flag.StringVar(&m, "m", "", "map data")
// 	flag.Parse()

// 	jsonData := fmt.Sprintf("{\"data\":%s}", m)
// 	jsonHolder := JsonHolder{}
// 	json.Unmarshal([]byte(jsonData), &jsonHolder)

// 	mp := jsonHolder.Data

// 	fmt.Printf("%#v\n", mp)

// }
