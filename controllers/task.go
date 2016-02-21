package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"quick/models"
	"strconv"
	"strings"
	"time"
)

type TaskController struct {
	beego.Controller
}

func (this *TaskController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	// c.Data["IsCategoryb"] = true
	this.Data["IsTask"] = true
	//3.取得客户端用户名
	sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		this.Data["Uname"] = v.(string)
	}
	this.TplName = "todo.html"
	tasks, tasks1, tasks2, err := models.GetAllTasks() //这里传入空字符串
	if err != nil {
		beego.Error(err.Error)
	} else {
		this.Data["Tasks"] = tasks
		this.Data["Tasks1"] = tasks1
		this.Data["Tasks2"] = tasks2
	}
	// this.Render()
}

func (this *TaskController) ShowDetails() {
	// this.TplName = "todo.html"
	tid := this.Input().Get("id")
	beego.Info(tid)
	tasks, err := models.GetDetails(tid) //这里传入空字符串
	if err != nil {
		beego.Error(err.Error)
	} else {
		// b, err := json.Marshal(tasks)
		// if err == nil {
		// this.Ctx.WriteString(string(b))
		this.Data["json"] = tasks //string(b)
		this.ServeJSON()
		// }
	}
}

func (this *TaskController) AddTask() {
	// if !checkAccount(this.Ctx) { //这里应该不需要
	// 	this.Redirect("/login", 302)
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
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	daterange := this.Input().Get("datefilter")
	array := strings.Split(daterange, " - ")
	// for _, v := range array {
	starttime1 := array[0]
	beego.Info(array[0])
	endtime1 := array[1]
	beego.Info(array[1])
	// }
	// starttime1 := this.Input().Get("starttime")
	// endtime1 := this.Input().Get("endtime")
	const lll = "2006-01-02" //"2006-01-02 15:04:05" //12-19-2015 22:40:24
	starttime, _ := time.Parse(lll, starttime1)
	endtime, _ := time.Parse(lll, endtime1)
	t1 := starttime.Add(-time.Duration(hours) * time.Hour)
	t2 := endtime.Add(-time.Duration(hours) * time.Hour)
	//12-19-2015 22:40:24
	// ck, err := this.Ctx.Request.Cookie("uname")
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
	this.Redirect("/todo", 302)
}

func (this *TaskController) Delete() {
	err := models.DeleteTask(this.Input().Get("tid")) //(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/todo", 302)
}
func (this *TaskController) Update() {
	err := models.UpdateTaskstate(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/todo", 302) //这里增加topic
}
func (this *TaskController) Update1() {
	err := models.UpdateTaskstate1(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/todo", 302) //这里增加topic
}

// Example:
//
//   req: GET /task/
//   res: 200 {"Tasks": [
//          {"ID": 1, "Title": "Learn Go", "Done": false},
//          {"ID": 2, "Title": "Buy bread", "Done": true}
//        ]}
func (this *TaskController) ListTasks() {
	res := struct{ Tasks []*models.Task }{models.DefaultTaskList.All()}
	this.Data["json"] = res
	this.ServeJSON()
}

// Examples:
//
//   req: POST /task/ {"Title": ""}
//   res: 400 empty title
//
//   req: POST /task/ {"Title": "Buy bread"}
//   res: 200
func (this *TaskController) NewTask() {
	req := struct{ Title string }{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("empty title"))
		return
	}
	// t, err := models.NewTask(req.Title)
	// if err != nil {
	// 	this.Ctx.Output.SetStatus(400)
	// 	this.Ctx.Output.Body([]byte(err.Error()))
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
func (this *TaskController) GetTask() {
	id := this.Ctx.Input.Param(":id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	t, ok := models.DefaultTaskList.Find(intid)
	beego.Info("Found", ok)
	if !ok {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("task not found"))
		return
	}
	this.Data["json"] = t
	this.ServeJSON()
}

// Example:
//
//   req: PUT /task/1 {"ID": 1, "Title": "Learn Go", "Done": true}
//   res: 200
//
//   req: PUT /task/2 {"ID": 2, "Title": "Learn Go", "Done": true}
//   res: 400 inconsistent task IDs
func (this *TaskController) UpdateTask() {
	id := this.Ctx.Input.Param(":id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	var t models.Task
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &t); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	if t.ID != intid {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("inconsistent task IDs"))
		return
	}
	if _, ok := models.DefaultTaskList.Find(intid); !ok {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("task not found"))
		return
	}
	models.DefaultTaskList.Save(&t)
}
