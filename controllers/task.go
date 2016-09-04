package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"hydrocms/models"
	"strconv"
	"strings"
	"time"
)

type TaskController struct {
	beego.Controller
}

func (c *TaskController) Get() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	// c.Data["IsCategoryb"] = true
	c.Data["IsTask"] = true
	uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	c.Data["Uname"] = uname
	c.TplName = "todo.html"
	tasks, tasks1, tasks2, err := models.GetAllTasks() //这里传入空字符串
	if err != nil {
		beego.Error(err.Error)
	} else {
		c.Data["Tasks"] = tasks
		c.Data["Tasks1"] = tasks1
		c.Data["Tasks2"] = tasks2
	}
	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/test.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "Viewtodo")
	logs.Close()
	// c.Render()
}

func (c *TaskController) ShowDetails() {
	// c.TplName = "todo.html"
	tid := c.Input().Get("id")
	beego.Info(tid)
	tasks, err := models.GetDetails(tid) //这里传入空字符串
	if err != nil {
		beego.Error(err.Error)
	} else {
		// b, err := json.Marshal(tasks)
		// if err == nil {
		// c.Ctx.WriteString(string(b))
		c.Data["json"] = tasks //string(b)
		c.ServeJSON()
		// }
	}
}

func (c *TaskController) AddTask() {
	// if !checkAccount(c.Ctx) { //这里应该不需要
	// 	c.Redirect("/login", 302)
	// 	return
	// }
	//解析表单     表示时间的变量和字段，应为time.Time类型
	type Duration int64
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	//seconds := 10
	//fmt.Print(time.Duration(seconds)*time.Second) // prints 10s
	hours := 8
	//	time.Duration(hours) * time.Hour
	// t1 := t.Add(time.Duration(hours) * time.Hour)
	// datestring = t1.Format(layout) //{{dateformat .Created "2006-01-02 15:04:05"}}
	// return

	var err error
	tid := c.Input().Get("tid")
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	daterange := c.Input().Get("datefilter")
	array := strings.Split(daterange, " - ")
	// for _, v := range array {
	starttime1 := array[0]
	// beego.Info(array[0])
	endtime1 := array[1]
	// beego.Info(array[1])
	// }
	// starttime1 := c.Input().Get("starttime")
	// endtime1 := c.Input().Get("endtime")
	const lll = "2006-01-02" //"2006-01-02 15:04:05" //12-19-2015 22:40:24
	starttime, _ := time.Parse(lll, starttime1)
	endtime, _ := time.Parse(lll, endtime1)
	t1 := starttime.Add(-time.Duration(hours) * time.Hour)
	t2 := endtime.Add(-time.Duration(hours) * time.Hour)
	//12-19-2015 22:40:24
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// beego.Error(err)
	// }
	// uname := ck.Value
	if len(tid) == 0 {
		err = models.AddTask(title, content, t1, t2)
		// beego.Info(attachment)
		// } else {
		// 	err = models.UpdateTask(tid, title, content)
	}
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/todo", 302)
	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/test.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "addtask")
	logs.Close()
}

func (c *TaskController) Delete() {

	//2.如果登录或ip在允许范围内，进行访问权限检查
	uname, role, _ := checkRolewrite(c.Ctx) //login里的
	rolename, _ := strconv.Atoi(role)
	c.Data["Uname"] = uname
	if rolename > 2 { //
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	err := models.DeleteTask(c.Input().Get("tid")) //(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/todo", 302)
	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/test.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "deletetask")
	logs.Close()
}
func (c *TaskController) Update() {
	err := models.UpdateTaskstate(c.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/todo", 302) //这里增加topic
}
func (c *TaskController) Update1() {
	err := models.UpdateTaskstate1(c.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/todo", 302) //这里增加topic
}

// Example:
//
//   req: GET /task/
//   res: 200 {"Tasks": [
//          {"ID": 1, "Title": "Learn Go", "Done": false},
//          {"ID": 2, "Title": "Buy bread", "Done": true}
//        ]}
func (c *TaskController) ListTasks() {
	res := struct{ Tasks []*models.Task }{models.DefaultTaskList.All()}
	c.Data["json"] = res
	c.ServeJSON()
}

// Examples:
//
//   req: POST /task/ {"Title": ""}
//   res: 400 empty title
//
//   req: POST /task/ {"Title": "Buy bread"}
//   res: 200
func (c *TaskController) NewTask() {
	req := struct{ Title string }{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("empty title"))
		return
	}
	// t, err := models.NewTask(req.Title)
	// if err != nil {
	// 	c.Ctx.Output.SetStatus(400)
	// 	c.Ctx.Output.Body([]byte(err.Error()))
	// 	return
	// }
	// models.DefaultTaskList.Save(t)
}

// Examples:
//
//   req: GET /task/1
//   res: 200 {"ID": 1, "Title": "Buy bread", "Done": true}
//
//   req: GET /task/42
//   res: 404 task not found
func (c *TaskController) GetTask() {
	id := c.Ctx.Input.Param(":id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	t, ok := models.DefaultTaskList.Find(intid)
	beego.Info("Found", ok)
	if !ok {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte("task not found"))
		return
	}
	c.Data["json"] = t
	c.ServeJSON()
}

// Example:
//
//   req: PUT /task/1 {"ID": 1, "Title": "Learn Go", "Done": true}
//   res: 200
//
//   req: PUT /task/2 {"ID": 2, "Title": "Learn Go", "Done": true}
//   res: 400 inconsistent task IDs
func (c *TaskController) UpdateTask() {
	id := c.Ctx.Input.Param(":id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	var t models.Task
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &t); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	if t.ID != intid {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("inconsistent task IDs"))
		return
	}
	if _, ok := models.DefaultTaskList.Find(intid); !ok {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("task not found"))
		return
	}
	models.DefaultTaskList.Save(&t)
}
